package day08

import "strconv"

var (
	regs  = make(map[string]int)
	conds = map[string]func(string, int) bool{
		">": func(r string, n int) bool {
			return regs[r] > n
		},
		"<": func(r string, n int) bool {
			return regs[r] < n
		},
		">=": func(r string, n int) bool {
			return regs[r] >= n
		},
		"<=": func(r string, n int) bool {
			return regs[r] <= n
		},
		"==": func(r string, n int) bool {
			return regs[r] == n
		},
		"!=": func(r string, n int) bool {
			return regs[r] != n
		},
	}

	ops = map[string]func(string, int){
		"inc": func(r string, n int) {
			regs[r] += n
		},
		"dec": func(r string, n int) {
			regs[r] -= n
		},
	}
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	commands, err := p.parse(data)
	if err != nil {
		return "", err
	}
	for _, c := range commands {
		if conds[c.Cond.Op](c.Cond.Reg, c.Cond.N) {
			ops[c.Do.Op](c.Do.Reg, c.Do.N)
		}
	}
	var l int
	for _, n := range regs {
		if n > l {
			l = n
		}
	}
	return strconv.Itoa(l), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	commands, err := p.parse(data)
	if err != nil {
		return "", err
	}
	var l int
	for _, c := range commands {
		if conds[c.Cond.Op](c.Cond.Reg, c.Cond.N) {
			ops[c.Do.Op](c.Do.Reg, c.Do.N)
		}
		for _, n := range regs {
			if n > l {
				l = n
			}
		}
	}
	return strconv.Itoa(l), nil
}

func (problem) parse(data []byte) ([]command, error) {
	var s int
	var e int
	var err error

	var commands []command

	for s < len(data)-1 {
		o := command{}

		for data[e] != ' ' {
			e++
		}
		o.Do.Reg = string(data[s:e])
		e++
		s = e
		for data[e] != ' ' {
			e++
		}
		o.Do.Op = string(data[s:e])
		e++
		s = e
		for data[e] != ' ' {
			e++
		}
		o.Do.N, err = strconv.Atoi(string(data[s:e]))
		if err != nil {
			return nil, err
		}
		e++
		for data[e] != ' ' {
			e++
		}
		e++
		s = e
		for data[e] != ' ' {
			e++
		}
		o.Cond.Reg = string(data[s:e])
		e++
		s = e
		for data[e] != ' ' {
			e++
		}
		o.Cond.Op = string(data[s:e])
		e++
		s = e
		for data[e] != '\n' {
			e++
		}
		o.Cond.N, err = strconv.Atoi(string(data[s:e]))
		e++
		s = e
		commands = append(commands, o)
	}
	return commands, nil
}

type group struct {
	Reg string
	Op  string
	N   int
}

type command struct {
	Do   group
	Cond group
}
