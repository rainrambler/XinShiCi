package main

import (
	"testing"
)

func TestWordSingle(t *testing.T) {
	s := `人`

	var cw ChineseWord
	cw.Chs = s

	if !cw.IsSingle() {
		t.Errorf("TestWordSingle failed: %v", cw.Chs)
	}
}

func TesSubChineseString(t *testing.T) {
	s := `壽丘惟舊跡，酆邑乃前基。`

	part := SubChineseString(s, 0, 2)

	if part != `壽丘` {
		t.Errorf("TesSubChineseString failed: %s", part)
	}

	part = SubChineseString(s, 1, 2)

	if part != `丘惟` {
		t.Errorf("TesSubChineseString failed: %s", part)
	}
}
