package main

import (
	"fmt"
	"strings"
)

func runYuanci() {
	var purifer YuanciPurifer
	purifer.ParseFile(`D:\tmp\Yuanci_demo.txt`)
}

type YuanciPurifer struct {
	basefile string
	allLines []string
	curPoet  string
	curTitle string
	curRow   int
	curLines []string
}

func (p *YuanciPurifer) AddLine(line string) {
	p.allLines = append(p.allLines, line)
}

func (p *YuanciPurifer) ParseFile(filename string) {
	p.basefile = filename

	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("[WARN]%v\n", err)
		return
	}

	for pos, line := range lines {
		p.curRow = pos + 1
		ln := strings.TrimSpace(line)
		if len(ln) == 0 {
			//p.AddLine("")
		} else {
			p.parseLine(pos, line)
		}
	}

	if len(p.curLines) != 0 {
		p.finishPoem()
	}

	WriteLines(p.allLines, filename+".txt")
}

func (p *YuanciPurifer) parseLine(pos int, line string) {
	if isVolume(line) {
		fmt.Printf("[DBG][%d]Volume: %s\n", pos+1, line)

		p.finishPoem()
		return
	}

	if isPoetTitle(line) {
		p.finishPoem()

		p.curPoet = parsePoetTitle(pos, line)
		p.AddLine("# " + p.curPoet)
		return
	}

	if isCiTitle(line) {
		p.finishPoem()

		rs := []rune(line)
		title := string(rs[1:])
		p.curTitle = title

		p.AddLine("【" + title + `】`)

	} else {
		p.curLines = append(p.curLines, line)
	}
}

func (p *YuanciPurifer) finishPoem() {
	linenum := len(p.curLines)
	switch linenum {
	case 0:
		{
			//fmt.Printf("[INFO][%d]Invalid poem!\n", p.curRow)
			p.curTitle = ""
		}
	case 1:
		{
			p.AddLine(p.curLines[0])
			p.AddLine("")
			p.curTitle = ""
		}
	default:
		{
			if linenum >= 4 {
				fmt.Printf("[DBG][%d]Lines: %d!\n", p.curRow, linenum)
			}

			p.judgeLines()
		}
	}

	p.curLines = []string{}
}

func (p *YuanciPurifer) judgeLines() {
	counts := []int{}
	for _, line := range p.curLines {
		counts = append(counts, getSentenceCount(line))
	}

	maxCount := 0
	for _, v := range counts {
		if v > maxCount {
			maxCount = v
		}
	}

	maxLine := ""
	for _, line := range p.curLines {
		curcount := getSentenceCount(line)
		if curcount != maxCount {
			p.AddLine("$ " + line)
		} else {
			maxLine = line
		}
	}

	p.AddLine(maxLine)
	p.AddLine("")
}

func getSentenceCount(line string) int {
	linenew := strings.TrimSpace(line)
	if len(linenew) == 0 {
		return 0
	}
	arr := strings.FieldsFunc(linenew, SplitSentence)
	return len(arr)
}

func isVolume(line string) bool {
	return strings.HasPrefix(line, `#`)
}

func isCiTitle(line string) bool {
	return strings.HasPrefix(line, `○`)
}

func isPoetTitle(s string) bool {
	sp := strings.TrimSpace(s)
	rs := []rune(sp)
	if len(rs) <= 2 {
		return false
	}

	return (rs[0] == '【') && (rs[len(rs)-1] == '】')
}

func parsePoetTitle(linenum int, s string) string {
	sp := strings.TrimSpace(s)
	rs := []rune(sp)

	rsp := rs[1 : len(rs)-1]
	sp2 := string(rsp)
	arr := strings.Split(sp2, `：`)
	if len(arr) != 2 {
		fmt.Printf("[INFO][%d]Invalid Poet: %s\n", linenum+1, s)
		return ""
	}

	return arr[0]
}
