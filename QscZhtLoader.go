// QscConvertor
package main

import (
	"fmt"
	"log"
	"strings"
)

type QscZhtLoader struct {
	curPoet   string
	curTitle  string
	allCipais Cipais
	allpoets  Poets
	allPoems  ChinesePoems
	runRhyme  bool
	prevPoet  bool
	preTitle  bool

	curContent   string
	curComment   string
	curLineNum   int
	titleLineNum int
}

func (p *QscZhtLoader) convertFile(srcFile string) {
	p.allCipais.Init("CiPaiZh.txt")
	p.allpoets.Init("SongPoetsZh.txt")
	p.allPoems.Init()
	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	lines := ReadTxtFile(srcFile)
	p.runRhyme = true
	p.parseLines(lines, srcFile+".txt")
	//p.allPoems.PrintResults()
}

func (p *QscZhtLoader) parseLines(lines []string, tofile string) {
	arr := []string{}
	totallines := len(lines)
	for i := 0; i < totallines; i++ {
		line := lines[i]
		//fmt.Printf("[DBG][%d]: %s\n", i+1, line)

		if IsEmptyLine(line) {
			arr = append(arr, line)
			continue
		}

		firstchar := GetFirstRune(line)
		switch firstchar {
		case '#':
			{
				p.beginNewPoet(line)
			}
		case '!':
			{
				p.curComment += line
			}
		case '$':
			{
				// sub-title
			}
		case '*':
			{
				// author desc
			}
		case '【':
			{
				// title
				p.beginNewPoem(line)
			}
		default:
			p.curContent += line + "\r\n"
		}

	}

	p.CommitPoem(totallines)
	WriteLines(arr, tofile)
}

func (p *QscZhtLoader) beginNewPoet(line string) {
	s := strings.Trim(line, " \t#")
	p.curPoet = s
}

func (p *QscZhtLoader) beginNewPoem(line string) {
	s := strings.Trim(line, " \t【】")
	p.curTitle = s
}

func (p *QscZhtLoader) CommitPoem(pos int) {
	if p.curPoet == "" {
		//fmt.Printf("DBG: Cannot find author in line: %d\n", pos)
		return
	}
	if p.curTitle == "" {
		//fmt.Printf("DBG: Cannot find title in line: %d\n", pos)
		return
	}
	if p.curContent == "" {
		log.Printf("DBG: Cannot find content in line: %d\n", pos)
		return
	}
	poetId := p.allpoets.FindPoet(p.curPoet)
	if poetId < 0 {
		fmt.Printf("DBG: [%d]Cannot find poet: %s\n", pos, p.curPoet)
		return
	}

	poemId := fmt.Sprintf("%d-%d", poetId, pos)
	cp := CreateQscPoem(poemId, p.curPoet, p.curTitle, p.curContent, p.curComment)

	if p.runRhyme {
		cp.analyseRhyme()
	}
	p.allPoems.AddPoem(cp)

	p.ClearCurrent()
}

func (p *QscZhtLoader) setNewTitle(pos int, line string) {
	p.curTitle = line
	p.titleLineNum = pos
	p.preTitle = true
}

func (p *QscZhtLoader) ClearCurrent() {
	p.curContent = ""
	p.curTitle = ""
	p.curComment = ""
	p.curLineNum = 0
	p.titleLineNum = 0
}
