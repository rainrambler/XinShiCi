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
	preTitle  bool

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
	p.parseLines(lines, srcFile+".txt")
	//p.allPoems.PrintResults()
}

func (p *QscZht) parseLines(lines []string, tofile string) {
	arr := []string{}
	totallines := len(lines)
	for i := 0; i < totallines; i++ {
		line := lines[i]

		if IsEmptyLine(line) {
			arr = append(arr, line)
			continue
		}

		if strings.HasPrefix(line, "!") {
			//fmt.Printf("Comment line at %d\n", i+1)
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
			p.prevPoet = false
			if p.preTitle {
				fmt.Printf("Line: [%d] repeat title: %s\n", i, line)
				//fmt.Printf("DBG: subtitle in line %d: %s\n", i+1, line)
				p.curComment = linenew
				arr = append(arr, `$ `+line) // sub-title
				p.preTitle = false
			} else {
				//fmt.Printf("DBG: Actual Cipai: %s\n", linenew)
				p.MakeNewPoem(i)
				p.setNewTitle(i, linenew)
				arr = append(arr, packCipai(linenew)) // Title
			}

		} else {
			if p.prevPoet {
				arr = append(arr, `* `+line) // Author Desc
				//fmt.Printf("DBG: Author desc: %s\n", line)
				p.prevPoet = false
				p.preTitle = false
			} else if ContainsChPunctions(linenew) {
				//fmt.Printf("DBG: content: %s\n", linenew)
				p.curContent += linenew
				arr = append(arr, line)

				if !strings.HasSuffix(linenew, "。") {
					fmt.Printf("Frag Line %d: %s\n", i+1, linenew)
				}
				p.preTitle = false
			} else {
				if p.preTitle {
					//fmt.Printf("DBG: sub-title in line %d: %s\n", i+1, line)
					p.curComment = linenew
					arr = append(arr, `$ `+line) // sub-title
					p.preTitle = false
				} else {
					iscipai, cipai := isLineSequenceCipai(linenew)
					if iscipai {
						//fmt.Printf("DBG: Parsed Cipai: %s in %s\n", cipai, linenew)
						arr = append(arr, packCipai(cipai)) // Title
						if linenew != cipai {
							arr = append(arr, `! `+line) // convert to comment
						}
						p.setNewTitle(i, linenew)
					} else {
						fmt.Printf("Possible cipai in %d: %s\n", i, linenew)
						arr = append(arr, line) // TODO confirm
					}
				}
			}
		}
	}

	p.MakeNewPoem(totallines)
	WriteLines(arr, tofile)
}

func (p *QscZht) setNewTitle(pos int, line string) {
	p.curTitle = line
	p.titleLineNum = pos
	p.preTitle = true
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
	if cipainame == "" {
		return false
	}
	if cipainame == `又` {
		return true
	}

	if cipainame == `同前` {
		return true
	}

	if cipainame == `同上` {
		return true
	}

	arr := []string{`一`, `二`, `三`, `四`, `五`, `六`, `七`, `八`, `九`, `十`, `十一`,
		`十二`, `十三`, `十四`, `十五`}
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

// `第三 蓬萊景` ==> true
func isLineSequenceCipai(line string) (bool, string) {
	arr := SplitBlank(line)
	cipai := ""
	switch len(arr) {
	case 0:
		fmt.Printf("Err empty line: %s\n", line)
		return false, line
	case 1:
		cipai = line
	case 2:
		cipai = arr[0]
	default:
		fmt.Printf("Err line: %s\n", line)
		return false, line
	}

	return isSequenceCipai(cipai), cipai
}
