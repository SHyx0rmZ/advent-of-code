package day06

type set struct {
	M map[string]int
}

func (s *set) Add(i string) {
	s.M[i] = len(s.M)
}

func (s *set) Contains(i string) bool {
	_, ok := s.M[i]
	return ok
}

func (s *set) LoopSize(i string) int {
	return len(s.M) - s.M[i]
}
