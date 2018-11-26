package day03

import (
	"bufio"
	"container/ring"
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code"
	"io"
)

type point struct {
	X, Y int
}

type problem struct{
	houses map[point]struct{}
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{
		houses: make(map[point]struct{}),
	}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	pt := point{0,0}
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

func (p problem) PartTwoWithReader(r io.Reader)  (string, error) {
	pt := ring.New(2)
	for i := 0; i < pt.Len(); i++ {
		pt.Value = &point{}
	}
	p.houses[point{}] = struct{}{}
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		switch s.Text() {
		case "^":
			pt.Value.(*point).Y--
		case "v":
			pt.Value.(*point).Y++
		case "<":
			pt.Value.(*point).X--
		case ">":
			pt.Value.(*point).X++
		}
		p.houses[*pt.Value.(*point)] = struct{}{}
		pt = pt.Next()
	}
	return fmt.Sprintf("%d", len(p.houses)), s.Err()
}
