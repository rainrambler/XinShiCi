package main

import (
	"log"
	"strings"
)

type Qts struct {
	ChinesePoems
}

func (p *Qts) ReadFile(filename string) {
	lines := ReadTxtFile(filename)

	for idx, line := range lines {
		qp := CreateQtsPoem(line, idx+1) // the first line is 1

		if qp != nil {
			p.AddPoem(qp)
		}
	}
}

func SplitLine(r rune) bool {
	return r == '\t' || r == ' '
}

// https://stackoverflow.com/questions/39862613/how-to-split-multiple-delimiter-in-golang
func CreateQtsPoem(line string, idx int) *ChinesePoem {
	arr := strings.FieldsFunc(line, SplitLine)

	if len(arr) != 4 {
		log.Printf("WARN: Format error in line [%d]: %s\n", idx, SubString(line, 0, 10))
		return nil
	}

	var poem ChinesePoem
	//poem.ID = arr[0] + "|" + arr[1] // ID + Title
	poem.ID = arr[0]
	poem.Title = arr[1]
	poem.Author = arr[2]
	poem.AllText = arr[3]
	poem.LineNumber = idx

	poem.ParseSentences()

	return &poem
}
