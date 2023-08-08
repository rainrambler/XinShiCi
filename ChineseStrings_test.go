package main

import (
	"testing"
)

func TestFindFirstStrangeEncoding(t *testing.T) {
	s := `夜凉沈水绣帘栊。酒香浓。雾濛濛。钗列吴娃，EC72袅带金虫。三十六宫蟾观冷，留不住，佩丁东。 `

	arr := FindFirstStrangeEncoding(s)

	if len(arr) != 1 {
		t.Errorf("TestFindFirstStrangeEncoding failed: %v", len(arr))
	}

	if arr[0] != "EC72" {
		t.Errorf("TestFindFirstStrangeEncoding failed: %v", arr[0])
	}
}

func TestFindFirstStrangeEncoding2(t *testing.T) {
	s := `夜凉沈水绣帘栊。酒香浓。雾濛濛。钗列吴娃，EC72`

	arr := FindFirstStrangeEncoding(s)

	if len(arr) != 1 {
		t.Errorf("TestFindFirstStrangeEncoding2 failed: %v", len(arr))
	}

	if arr[0] != "EC72" {
		t.Errorf("TestFindFirstStrangeEncoding2 failed: %v", arr[0])
	}
}

func TestFindFirstStrangeEncoding3(t *testing.T) {
	s := `爱小园、蜕箨E059B552碧`

	arr := FindFirstStrangeEncoding(s)

	if len(arr) != 2 {
		t.Errorf("TestFindFirstStrangeEncoding3 failed: %v", len(arr))
	}

	if arr[0] != "E059" {
		t.Errorf("TestFindFirstStrangeEncoding3 failed: %v", arr[0])
	}

	if arr[1] != "B552" {
		t.Errorf("TestFindFirstStrangeEncoding3 failed: %v", arr[0])
	}
}

func TestFindFirstStrangeEncoding1(t *testing.T) {
	s := `夜凉沈水绣帘栊。酒香浓。雾濛濛。钗列吴娃，EC72袅带金虫。BC39三十六宫蟾观冷，留不住，佩丁东。 `

	arr := FindFirstStrangeEncoding(s)

	if len(arr) != 2 {
		t.Errorf("TestFindFirstStrangeEncoding1 failed: %v", len(arr))
	}

	if arr[0] != "EC72" {
		t.Errorf("TestFindFirstStrangeEncoding1 failed: %v", arr[0])
	}

	if arr[1] != "BC39" {
		t.Errorf("TestFindFirstStrangeEncoding1 failed: %v", arr[1])
	}
}

func TestContainsChPunctions1(t *testing.T) {
	s := `城上風光鶯語亂。城下煙波春拍岸。`

	res := ContainsChPunctions(s)
	expected := true

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestContainsChPunctions2(t *testing.T) {
	s := `踏莎行 `

	res := ContainsChPunctions(s)
	expected := false

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestSplitZhString1(t *testing.T) {
	s := `失調名（般涉）`

	l, _ := SplitZhString(s, '（')
	res := l
	expected := `失調名`

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestSplitZhString2(t *testing.T) {
	s := `失調名_2`

	l, r := SplitZhString(s, '_')
	res := l
	expected := `失調名`

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}

	res = r
	expected = `2`

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestSplitZhString3(t *testing.T) {
	s := `竹枝_2`

	l, r := SplitZhString(s, '_')
	res := l
	expected := `竹枝`

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}

	res = r
	expected = `2`

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestOnlyContains1(t *testing.T) {
	tofind := `其一二三四五六七八九十百`
	s := `其一百三十四`

	res := OnlyContains(s, tofind)
	expected := true

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestOnlyContains2(t *testing.T) {
	tofind := `其一二三四五六七八九十百`
	s := `第三十`

	res := OnlyContains(s, tofind)
	expected := false

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
