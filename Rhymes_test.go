package main

import (
	"fmt"
	"testing"
)

func TestParseLine1(t *testing.T) {
	s := "【三扯】扯舸可坷惹喏舍者赭饿个和贺荷课锞社射麝赦猞蔗鹧柘"

	var cr ChineseRhymes
	cr.Init()
	cr.parseLine(1, s)

	for k, v := range cr.ZhChar2Rhyme {
		fmt.Printf("[%s]: [%s]\n", k, v.Desc)
	}

	if len(cr.ZhChar2Rhyme) != 24 {
		t.Errorf("TestParseLine1 failed: %v, want: 24", len(cr.ZhChar2Rhyme))
	}
}

func TestAnalyseRhyme1(t *testing.T) {
	rs := []rune("下中翁动守风空梦")
	strs := []string{}
	for _, zhch := range rs {
		strs = append(strs, string(zhch))
	}

	var cr ChineseRhymes
	cr.ImportFile("ShiYunXinBian.txt")

	cr.AnalyseRhyme(strs)
}
