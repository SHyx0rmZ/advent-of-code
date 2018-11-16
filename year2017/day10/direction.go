package day10

type Direction int

const (
	Forward Direction = iota
	Backward
)

func (d Direction) Next(m *Mark) *Mark {
	if d == Backward {
		return m.Prev
	}
	return m.Next
}

func (d Direction) Prev(m *Mark) *Mark {
	if d == Backward {
		return m.Next
	}
	return m.Prev
}

func (d Direction) String() string {
	if d == Backward {
		return "Backward"
	}
	return "Forward"
}

func (d *Direction) Toggle() {
	switch *d {
	case Forward:
		*d = Backward
	case Backward:
		*d = Forward
	}
}
