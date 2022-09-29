package main

import (
	"testing"
)

func TestPinyin1(t *testing.T) {
	s := `壽丘惟舊跡，酆邑乃前基。`
	pinyinstr := ""
	for _, ch := range s {
		res := pyf.FindPinyin(string(ch))

		if len(res) > 0 {
			pinyinstr += res + " "
		} else {
			pinyinstr += string(ch)
		}
	}

	if pinyinstr != "shou4 qiu1 wei2 jiu4 ji4 ，Feng1 yi4 nai3 qian2 ji1 。" {
		t.Errorf("TestPinyin1 failed: %v, original: %s", pinyinstr, s)
	}
}

func TestFindPingze2(t *testing.T) {
	s := `基`
	res := pyf.FindPingze2(s)
	expected := PingZePing

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestFindPingze3(t *testing.T) {
	s := `海`
	res := pyf.FindPingze2(s)
	expected := PingZeZe

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

var pyf PinyinFinder

func init() {
	pyf.Init("zht2py.txt")
}
