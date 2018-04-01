package main

import (
	"testing"
)

func TestPinyinCreate1(t *testing.T) {
	s := "deng1"
	py := CreatePinyin(s)

	if py.Shengdiao != "1" {
		t.Errorf("TestPinyinCreate1 failed: %v, want: 1", py.Shengdiao)
	}

	if py.Shengmu != "d" {
		t.Errorf("TestPinyinCreate1 failed: %v, want: d", py.Shengmu)
	}

	if py.Yunmu != "eng" {
		t.Errorf("TestPinyinCreate1 failed: %v, want: eng", py.Yunmu)
	}
}

func TestPinyinCreate2(t *testing.T) {
	s := "zi3"
	py := CreatePinyin(s)

	if py.Shengdiao != "3" {
		t.Errorf(" failed: %v, want: 3", py.Shengdiao)
	}

	if py.Shengmu != "z" {
		t.Errorf(" failed: %v, want: z", py.Shengmu)
	}

	if py.Yunmu != "i" {
		t.Errorf(" failed: %v, want: i", py.Yunmu)
	}
}
