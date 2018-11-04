package day21

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	return p.PartOneWithReader(bytes.NewReader(data))
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	ms, err := p.parse(r)
	if err != nil {
		return "", err
	}

	rules = ms

	m := Matrix{
		0,1,0,
		0,0,1,
		1,1,1,
	}

	for i := 0; i < 5; i++ {
		m = m.Enhance()
	}

	return fmt.Sprintf("%#v", m.Count()), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	return p.PartTwoWithReader(bytes.NewReader(data))
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	ms, err := p.parse(r)
	if err != nil {
		return "", err
	}

	rules = ms

	m := Matrix{
		0,1,0,
		0,0,1,
		1,1,1,
	}

	for i := 0; i < 18; i++ {
		m = m.Enhance()
	}

	return fmt.Sprintf("%#v", m.Count()), nil
}


func offset(c, d, x, y, dx, dy int) int {
	r := (y*c+dy)*(d*c) + (x * c) + dx
	return r
}

func bit(b byte) int {
	switch b {
	case '#':
		return 1
	case '.':
		return 0
	}
	panic("invalid bit byte")
}

type EnhancementRule struct {
	From Matrix
	To   Matrix
}

func (p problem) parse(r io.Reader) ([]EnhancementRule, error) {
	var rs []EnhancementRule
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		bs := scanner.Bytes()
		lp := bytes.Split(bs, []byte(" => "))
		rs = append(rs, EnhancementRule{
			From: MatrixFromBytes(lp[0]),
			To:   MatrixFromBytes(lp[1]),
		})
	}
	return rs, nil
}
