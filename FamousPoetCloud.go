package main

import (
	"fmt"
	"strings"
)

func CreateTangClouds() {
	var fps FamousPoets
	fps.Init()
	fps.CreateWordCloud()
}

type FamousPoets struct {
	allPoets []string
}

func (p *FamousPoets) Init() {
	p.allPoets = strings.Split(`李白,杜甫,白居易,李商隱,杜牧,高適`, ",")
}

func (p *FamousPoets) CreateWordCloud() {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	var pyf PinyinFinder
	pyf.Init(`zht2py.txt`)

	var poet2lines Poet2Lines

	for _, poem := range qtsInst.ID2Poems {
		for _, line := range poem.Sentences {
			poet2lines.AddLine(poem.Author, line)
		}
	}

	for _, poet := range p.allPoets {
		lines := poet2lines.GetLines(poet)
		fmt.Printf("[DBG]%d Lines of Poet %s.\n", len(lines), poet)

		poetpy := pyf.FindStrPinyin(poet)
		fmt.Printf("[DBG]Poet: %s\n", poetpy)
		createWordCloudByLines(lines, poetpy)
	}
}

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

	wc.SaveMultiFiles(outfile)
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
