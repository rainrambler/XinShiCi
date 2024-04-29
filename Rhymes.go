package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

type Rhyme struct {
	Desc   string
	rhymes map[string]int // key: rhymes
}

func CreateRhyme() *Rhyme {
	var rhy Rhyme
	rhy.rhymes = make(map[string]int)
	return &rhy
}

func (p *Rhyme) AddItem(rhy string) bool {
	if len(rhy) == 0 {
		return false
	}

	purified := strings.TrimSpace(rhy)
	if len(purified) == 0 {
		return false
	}

	p.Desc = p.Desc + rhy + "|"

	if _, ok := p.rhymes[rhy]; ok {

	} else {
		p.rhymes[rhy] = 1
	}

	return true
}

func (p *Rhyme) toDesc() string {
	return p.Desc
}

func (p *Rhyme) match(rhy string) bool {
	_, ok := p.rhymes[rhy]
	return ok
}

func (p *Rhyme) getAll() []string {
	var arr []string
	for k, _ := range p.rhymes {
		arr = append(arr, k)
	}
	return arr
}

var g_Rhymes ChineseRhymes

// 平声（诗韵新编）
const (
	P1Ma    = "P1"
	P2Bo    = "P2"
	P3Ge    = "P3"
	P4Jie   = "P4"
	P5Zhi   = "P5"
	P6Er    = "P6"
	P7Qi    = "P7"
	P8Wei   = "P8"
	P9Kai   = "P9"
	P10Gu   = "P10"
	P11Yu   = "P11"
	P12Hou  = "P12"
	P13Hao  = "P13"
	P14Han  = "P14"
	P15Hen  = "P15"
	P16Tang = "P16"
	P17Geng = "P17"
	P18Dong = "P18"
)

// 仄声（诗韵新编）
const (
	Z1Ba3     = "Z1"
	Z2Bo3     = "Z2"
	Z3Che3    = "Z3"
	Z4Jie3    = "Z4"
	Z5Chi3    = "Z5"
	Z6Er3     = "Z6"
	Z7Bi3     = "Z7"
	Z8Bei3    = "Z8"
	Z9Ai3     = "Z9"
	Z10Bu4    = "Z10"
	Z11Ju3    = "Z11"
	Z12Chou3  = "Z12"
	Z13Ao3    = "Z13"
	Z14An3    = "Z14"
	Z15Ben3   = "Z15"
	Z16Bang3  = "Z16"
	Z17Beng3  = "Z17"
	Z18Chong3 = "Z18"
)

// 入声（诗韵新编）
const (
	R1Ba  = "R1"
	R2Bo  = "R2"
	R3Ge  = "R3"
	R4Bie = "R4"
	R5Chi = "R5"
	R6Bi  = "R6"
	R7Chu = "R7"
	R8Qu  = "R8"
)

type ChineseRhymes struct {
	ZhChar2Rhyme map[rune]*Rhyme // "安" -> "【十四寒】"
}

func (p *ChineseRhymes) Init() {
	p.ZhChar2Rhyme = make(map[rune]*Rhyme)
	missedChars.Init()
}

func (p *ChineseRhymes) AddRhyme(zhch rune, rhyme string) {
	// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	if curRhyme, ok := p.ZhChar2Rhyme[zhch]; ok {
		// exists
		curRhyme.AddItem(rhyme)
		return
	}

	curRhyme := CreateRhyme()
	curRhyme.AddItem(rhyme)
	p.ZhChar2Rhyme[zhch] = curRhyme
}

// https://baike.baidu.com/item/%E8%AF%97%E9%9F%B5%E6%96%B0%E7%BC%96
func (p *ChineseRhymes) ImportFile(filename string) {
	p.Init()

	lines := ReadTxtFile(filename)

	for idx, line := range lines {
		p.parseLine(idx+1, line)
	}

	//p.checkMultivalue()
}

func splitShiyun(row int, line string) (string, string) {
	if len(line) == 0 {
		return "", ""
	}
	rs := []rune(line)
	firstchar := rs[0]
	if firstchar != '【' {
		log.Printf("WARN: Invalid line in Rhyme file (No Start): %d\n", row)
		return line, ""
	}

	pos := 1
	for rs[pos] != '】' {
		pos++

		if pos == len(line) {
			log.Printf("WARN: Invalid line in Rhyme file (No End): %d\n", row)
			return "", ""
		}
	}

	rhymestr := string(rs[1:pos])
	remain := string(rs[pos+1:])
	return rhymestr, remain
}

func (p *ChineseRhymes) parseLine(rownum int, line string) {
	rhymestr, remain := splitShiyun(rownum, line)
	if remain == "" {
		return
	}

	rs := []rune(remain)
	for _, r := range rs {
		p.AddRhyme(r, rhymestr)
		//fmt.Printf("[DBG]%s to {%s}\n", string(r), rhymestr)
	}
}

func (p *ChineseRhymes) checkMultivalue() {
	for k, v := range p.ZhChar2Rhyme {
		if len(v.rhymes) > 1 {
			fmt.Printf("%s:%s\n", string(k), v.Desc)
		}
	}
}

var missedChars Rhyme2Count

// 晴。清。明。蕖，盈。鷺，意，婷。箏。情。聽。收，靈。取，見，青。
// ==>
// 【十七庚】
func (p *ChineseRhymes) AnalyseRhyme(lastwords []rune) string {
	var rhy2count Rhyme2Count
	rhy2count.Init()

	for _, wd := range lastwords {
		if curRhyme, ok := p.ZhChar2Rhyme[wd]; ok {
			// exists
			for rhy, _ := range curRhyme.rhymes {
				rhy2count.Add(rhy)
			}
		} else {
			missedChars.Add(string(wd))
		}
	}

	return rhy2count.FindTop1()
}

func (p *ChineseRhymes) FindRhyme(chword rune) *Rhyme {
	if curRhyme, ok := p.ZhChar2Rhyme[chword]; ok {
		return curRhyme
	}
	return nil
}

type Rhyme2Count struct {
	rhy2Count map[string]int
}

func (p *Rhyme2Count) Init() {
	p.rhy2Count = make(map[string]int)
}

func (p *Rhyme2Count) Add(rhy string) {
	if len(rhy) == 0 {
		return
	}

	if _, ok := p.rhy2Count[rhy]; ok {
		// exists
		p.rhy2Count[rhy] += 1
	} else {
		p.rhy2Count[rhy] = 1
	}
}

type KeyValue struct {
	Key   string
	Value int
}

func (p *Rhyme2Count) SortByValue() []KeyValue {
	arrlen := len(p.rhy2Count)

	if arrlen == 0 {
		return []KeyValue{}
	}

	if arrlen == 1 {
		for k, v := range p.rhy2Count {
			return []KeyValue{KeyValue{k, v}}
		}
	}

	var kvarr []KeyValue
	for k, v := range p.rhy2Count {
		kvarr = append(kvarr, KeyValue{k, v})
	}

	sort.Slice(kvarr, func(i, j int) bool {
		return kvarr[i].Value > kvarr[j].Value
	})

	return kvarr
}

func (p *Rhyme2Count) FindTop1() string {
	kvs := p.SortByValue()

	if len(kvs) == 0 {
		return ""
	}

	return kvs[0].Key
}

func (p *Rhyme2Count) PrintSorted() {
	kvs := p.SortByValue()

	total := 0
	fmt.Println("-->Start")
	for _, kv := range kvs {
		if kv.Value > 1 {
			fmt.Printf("[%s]: %d, ", kv.Key, kv.Value)
			total++

			if (total % 10) == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println("<--End")
}
