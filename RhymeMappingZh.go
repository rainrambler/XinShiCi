package main

import (
	"fmt"
)

type RhymeMappingZh struct {
	Rhyme2Desc map[string]string
}

func (p *RhymeMappingZh) Init() {
	p.Rhyme2Desc = map[string]string{
		"一麻":  "P1",
		"二波":  "P2",
		"三歌":  "P3",
		"四皆":  "P4",
		"五支":  "P5",
		"六兒":  "P6",
		"七齊":  "P7",
		"八微":  "P8",
		"九開":  "P9",
		"十姑":  "P10",
		"十一魚": "P11",
		"十二侯": "P12",
		"十三豪": "P13",
		"十四寒": "P14",
		"十五痕": "P15",
		"十六唐": "P16",
		"十七庚": "P17",
		"十八東": "P18",
		"一把":  "Z1",
		"二跛":  "Z2",
		"三扯":  "Z3",
		"四解":  "Z4",
		"五齒":  "Z5",
		"六爾":  "Z6",
		"七比":  "Z7",
		"八北":  "Z8",
		"九矮":  "Z9",
		"十布":  "Z10",
		"十一舉": "Z11",
		"十二醜": "Z12",
		"十三襖": "Z13",
		"十四俺": "Z14",
		"十五本": "Z15",
		"十六榜": "Z16",
		"十七繃": "Z17",
		"十八寵": "Z18",
		"一八":  "R1",
		"二剝":  "R2",
		"三鴿":  "R3",
		"四鱉":  "R4",
		"五吃":  "R5",
		"六逼":  "R6",
		"七出":  "R7",
		"八曲":  "R8",
	}
}

func (p *RhymeMappingZh) FindDesc(rhyme string) string {
	if res, ok := p.Rhyme2Desc[rhyme]; ok {

		return res
	}

	// Not found
	fmt.Printf("[DBG]RhymeMappingZh: Cannot find rhyme: %s!\n", rhyme)
	return ""
}
