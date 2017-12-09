package day03

type position struct {
	X int
	Y int
}

func (p position) Add(o position) position {
	return position{
		X: p.X + o.X,
		Y: p.Y + o.Y,
	}
}
