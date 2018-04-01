package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestHasCipai2(t *testing.T) {
	s := `西江月（平山堂）`

	var allCipai Cipais

	allCipai.Init("CiPai.txt")

	if hascipai, retval := allCipai.HasCipai2(s); !hascipai {
		t.Errorf("TestHasCipai failed: %v, parsed: %s", s, retval)
	}

	s = `临江仙（冬日即事）`
	if hascipai, _ := allCipai.HasCipai2(s); !hascipai {
		t.Errorf("TestHasCipai failed: %v", s)
	}

	s = `戚氏（此词始终指意，言周穆王宾于西王母事）`
	if hascipai, _ := allCipai.HasCipai2(s); !hascipai {
		t.Errorf("TestHasCipai failed: %v", s)
	}

	s = `卜算子`
	if hascipai, _ := allCipai.HasCipai2(s); !hascipai {
		t.Errorf("TestHasCipai failed: %v", s)
	}

	s = `春事阑珊芳草歇。`
	if hascipai, _ := allCipai.HasCipai2(s); hascipai {
		t.Errorf("TestHasCipai failed: %v", s)
	}
}

func TestIsPossibleCipai(t *testing.T) {
	var qc QscConv
	qc.Init()

	qc.allCipais.Init("CiPai.txt")
	qc.allpoets.Init("SongPoets.txt")

	s := `西江月（平山堂）`

	if !qc.isPossibleCipai(s) {
		t.Errorf("TestIsPossibleCipai failed: %v", s)
	}

	s = `临江仙（冬日即事）`
	if !qc.isPossibleCipai(s) {
		t.Errorf("TestIsPossibleCipai failed: %v", s)
	}

	s = `临江仙 冬日即事`
	if !qc.isPossibleCipai(s) {
		t.Errorf("TestIsPossibleCipai failed: %v", s)
	}

	s = `戚氏（此词始终指意，言周穆王宾于西王母事）`
	if !qc.isPossibleCipai(s) {
		t.Errorf("TestIsPossibleCipai failed: %v", s)
	}

	s = `卜算子`
	if !qc.isPossibleCipai(s) {
		t.Errorf("TestIsPossibleCipai failed: %v", s)
	}

	s = `春事阑珊芳草歇。`
	if qc.isPossibleCipai(s) {
		t.Errorf("TestIsPossibleCipai failed: %v", s)
	}

	s = ` `
	if qc.isPossibleCipai(s) {
		t.Errorf("TestIsPossibleCipai failed: %v", s)
	}

	s = `桃园忆故人`
	if !qc.isPossibleCipai(s) {
		t.Errorf("TestIsPossibleCipai failed: %v", s)
	}
}

func TestConvertLines(t *testing.T) {
	s := `
	
	苏轼
	临江仙
	123, 234
	345, 456
	卜算子
	123, 234
	345, 456
	
	晏几道
	西江月
	123, 234
	345, 456
	
	`

	lines := strings.Split(s, "\n")
	var qc QscConv
	qc.Init()

	qc.convertLines(lines)

	for _, poem := range qc.allPoems.ID2Poems {
		fmt.Println(poem.toDesc())
	}

	actCount := qc.allPoems.Count()
	if actCount != 3 {
		t.Errorf("TestConvertLines failed: %v", actCount)
	}
}

func TestConvertLines2(t *testing.T) {
	s := `
	
苏轼
满庭芳
香叆雕盘，寒生冰箸，画堂别是风光。主人情重，开宴出红妆。腻玉圆搓素颈，藕丝嫩、新织仙裳。双歌罢，虚檐转月，余韵尚悠扬。 
人间，何处有，司空见惯，应谓寻常。坐中有狂客，恼乱愁肠。报道金钗坠也，十指露、春笋纤长。亲曾见，全胜宋玉，想像赋高唐。 
满庭芳
蜗角虚名，蝇头微利，算来著甚干忙。事皆前定，谁弱又谁强。且趁闲身未老，尽放我、些子疏狂。百年里，浑教是醉，三万六千场。 
思量。能几许，忧愁风雨，一半相妨，又何须，抵死说短论长。幸对清风皓月，苔茵展、云幕高张。江南好，千钟美酒，一曲满庭芳。 

	`

	lines := strings.Split(s, "\n")

	var qc QscConv
	qc.Init()

	qc.convertLines(lines)

	actCount := qc.allPoems.Count()
	if actCount != 2 {
		t.Errorf("TestConvertLines2 failed: %v", actCount)
	}

	ids := qc.allPoems.GetAllIDs()

	poem := qc.allPoems.GetPoem(ids[0])

	sentencesize := len(poem.Sentences)

	if sentencesize != 23 {
		t.Errorf("TestConvertLines2 failed: want 23, actual %v", sentencesize)
	}

	for _, sentence := range poem.Sentences {
		fmt.Printf("[%s]\n", sentence)
	}
}

func TestConvertLines3(t *testing.T) {
	s := `
	
姜个翁
霓裳中序第一（春晚旅寓）
园林罢组织。树树东风翠云滴。草满旧家行迹。
龟石。当年第一。也似老、人间风日。

	`

	lines := strings.Split(s, "\n")

	var qc QscConv
	qc.Init()

	qc.convertLines(lines)

	actCount := qc.allPoems.Count()
	if actCount != 1 {
		t.Errorf("TestConvertLines3 failed: %v", actCount)
	}

	for _, poem := range qc.allPoems.ID2Poems {
		fmt.Println(poem.toDesc())
	}

	ids := qc.allPoems.GetAllIDs()

	poem := qc.allPoems.GetPoem(ids[0])

	sentencesize := len(poem.Sentences)

	if sentencesize != 7 {
		t.Errorf("TestConvertLines3 failed: want 7, actual %v", sentencesize)
	}

	for _, sentence := range poem.Sentences {
		fmt.Printf("[%s]\n", sentence)
	}
}

func testUnicode2(t *testing.T) {
	var r rune
	r = 0xb75e

	s := string(r)

	if s != "3" {
		t.Errorf(" failed: %v, want: 3", s)
	}
}
