package day13

import "C"
import (
	"bufio"
	"fmt"
	"io"
	"sort"

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

type track int

const (
	empty track = iota
	vertical
	horizontal
	junction
	cornerLeft
	cornerRight
	any
)

type direction string

const (
	north direction = "north"
	east  direction = "east"
	south direction = "south"
	west  direction = "west"
)

type behavior string

const (
	left     behavior = "left"
	straight behavior = "straight"
	right    behavior = "right"
)

func (b *behavior) next() {
	switch *b {
	case left:
		*b = straight
	case straight:
		*b = right
	case right:
		*b = left
	}
}

var dbm = map[struct {
	direction
	behavior
}]direction{
	{north, left}:     west,
	{north, straight}: north,
	{north, right}:    east,
	{east, left}:      north,
	{east, straight}:  east,
	{east, right}:     south,
	{south, left}:     east,
	{south, straight}: south,
	{south, right}:    west,
	{west, left}:      south,
	{west, straight}:  west,
	{west, right}:     north,
}

func (d direction) add(b behavior) direction {
	return dbm[struct {
		direction
		behavior
	}{
		d,
		b,
	}]
}

type cart struct {
	point
	direction
	behavior
}

func (c *cart) String() string {
	return fmt.Sprint(*c)
}

var rs = []struct {
	left, top, right, bottom track
	replacement              track
}{
	{horizontal, any, horizontal, any, horizontal},
	{any, vertical, any, vertical, vertical},
	{empty, empty, horizontal, vertical, cornerLeft},
	{horizontal, vertical, empty, empty, cornerLeft},
	{empty, vertical, horizontal, empty, cornerRight},
	{horizontal, empty, empty, vertical, cornerRight},
	{horizontal, empty, any, empty, horizontal},
	{any, empty, horizontal, empty, horizontal},
	{empty, vertical, empty, any, vertical},
	{empty, any, empty, vertical, vertical},
	{horizontal, any, junction, horizontal, horizontal},
	{horizontal, horizontal, junction, any, horizontal},
	{junction, any, horizontal, horizontal, horizontal},
	{junction, horizontal, horizontal, any, horizontal},
	{vertical, vertical, any, junction, vertical},
	{any, vertical, vertical, junction, vertical},
	{vertical, junction, any, vertical, vertical},
	{any, junction, vertical, vertical, vertical},
	{junction, empty, junction, any, horizontal},
	{junction, any, junction, empty, horizontal},
	{empty, junction, any, junction, vertical},
	{any, junction, empty, junction, vertical},
	{horizontal, any, cornerLeft, empty, horizontal},
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	m := make(map[point]track)
	var cs []cart
	var y int
	for s.Scan() {
		for x, c := range s.Text() {
			p := point{x, y}
			switch c {
			case '/':
				m[p] = cornerLeft
			case '\\':
				m[p] = cornerRight
			case '|':
				m[p] = vertical
			case '-':
				m[p] = horizontal
			case '+':
				m[p] = junction
			case 'v':
				cs = append(cs, cart{p, south, left})
			case '>':
				cs = append(cs, cart{p, east, left})
			case '^':
				cs = append(cs, cart{p, north, left})
			case '<':
				cs = append(cs, cart{p, west, left})
			}
		}
		y++
	}
	sort.Slice(cs, func(i, j int) bool {
		if cs[i].Y == cs[j].Y {
			return cs[i].X < cs[j].X
		}
		return cs[i].Y < cs[j].Y
	})
	for _, c := range cs {
		l := m[point{c.X - 1, c.Y}]
		r := m[point{c.X + 1, c.Y}]
		t := m[point{c.X, c.Y - 1}]
		b := m[point{c.X, c.Y + 1}]
		var found bool
		for _, rp := range rs {
			if (rp.left == l || rp.left == any) && (rp.top == t || rp.top == any) && (rp.right == r || rp.right == any) && (rp.bottom == b || rp.bottom == any) {
				m[c.point] = rp.replacement
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Sprintf("%#v %d %d %d %d\n%v", c, l, r, t, b, cs))
		}
	}
	for {
		sort.Slice(cs, func(i, j int) bool {
			if cs[i].Y == cs[j].Y {
				return cs[i].X < cs[j].X
			}
			return cs[i].Y < cs[j].Y
		})
		for i, c := range cs {
			var np point
			switch c.direction {
			case north:
				np = point{c.X, c.Y - 1}
			case south:
				np = point{c.X, c.Y + 1}
			case west:
				np = point{c.X - 1, c.Y}
			case east:
				np = point{c.X + 1, c.Y}
			}
			switch m[np] {
			case junction:
				c.direction = c.direction.add(c.behavior)
				c.behavior.next()
			case cornerLeft:
				switch c.direction {
				case east:
					c.direction = north
				case south:
					c.direction = west
				case west:
					c.direction = south
				case north:
					c.direction = east
				}
			case cornerRight:
				switch c.direction {
				case north:
					c.direction = west
				case east:
					c.direction = south
				case south:
					c.direction = east
				case west:
					c.direction = north
				}
			case empty:
				panic("bug")
			}
			c.point = np
			for oi, oc := range cs {
				if c.point == oc.point && i != oi {
					return fmt.Sprint(c.point), nil
				}
			}
			cs[i] = c
		}
		fmt.Println(cs)
	}
	return "", nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	m := make(map[point]track)
	var cs []*cart
	var y int
	for s.Scan() {
		for x, c := range s.Text() {
			p := point{x, y}
			switch c {
			case '/':
				m[p] = cornerLeft
			case '\\':
				m[p] = cornerRight
			case '|':
				m[p] = vertical
			case '-':
				m[p] = horizontal
			case '+':
				m[p] = junction
			case 'v':
				cs = append(cs, &cart{p, south, left})
			case '>':
				cs = append(cs, &cart{p, east, left})
			case '^':
				cs = append(cs, &cart{p, north, left})
			case '<':
				cs = append(cs, &cart{p, west, left})
			}
		}
		y++
	}
	sort.Slice(cs, func(i, j int) bool {
		if cs[i].Y == cs[j].Y {
			return cs[i].X < cs[j].X
		}
		return cs[i].Y < cs[j].Y
	})
	for _, c := range cs {
		l := m[point{c.X - 1, c.Y}]
		r := m[point{c.X + 1, c.Y}]
		t := m[point{c.X, c.Y - 1}]
		b := m[point{c.X, c.Y + 1}]
		var found bool
		for _, rp := range rs {
			if (rp.left == l || rp.left == any) && (rp.top == t || rp.top == any) && (rp.right == r || rp.right == any) && (rp.bottom == b || rp.bottom == any) {
				m[c.point] = rp.replacement
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Sprintf("%#v %d %d %d %d\n%v", c, l, r, t, b, cs))
		}
	}
	var cc int
	for _, c := range cs {
		if c != nil {
			cc++
		}
	}
	for cc > 1 {
		sort.Slice(cs, func(i, j int) bool {
			if cs[i] == nil && cs[j] != nil {
				return false
			}
			if cs[i] != nil && cs[j] == nil {
				return true
			}
			if cs[i] == nil && cs[j] == nil {
				return false
			}
			if cs[i].Y == cs[j].Y {
				return cs[i].X < cs[j].X
			}
			return cs[i].Y < cs[j].Y
		})
		for i, c := range cs {
			if c == nil {
				continue
			}
			var np point
			switch c.direction {
			case north:
				np = point{c.X, c.Y - 1}
			case south:
				np = point{c.X, c.Y + 1}
			case west:
				np = point{c.X - 1, c.Y}
			case east:
				np = point{c.X + 1, c.Y}
			}
			switch m[np] {
			case junction:
				c.direction = c.direction.add(c.behavior)
				c.behavior.next()
			case cornerLeft:
				switch c.direction {
				case east:
					c.direction = north
				case south:
					c.direction = west
				case west:
					c.direction = south
				case north:
					c.direction = east
				}
			case cornerRight:
				switch c.direction {
				case north:
					c.direction = west
				case east:
					c.direction = south
				case south:
					c.direction = east
				case west:
					c.direction = north
				}
			case empty:
				panic("bug")
			}
			fmt.Println(np)
			c.point = np
			for oi, oc := range cs {
				if oc == nil {
					continue
				}
				if c.point == oc.point && i != oi {
					cs[i] = nil
					cs[oi] = nil
					break
				}
			}
		}
		fmt.Println(cs)
		cc = 0
		for _, c := range cs {
			if c != nil {
				cc++
			}
		}
	}
	return "", nil
}
