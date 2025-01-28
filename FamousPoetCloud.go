package main

import (
	"fmt"
	"strings"
)

func CreateTangClouds() {
	var fps FamousPoets
	fps.Init(`李白,杜甫,白居易,李商隱,杜牧,高適`)
	fps.CreateWordCloud(`qts_zht.txt`)
}

func CreateLiYuClouds() {
	var fps FamousPoets
	poetName := `李煜`
	fps.Init(poetName)

	filename := `LiYuCi.txt`
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("INFO: Cannot read file %s: %v!\n", filename, err)
		return
	}

	var poet2lines Poet2Lines
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		poet2lines.AddLine(poetName, line)
	}

	fps.convertLines(poet2lines)
}

type FamousPoets struct {
	allPoets []string
}

func (p *FamousPoets) Init(poets string) {
	p.allPoets = strings.Split(poets, ",")
}

func (p *FamousPoets) CreateWordCloud(contentfile string) {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile(contentfile)

	var poet2lines Poet2Lines

	for _, poem := range qtsInst.ID2Poems {
		for _, line := range poem.Sentences {
			poet2lines.AddLine(poem.Author, line)
		}
	}

	p.convertLines(poet2lines)
}

func (p *FamousPoets) convertLines(poet2lines Poet2Lines) {
	var pyf PinyinFinder
	pyf.Init(`zht2py.txt`)

	for _, poet := range p.allPoets {
		lines := poet2lines.GetLines(poet)
		fmt.Printf("[DBG]%d Lines of Poet %s.\n", len(lines), poet)

		poetpy := pyf.FindStrPinyin(poet)
		fmt.Printf("[DBG]Poet: %s\n", poetpy)
		createWordCloudByLines(lines, poetpy)
	}
}

const DESIRED_WORD_COUNT = 200

// outfile: "dufu" --> "dufu_2_30.html"
func createWordCloudByLines(lines []string, outfile string) {
	var wc WordCloud
	wc.InitParams()
	for _, line := range lines {
		wc.parseSentence(line)
	}

	//wc.PrintResult(500)
	if outfile == "" {
		return
	}

	wc.SaveFilesAutoCount(outfile, DESIRED_WORD_COUNT)
}

type Lines struct {
	allLines []string
}

func (p *Lines) AddLine(line string) {
	p.allLines = append(p.allLines, line)
}

type Poet2Lines struct {
	poet2lines map[string]*Lines
	inited     bool
}

func (p *Poet2Lines) AddLine(poet, line string) {
	if !p.inited {
		p.poet2lines = make(map[string]*Lines)
		p.inited = true
	}

	ls, exists := p.poet2lines[poet]
	if exists {
		ls.AddLine(line)
	} else {
		lsnew := new(Lines)
		lsnew.AddLine(line)
		p.poet2lines[poet] = lsnew
	}
}

func (p *Poet2Lines) GetLines(poet string) []string {
	poems, exists := p.poet2lines[poet]
	if exists {
		return poems.allLines
	}
	return []string{}
}
