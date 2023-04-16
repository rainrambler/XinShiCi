package main

type Poets struct {
	poet2id map[string]int
}

func (p *Poets) Init(filename string) {
	p.poet2id = make(map[string]int)

	lines := ReadTxtFile(filename)

	id := 100
	for _, line := range lines {
		if len(line) > 0 {
			p.poet2id[line] = id + 2
		}
	}
}

func (p *Poets) IsPoet(nm string) bool {
	if _, ok := p.poet2id[nm]; ok {
		return true
	}
	return false
}

func (p *Poets) FindPoet(nm string) int {
	if id, ok := p.poet2id[nm]; ok {
		return id
	}
	return -1
}

func (p *Poets) Count() int {
	return len(p.poet2id)
}
