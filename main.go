package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func readDictFile() {
	var cd ChineseDict
	cd.Init()
	ReadDict("/Users/anyu/goproj/Xindict/cedict.u8", &cd)

	singlecount := 0
	for _, cw := range cd.chs2Word {

		if cw.IsSingle2() {
			//fmt.Println(cw.toDesc())
			singlecount++
		}
	}

	fmt.Printf("Hanzi Count: %d\n", singlecount)
}

// VerboseMode
const (
	Brief    = 2
	Verbose  = 3
	VerbMore = 4
	Unknown  = 1
)

func findQscByChars() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s --find \"keyword\" --length [d] "+
			" -p [d] [--v] [--vv]\n", os.Args[0])
		fmt.Println("-p: pattern (1: Prefix, 2: Suffix, 3: Any, Other: invalid)")
		return
	}

	//fmt.Printf("[DBG]Args: %d\n", len(os.Args))

	var tofind string
	var length int
	var verbose bool
	var verbmore bool
	var pattern int

	flag.StringVar(&tofind, "find", "", "keyword to find")
	flag.IntVar(&length, "length", 0, "Sentence length")
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.BoolVar(&verbmore, "vv", false, "more verbose mode")
	flag.IntVar(&pattern, "p", POS_ANY, "match pattern (1: Prefix, 2: Suffix, 3: Any, Other: invalid)")

	flag.Parse()

	log.Printf("[INFO]%s\n", getCommandLine())

	starttime := time.Now()

	verbmode := Brief
	if verbose {
		if verbmore {
			log.Printf("WARN: Invalid parameters: %v and %v.\n", verbose, verbmore)
			return
		} else {
			verbmode = Verbose
		}
	} else if verbmore {
		verbmode = VerbMore
	}

	findQscKeyword(tofind, pattern, length, verbmode)
	//FindKeysInDbs(tofind, ignores, inBook, verbmode, pattern)

	elapsed := time.Since(starttime)
	log.Printf("Looking up took %s.\n", elapsed)
}

func findRelatedKeyword() {
	if len(os.Args) == 0 {
		fmt.Printf("Usage: %s --find \"keyword\" [--v] [--vv]\n", os.Args[0])
		return
	}

	var tofind string
	var verbose bool
	var verbmore bool

	flag.StringVar(&tofind, "find", "", "keyword to find")
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.BoolVar(&verbmore, "vv", false, "more verbose mode")

	flag.Parse()

	log.Printf("[INFO]%s\n", getCommandLine())
	log.Printf("[INFO]Keyword: %s.\n", tofind)

	starttime := time.Now()

	verbmode := Brief
	if verbose {
		if verbmore {
			log.Printf("WARN: Invalid parameters: %v and %v.\n", verbose, verbmore)
			return
		} else {
			verbmode = Verbose
		}
	} else if verbmore {
		verbmode = VerbMore
	}

	findRelated(tofind, verbmode)
	elapsed := time.Since(starttime)
	log.Printf("Looking up took %s.\n", elapsed)
}

func getCommandLine() string {
	s := ""
	for _, v := range os.Args {
		s += v + " "
	}
	return s
}

func main() {
	//CountPoemZhChars(`新綠`)
	//CountPoemZhChars(`紫府`)

	//findKeywords(`端午 客`)
	//findKeywords(`龍騰`)
	//findKeywords(`龍 雨`)
	//findKeywords(`紫府`)

	//findRepeatChChars()
	//ConvertRepeatCipai(`D:\Projects\Golang\NovelFinder\全唐五代詞.txt`)
	//AnalyseCipai()

	//findRelated(`新綠`, 1)
	//findRelated(`重錦`, 1)
	//findRelatedKeyword()

	//findAllRepeatWordInSentences()

	//readQts()
	//printHans(`allhans.txt`)

	//findQscCipai("定風波", Yun_Ing)
	findQscKeyword("", 0, 0, 0)
	//findQscByChars()

	//findQscDemo()

	//matchCipai()

	//ExportPoetry(`D:\tmp\poem\chinese-poetry-master\ci`, `D:\tmp\poem\exportdemo`)

	//ExportQsc(`D:\Projects\GitHubSync\XinShiCi\qsc.txt`, `D:\tmp\xsc`)

	//GenerateWordCloud(`qsc_zht_fmt.txt`)

	//qscRank()
}
