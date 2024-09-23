package main

import (
	"fmt"
	"sort"
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

type WordCloud struct {
	char2count map[string]int
	word2count map[string]int
}

func (p *WordCloud) InitParams() {
	p.char2count = make(map[string]int)
	p.word2count = make(map[string]int)
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

func (p *WordCloud) parseMultiLine(line string) {
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
		p.parseSentence(line)
		return
	}

	for _, item := range arr {
		p.parseSentence(item)
	}
}

func (p *WordCloud) parseSentence(line string) {
	rs := []rune(line)
	for _, r := range rs {
		p.AddChar(r)
	}

	rcount := len(rs)
	for i := 0; i < rcount-1; i++ {
		pair := rs[i : i+2]
		p.AddWord(string(pair))
	}
}

func (p *WordCloud) AddChar(r rune) {
	s := string(r)
	p.char2count[s] = p.char2count[s] + 1
}

func (p *WordCloud) AddWord(s string) {
	p.word2count[s] = p.word2count[s] + 1
}

// Print Top N values (sorted by value), -1 means all
func (p *WordCloud) PrintResult(topn int) {
	if len(p.word2count) == 0 {
		fmt.Println("No result!")
		return
	}
	PrintMapByValueTop(p.word2count, topn)
}

func (p *WordCloud) GetTopResult(topn int) map[string]int {
	if len(p.word2count) == 0 {
		fmt.Println("No result!")
		return map[string]int{}
	}

	k2v := make(map[string]int)

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range p.word2count {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	total := 0
	for _, kv := range ss {
		k2v[kv.Key] = kv.Value
		//df.AddLink(curWord, kv.Key, kv.Value)
		//fmt.Printf("[DBG][%d]%s -> %s (%d)\n", total, curWord, kv.Key, kv.Value)

		total++
		if total >= topn {
			break
		}
	}

	return k2v
}

func (p *WordCloud) CreateDot(curWord, filename string) {
	var df DotFile
	df.Init()

	k2v := p.GetTopResult(TOP_WORD)

	for k, v := range k2v {
		df.AddLink(curWord, k, v)
		//fmt.Printf("[DBG][%d]%s -> %s (%d)\n", total, curWord, kv.Key, kv.Value)
	}

	df.Generate(filename)
}

const Multiply = 30

func ConvertJsonHardCode(s2c map[string]int, margin int) string {
	s := ""
	for k, v := range s2c {
		if v > margin {
			line := fmt.Sprintf(`{name:"%s",value:%d},`, k, v)
			s += line
			//filtered = append(filtered, Char2Count{Title: k, Number: v * 50})
		}
	}

	s = s[:len(s)-1] // remove last comma
	s = "[" + s + "]"
	return s
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

func SplitSentence(r rune) bool {
	return IsPunctuationAll(r)
}
