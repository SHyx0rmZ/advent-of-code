package day10

type Iter struct {
	Direction
}

func (i *Iter) Next(m *Mark) *Mark {
	if i.enterNextToggle(m) {
		i.Toggle()
	}
	m = i.Direction.Next(m)
	if i.exitNextToggle(m) {
		i.Toggle()
	}
	return m
}

func (i *Iter) Prev(m *Mark) *Mark {
	if m.ToggleBackward {
		i.Toggle()
	}
	m = i.Direction.Prev(m)
	if m.ToggleBackward2 {
		i.Toggle()
	}
	return m
}

func (i Iter) enterNextToggle(m *Mark) bool {
	if i.Direction == Backward {
		return m.TNBS
	}
	return m.TNFS
}

func (i Iter) exitNextToggle(m *Mark) bool {
	if i.Direction == Backward {
		return m.TNFE
	}
	return m.TNBE
}
