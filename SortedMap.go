package main

import (
	"fmt"
	"sort"
)

// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
func PrintSortedMapByValue(m map[string]int) {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	fmt.Println("===>")
	for _, kv := range ss {
		fmt.Printf("%s:%d\n", kv.Key, kv.Value)
	}
	fmt.Println("<===")
}

// Print Top N values (sorted by value), -1 means all
func PrintMapByValueTop(m map[string]int, topn int) {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	total := 0
	fmt.Println("===>")
	for _, kv := range ss {
		fmt.Printf("%s:%d\n", kv.Key, kv.Value)
		total++

		if (total >= topn) && (topn > 0) {
			break
		}
	}
	fmt.Println("<===")
}

func PrintSortedMapByValueInt(m map[int]int) {
	type kv struct {
		Key   int
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	fmt.Println("===>")
	for _, kv := range ss {
		fmt.Printf("%d:%d\n", kv.Key, kv.Value)
	}
	fmt.Println("<===")
}

func PrintMapGroupByValue(m map[string]int) {
	v2count := make(map[int]int)
	for _, v := range m {
		v2count[v] = v2count[v] + 1
	}

	PrintSortedMapByValueInt(v2count)
}
