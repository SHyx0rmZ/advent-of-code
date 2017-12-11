package day11

import (
	"bytes"
	"fmt"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (problem) PartOne(data []byte) (string, error) {
	var p position
	for _, step := range bytes.Split(data, []byte(",")) {
		p = p.Add(directions[string(bytes.TrimSpace(step))])
	}
	return fmt.Sprintf("%+v", p.Distance(position{})), nil
}

func (problem) PartTwo(data []byte) (string, error) {
	var p position
	var m float64
	for _, step := range bytes.Split(data, []byte(",")) {
		p = p.Add(directions[string(bytes.TrimSpace(step))])
		if p.Distance(position{}) > m {
			m = p.Distance(position{})
		}
	}
	return fmt.Sprintf("%+v", m), nil
}
