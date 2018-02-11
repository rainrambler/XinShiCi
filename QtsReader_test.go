package main

import (
	"testing"
)

func TestCreatePoem(t *testing.T) {
	s := `1_5	幸武功慶善宮	李世民	壽丘惟舊跡，酆邑乃前基。`

	qp := CreateQtsPoem(s, 0)

	if qp.ID != "1_5" {
		t.Errorf("TestCreatePoem failed: %v", s)
	}

	if qp.Title != "幸武功慶善宮" {
		t.Errorf("TestCreatePoem failed: %v", s)
	}

	if qp.Author != "李世民" {
		t.Errorf("TestCreatePoem failed: %v", s)
	}

	if qp.Sentences[1] != "酆邑乃前基" {
		t.Errorf("TestCreatePoem failed: %v", s)
	}
}
