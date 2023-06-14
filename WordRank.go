package main

import (
	"fmt"
	"sort"
	"strings"
)

const DesiredLen = 1024

type WordRank struct {
	char2count map[rune]int
}

func (p *WordRank) Init() {
	p.char2count = make(map[rune]int)
}

func (p *WordRank) LoadFile(fileName string) {
	s, err := ReadTextFile(fileName)
	if err != nil {
		fmt.Printf("Cannot read file: %s: %v!\n", fileName, err)
		return
	}

	rs := []rune(s)
	for _, r := range rs {
		p.AddChar(r)
	}
}

func (p *WordRank) AddChar(r rune) {
	if IsInvalidChar(r) {
		return
	}
	p.char2count[r] = p.char2count[r] + 1
}

func (p *WordRank) PrintResult() {
	//PrintSortedMapByValue(p.char2count)
	type kv struct {
		Key   rune
		Value int
	}

	var ss []kv
	for k, v := range p.char2count {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	arr := []rune{}

	for _, kv := range ss {
		//fmt.Printf("%s:%d\n", kv.Key, kv.Value)
		arr = append(arr, kv.Key)

		if len(arr) > DesiredLen {
			break
		}
	}

	s := string(arr)
	fmt.Println(s)
}

func IsInvalidChar(r rune) bool {
	s := `；，。！？、（）《》·：` + `＿□` + ` ,.:;?()#` + `1234567890abcdefghijklmnopqrstuvwxyz`
	return strings.ContainsRune(s, r)
}

func qtsRank() {
	filename := `D:\Projects\GitHubSync\XinShiCi\qsc.txt`
	var wr WordRank
	wr.Init()
	wr.LoadFile(filename)

	wr.PrintResult()
}
