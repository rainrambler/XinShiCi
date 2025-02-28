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
	allpoets  AncientPoets
	allPoems  ChinesePoems
	runRhyme  bool
	prevPoet  bool
	preTitle  bool

	curContent   string
	curComment   string
	curLineNum   int
	titleLineNum int
}

func (p *QscZhtLoader) loadFile(srcFile string) {
	g_Rhymes.ImportFile("ShiYunXinBianZH.txt")

	p.allCipais.Init("CiPaiZh.txt")
	p.allpoets.Init(`Data/AncientAuthors.txt`)
	p.allPoems.Init()
	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	lines := ReadTxtFile(srcFile)
	p.runRhyme = true
	p.parseLines(lines)

	fmt.Printf("INFO: %d poems loaded.\n", p.allPoems.Count())
	//p.allPoems.PrintResults()
}

func (p *QscZhtLoader) parseLines(lines []string) {
	totallines := len(lines)
	for i := 0; i < totallines; i++ {
		line := lines[i]
		//fmt.Printf("[DBG][%d]: %s\n", i+1, line)

		if IsEmptyLine(line) {
			//p.curContent += "\r\n"
			continue
		}

		firstchar := GetFirstRune(line)
		switch firstchar {
		case '#':
			{
				p.CommitPoem(i - 1)
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
				p.CommitPoem(i - 1)
				p.beginNewPoem(line)
			}
		default:
			p.addLine(i, line)
		}
	}

	p.CommitPoem(totallines)
}

func (p *QscZhtLoader) addLine(pos int, line string) {
	linenew := TrimBlank(line)
	if len(linenew) == 0 {
		return
	}

	lastchar := GetLastRune(linenew)
	if !IsPunctuation(lastchar) {
		fmt.Printf("[%d]Possible sub-title: %s\n", pos, line)
		linenew += "。"
	}
	p.curContent += linenew
}

func TrimBlank(s string) string {
	return strings.Trim(s, " \t\r\n")
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
		//fmt.Printf("DBG: No title in line: %d, Poet: %s, [%s]%s\n", pos,
		//	p.curPoet, p.curTitle, p.curContent)
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
