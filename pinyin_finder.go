package main

import (
	"log"
	"strings"
)

type PinyinFinder struct {
	hz2pinyin map[string]string
	hz2py     map[rune]string
	hz2pz     map[rune]int
}

func (p *PinyinFinder) Init(filename string) {
	p.hz2pinyin = make(map[string]string)
	p.hz2py = make(map[rune]string)
	p.hz2pz = make(map[rune]int)

	lines := ReadTxtFile(filename)

	for idx, line := range lines {
		if !strings.Contains(line, "|") {
			log.Printf("WARN: Format error in dict line [%d]: %s\n",
				idx, SubString(line, 0, 10))
			continue
		}

		arr := strings.Split(line, "|")

		zhchar := arr[0]
		zhPy := arr[1]
		p.hz2pinyin[zhchar] = zhPy

		zhr := GetFirstRune(zhchar)
		p.hz2py[zhr] = zhPy

		shengdiao := zhPy[len(zhPy)-1:] // last char
		pingze := GetPingze(shengdiao)
		p.SetPingze2(zhr, pingze)
	}
}

func (p *PinyinFinder) SetPingze2(zhchar rune, pzval int) {
	if pzval == PingZeAny {
		log.Printf("[DBG]Possible error Pingze for %s: %d\n",
			string(zhchar), pzval)
		return
	}

	if pzval == PingZeUnknown {
		log.Printf("[DBG]Possible unknown Pingze for %s\n", string(zhchar))
		return
	}

	v, exists := p.hz2pz[zhchar]
	if !exists {
		p.hz2pz[zhchar] = pzval
		return
	}

	if v == pzval {
		return
	}

	p.hz2pz[zhchar] = PingZeAny
}

func (p *PinyinFinder) FindPinyin2(zhchar rune) string {
	if py, ok := p.hz2py[zhchar]; ok {
		return py
	}

	return "" // empty
}

func (p *PinyinFinder) FindPingze2(zhchar rune) int {
	if py, ok := p.hz2pz[zhchar]; ok {
		return py
	}

	return PingZeUnknown
}
