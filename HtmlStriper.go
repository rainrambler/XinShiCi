package main

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
)

type HtmlStriper2 struct {
	basedir string

	allLines []string
}

func (p *HtmlStriper2) ParseDir(dirname string) {
	p.basedir = dirname
	p.allLines = []string{}

	files := FindAllFilesInDir(dirname)
	for _, onefile := range files {
		p.ParseFile(onefile)
	}

	//WriteLines(p.allLines, "cleaned.txt")
}

// clean all html tags
func (p *HtmlStriper2) ParseFile(filename string) {
	//fmt.Printf("Parsing %s...\n", filename)
	// create from a file
	s, err := ReadTextFile(filename)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	policy := bluemonday.StrictPolicy()
	stripped := policy.Sanitize(s)
	p.allLines = append(p.allLines, stripped)
	WriteTextFile(filename+".txt", stripped)
}
