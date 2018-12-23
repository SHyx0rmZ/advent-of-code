package day22

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

type point struct {
	X, Y int
}

type terrain struct {
	geo int
	ero int
	typ
}

type typ int

const (
	rocky typ = iota
	wet
	narrow
)

const (
	mouth  typ = 4
	target typ = 8
	route  typ = 16
)

func (s typ) String() string {
	if (s & mouth) == mouth {
		return "M"
	}
	if (s & target) == target {
		return "T"
	}
	if (s & route) == route {
		return "%"
	}
	switch s & 3 {
	case rocky:
		return "."
	case wet:
		return "="
	case narrow:
		return "|"
	}
	panic("s ty string")
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	ch := make(chan string)
	go func() {
		for s.Scan() {
			switch s.Text() {
			case "depth:", "target:":
			default:
				ch <- s.Text()
			}
		}
	}()
	depthString := <-ch
	targetString := <-ch
	//depthString = "510"
	//targetString = "10,10"
	d, _ := strconv.Atoi(depthString)
	m := make(map[point]*terrain)
	//m[point{0, 0}] = terrain{
	//	geo: 0,
	//	ero: (0 + d) % 20183,
	//	typ: typ(((0+d)%20183)%3) | mouth,
	//}
	var tx, ty int
	ps := strings.Split(targetString, ",")
	tx, _ = strconv.Atoi(ps[0])
	ty, _ = strconv.Atoi(ps[1])
	//m[point{tx, ty}] = terrain{
	//	geo: 0,
	//	ero: (0 + d) % 20183,
	//	typ: typ(((0+d)%20183)%3) | target,
	//}
	for y := 0; y <= ty; y++ {
		for x := 0; x <= tx; x++ {
			t := new(terrain)
			if y == 0 {
				t.geo = x * 16807
				t.ero = (t.geo + d) % 20183
				t.typ = typ(t.ero % 3)
				m[point{x, y}] = t
				continue
			}
			if x == 0 {
				t.geo = y * 48271
				t.ero = (t.geo + d) % 20183
				t.typ = typ(t.ero % 3)
				m[point{x, y}] = t
				continue
			}
			t.geo = m[point{x - 1, y}].ero * m[point{x, y - 1}].ero
			if x == tx && y == ty {
				t.geo = 0
			}
			t.ero = (t.geo + d) % 20183
			t.typ = typ(t.ero % 3)
			m[point{x, y}] = t
		}
	}
	m[point{0, 0}].typ |= mouth
	m[point{tx, ty}].typ |= target
	var rs int
	for y := 0; y <= ty; y++ {
		for x := 0; x <= tx; x++ {
			rs += int(m[point{x, y}].typ & 3)
			fmt.Print(m[point{x, y}].typ)
		}
		fmt.Println()
	}
	return strconv.Itoa(rs), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	ch := make(chan string)
	go func() {
		for s.Scan() {
			switch s.Text() {
			case "depth:", "target:":
			default:
				ch <- s.Text()
			}
		}
	}()
	depthString := <-ch
	targetString := <-ch
	//depthString = "510"
	//targetString = "10,10"
	d, _ := strconv.Atoi(depthString)
	m := make(map[point]*terrain)
	var tx, ty int
	ps := strings.Split(targetString, ",")
	tx, _ = strconv.Atoi(ps[0])
	ty, _ = strconv.Atoi(ps[1])
	for y := 0; y <= ty+20; y++ {
		for x := 0; x <= tx+10; x++ {
			t := new(terrain)
			if y == 0 {
				t.geo = x * 16807
				t.ero = (t.geo + d) % 20183
				t.typ = typ(t.ero % 3)
				m[point{x, y}] = t
				continue
			}
			if x == 0 {
				t.geo = y * 48271
				t.ero = (t.geo + d) % 20183
				t.typ = typ(t.ero % 3)
				m[point{x, y}] = t
				continue
			}
			t.geo = m[point{x - 1, y}].ero * m[point{x, y - 1}].ero
			if x == tx && y == ty {
				t.geo = 0
			}
			t.ero = (t.geo + d) % 20183
			t.typ = typ(t.ero % 3)
			m[point{x, y}] = t
		}
	}
	m[point{0, 0}].typ |= mouth
	m[point{tx, ty}].typ |= target
	var rs int
	for _, p := range aStar(point{0, 0}, point{tx, ty}, m) {
		m[p].typ |= route
	}
	for y := 0; y <= ty+20; y++ {
		for x := 0; x <= tx+10; x++ {
			rs += int(m[point{x, y}].typ & 3)
			fmt.Print(m[point{x, y}].typ)
		}
		fmt.Println()
	}
	return strconv.Itoa(rs), nil
}

func aStar(s, g point, m map[point]*terrain) []point {
	vs := make(map[point]struct{})
	ts := make(map[point]struct{})
	ts[s] = struct{}{}
	for p := range m {
		//if t.terrain == cave && t.unit == nil {
		ts[p] = struct{}{}
		//}
	}
	d := func(a, b point) int {
		x := m[a]
		y := m[b]
		if x == nil || y == nil {
			return 999999999
		}
		return map[struct{ _1, _2 typ }]int{
			{rocky, rocky}:   1,
			{rocky, wet}:     1,
			{rocky, narrow}:  7,
			{wet, rocky}:     1,
			{wet, wet}:       1,
			{wet, narrow}:    7,
			{narrow, rocky}:  1,
			{narrow, wet}:    1,
			{narrow, narrow}: 7,
		}[struct{ _1, _2 typ }{x.typ & 3, y.typ & 3}]
	}
	h := func(a, b point) int {
		x := b.X - a.X
		y := b.Y - a.Y
		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}
		return x + y
	}
	cf := make(map[point]point)
	cm := make(map[point]int)
	cm[s] = 0
	fm := make(map[point]int)
	fm[s] = h(s, g)
	rp := func(c point) []point {
		var p []point
		for {
			f, ok := cf[c]
			if !ok {
				break
			}
			c = f
			p = append(p, c)
		}
		var rp []point
		for i := 0; i < len(p); i++ {
			rp = append(rp, p[len(p)-1-i])
		}
		rp = append(rp, g)
		return rp
	}
	for len(ts) > 0 {
		var mi struct {
			int
			point
		}
		mi.int = 999999999
		for t := range ts {
			fs, ok := fm[t]
			if ok && fs <= mi.int {
				mi.int = fs
				mi.point = t
			}
			if !ok && mi.int == 999999999 {
				mi.point = t
			}
		}
		cu := mi.point
		if cu == g {
			return rp(cu)
		}
		delete(ts, cu)
		vs[cu] = struct{}{}
		for _, ne := range []point{{cu.X, cu.Y - 1}, {cu.X - 1, cu.Y}, {cu.X + 1, cu.Y}, {cu.X, cu.Y + 1}} {
			if _, ok := vs[ne]; ok {
				continue
			}
			s := cm[cu] + d(cu, ne)
			cf[ne] = cu
			cm[ne] = s
			fm[ne] = s + h(ne, g)
		}
	}
	return nil
}
