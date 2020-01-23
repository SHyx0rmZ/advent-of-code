package day19

type point struct {
	X, Y int
}

func (p point) Less(o point) bool {
	if p.Y == o.Y {
		return p.X < o.X
	}
	return p.Y < o.Y
}
