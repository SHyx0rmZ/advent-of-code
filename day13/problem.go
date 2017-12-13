package day13

import (
	"bytes"
	"fmt"
	"strconv"
)

type layer struct {
	Depth int
	Range int
}

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	ls, err := p.parse(data)
	if err != nil {
		return "", err
	}
	var s int
	for _, l := range ls {
		if l.Depth%((l.Range-1)*2) == 0 {
			s += l.Depth * l.Range
		}
	}
	return fmt.Sprintf("%d", s), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	ls, err := p.parse(data)
	if err != nil {
		return "", err
	}
	var d int
	for {
		var c bool
		for _, l := range ls {
			if (l.Depth+d)%((l.Range-1)*2) == 0 {
				c = true
				break
			}
		}
		if !c {
			break
		}
		d++
	}
	return fmt.Sprintf("%d", d), nil
}

func (problem) parse(data []byte) ([]layer, error) {
	var ls []layer
	var err error
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		parts := bytes.Split(line, []byte(": "))
		var l layer
		l.Depth, err = strconv.Atoi(string(parts[0]))
		if err != nil {
			return nil, err
		}
		l.Range, err = strconv.Atoi(string(parts[1]))
		if err != nil {
			return nil, err
		}
		ls = append(ls, l)
	}
	return ls, nil
}
