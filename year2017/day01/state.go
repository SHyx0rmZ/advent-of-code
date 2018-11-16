package day01

import "container/ring"

type state struct {
	Ring *Ring
	Func func(*state) func(*ring.Ring)

	sum int
}

func (s *state) Count() int {
	if s.Func != nil {
		s.Ring.Do(s.Func(s))
	}
	return s.sum
}

func (s *state) half() func(*ring.Ring) {
	return func(r *ring.Ring) {
		c := r.Move(r.Len() / 2)
		if c.Value == r.Value {
			s.sum += int(r.Value.(rune) - '0')
		}
	}
}

func (s *state) next() func(*ring.Ring) {
	return func(r *ring.Ring) {
		c := r.Next()
		if c.Value == r.Value {
			s.sum += int(r.Value.(rune) - '0')
		}
	}
}
