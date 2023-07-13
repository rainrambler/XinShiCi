package main

import (
	"testing"
)

func TestGetPoetName1(t *testing.T) {
	s := `王子容`

	res := getPoetName(s)
	expected := s

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestGetPoetName2(t *testing.T) {
	s := `王子容 // aaa`

	res := getPoetName(s)
	expected := `王子容`

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
