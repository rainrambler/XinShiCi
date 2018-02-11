package main

import (
	"log"
)

type Qsc struct {
	ID2Poems map[string]*ChinesePoem
}

func (p *Qsc) Init() {
	p.ID2Poems = make(map[string]*ChinesePoem)
}

func (p *Qsc) AddPoem(poem *ChinesePoem) {
	if poem == nil {
		log.Printf("WARN: AddPoem: nil\n")
		return
	}

	if len(poem.ID) == 0 {
		log.Printf("WARN: AddPoem: empty ID!!\n")
		return
	}
	// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	if _, ok := p.ID2Poems[poem.ID]; ok {
		// exists
		log.Printf("WARN: AddPoem: exists ID: %s, Line: %d\n", poem.ID, poem.LineNumber)
		return
	}
	p.ID2Poems[poem.ID] = poem
}

func (p *Qsc) ReadFile(filename string) {

}
