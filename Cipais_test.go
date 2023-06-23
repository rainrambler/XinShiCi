package main

import (
	"testing"
)

func TestHasActualCipai2(t *testing.T) {
	var cp Cipais
	cp.Init(`CiPaiZh.txt`)

	tofind := `菩薩蠻`

	res := cp.HasActualCipai(tofind)
	expected := true

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
