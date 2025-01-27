package main

import (
	"fmt"
	"log"
	"strings"
)

type Cipai struct {
	Title        string
	AllSentences []*Sentence
}

func (p *Cipai) IsValid() bool {
	return p.Title != ""
}

func (p *Cipai) ParseString(content string) bool {
	arr := strings.FieldsFunc(content, SplitF)
	if len(arr) == 0 {
		log.Printf("WARN: Possible error format: %s!\n",
			content)
		return false
	}

	for _, item := range arr {
		oneSen := new(Sentence)
		oneSen.Parse(item)

		p.AllSentences = append(p.AllSentences, oneSen)
	}

	return true
}

func (p *Cipai) DbgPrint() {
	fmt.Printf("-->Sentenses for %s: \n", p.Title)
	for i, item := range p.AllSentences {
		fmt.Printf("[%d]%+v\n", i,
			item.pingzeArr)
	}
	fmt.Printf("<--Sentenses for %s: \n", p.Title)
}

func (p *Cipai) Match(content string) bool {
	arr := strings.FieldsFunc(content, SplitF)
	if len(arr) == 0 {
		log.Printf("WARN: Possible error content: %s!\n",
			content)
		return false
	}

	return p.compareArr(arr)
}

func (p *Cipai) compareArr(arr []string) bool {
	endPos := len(p.AllSentences) - len(arr)
	if endPos < 0 {
		return false
	}

	for startPos := 0; startPos < endPos; startPos++ {
		partSentences := p.AllSentences[startPos : startPos+len(arr)]

		if compareArrSameSize(arr, partSentences) {
			fmt.Printf("Found in %s. Line: %d. Total (%d) Chars\n",
				p.Title, startPos+1, p.CharSize())
			fmt.Println("----------------------------")
			return true
		}
	}

	return false
}

func compareArrSameSize(arr []string, sens []*Sentence) bool {
	for i := 0; i < len(arr); i++ {
		if ChcharLen(arr[i]) != sens[i].Length() {
			return false
		}

		if !sens[i].Match(arr[i]) {
			return false
		}
	}

	// Output result
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%s (%s)\n", arr[i], sens[i].ToDesc())
	}

	return true
}

func (p *Cipai) CharSize() int {
	totallen := 0
	for _, v := range p.AllSentences {
		totallen += v.Length()
	}
	return totallen
}

// See IsPunctuation()
func SplitF(r rune) bool {
	return r == '，' || r == '。' || r == '？' || r == '！' || r == '、'
}

type Sentence struct {
	pingzeArr []int
}

func (p *Sentence) Parse(s string) {
	p.pingzeArr = []int{}
	rs := []rune(s)
	for i, r := range rs {
		switch r {
		case '平':
			p.pingzeArr = append(p.pingzeArr, PingZePing)
		case '仄':
			p.pingzeArr = append(p.pingzeArr, PingZeZe)
		case '中':
			p.pingzeArr = append(p.pingzeArr, PingZeAny)
		default:
			log.Printf("[DBG]Unknown char in pos %d of %s!\n", i, s)
		}
	}
}

func (p *Sentence) Length() int {
	return len(p.pingzeArr)
}

func (p *Sentence) Match(s string) bool {
	rs := []rune(s)
	if len(rs) != len(p.pingzeArr) {
		fmt.Printf("[DBG]Diff size in Sentence match: %d vs %d\n", len(rs), len(p.pingzeArr))
		return false
	}

	arr := []int{}
	for _, r := range rs {
		pzval := pinyinInst.FindPingze2(r)
		arr = append(arr, pzval)
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == PingZeUnknown {
			fmt.Printf("[DBG]Unknown in pos %d of %v\n", i, arr)
			return false
		}

		if isFreePingze(arr[i]) {
			// either is ok
			continue
		}

		if isFreePingze(p.pingzeArr[i]) {
			// either is ok
			continue
		}

		if p.pingzeArr[i] != arr[i] {
			//fmt.Printf("[DBG]Not match: %v vs %v at %d\n",
			//	p.pingzeArr[i], arr[i], i)
			return false
		}
	}

	return true
}

func (p *Sentence) ToDesc() string {
	s := ""
	for _, item := range p.pingzeArr {
		switch item {
		case PingZePing:
			s += "平"
		case PingZeZe:
			s += "仄"
		case PingZeAny:
			s += "中"
		default:
			s += "??"
		}
	}
	return s
}

var pinyinInst PinyinFinder

func init() {
	pinyinInst.Init(`zht2py.txt`)
}
