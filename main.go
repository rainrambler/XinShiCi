// Xindict project main.go
package main

import (
	"fmt"
)

func demoText() {
	var cd ChineseDict
	cd.Init()
	ReadDict("/Users/anyu/goproj/Xindict/dictdemo.txt", &cd)

	for _, cw := range cd.chs2Word {
		fmt.Println(cw.toDesc())
	}
}

func readDictFile() {
	var cd ChineseDict
	cd.Init()
	ReadDict("/Users/anyu/goproj/Xindict/cedict.u8", &cd)

	singlecount := 0
	for _, cw := range cd.chs2Word {

		if cw.IsSingle() {
			//fmt.Println(cw.toDesc())
			singlecount++
		}
	}

	fmt.Printf("Hanzi Count: %d\n", singlecount)
}

func readQts() {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("/Users/anyu/goproj/Xindict/qtszht.txt")

	var pyf PinyinFinder
	pyf.Init("/Users/anyu/goproj/Xindict/zht2py.txt")

}

func main() {
	//readQts()
	//convertHanzi2Pinyin()

	var qc QscConv
	qc.Init()

	//missedChars.rhy2Count = make(map[string]int)

	//g_Rhymes.Init()
	//g_Rhymes.ImportFile("ShiYunXinBian.txt")
	g_ZhRhymes.Init()

	//qc.analyseStrangeEncoding("qsc.txt")
	//qc.analyseCipai("qsc.txt")
	qc.convertFile("qsc.txt")

	//qc.PrintRhyme()
	//qc.FindByCiPai("踏莎行")
	qc.FindByYayun("14")
	//qc.FindByYayun("2")
	//qc.FindByCiPaiYayun("临江仙", "15")
	//qc.FindByCiPaiYayun("鹊桥仙", "14")
	//qc.FindSentense(createQuery("艳似", POS_ANY, 0))

	//missedChars.DbgPrint()
}
