package main

import (
	"fmt"
	"testing"
)

func TestCollectLastWords(t *testing.T) {
	cp := CreateFakePoem()
	cp.ParseSentences()

	lastwords := cp.collectLastWords()
	arrlen := len(lastwords)
	if arrlen != 8 {
		t.Errorf("TestPinyin1 failed: %v, want: 8", arrlen)
	}

	s := ""
	for _, wd := range lastwords {
		s += wd
	}

	fmt.Println(s)
}

func CreateFakePoem() *ChinesePoem {
	cp := new(ChinesePoem)
	cp.ID = "1234"
	cp.Title = "西江月（平山堂）"
	cp.AllText = "三过平山堂下，半生弹指声中。十年不见老仙翁。壁上龙蛇飞动。欲吊文章太守，仍歌杨柳春风。休言万事转头空。未转头时皆梦。"

	return cp
}
