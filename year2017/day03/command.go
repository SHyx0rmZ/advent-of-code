package day03

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (problem) PartOne(data []byte) (string, error) {
	n, err := strconv.Atoi(string(bytes.TrimSpace(data)))
	if err != nil {
		return "", err
	}
	i := 1
	for p := range spiral() {
		if i == n {
			return fmt.Sprintf("%+v", math.Abs(float64(p.X))+math.Abs(float64(p.Y))), nil
		}
		i++
	}
	return "", errors.New("never reached")
}

func (problem) PartTwo(data []byte) (string, error) {
	n, err := strconv.Atoi(string(bytes.TrimSpace(data)))
	if err != nil {
		return "", err
	}
	v := values{}
	for p := range spiral() {
		if v.calculate(p) > n {
			return fmt.Sprintf("%d", v.calculate(p)), nil
		}
	}
	return "", errors.New("never reached")
}
