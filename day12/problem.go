package day12

import (
	"bytes"
	"fmt"
	"strconv"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (problem) PartOne(data []byte) (string, error) {
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
				return "", err
			}
			ids = append(ids, id)
		}
		id, err := strconv.Atoi(string(parts[0]))
		if err != nil {
			return "", err
		}
		ps = append(ps, program{
			ID:    id,
			Pipes: ids,
		})
	}
	cy := set{M: make(map[int]struct{})}
	cn := set{M: make(map[int]struct{})}
	for _, p := range ps {
		cn.Add(p.ID)
	}
	mod := true
	for mod {
		mod = false
		for _, p := range ps {
			if p.ID == 0 && cn.Contains(p.ID) {
				cn.Delete(p.ID)
				cy.Add(p.ID)
				mod = true
			}
			for _, id := range p.Pipes {
				if cy.Contains(id) && cn.Contains(p.ID) {
					cn.Delete(p.ID)
					cy.Add(p.ID)
					mod = true
				}
			}
		}
	}
	return fmt.Sprintf("%d", len(cy.M)), nil
}

func (problem) PartTwo(data []byte) (string, error) {
	return "", nil
}

type program struct {
	ID    int
	Pipes []int
}
