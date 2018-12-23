package day20

import (
	"bufio"
	"github.com/SHyx0rmZ/advent-of-code"
	"io"
	"sort"
	"strconv"
	"sync"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

type point struct {
	X, Y int
}

type grid map[string][]point

func (g *grid) Process(ch <-chan rune) {
	g.process(ch, "", 0, 0)
}

func (g *grid) process(ch <-chan rune, s string, x, y int) rune {
	for c := range ch {
		switch c {
		case 'N':
			s += string(c)
			y++
		case 'E':
			s += string(c)
			x++
		case 'S':
			s += string(c)
			y--
		case 'W':
			s += string(c)
			x--
		case '(':
			for {
				c = g.process(ch, s, x, y)
				if c != '|' {
					break
				}
			}
		case ')', '|':
			return c
		}
		var found bool
		p := point{x, y}
		for _, r := range (*g)[s] {
			if p == r {
				found = true
				break
			}
		}
		if !found {
			(*g)[s] = append((*g)[s], p)
		}
	}
	return 0
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	g := make(grid)
	ch := make(chan rune)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		g.Process(ch)
		wg.Done()
	}()
	for s.Scan() {
		for _, c := range s.Text() {
			ch <- c
		}
	}
	close(ch)
	wg.Wait()
	var ps []string
	for p := range g {
		ps = append(ps, p)
	}
	sort.Strings(ps)
	sort.SliceStable(ps, func(i, j int) bool {
		return len(ps[i]) < len(ps[j])
	})
	return strconv.Itoa(len(ps[len(ps)-1])), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	g := make(grid)
	ch := make(chan rune)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		g.Process(ch)
		wg.Done()
	}()
	for s.Scan() {
		for _, c := range s.Text() {
			ch <- c
		}
	}
	close(ch)
	wg.Wait()
	var ps []string
	for p := range g {
		ps = append(ps, p)
	}
	sort.Strings(ps)
	sort.SliceStable(ps, func(i, j int) bool {
		return len(ps[i]) < len(ps[j])
	})

	sr := make(map[point]struct{})
	var rs int
	for _, p := range ps {
		if len(p) >= 1000 {
			for _, r := range g[p] {
				if _, ok := sr[r]; !ok {
					rs++
					sr[r] = struct{}{}
				}
			}
		}
	}
	return strconv.Itoa(rs), nil
}
