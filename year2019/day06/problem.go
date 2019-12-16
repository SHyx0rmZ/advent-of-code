package day06

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	orbits, err := p.parse(r)
	if err != nil {
		return "", err
	}

	var i int
	for o := range orbits {
		i += dist(orbits, o)
	}

	return fmt.Sprintf("%d", i), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	orbits, err := p.parse(r)
	if err != nil {
		return "", err
	}

	y := orbits["YOU"]
	s := orbits["SAN"]

	return fmt.Sprintf("%d", dist(orbits, y)+dist(orbits, s)-2*dist(orbits, common(orbits, y, s))), nil
}

func common(m map[string]string, o1, o2 string) string {
	if o1 == o2 {
		return o1
	}
	s := make(map[string]struct{})
	for c := m[o1]; c != "COM"; c = m[c] {
		s[c] = struct{}{}
	}
	for c := m[o2]; c != "COM"; c = m[c] {
		if _, ok := s[c]; ok {
			return c
		}
	}
	return "COM"
}

func dist(m map[string]string, o string) int {
	i := 1
	for c := m[o]; c != "COM"; c = m[c] {
		i++
	}
	return i
}

func (problem) parse(r io.Reader) (map[string]string, error) {
	orbits := make(map[string]string)
	s := bufio.NewScanner(r)
	for s.Scan() {
		objects := strings.Split(s.Text(), ")")
		if _, ok := orbits[objects[1]]; ok {
			panic(objects[1])
		}
		orbits[objects[1]] = objects[0]
	}
	return orbits, nil
}
