package main

import (
	"fmt"

	"github.com/grokify/html-strip-tags-go" // => strip
	"github.com/microcosm-cc/bluemonday"
)

type HtmlStriper struct {
	basedir string

	allLines []string
}

func (p *HtmlStriper) ParseDir(dirname string) {
	p.basedir = dirname
	p.allLines = []string{}

	files := FindAllFilesInDir(dirname)
	for _, onefile := range files {
		p.ParseFile(onefile)
	}
}

// clean all html tags
func (p *HtmlStriper) ParseFile(filename string) {
	//fmt.Printf("Parsing %s...\n", filename)
	// create from a file
	s, err := ReadTextFile(filename)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	stripped := strip.StripTags(s)
	p.allLines = append(p.allLines, stripped)
	//WriteTextFile(filename+".txt", stripped)
}

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
