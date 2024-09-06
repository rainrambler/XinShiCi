package main

import (
	"testing"
)

func TestHasCipai2(t *testing.T) {
	s := `西江月（平山堂）`

	var allCipai Cipais

	allCipai.Init("CiPai.txt")

	if hascipai, retval := allCipai.HasCipai(s); !hascipai {
		t.Errorf("TestHasCipai2 failed: %v, parsed: %s", s, retval)
	}

	s = `临江仙（冬日即事）`
	if hascipai, _ := allCipai.HasCipai(s); !hascipai {
		t.Errorf("TestHasCipai2 failed: %v", s)
	}

	s = `戚氏（此词始终指意，言周穆王宾于西王母事）`
	if hascipai, _ := allCipai.HasCipai(s); !hascipai {
		t.Errorf("TestHasCipai2 failed: %v", s)
	}

	s = `卜算子`
	if hascipai, _ := allCipai.HasCipai(s); !hascipai {
		t.Errorf("TestHasCipai2 failed: %v", s)
	}

	s = `春事阑珊芳草歇。`
	if hascipai, _ := allCipai.HasCipai(s); hascipai {
		t.Errorf("TestHasCipai2 failed: %v", s)
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

func testLineFormat2(t *testing.T) {
	s := "庶有瘳乎。|<事见七发>| "
	trimed := lineFormat(s)
	wanted := "庶有瘳乎。"
	if trimed != wanted {
		t.Errorf(" failed: %v, want: %s", trimed, wanted)
	}
}

func testLineFormat1(t *testing.T) {
	s := "十巡今止。乐事要须防极喜。|<淳子髡曰：酒极则乱，乐极则悲>|烛影摇风。月落参横影子通。"
	trimed := lineFormat(s)
	wanted := "十巡今止。乐事要须防极喜。烛影摇风。月落参横影子通。"
	if trimed != wanted {
		t.Errorf(" failed: %v, want: %s", trimed, wanted)
	}
}
