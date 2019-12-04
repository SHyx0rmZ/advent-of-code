package day25

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

type key struct {
	X, Y int
}

func diagonals(n int) map[key]int {
	m := make(map[key]int)
	x := 20151125
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			m[key{i - j, j + 1}] = x
			x = (x * 252533) % 33554393
		}
	}
	return m
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	x := diagonals(6000)[key{2947, 3029}]
	return strconv.Itoa(x), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	return "", nil
}
