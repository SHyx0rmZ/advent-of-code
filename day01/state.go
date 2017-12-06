package day01

type state struct {
	Ring *Ring
	Func func(*state) func(*Ring)

	sum int
}

func (s *state) Count() int {
	if s.Func != nil {
		s.Ring.Do(s.Func(s))
	}
	return s.sum
}

func (s *state) half() func(*Ring) {
	return func(r *Ring) {
		c := r.Advance(r.Len() / 2)
		if c.Value == r.Value {
			s.sum += int(r.Value.(rune) - '0')
		}
	}
}

func (s *state) next() func(*Ring) {
	return func(r *Ring) {
		c := r.Next()
		if c.Value == r.Value {
			s.sum += int(r.Value.(rune) - '0')
		}
	}
}
