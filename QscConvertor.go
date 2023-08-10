package main

import (
	"fmt"
	"log"
	"strings"
)

type QscConv struct {
	curPoet    string
	curTitle   string
	curContent string
	curComment string
	curLineNum int

	allpoets  Poets
	allPoems  ChinesePoems
	allCipais Cipais
}

func (p *QscConv) Init() {
	p.allPoems.Init()
}

func (p *QscConv) convertFile(srcFile string) {
	lines := ReadTxtFile(srcFile)
	p.convertLines(lines, true)
}

func (p *QscConv) convertLines(lines []string, runRhyme bool) {
	p.allCipais.Init("CiPai.txt")
	p.allpoets.Init("SongPoets.txt")

	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	totallines := len(lines)

	prevBlank := false
	for i := 0; i < totallines; i++ {
		line := lines[i]
		linenew := lineFormat(line) // remove comment tag : |< >|

		if len(linenew) == 0 {
			if prevBlank {
				prevBlank = true
				continue
			} else {
				p.MakeNewPoem(i, runRhyme)
				prevBlank = true
				continue
			}
		}

		if prevBlank {
			prevBlank = false
			p.curPoet = linenew

			if !p.allpoets.IsPoet(linenew) {
				fmt.Printf("WARN: [%d]Cannot find poet: %s\n", i, linenew)
			}
			continue
		}

		hascipai, title := p.allCipais.HasCipai(linenew)
		if hascipai {
			// Start a new poem
			p.MakeNewPoem(i, runRhyme)

			p.curTitle = title
			continue
		}

		if IsCommentLine(linenew) {
			p.curComment += linenew + "\n"
		} else {
			p.curContent += linenew
			p.curLineNum++
		}
	}

	p.MakeNewPoem(totallines, runRhyme)
}

// remove prefix and suffix spaces and comments (tag : |< >|)
func lineFormat(line string) string {
	linenew := strings.TrimSpace(line)

	posStart := strings.Index(linenew, "|<")
	if posStart == -1 {
		return linenew
	}

	posEnd := strings.Index(linenew, ">|")
	if posEnd == -1 {
		log.Printf("INFO: Line only has start comment tag: %s\n", line)
		return linenew
	}

	leftPart := SubString(linenew, 0, posStart)
	rightPart := SubString(linenew, posEnd+2, len(linenew))

	return leftPart + rightPart
}

func (p *QscConv) MakeNewPoem(id int, runRhyme bool) {
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

	if runRhyme {
		cp.analyseRhyme()
	}
	p.allPoems.AddPoem(cp)

	p.ClearCurrent()
}

func getPartNumber(cipai string) int {
	switch cipai {
	case "莺啼序":
		return 4
	case "瑞龙吟":
		return 4
	case "十样花":
		return 7
	default:
		return 2
	}
}

func (p *QscConv) ClearCurrent() {
	p.curContent = ""
	p.curTitle = ""
	p.curLineNum = 0
}

// analyse and collect Cipai
func (p *QscConv) analyseCipai(srcFile string) {
	lines := ReadTxtFile(srcFile)

	p.allCipais.Init("CiPai.txt")
	p.allpoets.Init("SongPoets.txt")

	var k2lines Keyword2Lines
	k2lines.Init()

	for idx, line := range lines {
		linenew := strings.TrimSpace(line)

		isCipai, cipainame := p.isCipaiMissed(linenew)
		if isCipai {
			k2lines.AddLine(cipainame, idx+1)
		}
	}

	fmt.Println("-----------------")

	k2lines.DemoPrint()
}

func (p *QscConv) isCipaiMissed(s string) (bool, string) {
	chsize := ChcharLen(s)

	if chsize < 2 {
		return false, s
	}

	if IsCommentLine(s) {
		return false, s
	}

	pos := strings.Index(s, " ")
	if pos != -1 {
		// contains blank
		leftpart := s[:pos]

		if ChcharLen(leftpart) > CIPAI_MAX {
			return false, leftpart
		}

		return !p.allCipais.HasActualCipai(leftpart), leftpart
	}

	pos = strings.Index(s, "（")
	if pos != -1 {
		// contains blank
		leftpart := s[:pos]
		if ChcharLen(leftpart) > CIPAI_MAX {
			return false, leftpart
		}

		return !p.allCipais.HasActualCipai(leftpart), leftpart
	}

	if chsize <= 5 {
		if p.allpoets.IsPoet(s) {
			return false, s
		}

		return !p.allCipais.HasActualCipai(s), s
	} else {
		return false, s
	}
}

func (p *QscConv) PrintRhyme() {
	for _, v := range p.allPoems.ID2Poems {
		fmt.Printf("[%s]: %s\n", v.Rhyme, v.LeftChars(50))
	}
}

func (p *QscConv) FindByCiPai(cipai string) {
	resultcount := 0
	for _, v := range p.allPoems.ID2Poems {
		if v.Title == cipai {
			fmt.Printf("[%s]: %s\n", v.title(), v.LeftChars(75))
			resultcount++
		}
	}

	fmt.Printf("Total %d results.\n", resultcount)
}

func (p *QscConv) FindByYayun(yayun string) {
	resultcount := 0
	for _, v := range p.allPoems.ID2Poems {
		if v.Rhyme == yayun {
			fmt.Printf("[%s]: %s\n", v.title(), v.LeftChars(50))
			resultcount++
		}
	}
	fmt.Printf("Total %d results.\n", resultcount)
}

func (p *QscConv) FindByCiPaiYayun(cipai, yayun string) {
	resultcount := 0
	for _, v := range p.allPoems.ID2Poems {
		if (v.Rhyme == yayun) && (v.Title == cipai) {
			fmt.Printf("[%s]: %s\n", v.toDesc(), v.LeftChars(75))
			resultcount++
		}
	}
	fmt.Printf("Total %d results.\n", resultcount)
}

func (p *QscConv) FindByYayunLength(yayun string, chlen int) {
	for _, v := range p.allPoems.ID2Poems {
		arr := v.FindByYayunLength(yayun, chlen)

		for id, item := range arr {
			fmt.Printf("[%d][%s][%s]\n", id, v.Title, item)
		}
	}
}

// see: ZhRhymes
// chlen: 0 means any
func (p *QscConv) FindByYayunLengthPingze(yayun string, chlen, pztype int) {
	totalResults := 0
	for _, v := range p.allPoems.ID2Poems {
		arr := v.FindByYayunLengthPingze(yayun, chlen, pztype)

		for id, item := range arr {
			fmt.Printf("[%d][%s][%s]\n", id, v.Title, item)
			totalResults++
		}
	}
	fmt.Printf("Total %d results.\n", totalResults)
}

type ZhCharCount struct {
	r2c Rhyme2Count
}

func (p *ZhCharCount) Init() {
	p.r2c.Init()
}

func (p *ZhCharCount) AddPoem(poem *ChinesePoem) {
	for _, s := range poem.Sentences {
		rs := []rune(s)
		for i := 0; i < len(rs)-1; i++ {
			chpair := []rune{rs[i], rs[i+1]}
			p.r2c.Add(string(chpair))
		}
	}
}

func CreateQscPoem(id, author, title, content, comment string) *ChinesePoem {
	cp := new(ChinesePoem)
	cp.ID = id
	cp.Author = author
	cp.Title = title
	cp.AllText = content

	if len(comment) > 0 {
		cp.Comments = comment
	}

	cp.ParseSentences()
	//cp.analyseRhyme()

	return cp
}
