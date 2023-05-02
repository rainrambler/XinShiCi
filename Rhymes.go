package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

var g_Rhymes ChineseRhymes

type Rhyme struct {
	Desc   string
	rhymes []string
}

func (p *Rhyme) AddItem(rhy string) {
	if len(rhy) == 0 {
		return
	}

	p.Desc = p.Desc + rhy + "|"
	p.rhymes = append(p.rhymes, rhy)
}

func (p *Rhyme) toDesc() string {
	return p.Desc
}

type ChineseRhymes struct {
	ZhChar2Rhyme map[string]*Rhyme // "安" -> "【十四寒】"
}

func (p *ChineseRhymes) Init() {
	p.ZhChar2Rhyme = make(map[string]*Rhyme)
}

func (p *ChineseRhymes) AddRhyme(zhch, rhyme string) {
	// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	if curRhyme, ok := p.ZhChar2Rhyme[zhch]; ok {
		// exists
		curRhyme.AddItem(rhyme)
		return
	}

	curRhyme := new(Rhyme)
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
}

func (p *ChineseRhymes) parseLine(rownum int, line string) {
	if !strings.HasPrefix(line, "【") {
		log.Printf("WARN: Invalid line in Rhyme file (No Start): %d\n", rownum)
		return
	}

	pos := strings.Index(line, "】")

	if pos == -1 {
		log.Printf("WARN: Invalid line in Rhyme file (No End): %d\n", rownum)
		return
	}

	rhymestr := SubString(line, ZH_CHAR_LEN, pos-ZH_CHAR_LEN)
	zhchars := SubString(line, pos+ZH_CHAR_LEN, len(line))

	rs := []rune(zhchars)
	for _, zhch := range rs {
		p.AddRhyme(string(zhch), rhymestr)
	}

}

var missedChars Rhyme2Count

func (p *ChineseRhymes) AnalyseRhyme(lastwords []string) string {
	var rhy2count Rhyme2Count
	rhy2count.Init()

	for _, wd := range lastwords {
		if curRhyme, ok := p.ZhChar2Rhyme[wd]; ok {
			// exists
			//fmt.Printf("[%s] Rhyme: [%s]\n", wd, curRhyme.Desc)
			for _, rhy := range curRhyme.rhymes {
				rhy2count.Add(rhy)
			}
		} else {
			//fmt.Printf("WARN: Cannot find rhyme for: %s\n", wd)
			missedChars.Add(wd)
		}
	}

	return rhy2count.FindTop1()
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

	fmt.Println("-->Start")
	for _, kv := range kvs {
		if kv.Value > 1 {
			fmt.Printf("[%s] count: %d\n", kv.Key, kv.Value)
		}
	}
	fmt.Println("<--End")
}
