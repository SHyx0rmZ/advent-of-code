package day03

import (
	"bytes"
	"io"
	"io/ioutil"
	"math"
	"strconv"

	aoc "github.com/SHyx0rmZ/advent-of-code"
)

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

type move struct {
	dir
	n int
}

func (m move) from(x, y int) (int, int) {
	switch m.dir {
	case right:
		return x + m.n, y
	case down:
		return x, y - m.n
	case left:
		return x - m.n, y
	case up:
		return x, y + m.n
	}
	panic("invalid direction")
}

type dir byte

const (
	right dir = 'R'
	down  dir = 'D'
	left  dir = 'L'
	up    dir = 'U'
)

func parse(r io.Reader) (w1, w2 []move, err error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}
	ls := bytes.SplitN(bs, []byte{'\n'}, 2)
	for _, i := range bytes.Split(bytes.TrimSpace(ls[0]), []byte{','}) {
		n, err := strconv.Atoi(string(i[1:]))
		if err != nil {
			return nil, nil, err
		}
		for n > 0 {
			w1 = append(w1, move{dir(i[0]), 1})
			n--
		}
	}
	for _, i := range bytes.Split(bytes.TrimSpace(ls[1]), []byte{','}) {
		n, err := strconv.Atoi(string(i[1:]))
		if err != nil {
			return nil, nil, err
		}
		for n > 0 {
			w2 = append(w2, move{dir(i[0]), 1})
			n--
		}
	}
	return
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	w1, w2, err := parse(r)
	if err != nil {
		return "", err
	}
	m := make(map[[2]int]struct{})
	i := make(map[[2]int]struct{})
	x, y := 0, 0
	for _, mv := range w1 {
		x, y = mv.from(x, y)
		m[[2]int{x, y}] = struct{}{}
	}
	x, y = 0, 0
	for _, mv := range w2 {
		x, y = mv.from(x, y)
		if _, ok := m[[2]int{x, y}]; ok {
			i[[2]int{x, y}] = struct{}{}
		}
	}
	d := math.MaxInt32
	for k := range i {
		d = min(d, distance(k))
	}
	return strconv.Itoa(d), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	w1, w2, err := parse(r)
	if err != nil {
		return "", err
	}
	m := make(map[[2]int]int)
	d := math.MaxInt32
	x, y := 0, 0
	s := 0
	for _, mv := range w1 {
		x, y = mv.from(x, y)
		s++
		m[[2]int{x, y}] = s
	}
	x, y = 0, 0
	s = 0
	for _, mv := range w2 {
		x, y = mv.from(x, y)
		s++
		if s1, ok := m[[2]int{x, y}]; ok {
			d = min(d, s1+s)
		}
	}
	return strconv.Itoa(d), nil
}

func distance(p [2]int) int {
	return abs(p[0]) + abs(p[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
