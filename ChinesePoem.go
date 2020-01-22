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
	return p.ID + "|" + p.Author + "|" + p.Title + "|" + SubChineseString(p.AllText, 0, 5)
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
		if curlen != chlen {
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
