package day12

import (
	"bytes"
	"container/ring"
	"fmt"
	"strconv"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	ps, err := p.parse(data)
	if err != nil {
		return "", err
	}
	rs := p.construct(ps)
	return fmt.Sprintf("%d", rs[0].Len()), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	ps, err := p.parse(data)
	if err != nil {
		return "", err
	}
	rs := p.construct(ps)
	for _, r := range rs {
		for i := r.Len() - 1; i > 0; i-- {
			r = r.Next()
			delete(rs, r.Value.(program).ID)
		}
	}
	return fmt.Sprintf("%d", len(rs)), nil
}

func (problem) construct(ps []program) map[int]*ring.Ring {
	rs := make(map[int]*ring.Ring)
	for _, p := range ps {
		rs[p.ID] = &ring.Ring{
			Value: p,
		}
	}
	for _, p := range ps {
		for _, i := range p.Pipes {
			r := rs[p.ID]
			c := false
			r.Do(func(v interface{}) {
				if v.(program).ID == i {
					c = true
				}
			})
			if !c {
				r.Link(rs[i])
			}
		}
	}
	return rs
}

func (problem) parse(data []byte) ([]program, error) {
	var ps []program
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		parts := bytes.Split(line, []byte(" <-> "))
		var ids []int
		for _, n := range bytes.Split(parts[1], []byte(", ")) {
			id, err := strconv.Atoi(string(n))
			if err != nil {
				return nil, err
			}
			ids = append(ids, id)
		}
		id, err := strconv.Atoi(string(parts[0]))
		if err != nil {
			return nil, err
		}
		ps = append(ps, program{
			ID:    id,
			Pipes: ids,
		})
	}
	return ps, nil
}

type program struct {
	ID    int
	Pipes []int
}
