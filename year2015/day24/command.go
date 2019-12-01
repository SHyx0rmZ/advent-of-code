package day24

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func parse(r io.Reader) ([]int, error) {
	var ns []int
	s := bufio.NewScanner(r)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return ns, nil
}

func mix(ns []int) {
}

func m(ns, g1, g2, g3 []int, q int) int {
	if len(ns) == 0 {
		fmt.Println(g1, g2, g3, q)
		var t1, t2, t3 int
		t := 1
		for _, n := range g1 {
			t1 += n
			t *= n
		}
		for _, n := range g2 {
			t2 += n
		}
		for _, n := range g3 {
			t3 += n
		}
		if t1 != t2 || t2 != t3 {
			return q
		}
		if t < q && t > 0 {
			fmt.Println(g1, g2, g3, q, t)
			return t
		}
		return q
	}
	q = m(ns[1:], append(g1, ns[0]), g2, g3, q)
	q = m(ns[1:], g1, append(g2, ns[0]), g3, q)
	q = m(ns[1:], g1, g2, append(g3, ns[0]), q)
	return q
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	ns, err := parse(r)
	if err != nil {
		return "", err
	}
	sort.Ints(ns)
	x := m(ns, nil, nil, nil, math.MaxInt32)
	return strconv.Itoa(x), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	return "", nil
}
