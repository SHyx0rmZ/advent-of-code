package day18

import (
	"fmt"
)

type grid struct {
	M    [2]map[point]tile
	W, H int
	I    int
}

func (g *grid) add(p point, t tile) {
	g.M[g.I][p] = t
	if p.X > g.W {
		g.W = p.X
	}
	if p.Y > g.H {
		g.H = p.Y
	}
}

func (g *grid) count(p point) (nTrees, nLumberyards int) {
	for _, a := range []point{
		{p.X - 1, p.Y - 1},
		{p.X - 1, p.Y + 0},
		{p.X - 1, p.Y + 1},
		{p.X + 0, p.Y - 1},
		{p.X + 0, p.Y + 1},
		{p.X + 1, p.Y - 1},
		{p.X + 1, p.Y + 0},
		{p.X + 1, p.Y + 1},
	} {
		t := g.M[g.I][a]
		switch t {
		case trees:
			nTrees++
		case lumberyard:
			nLumberyards++
		}
	}
	return nTrees, nLumberyards
}

func (g *grid) update() {
	ni := (g.I + 1) % 2
	for y := 0; y <= g.H; y++ {
		for x := 0; x <= g.W; x++ {
			p := point{x, y}
			t, l := g.count(p)
			switch g.M[g.I][p] {
			case ground:
				if t >= 3 {
					g.M[ni][p] = trees
				} else {
					g.M[ni][p] = ground
				}
			case trees:
				if l >= 3 {
					g.M[ni][p] = lumberyard
				} else {
					g.M[ni][p] = trees
				}
			case lumberyard:
				if t >= 1 && l >= 1 {
					g.M[ni][p] = lumberyard
				} else {
					g.M[ni][p] = ground
				}
			}
		}
	}
	g.I = ni
}

func (g *grid) print() {
	for y := 0; y <= g.H; y++ {
		for x := 0; x <= g.W; x++ {
			fmt.Print(g.M[g.I][point{x, y}])
		}
		fmt.Println()
	}
}

func (g *grid) score() int {
	var t, l int
	for y := 0; y <= g.H; y++ {
		for x := 0; x <= g.W; x++ {
			switch g.M[g.I][point{x, y}] {
			case trees:
				t++
			case lumberyard:
				l++
			}
		}
	}
	return t * l
}

func (g *grid) copy() map[point]tile {
	c := make(map[point]tile)
	for p, t := range g.M[g.I] {
		c[p] = t
	}
	return c
}

func (g *grid) equal(c map[point]tile) bool {
	for p, t := range c {
		if g.M[g.I][p] != t {
			return false
		}
	}
	return true
}
