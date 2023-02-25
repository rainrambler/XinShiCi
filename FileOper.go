package main

import (
	"os"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// file does not exist
		return false
	}

	return true
}
