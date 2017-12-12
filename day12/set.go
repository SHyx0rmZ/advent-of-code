package day12

type set struct {
	M map[int]struct{}
}

func (s *set) Add(i int) {
	s.M[i] = struct{}{}
}

func (s *set) Contains(i int) bool {
	_, ok := s.M[i]
	return ok
}

func (s *set) Delete(i int) {
	delete(s.M, i)
}
