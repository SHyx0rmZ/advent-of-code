package day14

import (
	"bytes"
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

type node struct {
	bool
	int
}

func (p problem) PartTwo(data []byte) (string, error) {
	m := [128][128]node{}
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
			if m[x][y].bool && m[x][y].int == 0 {
				i += 1
			}
			det(&i, &m, x, y)
		}
	}
	return fmt.Sprintf("%d", i), nil
}

func (problem) parse(data []byte) ([]int, error) {
	return nil, nil
}

func det(i *int, m *[128][128]node, x, y int) {
	if  m[x][y].int != 0 { return }
	if !m[x][y].bool     { return }

	m[x][y].int = *i

	if x > 0   { det(i, m, x-1, y) }
	if y > 0   { det(i, m, x, y-1) }
	if x < 127 { det(i, m, x+1, y) }
	if y < 127 { det(i, m, x, y+1) }
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