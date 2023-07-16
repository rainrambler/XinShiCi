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
	prevPoet  bool

	curContent   string
	curComment   string
	curLineNum   int
	titleLineNum int

	cipai2count Rhyme2Count
}

func (p *QscZht) convertFile(srcFile string) {
	p.allCipais.Init("CiPaiZh.txt")
	p.allpoets.Init("SongPoetsZh.txt")
	p.allPoems.Init()
	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	p.cipai2count.Init()

	lines := ReadTxtFile(srcFile)
	p.runRhyme = true
	p.parseLines(lines, srcFile+".txt")
	//p.convertLines(lines)

	//p.allPoems.PrintResults()

	p.cipai2count.PrintSorted()
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
						p.cipai2count.Add(linenew)
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

		if len(line) == 0 {
			arr = append(arr, line)
			continue
		}

		if strings.HasPrefix(line, "!") {
			arr = append(arr, line)
			continue
		}

		if strings.HasPrefix(line, "#") {
			p.MakeNewPoem(i)
			// # 林逋
			linenew := strings.Trim(line, "\t #")
			p.curPoet = linenew
			p.prevPoet = true
			if !p.allpoets.IsPoet(linenew) {
				fmt.Printf("WARN: Cannot find poet [%s] in line: [%d]\n", linenew, i)
			}

			arr = append(arr, line)
			continue
		}

		linenew := strings.Trim(line, " 	\r\n")
		if p.allCipais.HasActualCipai(linenew) {
			//fmt.Printf("Line %d: Cipai %s\n", i, line)
			p.MakeNewPoem(i)
			p.curTitle = linenew
			p.titleLineNum = i
			arr = append(arr, `【`+linenew+`】`) // Title
		} else {
			if p.prevPoet {
				arr = append(arr, `* `+line) // Author Desc
				p.prevPoet = false
			} else if ContainsChPunctions(linenew) {
				p.curContent += linenew
				arr = append(arr, line)

				if !strings.HasSuffix(linenew, "。") {
					//fmt.Printf("Frag Line %d: %s\n", i+1, linenew)
				}
			} else {
				if i-1 == p.titleLineNum {
					p.curComment = linenew
					arr = append(arr, `$ `+line) // sub-title
				} else {
					if isSequenceCipai(linenew) {
						arr = append(arr, `【`+linenew+`】`) // Title
						p.curTitle = linenew
						p.titleLineNum = i
					} else {
						fmt.Printf("Possible cipai in %d: %s\n", i, linenew)
						p.cipai2count.Add(linenew)
					}
				}
			}
		}
	}

	p.MakeNewPoem(totallines)
	WriteLines(arr, tofile)
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

func isSequenceCipai(cipainame string) bool {
	if cipainame == `又` {
		return true
	}

	arr := []string{`一`, `二`, `三`, `四`, `五`, `六`, `七`, `八`, `九`, `十`, `十一`,
		`十二`, `十三`}
	for _, item := range arr {
		s := `其` + item
		if cipainame == s {
			return true
		}
	}

	// 二、三、……
	for i := 1; i < len(arr); i++ {
		if cipainame == arr[i] {
			return true
		}
	}

	// 第二
	for _, item := range arr {
		s := `第` + item
		if cipainame == s {
			return true
		}
	}

	for _, item := range arr {
		s := `右` + item
		if cipainame == s {
			return true
		}
	}

	return false
}
