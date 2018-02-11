package main

import (
	"log"
	"strings"
)

type PinyinFinder struct {
	hz2pinyin map[string]string
}

func (p *PinyinFinder) Init(filename string) {
	p.hz2pinyin = make(map[string]string)

	lines := ReadTxtFile(filename)

	for idx, line := range lines {
		if !strings.Contains(line, "|") {
			log.Printf("WARN: Format error in dict line [%d]: %s\n", idx, SubString(line, 0, 10))
		} else {
			arr := strings.Split(line, "|")

			p.hz2pinyin[arr[0]] = arr[1]
		}
	}
}

func (p *PinyinFinder) FindPinyin(zhchar string) string {
	if py, ok := p.hz2pinyin[zhchar]; ok {
		return py
	}

	return "" // empty
}
