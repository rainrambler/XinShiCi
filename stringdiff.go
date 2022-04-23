package main

import (
	"bytes"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func convdiffmatchpatch(diffs []diffmatchpatch.Diff) string {
	var buff bytes.Buffer
	for _, diff := range diffs {
		text := diff.Text

		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			_, _ = buff.WriteString("+[")
			_, _ = buff.WriteString(text)
			_, _ = buff.WriteString("]")
		case diffmatchpatch.DiffDelete:
			_, _ = buff.WriteString("-[")
			_, _ = buff.WriteString(text)
			_, _ = buff.WriteString("]")
		case diffmatchpatch.DiffEqual:
			_, _ = buff.WriteString(text)
		}
	}

	return buff.String()
}

func compString(text1, text2 string) string {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(text1, text2, false)

	return convdiffmatchpatch(diffs)
}
