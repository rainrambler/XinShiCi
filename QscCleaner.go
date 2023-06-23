// QscConvertor
package main

import (
	"fmt"
	"strings"
)

type QscCleaner struct {
	curPoet   string
	curTitle  string
	isPoet    bool
	allCipais Cipais
	allpoets  Poets
}

func (p *QscCleaner) convertFile(srcFile string) {
	lines := ReadTxtFile(srcFile)
	p.convertLines(lines, srcFile+".txt")
}

func (p *QscCleaner) convertLines(lines []string, tofile string) {
	p.allCipais.Init("CiPaiZh.txt")
	p.allpoets.Init("SongPoetsZh.txt")

	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	totallines := len(lines)
	arr := []string{}
	cleaned := []string{}

	for i := 0; i < totallines; i++ {
		line := lines[i]

		if IsCommentLine(line) {
			// # 林逋
			p.curPoet = line
			p.isPoet = true
			arr = append(arr, line)
			continue
		}

		if p.isPoet && (len(line) != 0) {
			// Poet Desc?
			linenew := strings.Trim(line, " 	\r\n")
			if p.allCipais.HasActualCipai(linenew) {
				fmt.Printf("Line %d: Cipai %s\n", i, line)
				arr = append(arr, line)
			} else {
				desc := fmt.Sprintf("[%d] %s", i, line)
				cleaned = append(cleaned, desc)
			}
			p.isPoet = false
		} else {
			arr = append(arr, line)
		}
	}

	WriteLines(arr, tofile)
	WriteLines(cleaned, "cleaned.txt")
}
