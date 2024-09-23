package main

import (
	"fmt"
	"strings"
)

type QtsPurifer struct {
	basefile  string
	allLines  []string
	poets     map[string]int
	curIndent int
}

func (p *QtsPurifer) AddLine(line string) {
	p.allLines = append(p.allLines, line)
}

func (p *QtsPurifer) ParseFile(filename string) {
	p.loadPoets(`.\Data\唐代詩人總表.txt`)

	p.basefile = filename

	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("[WARN]%v\n", err)
		return
	}

	for pos, line := range lines {
		res := p.parseLine(pos, line)
		p.AddLine(res)
	}

	WriteLines(p.allLines, filename+".txt")
}

func (p *QtsPurifer) parseLine(pos int, line string) string {
	if line == `全唐詩` {
		return ""
	}

	if isTitle(line) {
		title := TrimTitle(line)
		if p.isPoet(title) {
			titlenew := "# " + title
			fmt.Printf("[%d]%s to poet: %s\n", pos, line, titlenew)
			return titlenew
		} else {
			return line
		}
	}

	if p.isPoet(line) {
		titlenew := "# " + line
		fmt.Printf("[%d]%s to poet 2: %s\n", pos, line, titlenew)
		return titlenew
	} else {
		return line
	}
}

func (p *QtsPurifer) isPoet(name string) bool {
	if name == "" {
		return false
	}
	_, exists := p.poets[name]
	return exists
}

func (p *QtsPurifer) loadPoets(filename string) {
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("[WARN]%v\n", err)
		return
	}

	p.poets = make(map[string]int)

	for _, line := range lines {
		if line != "" {
			p.poets[line] = 1
		}
	}

	fmt.Printf("[INFO]%d poets loaded.\n", len(p.poets))
}

func (p *QtsPurifer) ParseFile_ErrBigFile(filename string) {
	s, err := ReadTextFile(filename)
	if err != nil {
		fmt.Printf("[WARN]%v\n", err)
		return
	}

	cleaned := CleanComment(s)
	WriteTextFile(filename+".txt", cleaned)
}

func (p *QtsPurifer) ParseFile3(filename string) {
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("[WARN]%v\n", err)
		return
	}

	arr := []string{}
	for _, line := range lines {
		cleaned := p.CleanCommentInLine(line)
		arr = append(arr, cleaned)
	}

	WriteLines(arr, filename+".txt")
}

func (p *QtsPurifer) CleanCommentInLine(s string) string {
	if s == "" {
		return ""
	}
	rs := []rune(s)
	arr := []rune{}
	for i, r := range rs {

		if r == '〈' {
			if p.curIndent > 0 {
				fmt.Printf("Nested: %s\n", string(rs[i-5:i+5]))
			}
			p.curIndent++
		} else if r == '〉' {
			if p.curIndent == 0 {
				fmt.Printf("Err closed: %s\n", string(rs[i-5:i+5]))
			} else {
				p.curIndent--
			}
		} else {
			if p.curIndent == 0 {
				arr = append(arr, r)
			}
		}
	}

	//fmt.Println("DBG:" + string(arr))
	return string(arr)
}

func TrimTitle(s string) string {
	return strings.Trim(s, " \t【】")
}

func CleanComment(s string) string {
	if s == "" {
		return ""
	}
	rs := []rune(s)
	arr := []rune{}
	indent := 0
	for i, r := range rs {
		if r == '〈' {
			if indent > 0 {
				fmt.Printf("Nested: %s\n", string(rs[i-5:i+5]))
			}
			indent++
		} else if r == '〉' {
			if indent == 0 {
				fmt.Printf("Err closed: %s\n", string(rs[i-5:i+5]))
			} else {
				indent--
			}
		} else {
			if indent == 0 {
				arr = append(arr, r)
			}
		}
	}

	//fmt.Println("DBG:" + string(arr))
	return string(arr)
}

func containsCommentStart(s string) bool {
	return strings.ContainsRune(s, '〈')
}

func containsCommentEnd(s string) bool {
	return strings.ContainsRune(s, '〉')
}

func containsComment(s string) bool {
	return containsCommentStart(s) || containsCommentEnd(s)
}

func isTitle(s string) bool {
	rs := []rune(s)
	if len(rs) <= 2 {
		return false
	}

	return (rs[0] == '【') && (rs[len(rs)-1] == '】')
}
