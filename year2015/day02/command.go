package day02

import (
	"bufio"
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/pkg/lib"
	"io"
	"sort"
	"strings"
)

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	var i int
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		sides, err := lib.StringsToInts(strings.SplitN(s.Text(), "x", 3))
		if err != nil {
			return "", err
		}
		sort.Ints(sides)
		i += 3 * sides[0] * sides[1] + 2 * sides[0] * sides[2] + 2 * sides[1] * sides[2]
	}
	return fmt.Sprintf("%d", i), s.Err()
}

func (p problem) PartTwoWithReader(r io.Reader)  (string, error) {
	var i int
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		sides, err := lib.StringsToInts(strings.SplitN(s.Text(), "x", 3))
		if err != nil {
			return "", err
		}
		sort.Ints(sides)
		i += 2 * (sides[0] + sides[1]) + sides[0] * sides[1] * sides[2]
	}
	return fmt.Sprintf("%d", i), s.Err()
}
