package main

import (
	"testing"
)

func TestChangeFileExt1(t *testing.T) {
	fn := `d:\\aaa.txt`
	ext := `.png`

	res := ChangeFileExt(fn, ext)

	wanted := `d:\\aaa.png`
	if res != wanted {
		t.Errorf(" failed: %v, want: %v", res, wanted)
	}
}
