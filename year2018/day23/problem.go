package day23

import (
	"bufio"
	"fmt"
	"io"
	"sort"
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
	X, Y, Z int
}

type pair [2]point

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	m := make(map[point]int)
	d := make(map[pair]int)
	var mr struct {
		int
		point
	}
	for s.Scan() {
		ps1 := strings.SplitN(s.Text(), "pos=<", 2)
		ps2 := strings.SplitN(ps1[1], ">, r=", 2)
		ps3 := strings.SplitN(ps2[0], ",", 3)
		var p point
		p.X, _ = strconv.Atoi(ps3[0])
		p.Y, _ = strconv.Atoi(ps3[1])
		p.Z, _ = strconv.Atoi(ps3[2])
		var r int
		r, _ = strconv.Atoi(ps2[1])
		m[p] = r
		if r > mr.int {
			mr.int = r
			mr.point = p
		}
	}
	var t int
	for po := range m {
		for p := range m {
			if _, ok := d[pair{po, p}]; !ok {
				v := dist(p, po)
				d[pair{po, p}] = v
				d[pair{p, po}] = v
			}
		}
	}
	for p := range m {
		if d[pair{p, mr.point}] <= mr.int {
			t++
		}
	}
	fmt.Println(mr)
	//return strconv.Itoa(dist(point{0, 0, 0}, mr.point)), nil
	return strconv.Itoa(t), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	m := make(map[point]int)
	d := make(map[pair]int)
	c := make(map[point]int)
	var mr struct {
		int
		point
	}
	var mi point
	var ma point
	for s.Scan() {
		ps1 := strings.SplitN(s.Text(), "pos=<", 2)
		ps2 := strings.SplitN(ps1[1], ">, r=", 2)
		ps3 := strings.SplitN(ps2[0], ",", 3)
		var p point
		p.X, _ = strconv.Atoi(ps3[0])
		p.Y, _ = strconv.Atoi(ps3[1])
		p.Z, _ = strconv.Atoi(ps3[2])
		var r int
		r, _ = strconv.Atoi(ps2[1])
		m[p] = r
		if r > mr.int {
			mr.int = r
			mr.point = p
		}
		if p.X < mi.X {
			mi.X = p.X
		}
		if p.X > ma.X {
			ma.X = p.X
		}
		if p.Y < mi.Y {
			mi.Y = p.Y
		}
		if p.Y > ma.Y {
			ma.Y = p.Y
		}
		if p.Z < mi.Z {
			mi.Z = p.Z
		}
		if p.Z > ma.Z {
			ma.Z = p.Z
		}
	}
	//c := make(map[point]int)
	for po := range m {
		for p := range m {
			if _, ok := d[pair{po, p}]; !ok {
				v := dist(p, po)
				//if v <= ro {
				//	c[po]++
				//}
				d[pair{po, p}] = v
				d[pair{p, po}] = v
			}
		}
	}
	fmt.Println(mi)
	fmt.Println(ma)
	//xm := make(map[int]int)
	//for x := mi.X; x <= ma.X; x++ {
	//	for p, r := range m {
	//		if dist(p, point{x, p.Y, p.Z}) <= r {
	//			xm[x]++
	//		}
	//	}
	//}
	//var mx struct {
	//	V int
	//	X int
	//}
	//for x, v := range xm {
	//	if v > mx.V {
	//		mx.V = v
	//		mx.X = x
	//	}
	//}
	//fmt.Println(mx)
	var ps []point
	for p := range m {
		ps = append(ps, p)
	}
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].Z == ps[j].Z {
			if ps[i].Y == ps[j].Y {
				return ps[i].X < ps[j].X
			}
			return ps[i].Y < ps[j].Y
		}
		return ps[i].Z < ps[j].Z
	})
	sort.Slice(ps, func(i, j int) bool {
		return c[ps[i]] < c[ps[j]]
	})
	//for _, p := range ps {
	//	fmt.Println(c[p], p)
	//}
	perm(c, m, ps, 0)
	var cs []point
	for p := range c {
		cs = append(cs, p)
	}
	sort.Slice(cs, func(i, j int) bool {
		return c[cs[i]] < c[cs[j]]
	})
	cp := cs[len(cs)-1]
	wc := c[cp]
	var cl struct {
		int
		point
	}
	cl.int = dist(point{}, cp)
	for _, p := range cs {
		if c[p] != wc {
			continue
		}
		v := dist(point{}, p)
		if v < cl.int {
			cl.int = v
			cl.point = p
		}
	}
	fmt.Println(cp, c[cp], dist(point{}, cp))
	fmt.Println(cl.point, c[cl.point], dist(point{}, cl.point))
	return strconv.Itoa(dist(point{0, 0, 0}, ps[len(ps)-1])), nil
}

func perm(c, os map[point]int, ps []point, l int) {
	fmt.Println(len(ps))
	if len(ps) < 2 {
		return
	}
	m := middle(ps)
	var t int
	for p, r := range os {
		if dist(p, m) <= r {
			t++
		}
	}
	c[m] = t
	if t < l {
		return
	}
	for i := 0; i < len(ps); i++ {
		perm(c, os, append(ps[:i], ps[i+1:]...), t)
	}
}

func middle(ps []point) point {
	var m point
	for _, p := range ps {
		m.X += p.X
		m.Y += p.Y
		m.Z += p.Z
	}
	m.X /= len(ps)
	m.Y /= len(ps)
	m.Z /= len(ps)
	return m
}

func dist(a, b point) int {
	x := a.X - b.X
	y := a.Y - b.Y
	z := a.Z - b.Z
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if z < 0 {
		z = -z
	}
	return x + y + z
}
