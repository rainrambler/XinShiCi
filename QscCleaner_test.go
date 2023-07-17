package main

import (
	"testing"
)

func TestChangeLineZh1(t *testing.T) {
	var qc QscChanger
	qc.allCipais.Init(`CiPaiZh.txt`)
	line := `元夕詞（鷓鴣天）`
	v := qc.changeLineZh(line)

	wanted := true
	if v != wanted {
		t.Errorf(" failed: %v, want: %v", v, wanted)
	}
}
