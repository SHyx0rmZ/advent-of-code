package day14

import (
	"bytes"
	"container/ring"
	"fmt"
	"math/bits"
	"strconv"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	var s int
	for i := 0; i < 128; i++ {
		hash, err := p.hash([]byte(fmt.Sprintf("%s-%d", string(bytes.TrimSpace(data)), i)))
		if err != nil {
			return "", err
		}
		for n := 0; n < 2; n++ {
			hv, err := strconv.ParseUint(hash[(n*16):(n*16+16)], 16, 64)
			if err != nil {
				return "", err
			}
			s += bits.OnesCount64(uint64(hv))
		}
	}
	return fmt.Sprintf("%d", s), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	m := [128][128]struct {
		bool
		int
		v bool
	}{}
	for i := 0; i < 128; i++ {
		hash, err := p.hash([]byte(fmt.Sprintf("%s-%d", string(bytes.TrimSpace(data)), i)))
		if err != nil {
			return "", err
		}
		for n := 0; n < 2; n++ {
			hv, err := strconv.ParseUint(hash[(n*16):(n*16+16)], 16, 64)
			if err != nil {
				return "", err
			}
			for b := 0; b < 64; b++ {
				m[i][b+n*64].bool = ((hv >> (63 - uint(b))) & 1) == 1
			}
		}
	}
	var i int
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			if m[y][x].bool == false {
				continue
			}
			det(&i, &m, x, y)
		}
	}
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			fmt.Printf(" [%5t,%3d]", m[y][x].bool, m[y][x].int)
		}
		fmt.Println()
	}
	return fmt.Sprintf("%d", i-1), nil
}

func (problem) parse(data []byte) ([]int, error) {
	return nil, nil
}

func det(i *int, m *[128][128]struct {
	bool
	int
	v bool
}, x int, y int) int {
	if m[x][y].int != 0 || m[x][y].v {
		return m[x][y].int
	}
	if !m[x][y].bool {
		return 0
	}
	m[x][y].v = true
	var vl, vd, vr, vu int
	if x > 0 {
		vl = det(i, m, x-1, y)
	}
	if y > 0 {
		vd = det(i, m, x, y-1)
	}
	if x < 127 {
		vr = det(i, m, x+1, y)
	}
	if y < 127 {
		vu = det(i, m, x, y+1)
	}
	switch {
	case vl != 0:
		m[x][y].int = vl
	case vd != 0:
		m[x][y].int = vd
	case vr != 0:
		m[x][y].int = vr
	case vu != 0:
		m[x][y].int = vu
	default:
		*i++
		m[x][y].int = *i
	}
	return m[x][y].int
}

type state struct {
	List  *Ring
	First *Ring
	skip  int
}

func State(n int) *state {
	r := NewRing(n)
	return &state{
		List:  r,
		First: r,
	}
}

func (p problem) hash(data []byte) (string, error) {
	lengths, err := p.parseBytes(data)
	if err != nil {
		return "", err
	}
	s := State(256)
	lengths = append(lengths, 17, 31, 73, 47, 23)
	for i := 0; i < 64; i++ {
		s.round(lengths)
	}
	var hash string
	for _, x := range s.dense() {
		hash += fmt.Sprintf("%02x", x)
	}
	return hash, nil
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
	s.List = s.List.Prev()
	r := (*Ring)((*ring.Ring)(s.List).Unlink(n))
	var ns []int
	for i := 0; i < n; i++ {
		ns = append(ns, r.Advance(i).Value.(int))
	}
	for i := 0; i < n; i++ {
		r.Advance(n - 1 - i).Value = ns[i]
	}
	(*ring.Ring)(s.List).Link((*ring.Ring)(r))
	s.List = s.List.Next()
}

func (s *state) round(lengths []int) {
	for _, l := range lengths {
		s.reverse(l)
		s.List = s.List.Advance(l + s.skip)
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
