package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type ChineseWord struct {
	Chs      string
	Cht      string
	Pinyin   string
	EngMean1 string
	EngMean2 string
}

// TODO https://stackoverflow.com/questions/19965795/go-golang-write-log-to-file
func (p *ChineseWord) toDesc() string {
	if len(p.EngMean2) == 0 {
		return fmt.Sprintf("%s %s [%s] /%s", p.Chs, p.Cht, p.Pinyin,
			p.EngMean1)
	}
	return fmt.Sprintf("%s %s [%s] /%s/%s", p.Chs, p.Cht, p.Pinyin,
		p.EngMean1, p.EngMean2)
}

func (p *ChineseWord) IsSingle2() bool {
	rs := []rune(p.Chs)
	return (len(rs) == 1) && (len(p.Chs) > 1) // filter "a", "1", etc.
}

type ChineseDict struct {
	chs2Word map[string]*ChineseWord
}

func (p *ChineseDict) Init() {
	p.chs2Word = make(map[string]*ChineseWord)
}

func (p *ChineseDict) AddChineseWord(wd *ChineseWord) {
	if wd == nil {
		return
	}
	if len(wd.Chs) == 0 {
		log.Printf("WARN: Word CHS empty: %s\n", wd.toDesc())
		return
	}
	p.chs2Word[wd.Chs] = wd
}

func (p *ChineseDict) FindChWord(chs string) *ChineseWord {
	if len(chs) == 0 {
		log.Printf("WARN: CHS empty in find\n")
		return nil
	}
	return p.chs2Word[chs]
}

// Format: cht chs [pinyin] /Eng1/Eng2/
func ParseLine(line string) *ChineseWord {

	pos := strings.Index(line, " ")

	if pos == -1 {
		log.Printf("WARN: Cannot find CHS in: [%s]\n", line)
		return nil
	}

	cw := new(ChineseWord)
	cw.Cht = line[:pos]

	part := line[pos+1:]
	pos = strings.Index(part, " ")
	if pos == -1 {
		log.Printf("WARN: Cannot find CHT in: [%s]\n", part)
		return cw
	}

	cw.Chs = part[:pos]
	part = part[pos+1:]

	posStart := strings.Index(part, "[")
	posEnd := strings.Index(part, "]")

	cw.Pinyin = part[posStart+1 : posEnd]
	part = part[posEnd+2:] // ']' and blank

	if part[0] != '/' {
		log.Printf("WARN: Err Format in Eng Desc: [%s]\n", line)
	}

	arr := SplitEngMeans(part)
	arrlen := len(arr)
	if arrlen == 1 {
		cw.EngMean1 = arr[0]
	} else if arrlen >= 2 {
		cw.EngMean1 = arr[0]
		cw.EngMean2 = arr[1]
	} else {
		log.Printf("WARN: Strange Format in Eng Desc, Count [%d]: [%s]\n",
			arrlen, line)
	}

	return cw
}

func SplitEngMeans(line string) []string {
	part := strings.Trim(line, "/")
	arr := []string{}

	if !strings.Contains(part, "/") {
		arr = append(arr, part)
		return arr
	}

	arr = strings.Split(part, "/")
	return arr
}

func ReadDict(filename string, cd *ChineseDict) {
	lines := ReadTxtFile(filename)

	for idx, line := range lines {
		if IsCommentLine(line) {
			continue
		}
		cw := ParseLine(line)

		if cw != nil {
			cd.AddChineseWord(cw)
		} else {
			log.Printf("WARN: Err reading line: [%d]\n", idx)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertHanzi2Pinyin() {
	var cd ChineseDict
	cd.Init()
	ReadDict("cedict.u8", &cd)

	var singles string
	for _, cw := range cd.chs2Word {

		if cw.IsSingle2() {
			singles += cw.Chs + "|" + cw.Pinyin + "\n"
		}
	}

	bs := []byte(singles)
	err := ioutil.WriteFile("zhs2py", bs, 0644)
	check(err)
}

func convertHanzi2Pinyin2() {
	var singles string
	lines := ReadTxtFile("cedict.u8")

	for idx, line := range lines {
		if IsCommentLine(line) {
			continue
		}
		cw := ParseLine(line)

		if cw != nil {
			if cw.IsSingle2() {
				singles += cw.Chs + "|" + cw.Pinyin + "\n"
			}
		} else {
			log.Printf("WARN: Err reading line: [%d]\n", idx)
		}
	}

	bs := []byte(singles)
	err := ioutil.WriteFile("zhs2py", bs, 0644)
	check(err)
}
