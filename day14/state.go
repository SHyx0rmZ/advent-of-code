package day14

import "container/ring"

type state struct {
	List  *ring.Ring
	First *ring.Ring
	skip  int
}

func State(n int) *state {
	r := ring.New(n)
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	return &state{
		List:  r,
		First: r,
	}
}

func (s *state) reverse(n int) {
	r := s.List
	for i := 0; i < n/2; i++ {
		pl := r.Move(i)
		pr := r.Move(n - 1 - i)
		pl.Value, pr.Value = pr.Value, pl.Value
	}
}

func (s *state) round(lengths []int) {
	for _, l := range lengths {
		s.reverse(int(l))
		s.List = s.List.Move(int(l) + s.skip)
		s.skip++
	}
}

func (s *state) dense() []int {
	v := make([]int, (s.First.Len()+(16-s.First.Len()%16)%16)/16)
	r := s.First
	for i := 0; i < s.First.Len(); i++ {
		v[i/16] ^= r.Value.(int)
		r = r.Next()
	}
	return v
}
