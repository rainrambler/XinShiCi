package main

import (
	"log"
)

const (
	PingZePing    = 1
	PingZeZe      = 2
	PingZeAny     = 0
	PingZeUnknown = -1
)

var g_ZhRhymes ZhRhymes

type ZhRhymes struct {
	ZhChar2Rhyme map[string]string // "an" -> "1"
	pyf          PinyinFinder
}

func (p *ZhRhymes) Init() {
	p.pyf.Init("zhs2py.txt")

	p.ZhChar2Rhyme = map[string]string{"a": "1", "ia": "1", "ua": "1",
		"o": "2", "uo": "2",
		"e":  "3",
		"ie": "4", "ue": "4",
		"i":  "19", // need more analysis
		"er": "6",
		"ei": "8", "ui": "8",
		"ai": "9", "uai": "9",
		"u":  "10",
		"v":  "11",
		"ou": "12", "iu": "12",
		"ao": "13",
		"an": "14", "ian": "14", "uan": "14",
		"en": "15", "in": "15", "un": "15", "vn": "15",
		"ang": "16", "iang": "16", "uang": "16",
		"eng": "17", "ing": "17",
		"ong": "18", "iong": "18",
	}
}

func (p *ZhRhymes) AnalyseRhyme(lastwords []string) string {
	var rhy2count Rhyme2Count
	rhy2count.Init()

	for _, wd := range lastwords {
		curRhyme := p.findRhyme(wd)
		if len(curRhyme) != 0 {
			rhy2count.Add(curRhyme)
		}
	}
	return rhy2count.FindTop1()
}

// `闲` ==> `14` (ian)
func (p *ZhRhymes) findRhyme(chword string) string {
	pystr := p.pyf.FindPinyin(chword)

	if pystr == "" {
		//log.Printf("DBG: Cannot find pinyin for %s!\n", chword)
		return ""
	}

	pyval := CreatePinyin(pystr)
	if pyval == nil {
		log.Printf("DBG: Cannot convert pinyin for %s! Pinyin: %s\n", chword, pystr)
		return ""
	}

	if curRhyme, ok := p.ZhChar2Rhyme[pyval.Yunmu]; ok {
		return curRhyme
	} else {
		return ""
	}
}

// eg: Input: (`闲`, PingZePing), Output:  `14` (ian)
// eg: Input: (`闲`, PingZeZe), Output:  ``
func (p *ZhRhymes) findRhymePingze(chword string, pztype int) string {
	pystr := p.pyf.FindPinyin(chword)

	if pystr == "" {
		//log.Printf("DBG: Cannot find pinyin for %s!\n", chword)
		return ""
	}

	pyval := CreatePinyin(pystr)
	if pyval == nil {
		log.Printf("DBG: Cannot convert pinyin for %s! Pinyin: %s\n", chword, pystr)
		return ""
	}

	if pztype != PingZeAny {
		if pyval.GetPingze() != pztype {
			//fmt.Printf("DBG: Pingze not match: [%s] to [%d]!\n", pyval.toDesc(),
			//	pztype)
			return ""
		}
	}

	if curRhyme, ok := p.ZhChar2Rhyme[pyval.Yunmu]; ok {
		return curRhyme
	} else {
		return ""
	}
}