package main

type PoemWords struct {
	word2cnt map[string]int
}

func (p *PoemWords) InitMaps() {
	p.word2cnt = make(map[string]int)
}

func (p *PoemWords) Init(filename string) {
	p.InitMaps()

	lines := ReadTxtFile(filename)

	for _, line := range lines {
		if len(line) > 0 {
			p.AddWord(line)
		}
	}
}

func (p *PoemWords) Contains(nm string) bool {
	if _, ok := p.word2cnt[nm]; ok {
		return true
	}
	return false
}

func (p *PoemWords) Count() int {
	return len(p.word2cnt)
}

func (p *PoemWords) AddWord(wd string) {
	if len(wd) == 0 {
		return
	}

	p.word2cnt[wd] = p.word2cnt[wd] + 1
}

func (p *PoemWords) Write(filename string) {
	lines := []string{}

	for k, _ := range p.word2cnt {
		lines = append(lines, k)
	}

	WriteLines(lines, filename)
}
