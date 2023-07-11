// QscConvertor
package main

import (
	"fmt"
	"log"
	"strings"
)

type QscZht struct {
	curPoet   string
	curTitle  string
	isPoet    bool
	allCipais Cipais
	allpoets  Poets
	allPoems  ChinesePoems
	runRhyme  bool

	curContent string
	curComment string
	curLineNum int
}

func (p *QscZht) convertFile(srcFile string) {
	p.allCipais.Init("CiPaiZh.txt")
	p.allpoets.Init("SongPoetsZh.txt")
	p.allPoems.Init()
	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	lines := ReadTxtFile(srcFile)
	p.runRhyme = true
	p.convertLines(lines)
}

func (p *QscZht) convertLines(lines []string) {
	totallines := len(lines)

	for i := 0; i < totallines; i++ {
		line := lines[i]

		if IsCommentLine(line) {
			// # 林逋
			linenew := strings.Trim(line, "\t #")
			p.curPoet = linenew
			p.isPoet = true

			if !p.allpoets.IsPoet(linenew) {
				fmt.Printf("WARN: Cannot find poet [%s] in line: [%d]\n", line, i)
			}
			continue
		}

		if p.isPoet && (len(line) != 0) {
			// Poet Desc?
			linenew := strings.Trim(line, " 	\r\n")
			if p.allCipais.HasActualCipai(linenew) {
				//fmt.Printf("Line %d: Cipai %s\n", i, line)
				p.MakeNewPoem(i)
				fmt.Printf("[DBG1][%d]%s\n", i, linenew)
			} else {
				fmt.Printf("[DBG2][%d]%s\n", i, linenew)
			}
			p.isPoet = false
		} else {
			fmt.Printf("[DBG3][%d]%s\n", i, line)
		}
	}
}

func (p *QscZht) MakeNewPoem(id int) {
	if p.curPoet == "" {
		//fmt.Printf("DBG: Cannot find author in line: %d\n", id)
		return
	}
	if p.curTitle == "" {
		//fmt.Printf("DBG: Cannot find title in line: %d\n", id)
		return
	}
	if p.curContent == "" {
		log.Printf("DBG: Cannot find content in line: %d\n", id)
		return
	}
	poetId := p.allpoets.FindPoet(p.curPoet)
	if poetId < 0 {
		log.Printf("DBG: [%d]Cannot find poet: %s\n", id, p.curPoet)
		return
	}

	if p.curLineNum > 3 {
		if getPartNumber(p.curTitle) <= 2 {
			/*
				// TODO confirm
				log.Printf("DBG: [%d][%s] Lines: (%d): %s\n",
					id, p.curTitle, p.curLineNum, SubChineseString(p.curContent, 0, 7))
			*/
		}
	}

	poemId := fmt.Sprintf("%d-%d", poetId, id)
	cp := CreateQscPoem(poemId, p.curPoet, p.curTitle, p.curContent, p.curComment)

	if p.runRhyme {
		cp.analyseRhyme()
	}
	p.allPoems.AddPoem(cp)

	p.ClearCurrent()
}

func (p *QscZht) ClearCurrent() {
	p.curContent = ""
	p.curTitle = ""
	p.curLineNum = 0
}
