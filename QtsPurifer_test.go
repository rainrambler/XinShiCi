package main

import (
	"testing"
)

func TestCleanComment1(t *testing.T) {
	s := "abc"
	res := CleanComment(s)
	expected := s

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestCleanComment2(t *testing.T) {
	s := "abc〈123〉"
	res := CleanComment(s)
	expected := "abc"

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestCleanComment3(t *testing.T) {
	s := "abc〈123〉def"
	res := CleanComment(s)
	expected := "abcdef"

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
