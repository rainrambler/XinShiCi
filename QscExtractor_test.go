package main

import (
	"testing"
)

func TestFindPoemIds(t *testing.T) {
	filename := `d:\aaa11.txt`
	val := findPoemId(filename)
	want := 11

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestFindPoemIds2(t *testing.T) {
	filename := `d:\aaa\bbb.txt`
	val := findPoemId(filename)
	want := 0

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestFindPoemIds3(t *testing.T) {
	filename := `d:\aaa\bbb1.txt`
	val := findPoemId(filename)
	want := 1

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestFindPoemIds4(t *testing.T) {
	filename := `d:\aaa\50.txt`
	val := findPoemId(filename)
	want := 50

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestFindPoemIds5(t *testing.T) {
	filename := `d:\aaa\西江月3.txt`
	val := findPoemId(filename)
	want := 3

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestPrefixCharSame1(t *testing.T) {
	s1 := `12345`
	s2 := `12abc`

	val := prefixCharSame(s1, s2, 2)
	want := true

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestPrefixCharSame2(t *testing.T) {
	s1 := `12345`
	s2 := `12abc`

	val := prefixCharSame(s1, s2, 3)
	want := false

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestPrefixCharSame3(t *testing.T) {
	s1 := `12345`
	s2 := `abcde`

	val := prefixCharSame(s1, s2, 2)
	want := false

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}

func TestCombineNewFile1(t *testing.T) {
	s1 := `d:\aaa\abc.txt`
	id := 20

	val := combineNewFile(s1, id)
	want := `d:\aaa\abc20.txt`

	if val != want {
		t.Errorf("failed: %v, want: %v", val, want)
	}
}
