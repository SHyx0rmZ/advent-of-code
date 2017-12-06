package day04

type set struct {
	M map[string]struct{}
}

func (s *set) Add(i string) {
	s.M[i] = struct{}{}
}

func (s *set) Contains(i string) bool {
	_, ok := s.M[i]
	return ok
}
