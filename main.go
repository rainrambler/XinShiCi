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

		if cw.IsSingle2() {
			//fmt.Println(cw.toDesc())
			singlecount++
		}
	}

	fmt.Printf("Hanzi Count: %d\n", singlecount)
}

func analyseQsc() {
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
	//qc.FindByCiPai(`阮郎归`)
	//qc.FindByCiPai(`醉桃源`)
	//qc.FindByCiPai(`醉花阴`)
	//qc.FindByCiPai("采桑子")

	//qc.countCiPai()

	//qc.FindByYayun("12") // ou
	//qc.FindByYayun("10") // ou
	//qc.FindByYayunLength("12", 4)
	//qc.FindByYayunLengthPingze(Yun_Ing, 0, PingZePing)
	//qc.FindByYayunLengthPingze(Yun_Ei, 0, PingZeZe)
	//qc.FindByYayunLength("8", 7)
	//qc.FindByYayunLengthPingze(Yun_U, 7, PingZeZe)
	//qc.FindByYayunLengthPingze(Yun_Ou3, 4, PingZeZe)

	//qc.FindByYayun("8")
	//qc.FindByCiPaiYayun("临江仙", "15")
	//qc.FindByCiPaiYayun("青玉案", Yun_Ou3)

	//qc.FindByCiPaiYayun("醉花阴", Yun_U)
	//qc.FindByCiPaiYayun("采桑子", Yun_Ong)

	//qc.FindSentense(createQuery("笙歌", POS_ANY, 0))
	//qc.FindSentense(createQuery("媚", POS_ANY, 0))
	//qc.FindSentense(createQuery("美", POS_SUFFIX, 0))
	//qc.FindSentense(createQuery("屏", POS_SUFFIX, 7))
	//qc.FindSentense(createQuery("秀", POS_SUFFIX, 4))
	//qc.FindSentense(createQuery("酒", POS_SUFFIX, 7))
	//qc.FindSentense(createQuery("翠", POS_SUFFIX, 0))
	//qc.FindSentense(createQuery("桐", POS_SUFFIX, 0))
	//qc.FindSentense(createQuery("柳", POS_SUFFIX, 0))
	//qc.FindSentense(createQuery("暖", POS_ANY, 4))
	//qc.FindSentense(createQuery("染", POS_ANY, 0))
	//qc.FindSentense(createQuery("路", POS_SUFFIX, 5))
	//missedChars.DbgPrint()
}

func FindSentence(keyword string, mode int, charcount int) {
}

func main() {
	//readQts()
	//loadPoetryFile(`./quantangshi/005.json`)
	//loadFileDemo(`./quantangshi/005.json`)
	//LoadPoetries(`./quantangshi/`)
	//printHans(`allhans.txt`)
	//analyseQsc()
	//convertHanzi2Pinyin2()

	//cipaiDemo()

	ExportQsc(`D:\Projects\GitHubSync\XinShiCi\qsc.txt`, `D:\tmp\xsc`)
}
