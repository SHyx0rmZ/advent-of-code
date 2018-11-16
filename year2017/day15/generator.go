package day15

type generator struct {
	State  uint
	Stride uint
	Check  uint
}

func (g *generator) Step() uint {
	v := g.step()
	for (v % g.Check) != 0 {
		v = g.step()
	}
	return v
}

func (g *generator) step() uint {
	g.State = (g.State * g.Stride) % 2147483647
	return g.State
}
