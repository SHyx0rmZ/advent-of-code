package day09

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
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	output := make(chan int, 1)
	input <- 1
	go prg.Run(input, output)
	return strconv.Itoa(<-output), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	output := make(chan int, 1)
	input <- 2
	go prg.Run(input, output)
	return strconv.Itoa(<-output), nil
}
