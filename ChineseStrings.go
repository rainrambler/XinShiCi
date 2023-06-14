package main

import (
	"strings"
)

const HEX_STRING = "0123456789ABCDEF"

func FindFirstStrangeEncoding(s string) []string {

	if !strings.ContainsAny(s, HEX_STRING) {
		return []string{}
	}

	firstpos := strings.IndexAny(s, HEX_STRING)
	firstpart := SubString(s, firstpos, 4) // Length = 4

	remain := SubString(s, firstpos+4, len(s))

	results := []string{firstpart}

	if len(remain) == 0 {
		return results
	}
	arr := FindFirstStrangeEncoding(remain)

	if len(arr) > 0 {
		results = append(results, arr...)
	}

	return results
}

func SubChineseString(s string, beginPos, size int) string {
	rs := []rune(s)
	slen := len(rs)

	if beginPos >= slen {
		return ""
	}

	if beginPos+size >= slen {
		return string(rs[beginPos:])
	}

	return string(rs[beginPos : beginPos+size])
}

const ZH_CHAR_LEN = 3

func ChcharLen(s string) int {
	rs := []rune(s)
	return len(rs)
}

// 昨夜 星辰 昨夜 风 ==> true
func HasRepeatWordsZh(sentense string) bool {
	rs := []rune(sentense)
	rlen := len(rs)

	maxlen := rlen / 2
	for i := 0; i < maxlen-1; i++ {
		for j := maxlen; j > i; j-- {
			subrs := rs[i:j]
			if len(subrs) <= 1 {
				return false
			}

			remain := rs[j+1:]
			if ContainsRunes(remain, subrs) {
				//fmt.Printf("[%d:%d]: Sub: %s, Remain: %s\n", i, j, string(subrs), string(remain))
				return true
			}
		}
	}

	return false
}

// 花 未 全开月 未 圆 ==> true
func HasRepeatCharsZh(sentense string) bool {
	rs := []rune(sentense)
	c2pos := make(map[rune]int)
	for i := 0; i < len(rs); i++ {
		curChar := rs[i]

		pos, exists := c2pos[curChar]
		if exists {
			if pos != i-1 {
				return true
			}
		} else {
			c2pos[curChar] = i
		}
	}

	return false
}

func ContainsRunes(r1, subr1 []rune) bool {
	return IndexRunes(r1, subr1) != -1
}

// https://github.com/tinygo-org/tinygo/blob/release/src/internal/bytealg/bytealg.go
// Index finds the base index of the first instance of the byte sequence b in a.
// If a does not contain b, this returns -1.
func IndexRunes(a, b []rune) int {
	for i := 0; i <= len(a)-len(b); i++ {
		if EqualRunes(a[i:i+len(b)], b) {
			return i
		}
	}
	return -1
}

func ContainsRune(r1 []rune, c1 rune) bool {
	for _, v := range r1 {
		if v == c1 {
			return true
		}
	}

	return false
}

// https://github.com/tinygo-org/tinygo/blob/release/src/internal/bytealg/bytealg.go
func EqualRunes(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
