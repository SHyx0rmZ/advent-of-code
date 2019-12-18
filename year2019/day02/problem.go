package day02

import (
	"io"
	"strconv"

	aoc "github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
)

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	ns, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	ns[1] = 12
	ns[2] = 2
	ns.RunInPlace(nil, make(chan int))
	return strconv.Itoa(ns[0]), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	ns, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	bu := make(intcode.Program, len(ns))
	copy(bu, ns)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(ns, bu)
			ns[1] = noun
			ns[2] = verb
			ns.RunInPlace(nil, make(chan int))
			if ns[0] == 19690720 {
				return strconv.Itoa(noun*100 + verb), nil
			}
		}
	}
	panic("no solution")
}
