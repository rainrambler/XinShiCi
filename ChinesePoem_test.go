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

func TestParseAllErrorText1(t *testing.T) {
	s := "  aa bb"

	arr := parseAllErrorText(s)
	res := len(arr)
	expected := 2

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}

	res1 := arr[1]
	expected1 := "bb"
	if res1 != expected1 {
		t.Errorf("Result: %v, want: %v", res1, expected1)
	}
}

func TestParseAllErrorText2(t *testing.T) {
	s := "  aa bb  "

	arr := parseAllErrorText(s)
	res := len(arr)
	expected := 2

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}

	res1 := arr[1]
	expected1 := "bb"
	if res1 != expected1 {
		t.Errorf("Result: %v, want: %v", res1, expected1)
	}
}

func TestHasRepeatWords1(t *testing.T) {
	s1 := `abcdeabcd`

	val := HasRepeatWords(s1)
	want := true

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestHasRepeatWords2(t *testing.T) {
	s1 := `abcdefgabchhh`

	val := HasRepeatWords(s1)
	want := true

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestHasRepeatWords3(t *testing.T) {
	s1 := `abcdefghijklmn`

	val := HasRepeatWords(s1)
	want := false

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestHasRepeatWords4(t *testing.T) {
	s1 := `昨夜星辰昨夜风`

	val := HasRepeatWords(s1)
	want := true

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}
