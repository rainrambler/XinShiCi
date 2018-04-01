package main

import (
	"log"
	"strings"
)

type Pinyin struct {
	Shengmu   string
	Yunmu     string
	Shengdiao string
}

func CreatePinyin(s string) *Pinyin {
	if len(s) == 0 {
		return nil
	}

	if len(s) == 1 {
		// ???
		log.Printf("WARN: Unknown pinyin: %s\n", s)
		return nil
	}

	py := new(Pinyin)
	py.Shengdiao = s[len(s)-1:]

	part := ""
	switch py.Shengdiao {
	case "0", "1", "2", "3", "4":
		part = s[0 : len(s)-1]
	default:
		py.Shengdiao = ""
		part = s
	}

	py.Shengmu, py.Yunmu = SplitPinyin(part)
	return py
}

func SplitPinyin(s string) (Shengmu, Yunmu string) {
	if len(s) == 0 {
		Shengmu = ""
		Yunmu = ""
		log.Printf("WARN: Empty pinyin\n")
		return
	}

	if len(s) == 1 {
		Shengmu = ""
		Yunmu = s
		return
	}

	allshengmu := "bpmfdtnlgkhjqxrzcs"

	if len(s) == 2 {
		firstchar := s[0:1]

		if strings.Contains(allshengmu, firstchar) {
			Shengmu = firstchar
			Yunmu = s[1:]
			return
		} else {
			Shengmu = ""
			Yunmu = s
			return
		}

	}

	if len(s) > 2 {
		first2char := s[0:2]
		switch first2char {
		case "zh", "ch", "sh":
			Shengmu = first2char
			Yunmu = s[2:]
			return
		}

		firstchar := s[0:1]

		if strings.Contains(allshengmu, firstchar) {
			Shengmu = firstchar
			Yunmu = s[1:]
			return
		} else {
			Shengmu = ""
			Yunmu = s
			return
		}
	}

	log.Printf("WARN: Cannot parse pinyin: %s\n", s)
	Shengmu = ""
	Yunmu = ""
	return
}
