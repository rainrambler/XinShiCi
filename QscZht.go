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
	allCipais Cipais
	allpoets  Poets
	allPoems  ChinesePoems
	runRhyme  bool

	curContent   string
	curComment   string
	curLineNum   int
	titleLineNum int
}

func (p *QscZht) convertFile(srcFile string) {
	p.allCipais.Init("CiPaiZh.txt")
	p.allpoets.Init("SongPoetsZh.txt")
	p.allPoems.Init()
	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	lines := ReadTxtFile(srcFile)
	p.runRhyme = true
	p.convertLines(lines)

	//p.allPoems.PrintResults()
}

func (p *QscZht) convertLines(lines []string) {
	totallines := len(lines)

	for i := 0; i < totallines; i++ {
		line := lines[i]

		if IsCommentLine(line) {
			p.MakeNewPoem(i)
			// # 林逋
			linenew := strings.Trim(line, "\t #")
			p.curPoet = linenew
			if !p.allpoets.IsPoet(linenew) {
				fmt.Printf("WARN: Cannot find poet [%s] in line: [%d]\n", linenew, i)
			}
			continue
		}

		if len(line) != 0 {
			linenew := strings.Trim(line, " 	\r\n")
			if p.allCipais.HasActualCipai(linenew) {
				//fmt.Printf("Line %d: Cipai %s\n", i, line)
				p.MakeNewPoem(i)
				p.curTitle = linenew
				p.titleLineNum = i
			} else {
				if ContainsChPunctions(linenew) {
					p.curContent += linenew
				} else {
					if i-1 == p.titleLineNum {
						p.curComment = linenew
					} else {
						fmt.Printf("Possible cipai in %d: %s\n", i, linenew)
					}
				}
			}
		} else {
		}
	}

	p.MakeNewPoem(totallines)
}

func (p *QscZht) parseLines(lines []string, tofile string) {
	arr := []string{}
	totallines := len(lines)
	for i := 0; i < totallines; i++ {
		line := lines[i]

		if IsCommentLine(line) {
			p.MakeNewPoem(i)
			// # 林逋
			linenew := strings.Trim(line, "\t #")
			p.curPoet = linenew
			if !p.allpoets.IsPoet(linenew) {
				fmt.Printf("WARN: Cannot find poet [%s] in line: [%d]\n", linenew, i)
			}

			arr = append(arr, line)
			continue
		}

		if len(line) != 0 {
			linenew := strings.Trim(line, " 	\r\n")
			if p.allCipais.HasActualCipai(linenew) {
				//fmt.Printf("Line %d: Cipai %s\n", i, line)
				p.MakeNewPoem(i)
				p.curTitle = linenew
				p.titleLineNum = i
				arr = append(arr, `【`+line+`】`)
			} else {
				if ContainsChPunctions(linenew) {
					p.curContent += linenew
				} else {
					if i-1 == p.titleLineNum {
						p.curComment = linenew
					} else {
						fmt.Printf("Possible cipai in %d: %s\n", i, linenew)
					}
				}
			}
		} else {
			arr = append(arr, line)
		}
	}

	p.MakeNewPoem(totallines)
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
		fmt.Printf("DBG: [%d]Cannot find poet: %s\n", id, p.curPoet)
		return
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
	p.curComment = ""
	p.curLineNum = 0
	p.titleLineNum = 0
}
