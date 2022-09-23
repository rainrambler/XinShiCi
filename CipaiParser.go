package main

import (
	"fmt"
	"strings"
)

func cipaiDemo() {
	var cp CipaiParser
	cp.Load(`Cipai/CipaiGelv.txt`)
	//cp.Match(`冰肌玉骨，自清凉无汗`)
	cp.Match(`冰肌玉骨，自清凉无汗`)
	//cp.DbgPrint()
}

type CipaiParser struct {
	AllCipais []*Cipai
}

func (p *CipaiParser) Load(filename string) bool {
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("WARN: Cannot parse file: %s: %v!\n", filename, err)
		return false
	}

	curCipai := new(Cipai)
	for row, line := range lines {
		purified := strings.TrimSpace(line)
		if strings.HasPrefix(purified, "#") {
			p.commitCipai(curCipai)
			curCipai = new(Cipai)
			curCipai.Title = strings.TrimSpace(purified[1:])
			continue
		}

		if strings.HasSuffix(purified, "|") {
			purified = purified[:len(purified)-1] // remove last char
			curCipai.ParseString(purified)
			p.commitCipai(curCipai)

			curTitle := curCipai.Title
			curCipai = new(Cipai)
			curCipai.Title = curTitle
			continue
		}

		if purified == "" {
			continue
		} else {
			fmt.Printf("[DBG]Possible format error: %s in row: %d!\n",
				line, row)
		}
	}
	return true
}

func (p *CipaiParser) commitCipai(aCipai *Cipai) {
	if !aCipai.IsValid() {
		return
	}

	if len(aCipai.AllSentences) == 0 {
		return
	}

	p.AllCipais = append(p.AllCipais, aCipai)
}

func (p *CipaiParser) Match(content string) bool {
	found := false
	for _, item := range p.AllCipais {
		if item.Match(content) {
			found = true
		}
	}
	return found
}

func (p *CipaiParser) DbgPrint() {
	for i, item := range p.AllCipais {
		fmt.Printf("[%d]%s: %d Sentences.\n", i,
			item.Title, len(item.AllSentences))

		item.DbgPrint()
	}
}
