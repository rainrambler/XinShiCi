package main

type PoemWords struct {
	poet2id map[string]int
}

func (p *PoemWords) InitMaps() {
	p.poet2id = make(map[string]int)
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
	if _, ok := p.poet2id[nm]; ok {
		return true
	}
	return false
}

func (p *PoemWords) Count() int {
	return len(p.poet2id)
}

func (p *PoemWords) AddWord(wd string) {
	if len(wd) == 0 {
		return
	}

	p.poet2id[wd] = 1
}

func (p *PoemWords) Write(filename string) {
	lines := []string{}

	for k, _ := range p.poet2id {
		lines = append(lines, k)
	}

	WriteLines(lines, filename)
}
