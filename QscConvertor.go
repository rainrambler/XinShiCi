package main

import (
	"log"
	"strings"
)

// remove prefix and suffix spaces and comments (tag : |< >|)
func lineFormat(line string) string {
	linenew := strings.TrimSpace(line)

	posStart := strings.Index(linenew, "|<")
	if posStart == -1 {
		return linenew
	}

	posEnd := strings.Index(linenew, ">|")
	if posEnd == -1 {
		log.Printf("INFO: Line only has start comment tag: %s\n", line)
		return linenew
	}

	leftPart := SubString(linenew, 0, posStart)
	rightPart := SubString(linenew, posEnd+2, len(linenew))

	return leftPart + rightPart
}

func getPartNumber(cipai string) int {
	switch cipai {
	case "莺啼序":
		return 4
	case "瑞龙吟":
		return 4
	case "十样花":
		return 7
	default:
		return 2
	}
}

type ZhCharCount struct {
	r2c Rhyme2Count
}

func (p *ZhCharCount) Init() {
	p.r2c.Init()
}

func (p *ZhCharCount) AddPoem(poem *ChinesePoem) {
	for _, s := range poem.Sentences {
		rs := []rune(s)
		for i := 0; i < len(rs)-1; i++ {
			// each two chars
			chpair := []rune{rs[i], rs[i+1]}
			p.r2c.Add(string(chpair))
		}
	}
}

func CreateQscPoem(id, author, title, content, comment string) *ChinesePoem {
	cp := new(ChinesePoem)
	cp.ID = id
	cp.Author = author
	cp.Title = title
	cp.AllText = content

	if len(comment) > 0 {
		cp.Comments = comment
	}

	cp.ParseSentences()
	//cp.analyseRhyme()

	return cp
}
