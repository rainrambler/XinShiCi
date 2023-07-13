package main

import (
	"testing"
)

func TestMatch2(t *testing.T) {
	s := "中平中仄仄平平"

	var sen Sentence
	sen.Parse(s)

	res := sen.Match(`平山闌檻倚晴風`)
	expected := true

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestParseString2(t *testing.T) {
	// 《朝中措·平山堂》
	s := "中平中仄仄平平，中仄仄平平。中仄中平中仄，中平中仄平平？"

	var cp Cipai
	res := cp.ParseString(s)
	cp.DbgPrint()

	expected := true

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestMatch1(t *testing.T) {
	s := "中平中仄仄平平，中仄仄平平。中仄中平中仄，中平中仄平平？"

	var cp Cipai
	res := cp.ParseString(s)
	//cp.DbgPrint()

	expected := true

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}

	res = cp.Match(`平山闌檻倚晴風，山色有無風`)
	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
