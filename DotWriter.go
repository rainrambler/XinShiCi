package main

import (
	"fmt"

	"github.com/jwangsadinata/go-multimap/setmultimap"
)

type DotFile struct {
	dotHeader []string
	dotFooter string
	Nodes     map[string]int // ignored value
	Links     []string

	Word2Word *setmultimap.MultiMap
}

func (p *DotFile) Init() {
	p.Word2Word = setmultimap.New()
	p.Nodes = make(map[string]int)
}

func (p *DotFile) AddNode(oneWord string) {
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

	p.Nodes[oneWord] = 1
}

func (p *DotFile) AddLink(word1 string, word2 string) {
	p.AddNode(word1)
	p.AddNode(word2)

	p.Word2Word.Put(word1, word2)
}

const TOP_WORD = 20

func (p *DotFile) Generate(filename string) {
	lines := []string{}

	header := `
	digraph "ChWordLinks" {
	graph [	fontname = "SimHei",
		fontsize = 36,
		label = "\n\n\n\nObject Oriented Graphs\nStephen North, 3/19/93",
		];
	node [	shape = polygon,
		sides = 4,
		distortion = "0.0",
		orientation = "0.0",
		skew = "0.0",
		color = Gray,
		style = filled,
		fontname = "Arial" ];
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
		line := fmt.Sprintf(`"%s" -> "%s";`, entry.Key.(string), entry.Value.(string))
		lines = append(lines, line)
	}

	footer := `}`
	lines = append(lines, footer)

	WriteLines(lines, filename)
}
