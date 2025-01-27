package main

import (
	"fmt"
	"strings"
)

func GenerateWordCloud(filename, outfilepart string) {
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("Cannot load file: %s!\n", filename)
		return
	}

	generateWordCloudByLines(lines, outfilepart)
}

// outfile: "dufu" --> "dufu_2_30.html"
func generateWordCloudByLines(lines []string, outfile string) {
	var wc CiCloud
	wc.Init()
	wc.parseLines(lines)

	wc.PrintResult(500)
	if outfile == "" {
		return
	}

	wc.SaveMultiFiles(outfile)
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
