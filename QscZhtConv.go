// QscConvertor
package main

import (
	"strings"
)

type QscZhtConv struct {
	curTitle   string
	curLineNum int
	allLines   []string
}

func (p *QscZhtConv) addLine(line string) {
	p.allLines = append(p.allLines, line)
}

func (p *QscZhtConv) convertFile(srcFile string) {
	lines := ReadTxtFile(srcFile)
	p.parseLines(lines)
	WriteLines(p.allLines, srcFile+".txt")
	//p.allPoems.PrintResults()
}

func (p *QscZhtConv) parseLines(lines []string) {
	totallines := len(lines)
	for i := 0; i < totallines; i++ {
		line := lines[i]
		//fmt.Printf("[DBG][%d]: %s\n", i+1, line)

		if IsEmptyLine(line) {
			p.addLine(line)
			continue
		}

		firstchar := GetFirstRune(line)
		switch firstchar {
		case '【':
			{
				// title
				title := strings.Trim(line, " \t【】")
				if isSequenceCipai(title) {
					p.addLine("【" + p.curTitle + "】")
				} else {
					p.addLine(line)
					p.curTitle = title
				}
			}
		default:
			p.addLine(line)
		}

	}
}
