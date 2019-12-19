package day10

import (
	"bufio"
	"fmt"
	aoc "github.com/SHyx0rmZ/advent-of-code"
	"io"
	"math"
	"sort"
	"strconv"
)

type point struct {
	X, Y float64
}

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	m := make(map[point]struct{})
	s := bufio.NewScanner(r)
	var y int
	for s.Scan() {
		if len(s.Text()) == 0 {
			continue
		}
		for x, b := range s.Text() {
			if b == '.' {
				continue
			}
			m[point{float64(x), float64(y)}] = struct{}{}
		}
		y++
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	var x struct {
		Vis int
		Pos point
	}
	for a := range m {
		ls := make(map[float64]struct{})
		for o := range m {
			if a == o {
				continue
			}
			bt := point{o.X-a.X, o.Y-a.Y}
			l := math.Atan2(bt.Y, bt.X)
			ls[l] = struct{}{}
		}
		if len(ls) > x.Vis {
			x.Vis = len(ls)
			x.Pos = a
		}
	}
	return strconv.Itoa(x.Vis), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	m := make(map[point]struct{})
	s := bufio.NewScanner(r)
	var y int
	for s.Scan() {
		if len(s.Text()) == 0 {
			continue
		}
		for x, b := range s.Text() {
			if b == '.' {
				continue
			}
			m[point{float64(x), float64(y)}] = struct{}{}
		}
		y++
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	var x struct {
		Vis int
		Pos point
	}
	for a := range m {
		ls := make(map[float64]struct{})
		for o := range m {
			if a == o {
				continue
			}
			bt := point{o.X-a.X, o.Y-a.Y}
			l := math.Atan2(bt.Y, bt.X)
			ls[l] = struct{}{}
		}
		if len(ls) > x.Vis {
			x.Vis = len(ls)
			x.Pos = a
		}
	}
	fmt.Println(">>>>", x)
	a := x.Pos
	ls := make(map[float64]point)
	for o := range m {
		if a == o {
			continue
		}
		bt := point{o.X-a.X, o.Y-a.Y}
		l := math.Atan2(-bt.Y, bt.X)
		l = l-math.Pi/2
		if l < 0 {
			l += math.Pi*2
		}
		ls[l] = o
	}
	var pt point
	for i := 0; i < 200 ; i++ {
		//if len(ls) == 0 {
		//	break
		//}
		var as []float64
		for a := range ls {
			as = append(as, a)
		}
		sort.Slice(as, func(i, j int) bool {
			l, r := as[i], as[j]
			//if l > math.Pi/2 { l+= 2*math.Pi}
			//if r > math.Pi/2 { r+= 2*math.Pi}
			return l > r
		})
		fmt.Println(i, ls[as[0]], as[0])
		pt = ls[as[0]]
		delete(ls, as[0])
	}
	return strconv.Itoa(int(pt.X *100 + pt.Y)), nil
}
