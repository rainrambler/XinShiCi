package main

import (
	"fmt"
)

func SubString(s string, beginPos, size int) string {
	slen := len(s)

	if beginPos >= slen {
		return ""
	}

	if beginPos+size >= slen {
		return s[beginPos:]
	}

	return s[beginPos : beginPos+size]
}

func GetLastZhChar(s string, count int) string {
	pos := ChcharLen(s) - count
	return SubChineseString(s, pos, count)
}

// abc, 5 ==> abc005
// https://zetcode.com/golang/string-format/
func combineFilename(prefix string, i int) string {
	return fmt.Sprintf(`%s%03d`, prefix, i)
}

func Arr2String(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	line := ""
	for _, s := range arr {
		line += s + ","
	}

	return line[:len(line)-1] // Remove last ","
}
