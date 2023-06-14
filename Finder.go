package main

func findKeywords(keywords string) {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	qtsInst.FindKeywords(keywords)

	var qc QscConv
	qc.Init()
	qc.convertFile("qsc.txt")

	qc.allPoems.FindKeywords(keywords)
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
	//qc.FindSentense(createQuery("自由", POS_ANY, 0))
	//missedChars.DbgPrint()

	//qc.CountPoemChars(createQuery("黄叶", POS_ANY, 0))
}

func findRepeatChChars() {
	var qtsInst Qts
	qtsInst.Init()
	qtsInst.ReadFile("qts_zht.txt")

	qtsInst.FindRepeatDiffs("repdiff_qts.txt")

	var qc QscConv
	qc.Init()
	qc.convertFile("qsc.txt")

	qc.allPoems.FindRepeatDiffs("repdiff_qsc.txt")
}
