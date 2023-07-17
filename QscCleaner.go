// QscConvertor
package main

import (
	"fmt"
	"strings"
)

type QscCleaner struct {
	curPoet   string
	curTitle  string
	isPoet    bool
	allCipais Cipais
	allpoets  Poets
}

func (p *QscCleaner) convertFile(srcFile string) {
	lines := ReadTxtFile(srcFile)
	p.convertLines(lines, srcFile+".txt")
}

func (p *QscCleaner) convertLines(lines []string, tofile string) {
	p.allCipais.Init("CiPaiZh.txt")
	p.allpoets.Init("SongPoetsZh.txt")

	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	totallines := len(lines)
	arr := []string{}
	cleaned := []string{}

	for i := 0; i < totallines; i++ {
		line := lines[i]

		if IsCommentLine(line) {
			// # 林逋
			p.curPoet = line
			p.isPoet = true
			arr = append(arr, line)
			continue
		}

		if p.isPoet && (len(line) != 0) {
			// Poet Desc?
			linenew := strings.Trim(line, " 	\r\n")
			if p.allCipais.HasActualCipai(linenew) {
				//fmt.Printf("Line %d: Cipai %s\n", i, line)
				arr = append(arr, line)
			} else {
				desc := fmt.Sprintf("[%d] %s", i, line)
				cleaned = append(cleaned, desc)
			}
			p.isPoet = false
		} else {
			arr = append(arr, line)
		}
	}

	WriteLines(arr, tofile)
	WriteLines(cleaned, "cleaned.txt")
}

type QscChanger struct {
	allCipais Cipais
	allpoets  Poets
	curLine   int
	curCipai  string
	leftPart  string
	contents  []string
}

func (p *QscChanger) ChangeSubtitle(srcFile string) {
	lines := ReadTxtFile(srcFile)
	p.contents = []string{}
	p.changeLines(lines, srcFile+".txt")
}

func (p *QscChanger) changeLines(lines []string, tofile string) {
	p.allCipais.Init("CiPaiZh.txt")
	p.allpoets.Init("SongPoetsZh.txt")

	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	totallines := len(lines)
	fmt.Printf("INFO: Total lines: %d\n", totallines)

	for i := 0; i < totallines; i++ {
		p.curLine = i + 1
		line := lines[i]
		//fmt.Println(line)

		if p.changeLineEn(line) {
			p.sumSubtitle()
		} else if p.changeLineZh(line) {
			p.sumSubtitle()
		} else {
			p.contents = append(p.contents, line)
		}
	}

	WriteLines(p.contents, tofile)
}

func (p *QscChanger) changeLineEn(line string) bool {
	startpos := strings.Index(line, "(")
	if startpos == -1 {
		return false
	}

	endpos := strings.Index(line, ")")
	if endpos == -1 {
		fmt.Printf("Err format: %s!\n", line)
		return false
	}

	part := line[startpos+1 : endpos]
	if p.allCipais.HasActualCipai(part) {
		p.curCipai = part
		p.leftPart = line[:startpos]

		//fmt.Printf("[DBG]EN: Cipai: %s, Left: %s in line: %s\n", p.curCipai, p.leftPart, line)
		return true
	}

	return false
}

func (p *QscChanger) changeLineZh(line string) bool {
	startpos := strings.Index(line, "（")
	if startpos == -1 {
		return false
	}

	endpos := strings.Index(line, "）")
	if endpos == -1 {
		fmt.Printf("Err format: %s in line %d!\n", line, p.curLine)
		return false
	}

	part := line[startpos+len("（") : endpos]
	p.curCipai = part
	p.leftPart = line[:startpos]

	if p.allCipais.HasActualCipai(part) {
		//fmt.Printf("[DBG]ZH: Cipai: %s, Left: %s in line: %s\n", p.curCipai, p.leftPart, line)
		return true
	}

	//fmt.Printf("Cipai not found: %s in line %d: %s\n", part, p.curLine, line)
	return false
}

func (p *QscChanger) sumSubtitle() {
	if len(p.contents) == 0 {
		p.contents = append(p.contents, p.curCipai)
		p.curCipai = ""
		p.contents = append(p.contents, p.leftPart)
		p.leftPart = ""
		return
	}
	pos := len(p.contents) - 1
	pos0 := pos
	for pos > 0 {
		s := p.contents[pos]

		if ContainsChPunctions(s) {
			break
		}
		pos--
	}

	if pos == pos0 {
		p.contents = append(p.contents, p.curCipai)
		p.curCipai = ""
		p.contents = append(p.contents, p.leftPart)
		p.leftPart = ""
		return
	}

	//fmt.Printf("[DBG]Pos: %d, len: %d\n", pos, pos0+1)

	arr := p.contents[pos+1:]
	p.contents = p.contents[:pos+1]
	//fmt.Printf("[DBG]Contents: %v\n", p.contents)
	//fmt.Printf("[DBG]Subtitle: %v\n", arr)

	s := ""
	for _, item := range arr {
		s += item
	}

	s += p.leftPart
	p.leftPart = ""
	p.contents = append(p.contents, p.curCipai)
	p.curCipai = ""
	p.contents = append(p.contents, s)
}
