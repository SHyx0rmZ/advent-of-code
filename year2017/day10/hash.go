package day10

import (
	"fmt"
	h "hash"
)

type digest struct {
	n     int
	nodes []*node
	skip  int
}

func (d *digest) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func (d0 *digest) Sum(b []byte) []byte {
	//d := *d0
	var s [16]byte
	return append(b, s[:]...)
}

func (d *digest) Reset() {
	d.skip = 0
}

func (d *digest) Size() int {
	return 16
}

func (d *digest) BlockSize() int {
	return 1
}

func New(n int) h.Hash {
	d := new(digest)
	d.n = n
	d.Reset()
	return d
}

type node struct {
	Value interface{}

	next *node
	prev *node

	toggle bool
}

type hash struct {
	length int
	nodes  *node
	iter   iter
	skip   int
}

/*
	toggle iterator?
	> replace iterator

	swap pointers, toggle iterator
	toggle toggle
	2 ops / length
*/

func newHash(l int) hash {
	n := &node{
		Value: 0,
	}
	p := n
	for i := 1; i < l; i++ {
		p.next = &node{
			Value: i,
			prev:  p,
		}
		p = p.next
	}
	p.next = n
	n.prev = p
	return hash{
		length: l,
		nodes:  n,
		iter:   next{},
	}
}

func (h hash) Digest() string {
	d := [16]byte{}
	s := ""
	for i := 0; i < len(d); i++ {
		s += fmt.Sprintf("%02x", d[i])
	}
	return s
}

func (h *hash) move(d int) *node {
	n := h.nodes
	b := h.iter
	for i := 0; i < d; i++ {
		n = h.Next(*n)
	}
	h.iter = b
	return n
}

func (h *hash) Next(o node) *node {
	if o.toggle {
		h.iter.Toggle(h)
	}
	return h.iter.Next(o)
}

func (h *hash) Update(p []byte) {

	//for i := 0; i < h.length; i++ {
	//	fmt.Printf(" %2d", h.nodes.Value)
	//	h.nodes = h.Next(*h.nodes)
	//}
	//fmt.Println()
	for _, b := range p {
		//fmt.Println(b)
		s := h.nodes
		e := h.move(int(b - 1))
		//s.toggle = !s.toggle
		//s.prev.toggle = !s.prev.toggle
		//fmt.Printf("U %p %p\n", s, e)
		if s.toggle {
			prev{}.Update(s, e)
		} else {
			next{}.Update(s, e)
		}
		//s.prev.next, e.next.prev = e.next.prev, s.prev.next
		//e.next, s.prev = s.prev, e.next
		if s.toggle {
			h.iter.Toggle(h)
		}
		h.nodes = e
		//h.move(int(b - 1) + h.skip)
		//h.skip++

		//for i := 0; i < h.length * 3; i++ {
		//	fmt.Printf("%p %+v\n", h.nodes, h.nodes)
		//	h.nodes = h.Next(*h.nodes)
		//}
		//fmt.Println()
		//break
	}
}

type iter interface {
	Next(node) *node
	NextPtr(*node) **node
	PrevPtr(*node) **node
	Update(s *node, e *node)
	Toggle(*hash)
}

type prev struct{}

func (prev) Next(n node) *node {
	return n.prev
}

func (prev) NextPtr(n *node) **node {
	return &n.prev
}

func (prev) PrevPtr(n *node) **node {
	return &n.next
}

func (prev) Update(s *node, e *node) {
	var si iter = prev{}
	var ei iter = prev{}
	//if s.toggle {
	//	si = next{}
	//}
	if e.toggle {
		ei = next{}
	}
	s.toggle = !s.toggle
	(*ei.NextPtr(e)).toggle = !(*ei.NextPtr(e)).toggle
	//e.next.toggle = !e.next.toggle
	//s.toggle = !s.toggle
	//(*si.PrevPtr(s)).toggle = !(*si.PrevPtr(s)).toggle
	//s.toggle = !s.toggle
	//
	//e.toggle = !e.toggle
	//(*ei.NextPtr(e)).toggle = !(*ei.NextPtr(e)).toggle //next
	(*si.PrevPtr(s)).next, (*ei.NextPtr(e)).prev = (*ei.NextPtr(e)).prev, (*si.PrevPtr(s)).next
	*ei.NextPtr(e), *si.PrevPtr(s) = *si.PrevPtr(s), *ei.NextPtr(e)
	//(*ei.PrevPtr(e)).toggle = !(*ei.PrevPtr(e)).toggle
	//(*si.NextPtr(s)).next, (*ei.PrevPtr(e)).prev = (*ei.PrevPtr(e)).prev, (*si.NextPtr(s)).next
	//*ei.PrevPtr(e), *si.NextPtr(s) = *si.NextPtr(s), *ei.PrevPtr(e)
}

func (prev) Toggle(h *hash) {
	h.iter = next{}
}

type next struct{}

func (next) Next(n node) *node {
	return n.next
}

func (next) NextPtr(n *node) **node {
	return &n.next
}

func (next) PrevPtr(n *node) **node {
	return &n.prev
}

func (next) Update(s *node, e *node) {
	//e.toggle = !e.toggle
	//e.next.toggle = !e.next.toggle
	//s.prev.next, e.next.prev = e.next.prev, s.prev.next
	//e.next, s.prev = s.prev, e.next
	var si iter = next{}
	var ei iter = next{}
	//if s.toggle {
	//	si = prev{}
	//}
	if e.toggle {
		ei = prev{}
	}
	e.toggle = !e.toggle
	(*ei.NextPtr(e)).toggle = !(*ei.NextPtr(e)).toggle
	//s.toggle = !s.toggle
	//(*si.PrevPtr(s)).toggle = !(*si.PrevPtr(s)).toggle
	(*si.PrevPtr(s)).next, (*ei.NextPtr(e)).prev = (*ei.NextPtr(e)).prev, (*si.PrevPtr(s)).next
	*ei.NextPtr(e), *si.PrevPtr(s) = *si.PrevPtr(s), *ei.NextPtr(e)
}

func (next) Toggle(h *hash) {
	h.iter = prev{}
}
