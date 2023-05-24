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

		if cw.IsSingle2() {
			//fmt.Println(cw.toDesc())
			singlecount++
		}
	}

	fmt.Printf("Hanzi Count: %d\n", singlecount)
}

func main() {
	findKeywords(`夏 茶`)
	//readQts()
	//loadPoetryFile(`./quantangshi/005.json`)
	//loadFileDemo(`./quantangshi/005.json`)
	//LoadPoetries(`./quantangshi/`)
	//printHans(`allhans.txt`)
	//analyseQsc()
	//convertHanzi2Pinyin2()

	//cipaiDemo()

	//ExportPoetry(`D:\tmp\poem\chinese-poetry-master\ci`, `D:\tmp\poem\exportdemo`)

	//ExportQsc(`D:\Projects\GitHubSync\XinShiCi\qsc.txt`, `D:\tmp\xsc`)

	//LoadTxt(`D:\2023\05\Nalan\nalan_all_purified.txt`)
	//LoadTxt(`D:\Projects\XinShiCi\qsc.txt`)
}
