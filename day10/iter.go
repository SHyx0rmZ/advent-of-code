package day10

import "github.com/SHyx0rmZ/advent-of-code/pkg/lib"

type Iter struct {
	Direction
	Set lib.GenSet
}

func (i *Iter) enter(n int) {
	i.Toggle()
	c := &counter{
		make(chan struct{}),
		n - 1,
	}
	i.Set.Add(c)
	go c.Run()
}

func (i *Iter) exit() {
	for _, c := range i.Set.Elements() {
		if c.(*counter).Done() {
			i.Set.Delete(c)
			i.Toggle()
		}
	}
}

func (i *Iter) Next(m *Mark) *Mark {
	if m.Fwd != 0 {
		i.enter(m.Fwd)
	}
	m = i.Direction.Next(m)
	i.exit()
	return m
}

func (i *Iter) Prev(m *Mark) *Mark {
	if m.Bkwd != 0 {
		i.enter(m.Bkwd)
	}
	m = i.Direction.Prev(m)
	i.exit()
	return m
}
