package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/PuerkitoBio/goquery"
)

type ZimParser struct {
	basedir string

	allLines []string
	authors  []string
}

func (p *ZimParser) ParseDir(dirname string) {
	p.basedir = dirname
	p.allLines = []string{}
	p.authors = []string{}

	files := FindAllFilesInDir(dirname)
	for _, onefile := range files {
		p.ParseFileQtc(onefile)
	}

	WriteLines(p.allLines, "res.txt")

}

// zim-tools dump result
func (p *ZimParser) ParseFileQtc(filename string) {
	fmt.Printf("Parsing %s...\n", filename)
	// create from a file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
		return
	}

	doc.Find("h2,h3,small,b,div[class=poem]").Each(func(i int, s *goquery.Selection) {

		//demoPrintSection(s)

		sn := goquery.NodeName(s)
		switch sn {
		case "h2":
			{
				// <h2 id="author_name">
				author, _ := s.Attr("id")
				p.AddLine("# " + author)
			}
		case "small":
			{
				authordesc := s.Text()
				p.AddLine("* " + authordesc)
			}
		case "h3":
			{
				title, _ := s.Attr("id")
				p.AddLine("【" + title + "】")
			}
		case "div":
			{
				poem := s.Text()
				p.AddLine(poem)
			}
		case "b":
			{
				content := s.Text()
				p.AddLine(content)
				//fmt.Printf("[DBG]b: %s\n", content)
			}
		default:
			fmt.Printf("Cannot parse node: %v\n", *s)
			demoPrintSection(s)
		}
	})
}

func (p *ZimParser) ParseDirQts(dirname string) {
	p.basedir = dirname
	p.allLines = []string{}

	for i := 1; i <= 900; i++ {
		filename := combineFilename(`卷`, i)
		onefile := path.Join(dirname, filename)

		p.ParseFileQts(onefile)
	}

	WriteLines(p.allLines, "res.txt")
	WriteLines(p.authors, "TangAuthors.txt")
}

// zim-tools dump result (全唐诗)
func (p *ZimParser) ParseFileQts(filename string) {
	fmt.Printf("Parsing %s...\n", filename)
	// create from a file
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
		return
	}

	h1count := countLabel(doc, "h1")
	if h1count == 0 {
		// no <h1>
		p.parseNonH1(doc)
	} else {
		p.parseH1(doc)
	}
}

func (p *ZimParser) parseH1(doc *goquery.Document) {
	doc.Find("h1,h2,h3,b,dd,div[class=poem]").Each(func(i int, s *goquery.Selection) {
		sn := goquery.NodeName(s)
		switch sn {
		case "h1":
			{
				// <h1 id="author_name">
				author, _ := s.Attr("id")
				p.AddLine("# " + author)
			}
		case "h2", "h3":
			{
				title, _ := s.Attr("id")
				p.AddLine("【" + title + "】")
			}
		case "div":
			{
				poem := s.Text()
				p.AddLine(poem)
			}
		case "b", "dd":
			{
				content := s.Text()
				p.AddLine(content)
				//fmt.Printf("[DBG]b: %s\n", content)
			}
		default:
			fmt.Printf("Cannot parse node: %v\n", *s)
			demoPrintSection(s)
		}
	})
}

func (p *ZimParser) parseNonH1(doc *goquery.Document) {
	doc.Find("h2,h3,b,dd,div[class=poem],p").Each(func(i int, s *goquery.Selection) {
		sn := goquery.NodeName(s)
		switch sn {
		case "h2":
			{
				// <h2 id="author_name">?
				author, _ := s.Attr("id")
				p.AddLine("# " + author)
			}
		case "h3":
			{
				title, _ := s.Attr("id")
				p.AddLine("【" + title + "】")
			}
		case "div":
			{
				poem := s.Text()
				p.AddLine(poem)
			}
		case "b", "dd", "p":
			{
				content := s.Text()
				p.AddLine(content)
				//fmt.Printf("[DBG]b: %s\n", content)
			}
		default:
			fmt.Printf("Cannot parse node: %v\n", *s)
			demoPrintSection(s)
		}
	})
}

// zim-tools dump result (全唐诗)
func (p *ZimParser) ParseQts(filename string) {
	fmt.Printf("Parsing %s...\n", filename)
	// create from a file
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
		return
	}

	doc.Find("a[title*=Author]").Each(func(i int, s *goquery.Selection) {
		p.authors = append(p.authors, s.Text())
	})

	doc.Find("h2,h3,b,dd,div[class=poem],a[title*=Author]").Each(func(i int, s *goquery.Selection) {
		sn := goquery.NodeName(s)
		switch sn {
		case "h2", "h3":
			{
				title, _ := s.Attr("id")
				p.AddLine("【" + title + "】")
			}
		case "div":
			{
				poem := s.Text()
				p.AddLine(poem)
			}
		case "a":
			{
				//poem := s.Text()
				//p.AddLine(poem)
				p.AddLine("# " + s.Text())
			}
		case "b", "dd":
			{
				content := s.Text()
				p.AddLine(content)
				//fmt.Printf("[DBG]b: %s\n", content)
			}
		default:
			fmt.Printf("Cannot parse node: %v\n", *s)
			demoPrintSection(s)
		}
	})
}

func countLabel(doc *goquery.Document, label string) int {
	if doc == nil {
		return 0
	}

	if label == "" {
		return 0
	}

	total := 0
	doc.Find(label).Each(func(i int, s *goquery.Selection) {
		total++
	})
	return total
}

func (p *ZimParser) AddLine(line string) {
	p.allLines = append(p.allLines, line)
}

func isParentSection(sel *goquery.Selection) bool {
	parent := sel.Parent()
	if parent == nil {
		return false
	}

	// TODO
	return true
}

func demoPrintSection(sel *goquery.Selection) {
	for _, n := range sel.Nodes {
		fmt.Printf("[DBG]%+v\n", *n)
	}
}
