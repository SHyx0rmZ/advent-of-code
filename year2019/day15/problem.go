package day15

import (
	aoc "github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type point struct {
	X, Y int
}

func (p point) Move(d direction) point {
	switch d {
	case north:
		return point{p.X, p.Y-1}
	case south:
		return point{p.X, p.Y+1}
	case west:
		return point{p.X-1, p.Y}
	case east:
		return point{p.X+1, p.Y}
	default:
		panic("invalid direction")
	}
}

type terrain int

const (
	unknown terrain = iota
	empty
	wall
	system
	droid
	route
	oxygen
)

type direction int

const (
	north direction = iota + 1
	south
	west
	east
)

func (d direction) Invert() direction {
	switch d {
	case north:
		return south
	case south:
		return north
	case west:
		return east
	case east:
		return west
	default:
		panic("invalid direction")
	}
}

type response int

const (
	hit response = iota
	moved
	arrived
)

type problem struct {}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	output := make(chan int, 1)
	go prg.Run(input, output)
	m := make(map[point]terrain)
	d := point{0, 0}
	m[d] = droid
	const scale = 5
	g := &gif.GIF{
		LoopCount:       -1,
		Config:          image.Config{
			ColorModel: color.Palette{
				color.RGBA{},
				color.RGBA{R: 128, G: 128, B: 128, A:255},
				color.RGBA{R: 255, G: 32, B: 16, A:255},
				color.RGBA{R: 0, G: 255, B: 0, A:255},
				color.RGBA{R: 255, G: 255, B: 0, A:255},
				color.RGBA{R: 255, G: 255, B: 255, A:255},
			},
			Width:      100* scale,
			Height:     100* scale,
		},
	}
	f, err := os.Create("day15.gif")
	if err != nil {
		panic(err)
	}
	defer func () {
		err = gif.EncodeAll(f, g)
		if err != nil {
			log.Println(err)
		}
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	rand.Seed(time.Now().Unix())
	layer := func() {
		i := image.NewPaletted(
			image.Rect(0,0,g.Config.Width, g.Config.Height),
			g.Config.ColorModel.(color.Palette),
		)
		for y := 0; y < g.Config.Height; y += scale {
			for x := 0; x < g.Config.Width; x += scale {
				pi := uint8(m[point{(x/ scale)-g.Config.Width/(2*scale), (y/ scale)-g.Config.Height/(2*scale)}])
				if pi == 0 {
					continue
				}
				for py := 0; py < scale; py++ {
					for px := 0; px < scale; px++ {
						i.SetColorIndex(x+px, y+py, pi)
					}
				}
			}
		}
		g.Image = append(g.Image, i)
		g.Delay = append(g.Delay, 0)
		g.Disposal = append(g.Disposal, gif.DisposalNone)
	}
	var s point
	var explore func(p point)
	explore = func(p point) {
		dirs := []direction{north, south, west, east}
		//perm := rand.Perm(4)
		//dirs = []direction{dirs[perm[0]], dirs[perm[1]], dirs[perm[2]], dirs[perm[3]]}
		for _, d := range dirs {
			np := p.Move(d)
			if _, ok := m[np]; ok {
				continue
			}
			input <- int(d)
			switch response(<-output) {
			case hit:
				m[np] = wall
			case moved:
				m[np] = empty
				layer()
				explore(np)
				input <- int(d.Invert())
				if response(<-output) != moved {
					panic("expected droid to have moved back")
				}
			case arrived:
				m[np] = system
				s = np
				layer()
				input <- int(d.Invert())
				if response(<-output) != moved {
					panic("expected droid to have moved back")
				}
				return
			}
		}
	}
	explore(d)
	layer()
	path := aStar(d, s, m)
	for _, p := range path {
		if p == d || p == s {
			continue
		}
		m[p] = route
	}
	layer()
	return strconv.Itoa(len(path)), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	output := make(chan int, 1)
	go prg.Run(input, output)
	m := make(map[point]terrain)
	d := point{0, 0}
	m[d] = droid
	const scale = 5
	g := &gif.GIF{
		LoopCount:       -1,
		Config:          image.Config{
			ColorModel: color.Palette{
				color.RGBA{},
				color.RGBA{R: 128, G: 128, B: 128, A:255},
				color.RGBA{R: 255, G: 32, B: 16, A:255},
				color.RGBA{R: 0, G: 255, B: 0, A:255},
				color.RGBA{R: 255, G: 255, B: 0, A:255},
				color.RGBA{R: 255, G: 255, B: 255, A:255},
				color.RGBA{R: 128, G: 128, B: 255, A:255},
			},
			Width:      100* scale,
			Height:     100* scale,
		},
	}
	f, err := os.Create("day15.gif")
	if err != nil {
		panic(err)
	}
	defer func () {
		err = gif.EncodeAll(f, g)
		if err != nil {
			log.Println(err)
		}
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	rand.Seed(time.Now().Unix())
	layer := func() {
		i := image.NewPaletted(
			image.Rect(0,0,g.Config.Width, g.Config.Height),
			g.Config.ColorModel.(color.Palette),
		)
		for y := 0; y < g.Config.Height; y += scale {
			for x := 0; x < g.Config.Width; x += scale {
				pi := uint8(m[point{(x/ scale)-g.Config.Width/(2*scale), (y/ scale)-g.Config.Height/(2*scale)}])
				if pi == 0 {
					continue
				}
				for py := 0; py < scale; py++ {
					for px := 0; px < scale; px++ {
						i.SetColorIndex(x+px, y+py, pi)
					}
				}
			}
		}
		g.Image = append(g.Image, i)
		g.Delay = append(g.Delay, 0)
		g.Disposal = append(g.Disposal, gif.DisposalNone)
	}
	var s point
	var explore func(p point)
	explore = func(p point) {
		dirs := []direction{north, south, west, east}
		//perm := rand.Perm(4)
		//dirs = []direction{dirs[perm[0]], dirs[perm[1]], dirs[perm[2]], dirs[perm[3]]}
		for _, d := range dirs {
			np := p.Move(d)
			if _, ok := m[np]; ok {
				continue
			}
			input <- int(d)
			switch response(<-output) {
			case hit:
				m[np] = wall
			case moved:
				m[np] = empty
				layer()
				explore(np)
				input <- int(d.Invert())
				if response(<-output) != moved {
					panic("expected droid to have moved back")
				}
			case arrived:
				m[np] = system
				s = np
				layer()
				input <- int(d.Invert())
				if response(<-output) != moved {
					panic("expected droid to have moved back")
				}
				return
			}
		}
	}
	explore(d)
	layer()
	path := aStar(d, s, m)
	for _, p := range path {
		if p == d || p == s {
			continue
		}
		m[p] = route
	}
	layer()
	var fill func(ps map[point]struct{}) int
	fill = func(ps map[point]struct{}) int {
		f := make(map[point]struct{})
		for p := range ps {
			for _, d := range []direction{north, south, west, east} {
				np := p.Move(d)
				switch m[np] {
				case empty, route, droid, system:
					m[np] = oxygen
					f[np] = struct{}{}
				}
			}
		}
		if len(f) == 0 {
			return 0
		}
		layer()
		return fill(f) + 1
	}
	m[s] = oxygen
	min := fill(map[point]struct{}{s:{}})
	layer()
	return strconv.Itoa(min), nil
}

func aStar(start, goal point, m map[point]terrain) []point {
	visited := make(map[point]struct{})
	toVisit := make(map[point]struct{})
	toVisit[start] = struct{}{}
	for p, t := range m {
		if t == empty || t == system {
			toVisit[p] = struct{}{}
		}
	}
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
		var ps []point
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
			score := scores[current] + distance(current, neighbor)
			cameFrom[neighbor] = current
			scores[neighbor] = score
			estimatedScores[neighbor] = score + heuristic(neighbor, goal)
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
