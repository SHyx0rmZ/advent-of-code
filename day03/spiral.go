package day03

import (
	"fmt"
	"math"
)

func Command() error {
	n := 368078
	i := 1
	for p := range spiral() {
		if i == n {
			fmt.Printf("%+v\n", math.Abs(float64(p.X))+math.Abs(float64(p.Y)))
			break
		}
		i++
	}
	v := values{}
	for p := range spiral() {
		if v.calculate(p) > n {
			fmt.Printf("%d\n", v.calculate(p))
			break
		}
	}
	return nil
}

var (
	deltas = []position{
		{X: -1, Y: -1},
		{X: -1, Y: +0},
		{X: -1, Y: +1},
		{X: +0, Y: -1},
		{X: +0, Y: +1},
		{X: +1, Y: -1},
		{X: +1, Y: +0},
		{X: +1, Y: +1},
	}
	directions = []position{
		{X: +1, Y: +0},
		{X: +0, Y: +1},
		{X: -1, Y: +0},
		{X: +0, Y: -1},
	}
)

type position struct {
	X int
	Y int
}

func (p position) Add(o position) position {
	return position{
		X: p.X + o.X,
		Y: p.Y + o.Y,
	}
}

type values struct {
	M map[position]int
}

func (v *values) calculate(p position) int {
	if v == nil {
		return 0
	}
	if v.M == nil {
		v.M = map[position]int{
			{X: 0, Y: 0}: 1,
		}
	}
	s, ok := v.M[p]
	if !ok {
		for _, d := range deltas {
			if a, ok := v.M[p.Add(d)]; ok {
				s += a
			}
		}
		v.M[p] = s
	}
	return s
}

func spiral() <-chan position {
	c := make(chan position)
	go func() {
		p := position{}
		g := 0
		for {
			for i := 0; i <= g/2; i++ {
				c <- p
				p = p.Add(directions[g%4])
			}
			g++
		}
	}()
	return c
}
