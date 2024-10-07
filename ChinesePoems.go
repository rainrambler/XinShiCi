package main

import (
	"fmt"
	"log"
	"strings"
)

type ChinesePoems struct {
	ID2Poems    map[string]*ChinesePoem
	dotFileInst *DotFile
	maxLayer    int
	commonWords *PoemWords
}

func (p *ChinesePoems) Init() {
	p.ID2Poems = make(map[string]*ChinesePoem)
	p.dotFileInst = new(DotFile)
	p.dotFileInst.Init()
	p.maxLayer = 0
	p.commonWords = new(PoemWords)
	p.commonWords.Init(`PoemWordsClean.txt`)
}

func (p *ChinesePoems) AddPoem(poem *ChinesePoem) {
	if poem == nil {
		log.Printf("WARN: AddPoem: nil\n")
		return
	}

	if len(poem.ID) == 0 {
		log.Printf("WARN: AddPoem: empty ID!!\n")
		return
	}
	// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	if res, ok := p.ID2Poems[poem.ID]; ok {
		// exists
		log.Printf("WARN: AddPoem: exists ID: %s, Line: %d whith Line: %d\n",
			poem.ID, poem.LineNumber, res.LineNumber)
		return
	}
	p.ID2Poems[poem.ID] = poem
}

func (p *ChinesePoems) Count() int {
	return len(p.ID2Poems)
}

func (p *ChinesePoems) GetAllIDs() []string {
	ids := []string{}

	for k, _ := range p.ID2Poems {
		ids = append(ids, k)
	}

	return ids
}

func (p *ChinesePoems) GetPoem(id string) *ChinesePoem {
	if cp, ok := p.ID2Poems[id]; ok {
		return cp
	}

	return nil
}

// "白日依山尽", ["白", "山"] ==> true
func (p *ChinesePoems) findByKeywords(keywords []string) *ChinesePoems {
	//fmt.Printf("Keys: %+v\n", keywords)
	var cp ChinesePoems
	cp.Init()

	for _, poem := range p.ID2Poems {
		arr := poem.MatchAllKeywords(keywords)
		if len(arr) == 0 {
			continue
		}
		pnew := poem.Clone()
		pnew.Segment = Arr2String(arr)
		cp.AddPoem(pnew)
	}

	return &cp
}

// "白日依山尽", "白 山" ==> true
func (p *ChinesePoems) FindKeywords(keywords string) *ChinesePoems {
	arr := strings.FieldsFunc(keywords, SplitLine) //blank or tab
	cp := p.findByKeywords(arr)
	return cp
}

func (p *ChinesePoems) PrintResults() {
	for _, v := range p.ID2Poems {
		fmt.Println(v.toFullDesc())
	}
}

func (p *ChinesePoems) FindRepeatDiffs(resultfile string) {
	allRes := []string{}
	totalResults := 0
	for _, v := range p.ID2Poems {
		arr := v.FindRepeatDiffs()
		if len(arr) > 0 {
			//fmt.Printf("[%s][%s][%s]\n", id, v.Title, v.toDesc())
			for _, founded := range arr {
				desc := fmt.Sprintf("[%s]:[%d][%s][%s]", founded,
					totalResults, v.Title, v.toDesc())
				allRes = append(allRes, desc)
				totalResults++
			}

			allRes = append(allRes, "")
		}
	}
	fmt.Printf("Total %d results.\n", totalResults)
	WriteLines(allRes, resultfile)
}

// Keys: [新綠]==>新綠:87, 東風:12, 闌幹:9, 年華:7
func (p *ChinesePoems) FindRelatedWords(keyword, tofile string) {
	cp := p.FindKeywords(keyword)

	var wc WordCloud
	wc.InitParams()
	for _, poem := range cp.ID2Poems {
		for _, oneLine := range poem.Sentences {
			wc.parseMultiLine(oneLine)
		}
	}

	wc.PrintResult(100)

	if tofile == "" {
		return
	}
	wc.CreateDot(keyword, tofile)
}

func (p *ChinesePoems) FindAllRelatedWords(keyword, tofile string, layers int) {
	// Check dependencies
	if p.commonWords.Count() == 0 {
		fmt.Println("[WARN]Typical Words not loaded!")
		return
	}

	p.maxLayer = layers

	p.FindRelatedWordsN(keyword, 0)

	if tofile == "" {
		return
	}

	p.dotFileInst.GenerateFull(tofile)

	CreatePngFromDot(tofile)
}

func (p *ChinesePoems) FindRelatedWordsN(keyword string, layers int) {
	if p.maxLayer <= layers {
		return
	}

	if !p.checkCommonWord(keyword) && (layers != 0) {
		fmt.Printf("[INFO]%s is not a common word.\n", keyword)
		return
	}

	cp := p.FindKeywords(keyword)

	var wc WordCloud
	wc.InitParams()
	for _, poem := range cp.ID2Poems {
		for _, oneLine := range poem.Sentences {
			wc.parseMultiLine(oneLine)
		}
	}

	wc.PrintResult(100)

	k2v := wc.GetTopResult(TOP_WORD * 2)

	validnum := 0
	for k, v := range k2v {
		if p.checkCommonWord(k) {
			p.dotFileInst.AddLink(keyword, k, v)
			validnum++

			if validnum >= TOP_WORD {
				// enough
				break
			}

			p.FindRelatedWordsN(k, layers+1)
		} else {
			if v >= 5 {
				fmt.Printf("[DBG]%s (Ref) is not a common word.\n", k)
			}
		}
	}
}

func (p *ChinesePoems) checkCommonWord(wd string) bool {
	return p.commonWords.Contains(wd)
}

func (p *ChinesePoems) FindSentense(qc *QueryCondition) {
	if qc == nil {
		return
	}
	totalResults := 0
	var wc WordCloud
	wc.InitParams()

	for _, v := range p.ID2Poems {
		for pos, sentence := range v.Sentences {
			if qc.ZhLen > 0 {
				if qc.ZhLen != ChcharLen(sentence) {
					continue
				}
			}

			founded := false
			switch qc.Pos {
			case POS_PREFIX:
				if strings.HasPrefix(sentence, qc.KeywordStr) {
					founded = true
				}
			case POS_SUFFIX:
				if strings.HasSuffix(sentence, qc.KeywordStr) {
					founded = true
				}
			case POS_MID:
				{
					zhlen := ChcharLen(sentence)
					if (zhlen % 2) == 1 {
						midchar := SubChineseString(sentence, zhlen/2, 1)
						if midchar == qc.KeywordStr {
							founded = true
						}
					}
				}
			case POS_ANY:
				if strings.Contains(sentence, qc.KeywordStr) {
					founded = true
				}
			default:

			}

			if founded {
				fmt.Printf("%s [%s]: %s\n", sentence, v.title(), v.FindContext(pos))
				lastchar := GetLastZhChar(sentence, 1)
				wc.AddWord(lastchar)

				totalResults++
			}
		}
	}

	fmt.Printf("Total %d results.\n", totalResults)
	wc.PrintResult(RhymeTop) // top chars. eg: ["夢": 30, "鏡":20]
}

// `黃葉` ==> 西風:15, 何處:7, 淒涼:5, 獨倚:4
// count all poems with specified keyword. eg. 靜夜思 (true) for "明月"
func (p *ChinesePoems) CountPoemChars(qc *QueryCondition) {
	totalResults := 0
	var c2c ZhCharCount
	c2c.Init()

	for _, v := range p.ID2Poems {
		for _, sentence := range v.Sentences {
			if qc.ZhLen > 0 {
				if qc.ZhLen != ChcharLen(sentence) {
					continue
				}
			}

			founded := false
			switch qc.Pos {
			case POS_PREFIX:
				if strings.HasPrefix(sentence, qc.KeywordStr) {
					founded = true
				}
			case POS_SUFFIX:
				if strings.HasSuffix(sentence, qc.KeywordStr) {
					founded = true
				}
			case POS_ANY:
				if strings.Contains(sentence, qc.KeywordStr) {
					founded = true
				}
			default:

			}

			if founded {
				//fmt.Printf("%s [%s]: %s\n", sentence, v.title(), v.FindContext(pos))
				c2c.AddPoem(v)
				totalResults++
			}
		}
	}

	fmt.Printf("Total %d results.\n", totalResults)
	c2c.r2c.PrintSorted()
}

// eg: 昨夜 星辰 昨夜 风
func (p *ChinesePoems) FindRepeatWords() {
	totalResults := 0
	for id, v := range p.ID2Poems {
		if v.hasRepeatWords() {
			fmt.Printf("[%s][%s][%s]\n", id, v.Title, v.toFullDesc())
			totalResults++
		}
	}
	fmt.Printf("Total %d results.\n", totalResults)
}

// eg: ["夢": 30, "鏡":20, ...]
const RhymeTop = 50

// see: ZhRhymes
// chlen: 0 means any
func (p *ChinesePoems) FindByYayunLengthPingze(yayun string, chlen, pztype int) {
	totalResults := 0
	var wc WordCloud
	wc.InitParams()

	for _, v := range p.ID2Poems {
		arr := v.FindByYayunLengthPingze(yayun, chlen, pztype)

		for _, item := range arr {
			fmt.Printf("[%s]|[%s]\n", item, v.toDesc())
			lastchar := GetLastZhChar(item, 1)
			wc.AddWord(lastchar)
			totalResults++
		}
	}
	fmt.Printf("Total %d results.\n", totalResults)
	wc.PrintResult(RhymeTop)
}

func (p *ChinesePoems) PrintRhyme() {
	for _, v := range p.ID2Poems {
		fmt.Printf("[%s]: %s\n", v.Rhyme, v.LeftChars(50))
	}
}

func (p *ChinesePoems) FindByCiPai(cipai string) {
	cipai2count := make(map[string]int)

	resultcount := 0
	for _, v := range p.ID2Poems {
		cipai2count[v.Title] = cipai2count[v.Title] + 1

		if v.Title == cipai {
			fmt.Printf("[%s]: %s\n", v.title(), v.LeftChars(75))
			resultcount++
		}
	}

	//PrintSortedMapByValue(cipai2count)

	fmt.Printf("Total %d results.\n", resultcount)
}

func (p *ChinesePoems) FindByYayun(yayun string) {
	resultcount := 0
	for _, v := range p.ID2Poems {
		if v.Rhyme == yayun {
			fmt.Printf("[%s]: %s\n", v.title(), v.LeftChars(50))
			resultcount++
		}
	}
	fmt.Printf("Total %d results.\n", resultcount)
}

func (p *ChinesePoems) FindByCiPaiYayun(cipai, yayun string) {
	//rhyme2c := make(map[string]int)

	resultcount := 0
	for _, v := range p.ID2Poems {
		//rhyme2c[v.Rhyme] = rhyme2c[v.Rhyme] + 1

		if (v.Rhyme == yayun) && (v.Title == cipai) {
			//fmt.Printf("[%s]: %s\n", v.toDesc(), v.LeftChars(75))
			fmt.Printf("[%s]\n", v.toFullDesc())
			resultcount++
		}
	}
	fmt.Printf("Total %d results.\n", resultcount)

	//PrintSortedMapByValue(rhyme2c)
}

func (p *ChinesePoems) FindByYayunLength(yayun string, chlen int) {
	for _, v := range p.ID2Poems {
		arr := v.FindByYayunLength(yayun, chlen)

		for id, item := range arr {
			fmt.Printf("[%d][%s][%s]\n", id, v.Title, item)
		}
	}
}
