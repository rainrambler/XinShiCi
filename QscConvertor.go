// QscConvertor
package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	CIPAI_MAX = 7
)

type Poets struct {
	poet2id map[string]int
}

func (p *Poets) Init(filename string) {
	p.poet2id = make(map[string]int)

	lines := ReadTxtFile(filename)

	id := 100
	for _, line := range lines {
		if len(line) > 0 {
			p.poet2id[line] = id + 2
		}
	}
}

func (p *Poets) IsPoet(nm string) bool {
	if _, ok := p.poet2id[nm]; ok {
		return true
	}
	return false
}

func (p *Poets) FindPoet(nm string) int {
	if id, ok := p.poet2id[nm]; ok {
		return id
	}
	return -1
}

func (p *Poets) Count() int {
	return len(p.poet2id)
}

type Cipais struct {
	item2id map[string]int
}

func (p *Cipais) Init(filename string) {
	p.item2id = make(map[string]int)

	lines := ReadTxtFile(filename)

	id := 100
	for _, line := range lines {
		if len(line) > 0 {
			p.item2id[line] = id + 2
		}
	}

	fmt.Printf("INFO: Total CiPai: %d\n", len(p.item2id))
}

func (p *Cipais) Exists(nm string) bool {
	if _, ok := p.item2id[nm]; ok {
		return true
	}
	return false
}

func (p *Cipais) Count() int {
	return len(p.item2id)
}

func (p *Cipais) HasCipai(line string) (bool, string) {
	s := line
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
			return false, s
		}

		return p.HasActualCipai(leftpart), leftpart
	}

	pos = strings.Index(s, "（")
	if pos != -1 {
		// contains bracket
		leftpart := s[:pos]
		if ChcharLen(leftpart) > CIPAI_MAX {
			return false, leftpart
		}

		return p.HasActualCipai(leftpart), leftpart
	}

	if chsize > CIPAI_MAX {
		return false, s
	}

	return p.HasActualCipai(s), s
}

func (p *Cipais) HasActualCipai(line string) bool {
	return p.Exists(line)
}

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
		log.Printf("DBG: [%d][%s] Lines: (%d): %s\n", id, p.curTitle, p.curLineNum, SubChineseString(p.curContent, 0, 7))
	}

	poemId := fmt.Sprintf("%d-%d", poetId, id)

	cp := CreateQscPoem(poemId, p.curPoet, p.curTitle, p.curContent, p.curComment)

	if runRhyme {
		cp.analyseRhyme()
	}
	p.allPoems.AddPoem(cp)

	p.ClearCurrent()
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
		fmt.Printf("[%s]: %s\n", v.Rhyme, SubChineseString(v.AllText, 0, 15))
	}
}

func (p *QscConv) FindByCiPai(cipai string) {
	for _, v := range p.allPoems.ID2Poems {
		if v.Title == cipai {
			fmt.Printf("[%s]: %s\n", v.toDesc(), SubChineseString(v.AllText, 0, 15))
		}
	}
}

func (p *QscConv) FindByYayun(yayun string) {
	for _, v := range p.allPoems.ID2Poems {
		if v.Rhyme == yayun {
			fmt.Printf("[%s]: %s\n", v.toDesc(), v.AllText)
			//fmt.Printf("[%s]: %s\n", v.toDesc(), SubChineseString(v.AllText, 0, 65))
		}
	}
}

func (p *QscConv) FindByCiPaiYayun(cipai, yayun string) {
	for _, v := range p.allPoems.ID2Poems {
		if (v.Rhyme == yayun) && (v.Title == cipai) {
			fmt.Printf("[%s]: %s\n", v.toDesc(), SubChineseString(v.AllText, 0, 75))
		}
	}
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
	for _, v := range p.allPoems.ID2Poems {
		arr := v.FindByYayunLengthPingze(yayun, chlen, pztype)

		for id, item := range arr {
			fmt.Printf("[%d][%s][%s]\n", id, v.Title, item)
		}
	}
}

func (p *QscConv) FindSentense(qc *QueryCondition) {
	for _, v := range p.allPoems.ID2Poems {
		for _, sentence := range v.Sentences {
			if qc.ZhLen > 0 {
				if qc.ZhLen != ChcharLen(sentence) {
					continue
				}
			}

			switch qc.Pos {
			case POS_PREFIX:
				if strings.HasPrefix(sentence, qc.KeywordStr) {
					fmt.Printf("%s [%s]\n", sentence, v.toDesc())
				}
			case POS_SUFFIX:
				if strings.HasSuffix(sentence, qc.KeywordStr) {
					fmt.Printf("%s [%s]\n", sentence, v.toDesc())
				}
			case POS_ANY:
				if strings.Contains(sentence, qc.KeywordStr) {
					fmt.Printf("%s [%s]\n", sentence, v.toDesc())
				}
			default:

			}
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
