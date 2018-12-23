package day15

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/SHyx0rmZ/advent-of-code"
)

const delay = 1 * time.Millisecond

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

type point struct {
	X, Y int
}

type terrain byte

const (
	wall terrain = '#'
	cave terrain = '.'
)

type faction byte

const (
	goblin faction = 'G'
	elf    faction = 'E'
)

type unit struct {
	point
	faction
	hp int
}

type tile struct {
	terrain
	*unit
	extra *byte
}

type pointSlice []point

func (p pointSlice) Len() int { return len(p) }
func (p pointSlice) Less(i, j int) bool {
	if p[i].Y == p[j].Y {
		return p[i].X < p[j].X
	}
	return p[i].Y < p[j].Y
}
func (p pointSlice) Swap(i, j int) { t := p[i]; p[i] = p[j]; p[j] = t }

type unitSlice []*unit

func (u unitSlice) Len() int { return len(u) }
func (u unitSlice) Less(i, j int) bool {
	if u[i].Y == u[j].Y {
		return u[i].X < u[j].X
	}
	return u[i].Y < u[j].Y
}
func (u unitSlice) Swap(i, j int) { t := u[i]; u[i] = u[j]; u[j] = t }

type sf func(m map[point]*tile, w, h int, us unitSlice) sf

var combat bool

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	m := make(map[point]*tile)
	r = strings.NewReader(`#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`)
	r = strings.NewReader(`#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`)
	r = strings.NewReader(`#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`)
	//	r = strings.NewReader(`#######
	//#E..EG#
	//#.#G.E#
	//#E.##E#
	//#G..#.#
	//#..E#.#
	//#######`)
	s := bufio.NewScanner(r)
	var w, h int
	var y int
	var us unitSlice
	var gs int
	var es int
	for s.Scan() {
		for x, c := range s.Text() {
			p := point{x, y}
			switch c {
			case rune(wall):
				m[p] = &tile{wall, nil, nil}
			case rune(cave):
				m[p] = &tile{cave, nil, nil}
			case rune(goblin):
				u := &unit{p, goblin, 200}
				m[p] = &tile{cave, u, nil}
				us = append(us, u)
				gs++
			case rune(elf):
				u := &unit{p, elf, 200}
				m[p] = &tile{cave, u, nil}
				us = append(us, u)
				es++
			}
			if x >= w {
				w = x
			}
		}
		y++
	}
	w++
	h = y
	sort.Sort(us)
	var t int
	combat = true
	for combat {
		for sf := findTargets(0); sf != nil; sf = sf(m, w, h, us) {
			for _, au := range us {
				if au.hp <= 0 {
					var ui int
					for i, u := range us {
						if u == au {
							ui = i
							break
						}
					}
					us = append(us[:ui], us[ui+1:]...)
					m[au.point].unit = nil
				}

			}
		}
		t++
		fmt.Printf("\033[%d;%dH%d   ", 1, w+3, t)
		time.Sleep(delay)
		gs = 0
		es = 0
		sort.Sort(us)
		for _, u := range us {
			if u.faction == goblin {
				gs++
			} else {
				es++
			}
		}
	}
	t--
	fmt.Print("\033[2J\033[1;1H")
	var hp int
	for _, u := range us {
		fmt.Println(*u)
		hp += u.hp
	}
	return fmt.Sprintf("%d %d %d\n", t, hp, t*hp), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	return "", nil
}

func cleanExtra(m map[point]*tile, w, h int) {
	//return
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			m[point{x, y}].extra = nil
		}
	}
}

func setExtra(m map[point]*tile, p point, b byte) {
	//return
	m[p].extra = &b
}

func end(m map[point]*tile, w, h int, us unitSlice) sf {
	combat = false
	return nil
}

func findTargets(i int) sf {
	return func(m map[point]*tile, w, h int, us unitSlice) sf {
		print(m, w, h, us)
		if i >= len(us) {
			return nil
		}
		cleanExtra(m, w, h)
		u := us[i]
		var es unitSlice
		for _, t := range us {
			if t.faction != u.faction {
				es = append(es, t)
			}
		}
		if len(es) == 0 {
			return end
		}
		return findOpenSquares(i, es)
	}
}

func findOpenSquares(i int, es unitSlice) sf {
	return func(m map[point]*tile, w, h int, us unitSlice) sf {
		var sq []point
		ff := func(x, y int) {
			p := point{x, y}
			t, ok := m[p]
			if !ok {
				return
			}
			if t.terrain == cave && t.unit == nil {
				sq = append(sq, p)
				setExtra(m, p, '?')
				m[p] = t
			}
		}
		for _, e := range es {
			ff(e.X, e.Y-1)
			ff(e.X-1, e.Y)
			ff(e.X+1, e.Y)
			ff(e.X, e.Y+1)
		}
		return findAttackable(i, es, sq)
	}
}

func findAttackable(i int, es unitSlice, sq []point) sf {
	return func(m map[point]*tile, w, h int, us unitSlice) sf {
		u := us[i]
		var at unitSlice
		for _, e := range es {
			if e.X == u.X && e.Y == u.Y-1 {
				at = append(at, e)
			}
			if e.X == u.X-1 && e.Y == u.Y {
				at = append(at, e)
			}
			if e.X == u.X+1 && e.Y == u.Y {
				at = append(at, e)
			}
			if e.X == u.X && e.Y == u.Y+1 {
				at = append(at, e)
			}
		}
		if len(sq) == 0 && len(at) == 0 {
			return findTargets(i + 1)
		}
		if len(at) > 0 {
			return attack(i, es, at)
		}
		return findReachable(i, es, sq)
	}
}

func findReachable(i int, es unitSlice, sq []point) sf {
	return func(m map[point]*tile, w, h int, us unitSlice) sf {
		rs := make(map[point][]point)
		u := us[i]
		for _, p := range sq {
			ap := aStar(u.point, p, m)
			if len(ap) > 0 && ap[len(ap)-1] == u.point {
				rs[p] = ap
			}
		}
		if len(rs) == 0 {
			return findTargets(i + 1)
		}
		cleanExtra(m, w, h)
		for p := range rs {
			setExtra(m, p, '@')
		}
		return findNearest(i, es, sq, rs)
	}
}

func findNearest(i int, es unitSlice, sq []point, rs map[point][]point) sf {
	return func(m map[point]*tile, w, h int, us unitSlice) sf {
		mi := 999
		for _, p := range rs {
			if len(p) < mi {
				mi = len(p)
			}
		}
		cleanExtra(m, w, h)
		ns := make(map[point][]point)
		for _, p := range sq {
			if ap, ok := rs[p]; ok && len(ap) == mi {
				ns[p] = ap
				setExtra(m, p, '!')
			}
		}
		return chooseStep(i, es, ns)
	}
}

func chooseStep(i int, es unitSlice, ns map[point][]point) sf {
	return func(m map[point]*tile, w, h int, us unitSlice) sf {
		var ps pointSlice
		for p := range ns {
			ps = append(ps, p)
		}
		//fmt.Println(ns, ps[0])
		sort.Sort(ps)
		cleanExtra(m, w, h)
		setExtra(m, ps[0], '+')
		return move(i, es, ps[0], ns[ps[0]])
	}
}

func move(i int, es unitSlice, p point, a []point) sf {
	return func(m map[point]*tile, w, h int, us unitSlice) sf {
		u := us[i]
		cleanExtra(m, w, h)

		fmt.Println(a)
		for pi, ps := range []point{{u.X, u.Y - 1}, {u.X - 1, u.Y}, {u.X + 1, u.Y}, {u.X, u.Y + 1}, {u.X, u.Y}} {
			if m[ps].terrain != cave {
				continue
			}
			ap := aStar(ps, p, m)
			if len(ap) == 0 {
				continue
			}
			fmt.Println(ap)
			if len(ap) == len(a)-1 || len(a) == 2 {
				m[u.point].unit = nil
				if len(a) > 1 {
					u.point = a[len(a)-2]
				} else {
					u.point = p
				}
				m[u.point].unit = u
			}
			if pi == 4 {
				m[u.point].unit = nil
				if len(a) > 1 {
					u.point = a[len(a)-2]
				} else {
					u.point = p
				}
				m[u.point].unit = u
			}
		}

		var at unitSlice
		for _, e := range es {
			if e.X == u.X && e.Y == u.Y-1 {
				at = append(at, e)
			}
			if e.X == u.X-1 && e.Y == u.Y {
				at = append(at, e)
			}
			if e.X == u.X+1 && e.Y == u.Y {
				at = append(at, e)
			}
			if e.X == u.X && e.Y == u.Y+1 {
				at = append(at, e)
			}
		}
		if len(at) == 0 {
			return findTargets(i + 1)
		}
		return attack(i, es, at)
		//for i, ap := range a {
		//	if i != 0 && i != len(a)-1 {
		//		setExtra(m, ap, '%')
		//	}
		//}
		print(m, w, h, us)
		fmt.Print("\033[1;40H")
		fmt.Println(us[i], p, a)
		panic("no step")
	}
}

func attack(i int, es unitSlice, at unitSlice) sf {
	return func(m map[point]*tile, w, h int, us unitSlice) sf {
		var mi *unit
		for _, u := range at {
			if mi == nil || u.hp < mi.hp {
				mi = u
			}
		}
		sort.Sort(at)
		for _, au := range at {
			if au.hp == mi.hp {
				if au.hp > 0 {
					au.hp -= 3
					if au.hp <= 0 {
						var ui int
						for i, u := range us {
							if u == au {
								ui = i
								break
							}
						}
						if ui < i {
							return findTargets(i)
						} else {
							return findTargets(i + 1)
						}
					}

				}
			}
		}
		return findTargets(i + 1)
	}
}

func print(m map[point]*tile, w, h int, us unitSlice) {
	fmt.Print("\033[2J")
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Printf("\033[%d;%dH", y+1, x+1)
			t := m[point{x, y}]
			//fmt.Printf("%#v %v", t, point{x, y})
			if t.extra != nil {
				fmt.Print(string(*t.extra))
			} else if t.unit != nil {
				fmt.Print(string(t.unit.faction))
			} else {
				if t.terrain == wall {
					fmt.Print(" ")
				} else {
					fmt.Print(string(t.terrain))
				}
			}
		}
		fmt.Println()
	}

	var gs int
	var es int
	for _, u := range us {
		if u.faction == goblin {
			fmt.Printf("\033[%d;%dH%3d", h+1+(gs%10), 2+(5*(gs/10)), u.hp)
			gs++
		} else {
			fmt.Printf("\033[%d;%dH%3d", h+1+(es%10), 22+(5*(es/10)), u.hp)
			es++
		}
	}
}

func aStar(s, g point, m map[point]*tile) []point {
	vs := make(map[point]struct{})
	ts := make(map[point]struct{})
	ts[s] = struct{}{}
	for p, t := range m {
		if t.terrain == cave && t.unit == nil {
			ts[p] = struct{}{}
		}
	}
	d := func(a, b point) int {
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
	h := func(x, y point) int {
		return d(x, y)
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
		return p
	}
	for len(ts) > 0 {
		var mi struct {
			int
			point
		}
		mi.int = 999999999
		var cs pointSlice
		for t := range ts {
			cs = append(cs, t)
		}
		sort.Sort(sort.Reverse(cs))
		for _, t := range cs {
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
