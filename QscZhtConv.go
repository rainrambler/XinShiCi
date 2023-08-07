// QscConvertor
package main

import (
	"fmt"
	"log"
	"strings"
)

type QscZhtConv struct {
	curPoet   string
	curTitle  string
	allCipais Cipais
	allpoets  Poets
	allPoems  ChinesePoems
	runRhyme  bool
	prevPoet  bool
	preTitle  bool

	curContent   string
	curComment   string
	curLineNum   int
	titleLineNum int

	allLines []string
}

func (p *QscZhtConv) loadFile(srcFile string) {
	p.allCipais.Init("CiPaiZh.txt")
	p.allpoets.Init("SongPoetsZh.txt")
	p.allPoems.Init()
	fmt.Printf("INFO: Total poets: %d\n", p.allpoets.Count())

	lines := ReadTxtFile(srcFile)
	p.runRhyme = true
	p.parseLines(lines)

	fmt.Printf("INFO: %d poems loaded.\n", p.allPoems.Count())

	WriteLines(p.allLines, srcFile+".txt")
}

func (p *QscZhtConv) parseLines(lines []string) {
	totallines := len(lines)
	for i := 0; i < totallines; i++ {
		line := lines[i]
		//fmt.Printf("[DBG][%d]: %s\n", i+1, line)

		if IsEmptyLine(line) {
			//p.curContent += "\r\n"
			p.appendLine(line)
			continue
		}

		firstchar := GetFirstRune(line)
		switch firstchar {
		case '#':
			{
				p.CommitPoem(i - 1)
				p.beginNewPoet(line)

				p.appendLine(line)
			}
		case '!':
			{
				p.curComment += line
				p.appendLine(line)
			}
		case '$':
			{
				// sub-title
				p.appendLine(line)
			}
		case '*':
			{
				// author desc
				p.appendLine(line)
			}
		case '【':
			{
				// title
				p.CommitPoem(i - 1)
				p.beginNewPoem(line)
				p.appendLine(line)
			}
		default:
			p.addLine(i, line)
		}
	}

	p.CommitPoem(totallines)
}

func (p *QscZhtConv) addLine(pos int, line string) {
	linenew := TrimBlank(line)
	if len(linenew) == 0 {
		return
	}
	//p.curContent += linenew + "\r\n" // TODO ??

	lastchar := GetLastRune(linenew)
	if !IsPunctuation(lastchar) {
		fmt.Printf("[%d]Possible sub-title: %s\n", pos, line)
		p.appendLine("$ " + linenew)
		linenew += "。"
	} else {
		p.curContent += linenew
		p.appendLine(line)
	}
}

func (p *QscZhtConv) appendLine(line string) {
	p.allLines = append(p.allLines, line)
}

func (p *QscZhtConv) beginNewPoet(line string) {
	s := strings.Trim(line, " \t#")
	p.curPoet = s
}

func (p *QscZhtConv) beginNewPoem(line string) {
	p.curTitle = TrimTitle(line)
}

func (p *QscZhtConv) CommitPoem(pos int) {
	if p.curPoet == "" {
		//fmt.Printf("DBG: Cannot find author in line: %d\n", pos)
		return
	}
	if p.curTitle == "" {
		//fmt.Printf("DBG: No title in line: %d, Poet: %s, [%s]%s\n", pos,
		//	p.curPoet, p.curTitle, p.curContent)
		return
	}
	if p.curContent == "" {
		log.Printf("DBG: Cannot find content in line: %d\n", pos)
		return
	}
	poetId := p.allpoets.FindPoet(p.curPoet)
	if poetId < 0 {
		fmt.Printf("DBG: [%d]Cannot find poet: %s\n", pos, p.curPoet)
		return
	}

	poemId := fmt.Sprintf("%d-%d", poetId, pos)

	cp := CreateQscPoem(poemId, p.curPoet, p.curTitle, p.curContent, p.curComment)

	if p.runRhyme {
		cp.analyseRhyme()
	}
	p.allPoems.AddPoem(cp)

	p.ClearCurrent()
}

func (p *QscZhtConv) setNewTitle(pos int, line string) {
	p.curTitle = line
	p.titleLineNum = pos
	p.preTitle = true
}

func (p *QscZhtConv) ClearCurrent() {
	p.curContent = ""
	p.curTitle = ""
	p.curComment = ""
	p.curLineNum = 0
	p.titleLineNum = 0
}

func TrimTitle(s string) string {
	return strings.Trim(s, " \t【】")
}

// '又'
func ConvertRepeatCipai(filename string) {
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("[WARN]Cannot open file: %s!\n", filename)
		return
	}

	linesnew := []string{}

	curTitle := ""
	for num, line := range lines {
		firstchar := GetFirstRune(line)
		switch firstchar {
		case '【':
			{
				// title
				curTitle = TrimTitle(line)
				linesnew = append(linesnew, line)
			}
		case '又':
			{
				if ChcharLen(line) == 1 {
					// repeated title
					title := packCipai(curTitle)
					linesnew = append(linesnew, title)
				} else {
					fmt.Printf("[%d]%s\n", num, line)
					linesnew = append(linesnew, line)
				}
			}
		default:
			linesnew = append(linesnew, line)
		}
	}

	WriteLines(linesnew, filename+".txt")
}

func ListCipai(filename string) {
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("[WARN]Cannot open file: %s!\n", filename)
		return
	}

	for _, line := range lines {
		firstchar := GetFirstRune(line)
		switch firstchar {
		case '【':
			{
				// title
				fmt.Println(TrimTitle(line))
			}

		default:
		}
	}
}

type CipaiPurifier struct {
	fromFile  string
	allCipais Cipais

	allLines []string
}

// Possible:
// 其一百三十四
// 其五_12
// 鳳歸雲〈閨怨〉
// 失調名（般涉）
// 失調名_2
// 又_15
// 馬藥方·第三十
func (p *CipaiPurifier) Purify(filename string) {
	p.allCipais.Init("CiPaiZh.txt")
	p.allLines = []string{}

	p.fromFile = filename

	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("[WARN]Cannot open file: %s!\n", filename)
		return
	}

	for num, line := range lines {
		firstchar := GetFirstRune(line)
		switch firstchar {
		case '【':
			{
				// title
				title := TrimTitle(line)
				if !p.parseTitle(title) {
					fmt.Printf("[DBG]Cannot parse [%d] %s!\n", num, line)
				}
			}

		default:
			p.allLines = append(p.allLines, line)
		}
	}

	WriteLines(p.allLines, filename+".txt")
}

func (p *CipaiPurifier) parseTitle(line string) bool {
	//rs := []rune(line)
	cipai, _ := SplitZhStringMulti(line, "〈（·")
	if cipai != line {
		if !p.allCipais.HasActualCipai(cipai) {
			fmt.Printf("Possible cipai: %s\n", cipai)
			return false
		}

		p.allLines = append(p.allLines, packCipai(cipai))
		return true
	}

	title, _ := SplitZhString(line, '_')
	if title != line {
		//fmt.Printf("[DBG]%s to %s\n", line, title)
		if p.allCipais.HasActualCipai(title) {
			p.allLines = append(p.allLines, packCipai(title))
			return true
		}
		if title == "又" {
			p.allLines = append(p.allLines, packCipai(title))
			return true
		}

		fmt.Printf("Cannot find title: %s\n", title)
		p.allLines = append(p.allLines, packCipai(title))
		return false
	}

	if StartWith(title, '其') {
		if !OnlyContains(title, `其一二三四五六七八九十百`) {
			fmt.Printf("Possible content: %s\n", line)
			p.allLines = append(p.allLines, line)
			return false
		} else {
			p.allLines = append(p.allLines, "【又】")
			return true
		}
	}

	p.allLines = append(p.allLines, packCipai(line))
	return true
}

func packCipai(title string) string {
	return "【" + title + "】"
}
