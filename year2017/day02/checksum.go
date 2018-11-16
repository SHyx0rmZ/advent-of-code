package day02

import (
	"bytes"
	"math"
	"strconv"
)

type state interface {
	Column(n int)
	Row() int
}

func ChecksumDivision(data []byte) (int, error) {
	return checksum(data, &stateDivision{
		Numbers: make([]int, 0),
	})
}

func ChecksumMinMax(data []byte) (int, error) {
	return checksum(data, &stateMinMax{
		High: math.MinInt64,
		Low:  math.MaxInt64,
	})
}

func checksum(data []byte, s state) (int, error) {
	sum := 0
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		for _, digits := range bytes.Split(line, []byte{'\t'}) {
			n, err := strconv.Atoi(string(digits))
			if err != nil {
				return 0, err
			}
			s.Column(n)
		}

		sum += s.Row()
	}
	return sum, nil
}
