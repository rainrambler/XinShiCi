// QscConvertor
package main

import (
	"fmt"
	"strings"
)

const (
	CIPAI_MAX = 7
)

type Cipais struct {
	item2id map[string]int
}

func (p *Cipais) Init(filename string) {
	p.item2id = make(map[string]int)

	lines := ReadTxtFile(filename)

	id := 100
	for _, line := range lines {
		if len(line) > 0 {
			p.item2id[line] = id + 2
		}
	}

	fmt.Printf("INFO: Total CiPai: %d\n", len(p.item2id))
}

func (p *Cipais) Exists(nm string) bool {
	if _, ok := p.item2id[nm]; ok {
		return true
	}
	return false
}

func (p *Cipais) Count() int {
	return len(p.item2id)
}

func (p *Cipais) HasCipai(line string) (bool, string) {
	s := line
	chsize := ChcharLen(s)
	if chsize < 2 {
		return false, s
	}

	if IsCommentLine(s) {
		return false, s
	}

	pos := strings.Index(s, " ")
	if pos != -1 {
		// contains blank
		leftpart := s[:pos]

		if ChcharLen(leftpart) > CIPAI_MAX {
			return false, s
		}

		return p.HasActualCipai(leftpart), leftpart
	}

	pos = strings.Index(s, "（")
	if pos != -1 {
		// contains bracket
		leftpart := s[:pos]
		if ChcharLen(leftpart) > CIPAI_MAX {
			return false, leftpart
		}

		return p.HasActualCipai(leftpart), leftpart
	}

	if chsize > CIPAI_MAX {
		return false, s
	}

	return p.HasActualCipai(s), s
}

func (p *Cipais) HasActualCipai(line string) bool {
	return p.Exists(line)
}
