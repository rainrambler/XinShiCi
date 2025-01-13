package main

import (
	"fmt"
	"strings"
)

func AnalyseCipaiBySentence() {
	matchCipai()
}

// `寒料峭，心似玉壺冰` --> `憶江南`，`小重山`，……
func matchCipai() {
	var cp CipaiParser
	cp.Load(`Cipai/CipaiGelv.txt`)
	//cp.Match(`冰肌玉骨，自清凉无汗`)
	//cp.Match(`玉骨冰肌，自清凉无汗`)
	//cp.Match(`冉冉秋光留不住，满径红叶暮`)
	//cp.Match(`此君一见如逢旧，愿今世，长相守`)
	//cp.Match(`千金难买相如赋，脉脉此情谁诉`)
	//cp.Match(`深深浅浅绿`)
	//cp.Match(`待得團圓時候，樽前問這時節。`)
	//cp.Match(`當年忠貞為國酬`)
	//cp.Match(`且待新歲譜新篇`)
	//cp.Match(`期許新歲譜新篇`)
	//cp.Match(`夜夜姑苏城外，当时月`)
	//cp.Match(`寒料峭，心似玉壶冰`)
	//cp.Match(`寒料峭，心似玉壺冰`)
	cp.Match(`功名半紙，風雪千山`)
	//cp.DbgPrint()
}

type CipaiParser struct {
	AllCipais []*Cipai
}

func (p *CipaiParser) Load(filename string) bool {
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("WARN: Cannot parse file: %s: %v!\n", filename, err)
		return false
	}

	curCipai := new(Cipai)
	for row, line := range lines {
		purified := strings.TrimSpace(line)
		if strings.HasPrefix(purified, "#") {
			p.commitCipai(curCipai)
			curCipai = new(Cipai)
			curCipai.Title = strings.TrimSpace(purified[1:])
			continue
		}

		if strings.HasSuffix(purified, "|") {
			purified = purified[:len(purified)-1] // remove last char
			curCipai.ParseString(purified)
			p.commitCipai(curCipai)

			curTitle := curCipai.Title
			curCipai = new(Cipai)
			curCipai.Title = curTitle
			continue
		}

		if purified == "" {
			continue
		} else {
			fmt.Printf("[DBG]Possible format error: %s in row: %d!\n",
				line, row)
		}
	}
	return true
}

func (p *CipaiParser) commitCipai(aCipai *Cipai) {
	if !aCipai.IsValid() {
		return
	}

	if len(aCipai.AllSentences) == 0 {
		return
	}

	p.AllCipais = append(p.AllCipais, aCipai)
}

// `玉骨冰肌，自清凉无汗` ==> "洞仙歌"
func (p *CipaiParser) Match(content string) bool {
	found := false
	for _, item := range p.AllCipais {
		if item.Match(content) {
			found = true
		}
	}
	return found
}

func (p *CipaiParser) DbgPrint() {
	for i, item := range p.AllCipais {
		fmt.Printf("[%d]%s: %d Sentences.\n", i,
			item.Title, len(item.AllSentences))

		item.DbgPrint()
	}
}
