package day23

import (
	"io"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	prg, err := parse(r)
	if err != nil {
		return "", err
	}
	c := cpu{Registers: make(map[register]int)}
	c.run(prg)
	return strconv.Itoa(c.Registers["b"]), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	prg, err := parse(r)
	if err != nil {
		return "", err
	}
	c := cpu{Registers: map[register]int{"a": 1}}
	c.run(prg)
	return strconv.Itoa(c.Registers["b"]), nil
}
