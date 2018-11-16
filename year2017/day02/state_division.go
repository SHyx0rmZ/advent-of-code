package day02

type stateDivision struct {
	Numbers []int
}

func (s *stateDivision) Column(n int) {
	s.Numbers = append(s.Numbers, n)
}

func (s *stateDivision) Row() int {
	for i, lp := range s.Numbers {
		if i+1 >= len(s.Numbers) {
			continue
		}
		for _, rp := range s.Numbers[(i + 1):] {
			if lp%rp == 0 {
				s.Numbers = make([]int, 0)
				return lp / rp
			}
			if rp%lp == 0 {
				s.Numbers = make([]int, 0)
				return rp / lp
			}
		}
	}
	s.Numbers = make([]int, 0)
	return 0
}
