package day09

import (
	"errors"
	"strconv"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (problem) PartOne(data []byte) (string, error) {
	m := Machine()
	score, err := m.Run(data)
	return strconv.Itoa(score), err
}

func (problem) PartTwo(data []byte) (string, error) {
	return "", errors.New("not yet implemented")
}
