package day12

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

type pf func(string) pf

type st struct {
	state  []bool
	offset int
	age    int
	rules  map[[5]bool]bool
	stable bool
}

func (s *st) parseInitial(l string) pf {
	ps := strings.SplitN(l, "initial state: ", 2)
	for i := range ps[1] {
		s.state = append(s.state, ps[1][i] == '#')
	}
	return s.parseRule
}

func (s *st) parseRule(l string) pf {
	if len(l) == 0 {
		return s.parseRule
	}
	ps := strings.SplitN(l, " => ", 2)
	if ps[1] == "." {
		return s.parseRule
	}
	var b [5]bool
	b[0] = ps[0][0] == '#'
	b[1] = ps[0][1] == '#'
	b[2] = ps[0][2] == '#'
	b[3] = ps[0][3] == '#'
	b[4] = ps[0][4] == '#'
	if s.rules == nil {
		s.rules = make(map[[5]bool]bool)
	}
	s.rules[b] = ps[1] == "#"
	return s.parseRule
}

func (s st) bit(i int) bool {
	if i < 0 || i >= len(s.state) {
		return false
	}
	return s.state[i]
}

func (s *st) grow() {
	if s.stable {
		s.age++
		s.offset++
		return
	}
	var new []bool
	for i := -2; i < len(s.state)+2; i++ {
		var b [5]bool
		b[0] = s.bit(i - 2)
		b[1] = s.bit(i - 1)
		b[2] = s.bit(i + 0)
		b[3] = s.bit(i + 1)
		b[4] = s.bit(i + 2)
		ok := s.rules[b]
		new = append(new, ok)
	}
	if !new[0] && !new[1] && !new[2] && !new[3] {
		new = new[1:]
		s.offset++
	}
	if !new[0] && !new[1] && !new[2] && !new[3] {
		new = new[1:]
		s.offset++
	}
	if !new[0] && !new[1] && !new[2] && !new[3] {
		new = new[1:]
		s.offset++
	}
	if !new[len(new)-4] && !new[len(new)-3] && !new[len(new)-2] && !new[len(new)-1] {
		new = new[:len(new)-1]
	}
	if !new[len(new)-4] && !new[len(new)-3] && !new[len(new)-2] && !new[len(new)-1] {
		new = new[:len(new)-1]
	}
	if !new[len(new)-4] && !new[len(new)-3] && !new[len(new)-2] && !new[len(new)-1] {
		new = new[:len(new)-1]
	}

	if s.same(s.state, new) {
		fmt.Println("same")
		s.stable = true
	}
	s.state = new
	s.age++
	s.offset -= 2
}

func (s st) print(st []bool) {
	for _, b := range st {
		if b {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func (s st) count() int {
	var t int
	for i, b := range s.state {
		if b {
			t += i + s.offset
		}
	}
	return t
}

func (s st) same(o, n []bool) bool {
	if len(o) != len(n) {
		return false
	}
	for i := 0; i < len(o); i++ {
		if o[i] != n[i] {
			return false
		}
	}
	return true
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	st := new(st)
	pf := st.parseInitial
	for s.Scan() {
		pf = pf(s.Text())
	}
	st.print(st.state)
	for i := 0; i < 20; i++ {
		st.grow()
		st.print(st.state)
	}
	return strconv.Itoa(st.count()), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	st := new(st)
	pf := st.parseInitial
	for s.Scan() {
		pf = pf(s.Text())
	}
	st.print(st.state)
	for i := 0; i < 50000000000; i++ {
		st.grow()
		if st.stable {
			st.offset += 50000000000 - i - 1
			break
		}
	}
	return strconv.Itoa(st.count()), nil
}
