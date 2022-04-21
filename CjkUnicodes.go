package main

// http://www.unicode.org/versions/Unicode5.0.0/ch12.pdf#G12159
func getAllHans() string {
	s := ""
	var r rune

	// Common
	for r = 0x4e00; r <= 0x9fff; r++ {
		s += string(r)
	}

	// Rare
	for r = 0x3400; r <= 0x4dff; r++ {
		s += string(r)
	}

	// Rare, historic
	for r = 0x20000; r <= 0x2a6df; r++ {
		s += string(r)
	}

	// Duplicates, unifiable variants, corporate characters
	for r = 0xf900; r <= 0xfaff; r++ {
		s += string(r)
	}

	// Unifiable variants
	for r = 0x2f800; r <= 0x2fa1f; r++ {
		s += string(r)
	}

	// https://stackoverflow.com/questions/1366068/whats-the-complete-range-for-chinese-characters-in-unicode
	// Rare, historic
	for r = 0x2A700; r <= 0x2B73F; r++ {
		s += string(r)
	}

	// Uncommon, some in current use
	for r = 0x2B740; r <= 0x2B81F; r++ {
		s += string(r)
	}

	// Rare, historic
	for r = 0x2B820; r <= 0x2CEAF; r++ {
		s += string(r)
	}

	return s
}

func printHans(filename string) {
	s := getAllHans()
	WriteTextFile(filename, s)
}
