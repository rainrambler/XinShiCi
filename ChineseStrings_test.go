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
