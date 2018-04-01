// QscConvertor
package main

import (
	"fmt"
)

type Lines struct {
	AllLines []int
}

func (p *Lines) ToOneLine() string {
	s := ""
	for _, item := range p.AllLines {
		s += fmt.Sprintf("%d", item) + ", "
	}

	return s
}

type Keywords struct {
	AllItems []string
}

func (p *Keywords) ToOneLine() string {
	s := ""
	for _, item := range p.AllItems {
		s += item + ", "
	}

	return s
}

func (p *Lines) Count() int {
	return len(p.AllLines)
}

type Keyword2Lines struct {
	Key2Lines map[string]*Lines
}

func (p *Keyword2Lines) Init() {
	p.Key2Lines = make(map[string]*Lines)
}

func (p *Keyword2Lines) AddLine(keyword string, line int) {
	if keyword == "" {
		return
	}

	ls, exists := p.Key2Lines[keyword]

	if exists {
		ls.AllLines = append(ls.AllLines, line)
	} else {
		lsnew := new(Lines)
		lsnew.AllLines = append(lsnew.AllLines, line)

		p.Key2Lines[keyword] = lsnew
	}
}

func (p *Keyword2Lines) DemoPrint() {
	count2lines := map[int]*Keywords{}

	totallines := 0
	for k, v := range p.Key2Lines {
		totallines += v.Count()

		kwds, exists := count2lines[v.Count()]

		if exists {
			kwds.AllItems = append(kwds.AllItems, k)
		} else {
			kwdsnew := new(Keywords)
			kwdsnew.AllItems = append(kwdsnew.AllItems, k)

			count2lines[v.Count()] = kwdsnew
		}
	}

	for i := 60; i > 0; i-- {
		kwds, exists := count2lines[i]

		if exists {

			fmt.Printf("%d: %s\n", i, kwds.ToOneLine())
			for _, item := range kwds.AllItems {
				ls := p.Key2Lines[item]

				fmt.Printf("%s: %s\n", item, ls.ToOneLine())
			}

			fmt.Println("---------------------")
		}

	}

	fmt.Printf("INFO: Total %d, unique: %d\n", totallines, len(p.Key2Lines))

}
