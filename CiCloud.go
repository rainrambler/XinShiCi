package main

import (
	"fmt"
	"strings"
)

func GenerateWordCloud(filename string) {
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("Cannot load file: %s!\n", filename)
		return
	}

	var wc CiCloud
	wc.Init()
	wc.parseLines(lines)

	wc.PrintResult(500)
	//PrintMapGroupByValue(wc.char2count)
	//wc.ConvertJsonHardCode()
	//wc.SaveFile(`nalan`)
}

type CiCloud struct {
	WordCloud
	allCipais Cipais
}

func (p *CiCloud) Init() {
	p.InitParams()
	p.allCipais.Init("CiPai.txt")
}

func (p *CiCloud) parseLines(lines []string) {
	for _, line := range lines {
		p.parseOneLine(line)
	}
}

func (p *CiCloud) parseOneLine(line string) {
	linenew := strings.TrimSpace(line)
	if len(linenew) == 0 {
		return
	}
	arr := strings.FieldsFunc(linenew, SplitSentence)
	sencount := len(arr)
	if sencount == 0 {
		fmt.Printf("Err format: %s\n", line)
		return
	}

	if sencount == 1 {
		if !p.allCipais.Exists(line) {
			//fmt.Printf("subtitle: [%s]\n", line)
		}

		// Do NOT use title or subtitle
		return
	}

	for _, item := range arr {
		p.parseSentence(item)
	}
}

func (p *CiCloud) SaveFile(filename string) {
	tmpl, err := ReadTextFile(`./doc/wordcloudtempl.html`)
	if err != nil {
		fmt.Println("Cannot read template file!")
		return
	}

	for i := 4; i < 30; i++ {
		s := ConvertJsonHardCode(p.char2count, i)
		content := strings.Replace(tmpl, `[$REALDATA$]`, s, 1)
		fullfname := fmt.Sprintf("%s_1_%d.html", filename, i)
		WriteTextFile(fullfname, content)
	}

	for i := 4; i < 30; i++ {
		s := ConvertJsonHardCode(p.word2count, i)
		content := strings.Replace(tmpl, `[$REALDATA$]`, s, 1)
		fullfname := fmt.Sprintf("%s_2_%d.html", filename, i)
		WriteTextFile(fullfname, content)
	}
}
