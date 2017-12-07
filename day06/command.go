package day06

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code/input"
)

func Command() error {
	if len(os.Args) < 4 {
		panic("not enough arguments")
	}

	c, err := input.ReadInput(os.Args[3])
	if err != nil {
		return err
	}

	var banks []int
	var x int

	for len(c) > 0 {
		i := bytes.IndexAny(c, "\t \n")
		if i < 1 {
			// if there are no bytes before the whitespace,
			// we must have reached the end of the line
			break
		}
		// parse column
		b, err := strconv.Atoi(string(c[0:i]))
		if err != nil {
			return err
		}
		banks = append(banks, b)
		for c[i] == '\t' || c[i] == ' ' {
			i++
		}
		c = c[i:]
		x++
	}

	steps := BalanceState(banks)

	switch os.Args[2] {
	case "steps":
		_, err = fmt.Printf("%d\n", steps)
	default:
		panic("unknown sub-command: " + os.Args[2])
	}

	return err
}

func BalanceState(banks []int) int {
	m := &memory{}
	m.Reload(banks)
	var steps int
	s := set{
		M: make(map[string]struct{}),
	}
	for {
		steps++
		m.Balance()
		if s.Contains(m.String()) {
			break
		}
		s.Add(m.String())
	}
	return steps
}
