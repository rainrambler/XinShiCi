package main

import (
	"fmt"
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
	p.Sentences = strings.FieldsFunc(p.AllText, IsPunctuation)
}

func (p *ChinesePoem) toDesc() string {
	return p.ID + "|" + p.Author + "|" + p.Title + "|" + p.LeftChars(20)
}

func (p *ChinesePoem) title() string {
	s := p.ID + "|" + p.Author + "|" + p.Title
	if len(p.Sentences) > 0 {
		s += "|" + p.Sentences[0]
	}
	return s
}

func (p *ChinesePoem) toFullDesc() string {
	return p.ID + "|" + p.Author + "|" + p.Title + "|" + p.AllText
}

func IsPunctuation(r rune) bool {
	return r == '；' || r == '，' || r == '。' || r == '！' || r == '？' || r == '、'
}

func IsPunctuationAll(r rune) bool {
	if IsPunctuation(r) {
		return true
	}

	switch r {
	case ';', ',', '（', '）', '《', '》', '(', ')':
		return true
	default:
		return false
	}
}

func (p *ChinesePoem) analyseRhyme() {
	lastwords := p.collectLastWords()
	s := g_ZhRhymes.AnalyseRhyme(lastwords)

	p.Rhyme = s
}

func (p *ChinesePoem) collectLastWords() []rune {
	if len(p.Sentences) == 0 {
		return []rune{}
	}

	lastwords := []rune{}

	for _, sentence := range p.Sentences {
		lastchar := GetLastRune(sentence)

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

		lastchar := GetLastRune(sentence)
		curRhyme := g_ZhRhymes.FindRhyme2(lastchar)
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

		lastchar := GetLastRune(sentence)
		curRhyme := g_ZhRhymes.findRhymePingze(lastchar, pztype)
		if curRhyme == yayun {
			arr = append(arr, sentence)
		}
	}

	return arr
}

func (p *ChinesePoem) LeftChars(n int) string {
	if len(p.AllText) <= n {
		return p.AllText
	}
	return SubChineseString(p.AllText, 0, n) + ".."
}

// [StartPos, EndPos]
func (p *ChinesePoem) FindContext(id int) string {
	startPos := id - 2
	endPos := id + 2
	if startPos < 0 {
		startPos = 0
		endPos = startPos + 5
	}
	if endPos >= len(p.Sentences) {
		endPos = len(p.Sentences) - 1

		startPos = endPos - 5
		if startPos < 0 {
			startPos = 0
		}
	}

	s := ""
	for i := startPos; i <= endPos; i++ {
		s += p.Sentences[i] + ","
	}

	if len(s) <= 1 {
		return s
	}

	return s[:len(s)-1]
}

func getLastZhChar(s string) string {
	return SubChineseString(s, ChcharLen(s)-1, 1)
}

func (p *ChinesePoem) hasRepeatChar() bool {
	if len(p.Sentences) == 0 {
		return false
	}

	for _, sentence := range p.Sentences {
		// eg: 梨花院落 溶溶 月
		if HasRepeatChars(sentence) {
			return true
		}
	}

	return false
}

func HasRepeatChars(sentense string) bool {
	rs := []rune(sentense)

	totallen := len(rs)
	for i := 0; i < totallen-1; i++ {
		if rs[i] == rs[i+1] {
			return true
		}
	}

	return false
}

func (p *ChinesePoem) hasRepeatWords() bool {
	if len(p.Sentences) == 0 {
		return false
	}

	for _, sentence := range p.Sentences {
		// eg: 昨夜 星辰 昨夜 风
		if HasRepeatWordsZh(sentence) {
			fmt.Println(sentence)
			return true
		}
	}

	return false
}

// 花 未 全开月 未 圆 ==> true
func (p *ChinesePoem) hasRepeatDiffs() bool {
	if len(p.Sentences) == 0 {
		return false
	}

	for _, sentence := range p.Sentences {
		if HasRepeatCharsZh(sentence) {
			fmt.Println(sentence)
			return true
		}
	}

	return false
}

// 花 未 全开月 未 圆
func (p *ChinesePoem) FindRepeatDiffs2() []string {
	arr := []string{}
	if len(p.Sentences) == 0 {
		return arr
	}

	for _, sentence := range p.Sentences {
		if HasRepeatCharsZh(sentence) {
			//return sentence
			arr = append(arr, sentence)
		}
	}

	return arr
}

// Not for Chinese string
func HasRepeatWords(sentense string) bool {
	totallen := len(sentense)
	if totallen <= 3 {
		return false
	}
	maxlen := totallen / 2
	for i := 0; i < maxlen-1; i++ {
		for j := maxlen; j > i; j-- {
			substr := sentense[i:j]
			if len(substr) <= 1 {
				return false
			}

			remain := sentense[j+1:]
			if strings.Contains(remain, substr) {
				//fmt.Printf("[%d:%d]: Sub: %s, Remain: %s\n", i, j, substr, remain)
				return true
			}
		}
	}

	return false
}

func (p *ChinesePoem) ContainsAll(keywords []string) bool {
	for _, key := range keywords {
		if len(key) == 0 {
			continue
		}
		if !strings.Contains(p.AllText, key) {
			return false
		}
	}

	return true
}
