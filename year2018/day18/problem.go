package day18

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/SHyx0rmZ/advent-of-code"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	g := grid{
		M: [...]map[point]tile{
			make(map[point]tile),
			make(map[point]tile),
		},
	}
	var y int
	s1 := bufio.NewScanner(r)
	for s1.Scan() {
		var x int
		s2 := bufio.NewScanner(strings.NewReader(s1.Text()))
		s2.Split(bufio.ScanRunes)
		for s2.Scan() {
			switch tile(s2.Text()) {
			case ground, trees, lumberyard:
				g.add(point{x, y}, tile(s2.Text()))
			}
			x++
		}
		y++
	}
	for i := 0; i < 10; i++ {
		g.update()
	}
	g.print()
	return strconv.Itoa(g.score()), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	g := grid{
		M: [...]map[point]tile{
			make(map[point]tile),
			make(map[point]tile),
		},
	}
	var y int
	s1 := bufio.NewScanner(r)
	for s1.Scan() {
		var x int
		s2 := bufio.NewScanner(strings.NewReader(s1.Text()))
		s2.Split(bufio.ScanRunes)
		for s2.Scan() {
			switch tile(s2.Text()) {
			case ground, trees, lumberyard:
				g.add(point{x, y}, tile(s2.Text()))
			}
			x++
		}
		y++
	}
	ss := make(map[int]int)
	gs := make(map[int]map[point]tile)
	for i := 0; i < 1000000000; i++ {
		g.update()
		s := g.score()
		if _, ok := ss[s]; ok {
			if g.equal(gs[ss[s]]) {
				fmt.Println(i)
				i += (1000000000 - i) / (i - ss[s]) * (i - ss[s])
				fmt.Println(i)
			}
		}
		ss[s] = i
		gs[i] = g.copy()
	}
	g.print()
	return strconv.Itoa(g.score()), nil
}
