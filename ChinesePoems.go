package main

import (
	"log"
)

type ChinesePoems struct {
	ID2Poems map[string]*ChinesePoem
}

func (p *ChinesePoems) Init() {
	p.ID2Poems = make(map[string]*ChinesePoem)
}

func (p *ChinesePoems) AddPoem(poem *ChinesePoem) {
	if poem == nil {
		log.Printf("WARN: AddPoem: nil\n")
		return
	}

	if len(poem.ID) == 0 {
		log.Printf("WARN: AddPoem: empty ID!!\n")
		return
	}
	// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	if res, ok := p.ID2Poems[poem.ID]; ok {
		// exists
		log.Printf("WARN: AddPoem: exists ID: %s, Line: %d whith Line: %d\n",
			poem.ID, poem.LineNumber, res.LineNumber)
		return
	}
	p.ID2Poems[poem.ID] = poem
}

func (p *ChinesePoems) Count() int {
	return len(p.ID2Poems)
}

func (p *ChinesePoems) GetAllIDs() []string {
	ids := []string{}

	for k, _ := range p.ID2Poems {
		ids = append(ids, k)
	}

	return ids
}

func (p *ChinesePoems) GetPoem(id string) *ChinesePoem {
	if cp, ok := p.ID2Poems[id]; ok {
		return cp
	}

	return nil
}