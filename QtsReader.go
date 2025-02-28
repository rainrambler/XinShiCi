package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func readQts() {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	var pyf PinyinFinder
	pyf.Init("zht2py.txt")

	//qtsInst.findLongNTitle(50) // Top 50
	//qtsInst.findRepeatChar()
	//qtsInst.dbgPrintMaxId()
	//qtsInst.findMalformedText()
	qtsInst.findRepeatWords()
	//qtsInst.exportTitles(`D:\qtstitles.txt`)
	//qtsInst.exportAuthors(`D:\qtsauthors.txt`)
}

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

// Find titles which length larger than maxLength
func (p *Qts) findLongTitle(maxLength int) {
	for _, poem := range p.ID2Poems {
		if ChcharLen(poem.Title) > maxLength {
			fmt.Printf("[%d][%s]: %s\n",
				len(poem.Title),
				poem.Title, poem.LeftChars(20))
		}
	}
}

// Find the longest title (num: number of poems)
func (p *Qts) findLongNTitle(num int) {
	toparr := []*KeyValue{}
	for _, poem := range p.ID2Poems {
		kv := new(KeyValue)
		kv.Key = poem.Title
		kv.Value = len(poem.Title)
		toparr = append(toparr, kv)
	}

	topn := len(toparr)
	if num < topn {
		topn = num
	}

	sort.Slice(toparr, func(i, j int) bool {
		return toparr[i].Value > toparr[j].Value
	})

	for i := 0; i < num; i++ {
		pair := toparr[i]
		fmt.Printf("[%d][%s]\n",
			pair.Value, pair.Key)
	}
}

// Find repeat Chinese chars in sentence (eg: 梨花院落 溶溶 月)
func (p *Qts) findRepeatChar() {
	for _, poem := range p.ID2Poems {

		if poem.hasRepeatChar() {
			fmt.Println(poem.toFullDesc())
		}
	}
}

// Find repeat Chinese words in sentence (eg: 昨夜星辰昨夜风)
func (p *Qts) findRepeatWords() {
	results := 0
	for _, poem := range p.ID2Poems {

		if poem.hasRepeatWords() {
			fmt.Println(poem.toFullDesc())
			results++
		}
	}

	fmt.Printf("Total %d results.\n", results)
}

// Find malformed poems
func (p *Qts) findMalformedText() {
	k2count := make(map[string]int)
	poem2number := make(map[string]int)
	for id, poem := range p.ID2Poems {
		arr, res := findAllErrorText(poem)
		for _, v := range arr {
			k2count[v] = k2count[v] + 1
		}

		if res != "" {
			fmt.Println(res + "|" + poem.toDesc())
		}

		if len(arr) > 0 {
			poem2number[id] = len(arr)
		}
	}
	PrintSortedMapByValue(k2count)
	PrintSortedMapByValue(poem2number)
}

// format: Author|Title
func (p *Qts) exportTitles(filename string) {
	lines := []string{}
	for _, poem := range p.ID2Poems {
		lines = append(lines, poem.Author+"|"+poem.Title)
	}
	WriteLines(lines, filename)
}

// format: each author a line
func (p *Qts) exportAuthors(filename string) {
	lines := []string{}
	authors := make(map[string]int)
	for _, poem := range p.ID2Poems {
		if poem.Author != "" {
			authors[poem.Author] = 1
		}
	}

	for k, _ := range authors {
		lines = append(lines, k)
	}
	WriteLines(lines, filename)
}

type VolumeIdFinder struct {
	volume2maxid map[int]int
}

func (p *VolumeIdFinder) addId(id string) bool {
	arr := strings.Split(id, "_")

	if len(arr) != 2 {
		return false
	}

	volnum, err := strconv.Atoi(arr[0])
	if err != nil {
		fmt.Printf("DBG: Cannot convert volume in %s!\n", id)
		return false
	}

	idnum, err := strconv.Atoi(arr[1])
	if err != nil {
		fmt.Printf("DBG: Cannot convert id in %s!\n", id)
		return false
	}

	curnum, exists := p.volume2maxid[volnum]
	if exists {
		if curnum < idnum {
			p.volume2maxid[volnum] = idnum // find max
		}
	} else {
		p.volume2maxid[volnum] = idnum
	}
	return true
}

func (p *VolumeIdFinder) printResult() {
	for i := 1; i < 904; i++ {
		fmt.Printf("Vol [%d]: [%d]\n", i, p.volume2maxid[i])
	}
}

func (p *Qts) dbgPrintMaxId() {
	var vif VolumeIdFinder
	vif.volume2maxid = make(map[int]int)

	for _, poem := range p.ID2Poems {
		vif.addId(poem.ID)
	}

	vif.printResult()
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
	poem.ID = arr[0]
	poem.Title = arr[1]
	poem.Author = arr[2]
	poem.AllText = arr[3]
	poem.LineNumber = idx

	poem.ParseSentences()

	return &poem
}

func convertQts(filename string) {
	allLines, err := ReadLines(filename)
	if err != nil {
		fmt.Printf("WARN: Cannot parse file: %s: %v!\n", filename, err)
		return
	}

	linesnew := []string{}
	id := 100
	for _, line := range allLines {
		if len(line) == 0 {
			continue
		}
		s := fmt.Sprintf("837_%d\t山居詩二十四首\t貫休\t%s", id, line)
		id++

		linesnew = append(linesnew, s)
	}

	WriteLines(linesnew, filename+".txt")
}

func countPoemLength() {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	poem2len := make(map[string]int)
	for _, poem := range qtsInst.ChinesePoems.ID2Poems {
		poem2len[poem.Title] = len(poem.AllText)
	}

	PrintMapByValueTop(poem2len, 50)
}
