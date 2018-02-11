package main

import (
	"testing"
)

func TestPinyin1(t *testing.T) {
	s := `壽丘惟舊跡，酆邑乃前基。`

	var pyf PinyinFinder
	pyf.Init("/Users/anyu/goproj/Xindict/zht2py.txt")
	//fmt.Printf("Total Hanzi: %d\n", len(pyf.hz2pinyin))

	pinyinstr := ""
	//for pos := 0; pos < len(s); pos++ {
	for _, ch := range s {
		//ch := SubString(s, pos, 1)
		//fmt.Printf("To find: %v\n", ch)
		res := pyf.FindPinyin(string(ch))

		if len(res) > 0 {
			pinyinstr += res + " "
		} else {
			pinyinstr += string(ch)
		}
	}

	if pinyinstr != "shou4 qiu1 wei2 jiu4 ji4 ，Feng1 yi4 nai3 qian2 ji1 。" {
		t.Errorf("TestPinyin1 failed: %v, original: %s", pinyinstr, s)
	}
}
