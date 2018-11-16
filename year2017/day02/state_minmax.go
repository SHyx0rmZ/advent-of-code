package day02

import "math"

type stateMinMax struct {
	High int
	Low  int
}

func (s *stateMinMax) Column(n int) {
	if s != nil {
		if n > s.High {
			s.High = n
		}
		if n < s.Low {
			s.Low = n
		}
	}
}

func (s *stateMinMax) Row() int {
	diff := s.High - s.Low
	s.High = math.MinInt64
	s.Low = math.MaxInt64
	return diff
}
