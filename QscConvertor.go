// QscConvertor
package main

import (
	"fmt"
	"log"
	"strings"
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

func (p *Cipais) HasCipai(line string) bool {
	firstpart := SubChineseString(line, 0, 3)
	//fmt.Printf("DBG:First part: [%s]\n", firstpart)
	if p.Exists(firstpart) {
		return true
	}

	if ChcharLen(line) >= 5 {
		firstpart = SubChineseString(line, 0, 5)
		if p.Exists(firstpart) {
			return true
		}
	}

	if ChcharLen(line) >= 4 {
		firstpart = SubChineseString(line, 0, 4)
		if p.Exists(firstpart) {
			return true
		}
	}

	firstpart = SubChineseString(line, 0, 2)
	if p.Exists(firstpart) {
		return true
	}

	return false
}

type QscConv struct {
	curPoet    string
	curTitle   string
	curLine    int
	curContent string
	curComment string

	allpoets Poets
	allPoems ChinesePoems
}

func (p *QscConv) Init() {
	p.allPoems.Init()
}

func (p *QscConv) convertFile(srcFile string) {
	lines := ReadTxtFile(srcFile)
	p.convertLines(lines)
}

func (p *QscConv) convertLines(lines []string) {
	var allCipai Cipais

	allCipai.Init("CiPai.txt")
	p.allpoets.Init("SongPoets.txt")

	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	totallines := len(lines)

	prevBlank := false
	for i := 0; i < totallines; i++ {
		line := lines[i]
		linenew := strings.TrimSpace(line)

		//fmt.Printf("Before: [%s], After: [%s]\n", line, linenew)

		if len(linenew) == 0 {
			prevBlank = true
			continue
		}

		if prevBlank {
			prevBlank = false

			p.MakeNewPoem(i)

			p.curPoet = linenew

			if !p.allpoets.IsPoet(linenew) {
				fmt.Printf("WARN: Cannot find poet: %s\n", linenew)
			}
			continue
		}

		if allCipai.HasCipai(linenew) {

			p.MakeNewPoem(i)

			p.curTitle = linenew
			continue
		}

		//p.curContent += linenew + "\n"
		if IsCommentLine(linenew) {
			p.curComment += linenew + "\n"
		} else {
			p.curContent += linenew
		}
	}

	p.MakeNewPoem(totallines)
}

func (p *QscConv) MakeNewPoem(id int) {
	if p.curPoet == "" {
		return
	}
	if p.curTitle == "" {
		return
	}
	if p.curContent == "" {
		return
	}
	poetId := p.allpoets.FindPoet(p.curPoet)
	if poetId < 0 {
		log.Printf("DBG: Cannot find poet: %s\n", p.curPoet)
		return
	}

	poemId := fmt.Sprintf("%d-%d", poetId, id)

	cp := CreateQscPoem(poemId, p.curPoet, p.curTitle, p.curContent, p.curComment)
	p.allPoems.AddPoem(cp)

	p.ClearCurrent()
}

func (p *QscConv) ClearCurrent() {
	p.curContent = ""
	p.curLine = 0
	p.curTitle = ""
}

// analyse and collect Cipai
func (p *QscConv) analyseCipai(srcFile string) {
	lines := ReadTxtFile(srcFile)

	var allCipai Cipais

	allCipai.Init("CiPai.txt")
	p.allpoets.Init("SongPoets.txt")

	linenums := []int{}
	idx2nm := map[int]string{}
	nm2idx := map[string]int{}

	for idx, line := range lines {
		linenew := strings.TrimSpace(line)

		chsize := ChcharLen(linenew)

		if (chsize < 2) || (chsize > 5) {
			continue
		}

		if p.allpoets.IsPoet(linenew) {
			continue
		}

		if allCipai.HasCipai(linenew) {
			continue
		}

		if IsCommentLine(linenew) {
			continue
		}

		idx2nm[idx] = linenew
		nm2idx[linenew] = idx
		linenums = append(linenums, idx)
	}

	for _, curLine := range linenums {
		fmt.Printf("[%d]%s\n", curLine+1, idx2nm[curLine])
	}

	fmt.Printf("Total %d lines, %d are unique.\n", len(linenums), len(nm2idx))
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

	return cp
}
