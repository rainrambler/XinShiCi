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
		fmt.Printf("[%s]: [%s]\n", string(k), v.Desc)
	}

	if len(cr.ZhChar2Rhyme) != 24 {
		t.Errorf("TestParseLine1 failed: %v, want: 24", len(cr.ZhChar2Rhyme))
	}
}

func TestAnalyseRhyme1(t *testing.T) {
	rs := []rune("空中李風相字鐘少翁") // 任昉 朝中措
	strs := []rune{}
	for _, zhch := range rs {
		strs = append(strs, zhch)
	}

	var cr ChineseRhymes
	cr.ImportFile("ShiYunXinBianZH.txt")

	missedChars.Init()

	cr.AnalyseRhyme(strs)
}

func TestSplitShiyun1(t *testing.T) {
	rs := "【ab】123"
	left, right := splitShiyun(1, rs)

	if right != "123" {
		t.Errorf("failed: %v, want: 24", right)
	}

	if left != "ab" {
		t.Errorf("failed: %v, want: 24", right)
	}
}
