package day10

import (
	"bytes"
	"container/ring"
	"fmt"
	"strconv"
)

type state struct {
	List  *ring.Ring
	First *ring.Ring
	skip  int
}

func State(n int) *state {
	r := NewRing(n)
	return &state{
		List:  r,
		First: r,
	}
}

type problem struct {
	state *state
}

func Problem(n int) *problem {
	return &problem{
		state: State(n),
	}
}

func (p problem) PartOne(data []byte) (string, error) {
	lengths, err := p.parseInts(data)
	if err != nil {
		return "", err
	}
	p.state.round(lengths)
	return fmt.Sprintf("%d", p.state.First.Value.(int)*p.state.First.Next().Value.(int)), nil
}

func (p problem) PartOne2(data []byte, n int) (string, error) {
	lengths, err := p.parseInts(data)
	if err != nil {
		return "", err
	}
	c := Circle{
		Marks: NewMarks(n),
	}
	c.Round(lengths)
	return fmt.Sprintf("%d", c.Marks[0].Value*c.Marks[1].Value), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	lengths, err := p.parseBytes(data)
	if err != nil {
		return "", err
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)
	for i := 0; i < 64; i++ {
		p.state.round(lengths)
	}
	var hash string
	for _, x := range p.state.dense() {
		hash += fmt.Sprintf("%02x", x)
	}
	return hash, nil
}

func (p problem) PartThree(data []byte) (string, error) {
	lengths := data
	//if err != nil {
	//	return "", err
	//}
	lengths = append(lengths, 17, 31, 73, 47, 23)
	hash := newHash(p.state.List.Len())
	for i := 0; i < 64; i++ {
		hash.Update(lengths)
	}
	return hash.Digest(), nil
}

func (problem) parseBytes(data []byte) ([]int, error) {
	var lengths []int
	for _, b := range bytes.TrimSpace(data) {
		lengths = append(lengths, int(b))
	}
	return lengths, nil
}

func (problem) parseInts(data []byte) ([]int, error) {
	var lengths []int
	for _, b := range bytes.Split(data, []byte(",")) {
		l, err := strconv.Atoi(string(bytes.TrimSpace(b)))
		if err != nil {
			return nil, err
		}
		lengths = append(lengths, l)
	}
	return lengths, nil
}

func (s *state) reverse(n int) {
	//s.List = s.List.Prev()
	//r := s.List.Unlink(n)
	//ns := make([]int, n)
	//for i := 0; i < n; i++ {
	//	ns[i] = r.Move(i).Value.(int)
	//}
	//for i := 0; i < n; i++ {
	//	r.Move(n - 1 - i).Value = ns[i]
	//}
	var pl, pr *ring.Ring
	for i := 0; i < n/2; i++ {
		pl = s.List.Move(i)
		pr = s.List.Move(n - 1 - i)
		pl.Value, pr.Value = pr.Value, pl.Value
	}
	//s.List.Link(r)
	//s.List = s.List.Next()
}

func (s *state) round(lengths []int) {
	for _, l := range lengths {
		s.reverse(l)
		s.List = s.List.Move(l + s.skip)
		s.skip++
	}
}

func (s *state) round2(lengths []byte) {
	for _, l := range lengths {
		s.reverse(int(l))
		s.List = s.List.Move(int(l) + s.skip)
		s.skip++
	}
}

func (s *state) dense() []int {
	v := make([]int, (s.First.Len()+(16-s.First.Len()%16)%16)/16)
	r := s.First
	for i := 0; i < s.First.Len(); i++ {
		v[i/16] ^= r.Value.(int)
		r = r.Next()
	}
	return v
}
