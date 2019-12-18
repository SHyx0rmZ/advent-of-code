package day05

import (
	"io"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	ns, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	input <- 1
	output := make(chan int, 1)
	go ns.Run(input, output)
	var tests bool
	var result int
	for {
		select {
		case v, ok := <-output:
			if !ok {
				if !tests {
					panic("no tests")
				}
				return strconv.Itoa(result), nil
			}
			if v == 0 {
				tests = true
			} else {
				if !tests {
					panic("no tests")
				}
				result = v
			}
		}
	}
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	ns, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	input <- 5
	output := make(chan int, 1)
	go ns.Run(input, output)
	return strconv.Itoa(<-output), nil
}
