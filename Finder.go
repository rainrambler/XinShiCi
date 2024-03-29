package main

import (
	"fmt"
)

func findKeywords(keywords string) {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	cp := qtsInst.FindKeywords(keywords)
	cp.PrintResults()

	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)
	//qc.loadFile(`TangSongCiZh.txt`)

	cp = qc.allPoems.FindKeywords(keywords)
	cp.PrintResults()
}

func findQscKeyword(keyword string, pattern, length int, verbmode int) {
	//var qc QscConv
	//qc.Init()
	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

	g_ZhRhymes.Init()

	//qc.analyseCipai("qsc.txt")
	//qc.convertFile("qsc.txt")

	//qc.FindRepeatWords()
	//qc.FindRepeatDiffs()

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
	//qc.FindByCiPaiYayun("诉衷情", Yun_Ong)

	qc.allPoems.FindSentense(createQuery(keyword, pattern, length))
	//qc.FindSentense(createQuery("媚", POS_ANY, 0))
	//qc.FindSentense(createQuery("美", POS_SUFFIX, 0))
	//qc.FindSentense(createQuery("屏", POS_SUFFIX, 7))
	//qc.FindSentense(createQuery("风", POS_SUFFIX, 4))
	//qc.FindSentense(createQuery("酒", POS_SUFFIX, 7))
	//qc.FindSentense(createQuery("翠", POS_SUFFIX, 0))
	//qc.FindSentense(createQuery("桐", POS_SUFFIX, 0))
	//qc.FindSentense(createQuery("柳", POS_SUFFIX, 0))
	//qc.FindSentense(createQuery("暖", POS_ANY, 4))
	//qc.FindSentense(createQuery("染", POS_ANY, 0))
	//qc.FindSentense(createQuery("自由", POS_ANY, 0))
	//missedChars.DbgPrint()
}

func findRepeatChChars() {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	qtsInst.FindRepeatDiffs("repdiff_qts.txt")

	//var qc QscConv
	//qc.Init()
	//qc.convertFile("qsc.txt")
	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

	qc.allPoems.FindRepeatDiffs("repdiff_qsc.txt")
}

func findRelated(keyword string, verbmode int) {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	qtsInst.FindRelatedWords(keyword)

	//var qc QscConv
	//qc.Init()
	//qc.convertFile("qsc.txt")

	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

	qc.allPoems.FindRelatedWords(keyword)
}

// `黃葉` ==> 西風:15, 何處:7, 淒涼:5, 獨倚:4
func CountPoemZhChars(keyword string) {
	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

	qc.allPoems.CountPoemChars(createQuery(keyword, POS_ANY, 0))
}
func AnalyseCipai() {
	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

	k2count := make(map[string]int)
	for _, v := range qc.allPoems.ID2Poems {
		k2count[v.Title] += 1
	}

	PrintSortedMapByValue(k2count)

	fmt.Println("------------------------")

	for k, _ := range qc.allCipais.item2id {
		_, exists := k2count[k]

		if !exists {
			fmt.Printf("No instance for: %s\n", k)
		}
	}
}
