package day01

import "container/ring"

type Ring struct {
	*ring.Ring
}

func NewRing(s string) *Ring {
	r := ring.New(len(s))
	for _, c := range s {
		r.Value = c
		r = r.Next()
	}
	return &Ring{r}
}

func (r *Ring) Do(f func(*ring.Ring)) {
	if r != nil {
		f(r.Ring)
		for p := r.Next(); p != r.Ring; p = p.Next() {
			f(p)
		}
	}
}
