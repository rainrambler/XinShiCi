package main

import (
	"strings"
)

func hasErrorTitle(poem *ChinesePoem) bool {
	if strings.Contains(poem.Title, `《`) {
		return true
	}

	if strings.Contains(poem.Title, `》`) {
		return true
	}

	return false
}

const ErrorChars = `abcdefghijklmnopqrstuvwxyz` +
	`ABCDEFGHIJKLMNOPQRSTUVWXYZ` + `0123456789`

// English and number chars in Chinese text
func hasErrorText(text string) bool {
	return strings.ContainsAny(text, ErrorChars)
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

func findAllErrorText(poem *ChinesePoem) ([]string, string) {
	res := ""
	arr := []string{}
	for _, line := range poem.Sentences {
		subarr := parseAllErrorText(line)
		if len(subarr) > 0 {
			arr = append(arr, subarr...)
			res = line + "#"
		}
	}

	return arr, res
}

func parseAllErrorText(line string) []string {
	arr := []string{}
	rs := []rune(line)

	curPart := ""
	for _, r := range rs {
		if strings.ContainsRune(ErrorChars, r) {
			curPart += string(r)
		} else {
			if curPart != "" {
				arr = append(arr, curPart)
				curPart = ""
			}
		}
	}

	if curPart != "" {
		arr = append(arr, curPart)
	}

	return arr
}
