package day07

import (
	"fmt"
	"io"
)

type problem struct {}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	_, err := p.parse(r)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", 0), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	_, err := p.parse(r)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", 0), nil
}

func (problem) parse(r io.Reader) ([]int, error) {
	var es []int
	return es, nil
}
