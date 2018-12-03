package day01

import (
	"bufio"
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code"
	"io"
)

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	var i int
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		switch s.Text() {
		case "(":
			i++
		case ")":
			i--
		default:
			panic("unexpected")
		}
	}
	return fmt.Sprintf("%d", i), s.Err()
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	var i int
	var j int
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		j++
		switch s.Text() {
		case "(":
			i++
		case ")":
			i--
			if i < 0 {
				return fmt.Sprintf("%d", j), s.Err()
			}
		default:
			panic("unexpected")
		}
	}
	panic("unexpected")
}
