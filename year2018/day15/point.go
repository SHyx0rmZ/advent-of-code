package day15

type point struct {
	X, Y int
}

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
	t := p[i]
	p[i] = p[j]
	p[j] = t
}
