package main

import (
	"fmt"
	"testing"
)

func TestAnalyseZhRhyme1(t *testing.T) {
	rs := []rune("下中翁动守风空梦")

	var cr ZhRhymes
	cr.Init()

	rhyval := cr.AnalyseRhyme(rs)
	fmt.Printf("TestAnalyseZhRhyme1 Result: %s\n", rhyval)

	wanted := "18"
	if rhyval != wanted {
		t.Errorf(" failed: %v, want: %v", rhyval, wanted)
	}
}

func TestFindRhymePingze1(t *testing.T) {
	var cr ZhRhymes
	cr.Init()

	r := GetFirstRune(`酒`)
	val := cr.findRhymePingze(r, PingZePing)
	want := ""

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}
