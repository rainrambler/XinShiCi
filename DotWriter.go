package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/jwangsadinata/go-multimap/setmultimap"
)

type DotFile struct {
	dotHeader  []string
	dotFooter  string
	Nodes      map[string]int // ignored value
	pair2count map[string]int // "紅葉|江南:20"
	Links      []string

	Word2Word *setmultimap.MultiMap
}

func (p *DotFile) Init() {
	p.Word2Word = setmultimap.New()
	p.Nodes = make(map[string]int)
	p.pair2count = make(map[string]int)
}

func (p *DotFile) AddLink(word1 string, word2 string, cnt int) {
	p.addNode(word1)
	p.addNode(word2)

	p.Word2Word.Put(word1, word2)
	p.addLinkCount(word1, word2, cnt)
}

func (p *DotFile) addNode(oneWord string) {
	if oneWord == "" {
		return
	}

	rs := []rune(oneWord)
	if len(rs) != 2 {
		fmt.Printf("[DBG]Error word: %s\n", oneWord)
	}

	_, exists := p.Nodes[oneWord]
	if exists {
		return
	}

	p.Nodes[oneWord] = p.Nodes[oneWord] + 1
	//fmt.Printf("[DBG]Node added: %s, Total: %d\n", oneWord, len(p.Nodes))
}

func createPair(wd1, wd2 string) string {
	pair := wd1 + "|" + wd2
	return pair
}

func (p *DotFile) addLinkCount(word1 string, word2 string, cnt int) {
	pair := createPair(word1, word2)

	_, exists := p.pair2count[pair]
	if exists {
		return
	}

	p.pair2count[pair] = cnt
}

func (p *DotFile) getLinkCount(word1 string, word2 string) int {
	pair := createPair(word1, word2)
	v, exists := p.pair2count[pair]
	if !exists {
		return 0
	}

	return v
}

const TOP_WORD = 8

func (p *DotFile) Generate(filename string) {
	lines := []string{}

	header := `digraph ChWordLinks {
	fontname = "SimHei";

	node [	shape = polygon,
		fontname="SimHei,SimSun"
		sides = 4,
		distortion = "0.0",
		orientation = "0.0",
		skew = "0.0",
		color = Gray,
		style = filled];
	`

	lines = append(lines, header)

	total := 0
	for node, _ := range p.Nodes {
		line := node + " " + `[sides=6, distortion="0.936354", orientation=28, skew="-0.126818", color=salmon2];`
		lines = append(lines, line)
		total++
		if total >= TOP_WORD {
			break
		}
	}

	for _, entry := range p.Word2Word.Entries() {
		wd1 := entry.Key.(string)
		wd2 := entry.Value.(string)

		cnt := p.getLinkCount(wd1, wd2)

		if cnt == 0 {
			fmt.Printf("[INFO]Pair not found: %s and %s!\n", wd1, wd2)
		} else {
			line := fmt.Sprintf(`"%s" -> "%s" [label="%d" color="coral1"];`,
				entry.Key.(string), entry.Value.(string), cnt)
			lines = append(lines, line)
		}
	}

	footer := `}`
	lines = append(lines, footer)

	WriteLines(lines, filename)
}
func (p *DotFile) GenerateFull(filename string) {
	lines := []string{}

	header := `
	digraph ChWordLinks {
	fontname = "SimHei";

	node [	shape = polygon,
		fontname="SimHei,SimSun"
		sides = 4,
		distortion = "0.0",
		orientation = "0.0",
		skew = "0.0",
		color = Gray,
		style = filled];
	`

	lines = append(lines, header)

	for node, _ := range p.Nodes {
		line := node + " " + `[sides=6, distortion="0.936354", orientation=28, skew="-0.126818", color=salmon2];`
		lines = append(lines, line)
	}

	for _, entry := range p.Word2Word.Entries() {
		wd1 := entry.Key.(string)
		wd2 := entry.Value.(string)

		cnt := p.getLinkCount(wd1, wd2)

		if cnt == 0 {
			fmt.Printf("[INFO]Pair not found: %s and %s!\n", wd1, wd2)
		} else {
			line := fmt.Sprintf(`"%s" -> "%s" [label="%d" color="coral1"];`,
				entry.Key.(string), entry.Value.(string), cnt)
			lines = append(lines, line)
		}
	}

	footer := `}`
	lines = append(lines, footer)

	WriteLines(lines, filename)
}

func CreatePngFromDot(dotname string) {
	pngname := ChangeFileExt(dotname, ".png")
	//fmt.Printf("[DBG]PNG: %s\n", pngname)

	dotpath := `D:\Soft\graphviz\Graphviz-12.1.1-win64\bin\dot.exe`
	cmd := exec.Command(dotpath, dotname, "-Tpng", "-o", pngname)

	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("[INFO]PNG %s created. Result: %s\n", pngname, out)
}
