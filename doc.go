// Xindict project doc.go

/*
Xindict document
*/
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

// TODO len(chinese word) == 3??
func (p *ChineseWord) IsSingle() bool {
	wdlen := len(p.Chs)
	//fmt.Printf("Len: %d\n", wdlen)
	return wdlen == 3
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

func ReadTxtFile(filename string) []string {
	linearr := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linearr = append(linearr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return linearr
}

func IsCommentLine(line string) bool {
	if strings.HasPrefix(line, "#") {
		return true
	}

	if strings.HasPrefix(line, "%") {
		return true
	}

	return false
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

func SplitEngMeans2(line string) []string {
	arr := []string{}
	if line[0] != '/' {
		log.Printf("WARN: Err Format in Eng Desc: [%s]\n", line)
	}

	part := line[1:] // remove the first / char
	pos := strings.Index(part, "/")
	if pos == -1 {
		arr = append(arr, part)
		return arr
	} else if pos == len(part) {
		tmp := part[:len(part)-1]
		arr = append(arr, tmp)
		return arr
	}

	leftpart := part[:pos]
	arr = append(arr, leftpart)
	part = part[pos+1:]

	if len(part) == 0 {
		return arr
	}

	return arr
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

	/*
		arrnoempty := []string{}

		for _, item := range arr {
			itemNew := strings.TrimSpace(item)

			if itemNew != "" {
				arrnoempty = append(arrnoempty, itemNew)
			}
		}

		return arrnoempty
	*/
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
	ReadDict("/Users/anyu/goproj/Xindict/cedict.u8", &cd)

	//singles := []string{}
	var singles string
	for _, cw := range cd.chs2Word {

		if cw.IsSingle() {
			//fmt.Println(cw.toDesc())
			//singlecount++
			singles += cw.Chs + "|" + cw.Pinyin + "\n"
		}
	}

	bs := []byte(singles)
	err := ioutil.WriteFile("hanzi2pinyin", bs, 0644)
	check(err)
	//fmt.Printf("Hanzi Count: %d\n", singlecount)
}

func SubString(s string, beginPos, size int) string {
	slen := len(s)

	if beginPos >= slen {
		return ""
	}

	if beginPos+size >= slen {
		return s[beginPos:]
	}

	return s[beginPos : beginPos+size]
}

func SubChineseString(s string, beginPos, size int) string {
	rs := []rune(s)
	slen := len(rs)

	if beginPos >= slen {
		return ""
	}

	if beginPos+size >= slen {
		return string(rs[beginPos:])
	}

	return string(rs[beginPos : beginPos+size])
}

func ChcharLen(s string) int {
	rs := []rune(s)
	return len(rs)
}
