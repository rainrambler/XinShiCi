package main

import (
	"log"
	"strings"
)

type Rhyme struct {
	Desc   string
	rhymes []string
}

func (p *Rhyme) AddItem(rhy string) {
	if len(rhy) == 0 {
		return
	}

	p.Desc = p.Desc + rhy + "|"
	p.rhymes = append(p.rhymes, rhy)
}

func (p *Rhyme) toDesc() string {
	return p.Desc
}

type ChineseRhymes struct {
	ZhChar2Rhyme map[string]*Rhyme // "安" -> "【十四寒】"
}

func (p *ChineseRhymes) Init() {
	p.ZhChar2Rhyme = make(map[string]*Rhyme)
}

func (p *ChineseRhymes) AddRhyme(zhch, rhyme string) {
	// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	if _, ok := p.ZhChar2Rhyme[zhch]; ok {
		// exists
		curRhyme := p.ZhChar2Rhyme[zhch]
		curRhyme.AddItem(rhyme)
		return
	}

	curRhyme := new(Rhyme)
	curRhyme.AddItem(rhyme)
	p.ZhChar2Rhyme[zhch] = curRhyme
}

func (p *ChineseRhymes) ImportFile(filename string) {
	p.Init()

	lines := ReadTxtFile(filename)

	for idx, line := range lines {
		p.parseLine(idx+1, line)
	}
}

func (p *ChineseRhymes) parseLine(rownum int, line string) {
	if !strings.HasPrefix(line, "【") {
		log.Printf("WARN: Invalid line in Rhyme file (No Start): %d\n", rownum)
		return
	}

	pos := strings.Index(line, "】")

	if pos == -1 {
		log.Printf("WARN: Invalid line in Rhyme file (No End): %d\n", rownum)
		return
	}

	rhymestr := SubString(line, ZH_CHAR_LEN, pos-ZH_CHAR_LEN)
	zhchars := SubString(line, pos+ZH_CHAR_LEN, len(line))

	rs := []rune(zhchars)
	for _, zhch := range rs {
		p.AddRhyme(string(zhch), rhymestr)
	}

}
