package main

import (
	"fmt"
)

func isSequenceCipai(cipainame string) bool {
	if cipainame == "" {
		return false
	}
	if cipainame == `又` {
		return true
	}

	if cipainame == `同前` {
		return true
	}

	if cipainame == `同上` {
		return true
	}

	arr := []string{`一`, `二`, `三`, `四`, `五`, `六`, `七`, `八`, `九`, `十`, `十一`,
		`十二`, `十三`, `十四`, `十五`}
	for _, item := range arr {
		s := `其` + item
		if cipainame == s {
			return true
		}
	}

	// 二、三、……
	for i := 1; i < len(arr); i++ {
		if cipainame == arr[i] {
			return true
		}
	}

	// 第二
	for _, item := range arr {
		s := `第` + item
		if cipainame == s {
			return true
		}
	}

	for _, item := range arr {
		s := `右` + item
		if cipainame == s {
			return true
		}
	}

	return false
}

// `第三 蓬萊景` ==> true
func isLineSequenceCipai(line string) (bool, string) {
	arr := SplitBlank(line)
	cipai := ""
	switch len(arr) {
	case 0:
		fmt.Printf("Err empty line: %s\n", line)
		return false, line
	case 1:
		cipai = line
	case 2:
		cipai = arr[0]
	default:
		fmt.Printf("Err line: %s\n", line)
		return false, line
	}

	return isSequenceCipai(cipai), cipai
}
