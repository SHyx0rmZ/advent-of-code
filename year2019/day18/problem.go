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
	"strconv"
	"sync"

	aoc "github.com/SHyx0rmZ/advent-of-code"
)

type between [2]byte

//type between struct{ l, r byte }

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
	//pprof.StartCPUProfile(os.Stderr)
	//defer pprof.StopCPUProfile()
	//timer := make(chan struct{})
	//time.AfterFunc(10*time.Second, func() { close(timer) })
	b := make(board)
	keys := make(map[byte]point)
	doors := make(map[byte]point)
	var entryPoint point
	var y int
	s := bufio.NewScanner(r)
	for s.Scan() {
		for x, c := range s.Bytes() {
			p := point{x, y}
			switch {
			case c >= 'a' && c <= 'z':
				keys[c] = p
			case c >= 'A' && c <= 'Z':
				doors[c] = p
			case c == '@':
				entryPoint = p
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
	//for c := range keys {
	//	i := image.NewPaletted(b.Bounds(), g.Config.ColorModel.(color.Palette))
	//	i.SetColorIndex(keys[c].X, keys[c].Y, 3)
	//	i.SetColorIndex(doors[c^0x20].X, doors[c^0x20].Y, 4)
	//	g.Image = append(g.Image, i)
	//	g.Delay = append(g.Delay, 100)
	//	g.Disposal = append(g.Disposal, gif.DisposalNone)
	//}

	seen := make(map[point]struct{})
	var reduce func([]point)
	reduce = func(points []point) {
		if len(points) == 0 {
			return
		}
		var targets pointSlice
		for _, p := range points {
			if _, ok := seen[p]; ok {
				continue
			}
			seen[p] = struct{}{}
			for _, newPoint := range []point{
				{p.X, p.Y + 1},
				{p.X, p.Y - 1},
				{p.X + 1, p.Y},
				{p.X - 1, p.Y},
			} {
				if b[newPoint] != '#' {
					targets = append(targets, newPoint)
				}
			}
		}
		reduce(targets)
		b.Print(g)
		for _, p := range points {
			if b[p] != '.' {
				continue
			}
			var n int
			for _, newPoint := range []point{
				{p.X, p.Y + 1},
				{p.X, p.Y - 1},
				{p.X + 1, p.Y},
				{p.X - 1, p.Y},
			} {
				if b[newPoint] == '#' {
					n++
				}
			}
			// Three of p's neighbors are walls, so p
			// is a dead end. Convert it to a wall.
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
	reduce([]point{entryPoint})
	b.Print(g)
	routes := make(map[between][]point)
	var routesFast [27][27][]point
	for key1 := range keys {
		for key2 := range keys {
			if key2 == key1 {
				continue
			}
			route, ok := routes[between{key2, key1}]
			if ok {
				continue
			}
			route = aStar(keys[key1], keys[key2], b)
			if len(route) == 0 {
				return "", fmt.Errorf("no path")
			}
			routes[between{key1, key2}] = route
			routesFast[key1|0x20-'`'][key2|0x20-'`'] = route
			routesFast[key2|0x20-'`'][key1|0x20-'`'] = route
			//fmt.Println(string(key1), " => ", string(key2), " :", len(doors))
		}
	}

	for key2 := range keys {
		route, ok := routes[between{key2, '@'}]
		if ok {
			continue
		}
		route = aStar(entryPoint, keys[key2], b)
		if len(route) == 0 {
			return "", fmt.Errorf("no path")
		}
		routes[between{'@', key2}] = route
	}
	//doorsPerRoute := make(map[between]int32)
	var doorsPerRoute [27][27]int32
	for key, route := range routes {
		for _, point := range route {
			if 'A' <= b[point] && b[point] <= 'Z' {
				doorsPerRoute[key[0]|0x20-'`'][key[1]|0x20-'`'] |= 1 << (b[point] - 'A')
			}
		}
	}
	// from here on, just try all combinations?
	// limit choices by keys picked up (0 initially)
	var rspk [26][]between
	for k := range routes {
		l := k[0] - 'a'
		//l := int32(1) << (k.l - 'a')
		r := k[1] - 'a'
		//r := int32(1) << (k.r - 'a')
		if k[0] != '@' {
			//if k.l != '@' {
			rspk[l] = append(rspk[l], k)
			//rspk[l] = append(rspk[l], k)
		}
		if k[1] != '@' {
			//if k.r != '@' {
			rspk[r] = append(rspk[r], k)
		}
	}
	var lookup [27][27][27]byte
	for points := range routes {
		lookup[points[0]|0x20-'`'][points[1]|0x20-'`'][points[0]|0x20-'`'] = points[1]
		lookup[points[0]|0x20-'`'][points[1]|0x20-'`'][points[1]|0x20-'`'] = points[0]
	}
	goal := func(s byte, p between) byte {
		//return lookup[[3]byte{
		//	s,
		//	p[0],
		//	p[1],
		//}]
		return lookup[p[0]|0x20-'`'][p[1]|0x20-'`'][s|0x20-'`']
	}
	mp := sync.Pool{New: func() interface{} {
		return make([]between, 0, 26*27)
	}}
	possible := func(start byte, _ []between, keys int32) []between {
		choices := mp.Get().([]between)
		for i := 0; i < 26; i++ {
			if (start-'a') != byte(i) && keys&(1<<i) != 0 {
				continue
			}
			for _, points := range rspk[i] {
				switch {
				case points[0] == start:
					//case points.l == start:
					if keys&(1<<(points[1]-'a')) != 0 {
						//if keys&(1<<(points.r-'a')) != 0 {
						continue
					}
				case points[1] == start:
					//case points.r == start:
					if keys&(1<<(points[0]-'a')) != 0 {
						//if keys&(1<<(points.l-'a')) != 0 {
						continue
					}
				default:
					continue
				}
				if doorsPerRoute[points[0]|0x20-'`'][points[1]|0x20-'`']&^keys != 0 {
					continue
				}
				choices = append(choices, points)
			}
		}
		return choices
	}
	copyThings := func(start, goal byte, doorsPerRoute []between, collected int32) ([]between, int32) {
		//newDoorsPerRoute := make([][2]byte, 0, len(doorsPerRoute))
		//for _, db := range doorsPerRoute {
		//	if db[0] == start || db[1] == start {
		//		continue
		//	}
		//	newDoorsPerRoute = append(newDoorsPerRoute, db)
		//}
		newDoorsPerRoute := doorsPerRoute
		newCollected := collected
		newCollected |= 1 << (goal - 'a')
		return newDoorsPerRoute, newCollected
	}
	//var xx int
	var shortest func(start byte, paths []between, collected int32, cost int) int
	shortest = func(start byte, paths []between, collected int32, cost int) int {
		if collected == 0b11111111111111111111111111 {
			return cost
		}
		//select {
		//case <-timer:
		//	fmt.Println(xx)
		//	return cost
		//default:
		//	xx++
		//	fmt.Println(bits.OnesCount32(uint32(collected)))
		//}
		choices := possible(start, paths, collected)
		if len(choices) == 0 {
			return math.MaxInt32
		}
		fewerPaths, biggerKeyCollection := copyThings(start, goal(start, choices[0]), paths, collected)
		minCost := shortest(goal(start, choices[0]), fewerPaths, biggerKeyCollection, len(routesFast[choices[0][0]|0x20-'`'][choices[0][1]|0x20-'`'])+cost)
		for _, choice := range choices[1:] {
			fewerPaths, biggerKeyCollection = copyThings(start, goal(start, choice), paths, collected)
			currentCost := shortest(goal(start, choice), fewerPaths, biggerKeyCollection, len(routesFast[choice[0]|0x20-'`'][choice[1]|0x20-'`'])+cost)
			if currentCost < minCost {
				minCost = currentCost
			}
		}
		mp.Put(choices[:0])
		return minCost
	}
	var paths []between
	for key := range routes {
		paths = append(paths, key)
	}
	return strconv.Itoa(shortest('@', paths, 0, 0)), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	return "", nil
}
