package main

import (
	"log"
	"strings"
)

type PinyinFinder struct {
	hz2pinyin map[string]string
	hz2pingze map[string]int
}

func (p *PinyinFinder) Init(filename string) {
	p.hz2pinyin = make(map[string]string)
	p.hz2pingze = make(map[string]int)

	lines := ReadTxtFile(filename)

	for idx, line := range lines {
		if !strings.Contains(line, "|") {
			log.Printf("WARN: Format error in dict line [%d]: %s\n", idx, SubString(line, 0, 10))
		} else {
			arr := strings.Split(line, "|")

			zhchar := arr[0]
			zhPy := arr[1]
			p.hz2pinyin[zhchar] = zhPy

			shengdiao := zhPy[len(zhPy)-1:] // last char
			pingze := GetPingze(shengdiao)

			p.SetPingze(zhchar, pingze)
		}
	}
}

func (p *PinyinFinder) SetPingze(zhchar string, pzval int) {
	if pzval == PingZeAny {
		log.Printf("[DBG]Possible error Pingze for %s: %d\n", zhchar, pzval)
		return
	}

	if pzval == PingZeUnknown {
		log.Printf("[DBG]Possible unknown Pingze for %s: %d\n", zhchar, pzval)
		return
	}

	v, exists := p.hz2pingze[zhchar]
	if !exists {
		p.hz2pingze[zhchar] = pzval
		return
	}

	if v == pzval {
		return
	}

	p.hz2pingze[zhchar] = PingZeAny
}

func (p *PinyinFinder) FindPinyin(zhchar string) string {
	if py, ok := p.hz2pinyin[zhchar]; ok {
		return py
	}

	return "" // empty
}

func (p *PinyinFinder) FindPingze(zhchar string) int {
	if py, ok := p.hz2pinyin[zhchar]; ok {
		shengdiao := py[len(py)-1:]

		return GetPingze(shengdiao)
	}

	return PingZeUnknown
}

func (p *PinyinFinder) FindPingze2(zhchar string) int {
	if py, ok := p.hz2pingze[zhchar]; ok {
		return py
	}

	return PingZeUnknown
}
