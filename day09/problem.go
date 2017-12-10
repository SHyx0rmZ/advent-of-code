package day09

import "strconv"

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (problem) PartOne(data []byte) (string, error) {
	m := Machine()
	err := m.Run(data)
	return strconv.Itoa(m.Score()), err
}

func (problem) PartTwo(data []byte) (string, error) {
	m := Machine()
	err := m.Run(data)
	return strconv.Itoa(m.Ignored()), err
}
