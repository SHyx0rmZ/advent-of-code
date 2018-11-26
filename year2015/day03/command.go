package day03

import (
	"bufio"
	"fmt"
	"io"

	"github.com/SHyx0rmZ/advent-of-code"
)

type point struct {
	X, Y int
}

type problem struct {
	houses map[point]struct{}
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{
		houses: make(map[point]struct{}),
	}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	pt := point{0, 0}
	p.houses[pt] = struct{}{}
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		switch s.Text() {
		case "^":
			pt.Y--
		case "v":
			pt.Y++
		case "<":
			pt.X--
		case ">":
			pt.X++
		}
		p.houses[pt] = struct{}{}
	}
	return fmt.Sprintf("%d", len(p.houses)), s.Err()
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	pts := [...]point{
		{}, // santa
		{}, // robot
	}
	var i int
	var pt = &pts[i]
	p.houses[point{}] = struct{}{}
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		switch s.Text() {
		case "^":
			pt.Y--
		case "v":
			pt.Y++
		case "<":
			pt.X--
		case ">":
			pt.X++
		}
		p.houses[*pt] = struct{}{}
		i ^= 1
		pt = &pts[i]
	}
	return fmt.Sprintf("%d", len(p.houses)), s.Err()
}
