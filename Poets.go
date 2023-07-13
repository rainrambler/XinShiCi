package main

import (
	"strings"
)

type Poets struct {
	poet2id map[string]int
}

func (p *Poets) Init(filename string) {
	p.poet2id = make(map[string]int)

	lines := ReadTxtFile(filename)

	id := 100
	for _, line := range lines {
		if len(line) > 0 {
			poetname := getPoetName(line)
			p.poet2id[poetname] = id + 2
		}
	}
}

func (p *Poets) IsPoet(nm string) bool {
	if _, ok := p.poet2id[nm]; ok {
		return true
	}
	return false
}

func (p *Poets) FindPoet(nm string) int {
	if id, ok := p.poet2id[nm]; ok {
		return id
	}
	return -1
}

func (p *Poets) Count() int {
	return len(p.poet2id)
}

// 衛芳華 // 詞綜 (四庫全書本)卷25
func getPoetName(line string) string {
	pos := strings.Index(line, `//`)
	if pos == -1 {
		return line
	}

	return strings.TrimSpace(line[:pos])
}
