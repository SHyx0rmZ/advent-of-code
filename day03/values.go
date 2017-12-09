package day03

var deltas = []position{
	{X: -1, Y: -1},
	{X: -1, Y: +0},
	{X: -1, Y: +1},
	{X: +0, Y: -1},
	{X: +0, Y: +1},
	{X: +1, Y: -1},
	{X: +1, Y: +0},
	{X: +1, Y: +1},
}

type values struct {
	M map[position]int
}

func (v *values) calculate(p position) int {
	if v == nil {
		return 0
	}
	if v.M == nil {
		v.M = map[position]int{
			{X: 0, Y: 0}: 1,
		}
	}
	s, ok := v.M[p]
	if !ok {
		for _, d := range deltas {
			if a, ok := v.M[p.Add(d)]; ok {
				s += a
			}
		}
		v.M[p] = s
	}
	return s
}
