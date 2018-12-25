package day15

import (
	"fmt"
	"sort"
	"time"
)

type stateFn func() stateFn

type caves struct {
	Width  int
	Height int
	Map    map[point]*tile
	Units  unitSlice
}

func (c *caves) AddTerrain(p point, t terrain) {
	c.Map[p] = &tile{
		terrain: t,
	}
}

func (c *caves) AddUnit(p point, f faction) {
	u := &unit{
		point:   p,
		faction: f,
		hp:      200,
	}
	c.Map[p] = &tile{
		terrain: cave,
		unit:    u,
	}
	c.Units = append(c.Units)
}

func (c *caves) CleanExtra() {
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			c.Map[point{x, y}].extra = nil
		}
	}
}

func (c *caves) SetExtra(p point, b byte) {
	c.Map[p].extra = &b
}

func (c *caves) RemoveUnit(i int) {
	u := c.Units[i]
	c.Units = append(c.Units[:i], c.Units[i+1:]...)
	c.Map[u.point].unit = nil
}

func (c *caves) findTargets(i int) stateFn {
	return func() stateFn {
		//print(m, w, h, us)
		if i >= len(c.Units) {
			return nil
		}
		//cleanExtra(m, w, h)
		es := c.Units.EnemiesOf(c.Units[i])
		if len(es) == 0 {
			return end
		}
		return c.findOpenSquares(i, es)
	}
}

func (c *caves) findOpenSquares(i int, es unitSlice) stateFn {
	return func() stateFn {
		var sq []point
		ff := func(x, y int) {
			p := point{x, y}
			t, ok := c.Map[p]
			if !ok {
				return
			}
			if t.terrain == cave && t.unit == nil {
				sq = append(sq, p)
				c.SetExtra(p, '?')
				c.Map[p] = t
			}
		}
		for _, e := range es {
			ff(e.X, e.Y-1)
			ff(e.X-1, e.Y)
			ff(e.X+1, e.Y)
			ff(e.X, e.Y+1)
		}
		return c.findAttackable(i, es, sq)
	}
}

func (c *caves) findAttackable(i int, es unitSlice, sq []point) sf {
	return func(m map[point]*tile, w, h int, us unitSlice) sf {
		u := us[i]
		at := us.UnitsAttackableBy(u)
		if len(sq) == 0 && len(at) == 0 {
			return findTargets(i + 1)
		}
		if len(at) > 0 {
			return attack(i, es, at)
		}
		return findReachable(i, es, sq)
	}
}

func (c *caves) do() (string, error) {
	sort.Sort(c.Units)
	var t int
	combat = true
	for combat {
		for sf := findTargets(0); sf != nil; sf = sf(c.Map, c.Width, c.Height, c.Units) {
			for _, au := range c.Units {
				if au.hp <= 0 {
					var ui int
					for i, u := range c.Units {
						if u == au {
							ui = i
							break
						}
					}
					c.Units = append(c.Units[:ui], c.Units[ui+1:]...)
					c.Map[au.point].unit = nil
				}

			}
		}
		t++
		fmt.Printf("\033[%d;%dH%d   ", 1, c.Width+3, t)
		time.Sleep(delay)
		gs := 0
		es := 0
		sort.Sort(c.Units)
		for _, u := range c.Units {
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
	for _, u := range c.Units {
		fmt.Println(*u)
		hp += u.hp
	}
	return fmt.Sprintf("%d %d %d\n", t, hp, t*hp), nil
}
