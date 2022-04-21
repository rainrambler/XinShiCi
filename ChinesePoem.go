package main

import (
	"strings"
)

type ChinesePoem struct {
	ID         string
	Title      string
	Comments   string
	Author     string
	Sentences  []string
	AllText    string
	LineNumber int
	Rhyme      string // 韵脚
}

func (p *ChinesePoem) ParseSentences() {
	p.Sentences = strings.FieldsFunc(p.AllText, SplitPoem)
}

func (p *ChinesePoem) toDesc() string {
	return p.ID + "|" + p.Author + "|" + p.Title + "|" + SubChineseString(p.AllText, 0, 20)
}

func (p *ChinesePoem) toFullDesc() string {
	return p.ID + "|" + p.Author + "|" + p.Title + "|" + p.AllText
}

func SplitPoem(r rune) bool {
	return r == '；' || r == '，' || r == '。' || r == '！' || r == '？' || r == '、'
}

func (p *ChinesePoem) analyseRhyme() {
	lastwords := p.collectLastWords()
	s := g_ZhRhymes.AnalyseRhyme(lastwords)

	p.Rhyme = s
}

func (p *ChinesePoem) collectLastWords() []string {
	if len(p.Sentences) == 0 {
		return []string{}
	}

	lastwords := []string{}

	for _, sentence := range p.Sentences {
		lastchar := getLastZhChar(sentence)

		lastwords = append(lastwords, lastchar)
	}

	return lastwords
}

func (p *ChinesePoem) FindByYayunLength(yayun string, chlen int) []string {
	arr := []string{}

	for _, sentence := range p.Sentences {
		curlen := ChcharLen(sentence)
		if curlen != chlen {
			continue
		}

		lastchar := getLastZhChar(sentence)
		curRhyme := g_ZhRhymes.findRhyme(lastchar)
		if curRhyme == yayun {
			arr = append(arr, sentence)
		}
	}

	return arr
}

func (p *ChinesePoem) FindByYayunLengthPingze(yayun string,
	chlen, pztype int) []string {
	arr := []string{}

	for _, sentence := range p.Sentences {
		curlen := ChcharLen(sentence)

		if (curlen != chlen) && (chlen != 0) {
			continue
		}

		lastchar := getLastZhChar(sentence)
		curRhyme := g_ZhRhymes.findRhymePingze(lastchar, pztype)
		if curRhyme == yayun {
			arr = append(arr, sentence)
		}
	}

	return arr
}

func getLastZhChar(s string) string {
	return SubChineseString(s, ChcharLen(s)-1, 1)
}

func (p *ChinesePoem) hasRepeatChar() bool {
	if len(p.Sentences) == 0 {
		return false
	}

	for _, sentence := range p.Sentences {
		if isRepeatChar(sentence) {
			return true
		}
	}

	return false
}

func isRepeatChar(sentense string) bool {
	rs := []rune(sentense)

	totallen := len(rs)
	for i := 0; i < totallen-1; i++ {
		if rs[i] == rs[i+1] {
			return true
		}
	}

	return false
}

func hasErrorTitle(poem *ChinesePoem) bool {
	if strings.Contains(poem.Title, `《`) {
		return true
	}

	if strings.Contains(poem.Title, `》`) {
		return true
	}

	return false
}

// English and number chars in Chinese text
func hasErrorText(text string) bool {
	return strings.ContainsAny(text, `abcdefghijklmnopqrstuvwxyz`+
		`ABCDEFGHIJKLMNOPQRSTUVWXYZ`+`0123456789`)
}

// English and number chars in Chinese text
func findErrorText(poem *ChinesePoem) string {
	res := ""
	for _, line := range poem.Sentences {
		if hasErrorText(line) {
			res = line + "#"
		}
	}

	return res
}
