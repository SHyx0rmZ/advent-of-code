package lib

import "reflect"

type set struct {
	m map[interface{}]struct{}
	k reflect.Kind
}

func Set() *set {
	return &set{
		m: make(map[interface{}]struct{}),
	}
}

func (s *set) Add(v interface{}) {
	k := reflect.ValueOf(v).Kind()
	if s.k == reflect.Invalid {
		s.k = k
	}
	if s.k != k {
		panic("adding a " + k.String() + " to a set of " + s.k.String())
	}
	s.m[v] = struct{}{}
}

func (s set) Contains(v interface{}) bool {
	_, ok := s.m[v]
	return ok
}

func (s *set) Delete(v interface{}) {
	delete(s.m, v)
}

func (s *set) Elements() []interface{} {
	var es []interface{}
	for e := range s.m {
		es = append(es, e)
	}
	return es
}

/*
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
*/
