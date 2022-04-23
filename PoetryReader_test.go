package main

import (
	"testing"
)

func TestChineseToNumber1(t *testing.T) {
	s := "五百"
	res := ChineseToNumber(s)
	expected := 500

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestChineseToNumber2(t *testing.T) {
	s := "五百零一"
	res := ChineseToNumber(s)
	expected := 501

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestChineseToNumber3(t *testing.T) {
	s := "五百一十"
	res := ChineseToNumber(s)
	expected := 510

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestChineseToNumber4(t *testing.T) {
	s := "五百一十二"
	res := ChineseToNumber(s)
	expected := 512

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestChineseToNumber5(t *testing.T) {
	s := "一十二"
	res := ChineseToNumber(s)
	expected := 12

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestChineseToNumber6(t *testing.T) {
	s := "二"
	res := ChineseToNumber(s)
	expected := 2

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestChineseToNumber7(t *testing.T) {
	s := "一十"
	res := ChineseToNumber(s)
	expected := 10

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestChineseToNumber8(t *testing.T) {
	s := "二十"
	res := ChineseToNumber(s)
	expected := 20

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
