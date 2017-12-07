package day07

type set struct {
	M map[string]struct{}
}

func (s *set) Add(i string) {
	s.M[i] = struct{}{}
}

func (s *set) Del(i string) {
	delete(s.M, i)
}

func (s *set) Elements() []string {
	var es []string
	for e := range s.M {
		es = append(es, e)
	}
	return es
}
