package main

import (
	"fmt"
	"strings"
)

func LoadTxt(filename string) {
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("Cannot load file: %s!\n", filename)
		return
	}

	var wc WordCloud
	wc.Init()
	wc.parseLines(lines)

	//wc.PrintResult()
	PrintMapGroupByValue(wc.char2count)
	//wc.ConvertJsonHardCode()
	wc.SaveFile(`nalan`)
}

type WordCloud struct {
	char2count map[string]int
	word2count map[string]int
	allCipais  Cipais
}

func (p *WordCloud) Init() {
	p.char2count = make(map[string]int)
	p.word2count = make(map[string]int)

	p.allCipais.Init("CiPai.txt")
}

func (p *WordCloud) parseLines(lines []string) {
	for _, line := range lines {
		p.parseOneLine(line)
	}
}

func (p *WordCloud) parseOneLine(line string) {
	arr := strings.FieldsFunc(line, SplitSentence)
	sencount := len(arr)
	if sencount == 0 {
		fmt.Printf("Err format: %s\n", line)
		return
	}

	if sencount == 1 {
		if !p.allCipais.Exists(line) {
			fmt.Printf("subtitle: [%s]\n", line)
		}

		// Do NOT use title or subtitle
		return
	}

	for _, item := range arr {
		p.parseSentence(item)
	}
}

func (p *WordCloud) parseSentence(line string) {
	rs := []rune(line)
	for _, r := range rs {
		p.AddChar(r)
	}

	rcount := len(rs)
	for i := 0; i < rcount-1; i++ {
		pair := rs[i : i+2]
		p.AddWord(string(pair))
	}
}

func (p *WordCloud) AddChar(r rune) {
	s := string(r)
	p.char2count[s] = p.char2count[s] + 1
}

func (p *WordCloud) AddWord(s string) {
	p.word2count[s] = p.word2count[s] + 1
}

func (p *WordCloud) PrintResult() {
	PrintSortedMapByValue(p.char2count)
	PrintSortedMapByValue(p.word2count)
}

type Char2Count struct {
	Title  string `json:"name"`
	Number int    `json:"value"`
}

func (p *WordCloud) ConvertJsonHardCode(margin int) string {
	s := ""
	for k, v := range p.word2count {
		if v > margin {
			line := fmt.Sprintf(`{name:"%s",value:%d},`, k, v*30)
			s += line
			//filtered = append(filtered, Char2Count{Title: k, Number: v * 50})
		}
	}

	s = s[:len(s)-1] // remove last comma
	s = "[" + s + "]"
	return s
}

func (p *WordCloud) SaveFile(filename string) {
	tmpl, err := ReadTextFile(`./doc/wordcloudtempl.html`)
	if err != nil {
		fmt.Println("Cannot read template file!")
		return
	}

	for i := 4; i < 30; i++ {
		s := p.ConvertJsonHardCode(i)
		content := strings.Replace(tmpl, `[$REALDATA$]`, s, 1)
		fullfname := fmt.Sprintf("%s%d.html", filename, i)
		WriteTextFile(fullfname, content)
	}
}

func SplitSentence(r rune) bool {
	return IsPunctuationAll(r)
}
