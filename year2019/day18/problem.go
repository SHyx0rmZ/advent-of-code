package day18

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"os"
	"sort"

	aoc "github.com/SHyx0rmZ/advent-of-code"
)

type point struct {
	X, Y int
}

func (p point) Less(o point) bool {
	if p.Y == o.Y {
		return p.X < o.X
	}
	return p.Y < o.Y
}

type board map[point]byte

func (b board) Bounds() image.Rectangle {
	if len(b) == 0 {
		return image.Rectangle{}
	}
	ps := make([]point, 0, len(b))
	for p := range b {
		ps = append(ps, p)
	}
	min := ps[0]
	max := ps[0]
	for _, p := range ps[1:] {
		if p.Less(min) {
			min = p
		}
		if max.Less(p) {
			max = p
		}
	}
	return image.Rect(min.X, min.Y, max.X+1, max.Y+1)
}

func (b board) Print(g *gif.GIF) {
	bounds := b.Bounds()
	i := image.NewPaletted(bounds, g.Config.ColorModel.(color.Palette))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			t := b[point{x, y}]
			switch {
			case t == '#':
				i.SetColorIndex(x, y, 1)
			case t == '@':
				i.SetColorIndex(x, y, 2)
			case t >= 'a' && t <= 'z':
				i.SetColorIndex(x, y, 5+uint8((t|0x20)-'a'))
			case t >= 'A' && t <= 'Z':
				i.SetColorIndex(x, y, 5+uint8((t|0x20)-'a'))
			}
		}
	}
	g.Image = append(g.Image, i)
	g.Delay = append(g.Delay, 1)
	g.Disposal = append(g.Disposal, gif.DisposalBackground)
}

const entrance = '@'
const passage = '.'
const wall = '#'

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	b := make(board)
	ks := make(map[byte]point)
	ds := make(map[byte]point)
	var e point
	var y int
	s := bufio.NewScanner(r)
	for s.Scan() {
		for x, c := range s.Bytes() {
			p := point{x, y}
			switch {
			case c >= 'a' && c <= 'z':
				ks[c] = p
			case c >= 'A' && c <= 'Z':
				ds[c] = p
			case c == '@':
				e = p
			}
			b[p] = c
		}
		y++
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	g := &gif.GIF{
		LoopCount: -1,
		Config: image.Config{
			ColorModel: color.Palette{
				color.Black,
				color.Gray{Y: 64},
				color.Gray{Y: 32},
				color.Gray{Y: 255},
				color.Gray{Y: 128},
				hsl(360.0/12.0*0, 0.5, 0.25),
				hsl(360.0/12.0*1, 0.5, 0.25),
				hsl(360.0/12.0*2, 0.5, 0.25),
				hsl(360.0/12.0*3, 0.5, 0.25),
				hsl(360.0/12.0*4, 0.5, 0.25),
				hsl(360.0/12.0*5, 0.5, 0.25),
				hsl(360.0/12.0*6, 0.5, 0.25),
				hsl(360.0/12.0*7, 0.5, 0.25),
				hsl(360.0/12.0*8, 0.5, 0.25),
				hsl(360.0/12.0*9, 0.5, 0.25),
				hsl(360.0/12.0*10, 0.5, 0.25),
				hsl(360.0/12.0*11, 0.5, 0.25),
				hsl(360.0/12.0*12, 0.5, 0.25),
				hsl(360.0/12.0*0, 1, 0.5),
				hsl(360.0/12.0*1, 1, 0.5),
				hsl(360.0/12.0*2, 1, 0.5),
				hsl(360.0/12.0*3, 1, 0.5),
				hsl(360.0/12.0*4, 1, 0.5),
				hsl(360.0/12.0*5, 1, 0.5),
				hsl(360.0/12.0*6, 1, 0.5),
				hsl(360.0/12.0*7, 1, 0.5),
				hsl(360.0/12.0*8, 1, 0.5),
				hsl(360.0/12.0*9, 1, 0.5),
				hsl(360.0/12.0*10, 1, 0.5),
				hsl(360.0/12.0*11, 1, 0.5),
				hsl(360.0/12.0*12, 1, 0.5),
			},
			Width:  b.Bounds().Dx(),
			Height: b.Bounds().Dy(),
		},
	}
	//for c := range ks {
	//	i := image.NewPaletted(b.Bounds(), g.Config.ColorModel.(color.Palette))
	//	i.SetColorIndex(ks[c].X, ks[c].Y, 3)
	//	i.SetColorIndex(ds[c^0x20].X, ds[c^0x20].Y, 4)
	//	g.Image = append(g.Image, i)
	//	g.Delay = append(g.Delay, 100)
	//	g.Disposal = append(g.Disposal, gif.DisposalNone)
	//}

	sr := make(map[point]struct{})
	var reduce func(ps []point)
	reduce = func(ps []point) {
		if len(ps) == 0 {
			return
		}
		var tr []point
		for _, p := range ps {
			if _, ok := sr[p]; ok {
				continue
			}
			sr[p] = struct{}{}
			for _, np := range []point{{p.X, p.Y + 1}, {p.X, p.Y - 1}, {p.X + 1, p.Y}, {p.X - 1, p.Y}} {
				if b[np] != '#' {
					tr = append(tr, np)
				}
			}
		}
		reduce(tr)
		b.Print(g)
		for _, p := range ps {
			if b[p] != '.' {
				continue
			}
			var n int
			for _, np := range []point{
				{p.X, p.Y + 1},
				{p.X, p.Y - 1},
				{p.X + 1, p.Y},
				{p.X - 1, p.Y},
			} {
				if b[np] == '#' {
					n++
				}
			}
			if n == 3 {
				b[p] = '#'
			}
		}
	}

	f, err := os.Create("day18.gif")
	if err != nil {
		return "", err
	}
	defer f.Close()
	defer gif.EncodeAll(f, g)
	reduce([]point{e})
	b.Print(g)
	rs := make(map[[2]byte][]point)
	for k1 := range ks {
		for k2 := range ks {
			if k2 == k1 {
				continue
			}
			r, ok := rs[[2]byte{k2, k1}]
			if ok {
				continue
			}
			r = aStar(ks[k1], ks[k2], b)
			if len(r) == 0 {
				return "", fmt.Errorf("no path")
			}
			rs[[2]byte{k1, k2}] = r
			//fmt.Println(string(k1), " => ", string(k2), " :", len(r))
		}
	}
	for k2 := range ks {
		r, ok := rs[[2]byte{k2, '@'}]
		if ok {
			continue
		}
		r = aStar(e, ks[k2], b)
		if len(r) == 0 {
			return "", fmt.Errorf("no path")
		}
		rs[[2]byte{'@', k2}] = r
		//fmt.Println(string(k1), " => ", string(k2), " :", len(r))
	}
	d := make(map[[2]byte][]byte)
	for k, r := range rs {
		d[k] = nil
		for _, p := range r {
			if 'A' <= b[p] && b[p] <= 'Z' {
				d[k] = append(d[k], b[p])
			}
		}
	}
	for k, dx := range d {
		fmt.Println(string(k[:]), ":", string(dx))
	}
	// from here on, just try all combinations?
	// limit choices by keys picked up (0 initially)

	return "", nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	return "", nil
}

func hsl(hue, saturation, lightness float64) color.RGBA {
	c := (1 - math.Abs(2.0*lightness-1)) * saturation
	h := hue / 60.0
	x := c * (1 - math.Abs(math.Mod(h, 2)-1))
	var r, g, b float64
	switch {
	case 0 <= h && h <= 1:
		r, g = c, x
	case 1 <= h && h <= 2:
		r, g = x, c
	case 2 <= h && h <= 3:
		g, b = c, x
	case 3 <= h && h <= 4:
		g, b = x, c
	case 4 <= h && h <= 5:
		r, b = x, c
	case 5 <= h && h <= 6:
		r, b = c, x
	}
	m := lightness - c/2
	return color.RGBA{
		R: uint8((r + m) * 255.0),
		G: uint8((g + m) * 255.0),
		B: uint8((b + m) * 255.0),
		A: 255,
	}
}

// this should be somewhat more correct than the other versions
func aStar(start, goal point, m board) []point {
	visited := make(map[point]struct{})
	toVisit := make(map[point]struct{})
	toVisit[start] = struct{}{}
	//for p, t := range m {
	//	if t != '#' {
	//		toVisit[p] = struct{}{}
	//	}
	//}
	distance := func(a, b point) int {
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
	heuristic := distance
	cameFrom := make(map[point]point)
	scores := make(map[point]int)
	scores[start] = 0
	estimatedScores := make(map[point]int)
	estimatedScores[start] = heuristic(start, goal)
	reconstructPath := func(current point) []point {
		ps := []point{goal}
		for {
			p, ok := cameFrom[current]
			if !ok {
				break
			}
			current = p
			ps = append(ps, current)
		}
		return ps
	}
	for len(toVisit) > 0 {
		var minimum struct {
			Score int
			Point point
		}
		minimum.Score = infinity
		var ps pointSlice
		for t := range toVisit {
			ps = append(ps, t)
		}
		sort.Sort(sort.Reverse(ps))
		for _, p := range ps {
			estimatedScore, ok := estimatedScores[p]
			if ok && estimatedScore <= minimum.Score {
				minimum.Score = estimatedScore
				minimum.Point = p
			}
			if !ok && minimum.Score == infinity {
				minimum.Point = p
			}
		}
		current := minimum.Point
		if current == goal {
			return reconstructPath(current)
		}
		delete(toVisit, current)
		visited[current] = struct{}{}
		for _, neighbor := range []point{
			{current.X, current.Y - 1},
			{current.X - 1, current.Y},
			{current.X + 1, current.Y},
			{current.X, current.Y + 1},
		} {
			if _, ok := visited[neighbor]; ok {
				continue
			}
			if m[neighbor] == '#' {
				continue
			}
			score := scores[current] + distance(current, neighbor)
			cameFrom[neighbor] = current
			scores[neighbor] = score
			estimatedScores[neighbor] = score + heuristic(neighbor, goal)
			if _, ok := toVisit[neighbor]; !ok {
				toVisit[neighbor] = struct{}{}
			}
		}
	}
	return nil
}

const infinity = 999999999

type pointSlice []point

func (p pointSlice) Len() int {
	return len(p)
}

func (p pointSlice) Less(i, j int) bool {
	if p[i].Y == p[j].Y {
		return p[i].X < p[j].X
	}
	return p[i].Y < p[j].Y
}

func (p pointSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
