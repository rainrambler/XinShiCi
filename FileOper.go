package main

import (
	"os"
	"strings"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// file does not exist
		return false
	}

	return true
}

// d:\aaa.txt, .png ==> d:\aaa.png
func ChangeFileExt(filename, extnew string) string {
	revpos := strings.LastIndex(filename, ".")
	if revpos == -1 {
		return filename + extnew
	}

	noext := filename[:revpos]
	return noext + extnew
}
