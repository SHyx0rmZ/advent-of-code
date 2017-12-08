package day08

import (
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code/input"
	"os"
	"strconv"
)

func Command() error {
	if len(os.Args) < 3 {
		panic("not enough arguments")
	}

	c, err := input.ReadInput(os.Args[2])
	if err != nil {
		return err
	}

	type group struct {
		Register  string
		Operation string
		Amount    int
	}

	type command struct {
		Do        group
		Condition group
	}

	var s int
	var e int

	var commands []command

	for s < len(c)-1 {
		o := command{}

		for c[e] != ' ' {
			e++
		}
		o.Do.Register = string(c[s:e])
		e++
		s = e
		for c[e] != ' ' {
			e++
		}
		o.Do.Operation = string(c[s:e])
		e++
		s = e
		for c[e] != ' ' {
			e++
		}
		o.Do.Amount, err = strconv.Atoi(string(c[s:e]))
		if err != nil {
			return err
		}
		e++
		for c[e] != ' ' {
			e++
		}
		e++
		s = e
		for c[e] != ' ' {
			e++
		}
		o.Condition.Register = string(c[s:e])
		e++
		s = e
		for c[e] != ' ' {
			e++
		}
		o.Condition.Operation = string(c[s:e])
		e++
		s = e
		for c[e] != '\n' {
			e++
		}
		o.Condition.Amount, err = strconv.Atoi(string(c[s:e]))
		e++
		s = e
		commands = append(commands, o)
	}

	regs := make(map[string]int)

	conds := map[string]func(string, int) bool{
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

	ops := map[string]func(string, int){
		"inc": func(r string, n int) {
			regs[r] += n
		},
		"dec": func(r string, n int) {
			regs[r] -= n
		},
	}

	for _, command := range commands {
		if conds[command.Condition.Operation](command.Condition.Register, command.Condition.Amount) {
			ops[command.Do.Operation](command.Do.Register, command.Do.Amount)
		}
	}

	l := 0

	for _, n := range regs {
		if n > l {
			l = n
		}
	}

	_, err = fmt.Printf("%d\n", l)

	//switch os.Args[4] {
	//case "steps":
	//	_, err = fmt.Printf("%d\n", steps)
	//case "loopsize":
	//	_, err = fmt.Printf("%d\n", length)
	//default:
	//	panic("unknown sub-command: " + os.Args[2])
	//}

	return err
}
