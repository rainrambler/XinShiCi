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
