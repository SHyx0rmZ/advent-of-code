package day01

import "container/ring"

type Ring ring.Ring

func NewRing(s string) *Ring {
	r := ring.New(len(s))
	for _, c := range s {
		r.Value = c
		r = r.Next()
	}
	return (*Ring)(r)
}

func (r *Ring) Advance(n int) *Ring {
	p := r
	if p != nil {
		for i := 0; i < n; i++ {
			p = p.Next()
		}
	}
	return p
}

func (r *Ring) Do(f func(*Ring)) {
	if r != nil {
		f(r)
		for p := r.Next(); p != r; p = p.Next() {
			f(p)
		}
	}
}

func (r *Ring) Len() int {
	return (*ring.Ring)(r).Len()
}

func (r *Ring) Next() *Ring {
	return (*Ring)((*ring.Ring)(r).Next())
}

func (r *Ring) Prev() *Ring {
	return (*Ring)((*ring.Ring)(r).Prev())
}
