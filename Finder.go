package main

import (
	"fmt"
)

// Keys: [端午 客] ==> 拜跪題封向端午,漂零已是滄浪客|杜甫|惜別行
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
	g_ZhRhymes.Init()
	//var qc QscConv
	//qc.Init()
	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

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
	//qc.allPoems.FindByCiPai("鷓鴣天")

	//qc.countCiPai()

	//qc.FindByYayun("12") // ou
	//qc.FindByYayun("10") // ou
	//qc.FindByYayunLength("12", 4)
	//qc.FindByYayunLengthPingze(Yun_Ing, 0, PingZePing)
	//qc.FindByYayunLengthPingze(Yun_Ei, 0, PingZeZe)
	//qc.FindByYayunLength("8", 7)
	//qc.allPoems.FindByYayunLength("8", 7)
	//qc.allPoems.FindByYayunLengthPingze(Yun_Ing, 7, PingZePing)
	//qc.allPoems.FindByYayunLengthPingze(Yun_Ing, 2, PingZeZe)
	//qc.FindByYayunLengthPingze(Yun_Ou3, 4, PingZeZe)

	//qc.FindByYayun("8")
	//qc.FindByCiPaiYayun("临江仙", "15")
	//qc.FindByCiPaiYayun("青玉案", Yun_Ou3)

	//qc.FindByCiPaiYayun("醉花阴", Yun_U)
	//qc.FindByCiPaiYayun("诉衷情", Yun_Ong)
	//qc.allPoems.FindByCiPaiYayun("鷓鴣天", Yun_Ong)

	//qc.allPoems.FindSentense(createQuery(keyword, pattern, length))
	//qc.allPoems.FindSentense(createQuery("解語", POS_ANY, 0))
	qc.allPoems.FindSentense(createQuery("聲", POS_SUFFIX, 7))

	//qc.allPoems.FindSentense(createQuery("酒", POS_ANY, 7))

	//qc.FindSentense(createQuery("风", POS_SUFFIX, 4))
	//qc.allPoems.FindSentense(createQuery("處", POS_SUFFIX, 7))
	//qc.FindSentense(createQuery("翠", POS_SUFFIX, 0))
	//qc.FindSentense(createQuery("桐", POS_SUFFIX, 0))
	//qc.FindSentense(createQuery("柳", POS_SUFFIX, 0))
	//qc.allPoems.FindSentense(createQuery("暖", POS_ANY, 4))
	//qc.FindSentense(createQuery("染", POS_ANY, 0))
	//qc.FindSentense(createQuery("自由", POS_ANY, 0))
}

func findQscDemo() {
	g_ZhRhymes.Init()

	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

	//qc.allPoems.FindByYayunLengthPingze(Yun_Ing, 3, PingZePing)
	//qc.allPoems.FindByYayunLength(Yun_Ou3, 4)
	//qc.FindByYayunLength(Yun_Ou3, 4)
	//qc.allPoems.FindByYayunLengthPingze(Yun_Ing, 7, PingZePing)
}

func findQscCipai(cipai, yayun string) {
	g_ZhRhymes.Init()
	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

	qc.allPoems.FindByCiPaiYayun(cipai, yayun)
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

// Keys: [新綠]==>新綠:87, 東風:12, 闌幹:9, 年華:7
func findRelated(keyword string, verbmode int) {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	if verbmode != 0 {
		fmt.Println("Finding in all poems in Tang Dynasty...")
	}

	qtsInst.FindRelatedWords(keyword)

	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

	if verbmode != 0 {
		fmt.Println("Finding in all poems in Song Dynasty...")
	}

	qc.allPoems.FindRelatedWords(keyword)
}

// Same to findRelated(), only in QSC
// `黃葉` ==> 西風:15, 何處:7, 淒涼:5, 獨倚:4
func CountPoemZhChars(keyword string) {
	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

	qc.allPoems.CountPoemChars(createQuery(keyword, POS_ANY, 0))
}

// eg: 昨夜 星辰 昨夜 风
func findAllRepeatWordInSentences() {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	qtsInst.FindRepeatWords()

	var qc QscZhtLoader
	qc.loadFile(`qsc_zht_fmt.txt`)

	qc.allPoems.FindRepeatWords()
}

// 浣溪沙:774, 水調歌頭:750
// No instance for: 倚樓人, 花幕暗
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
