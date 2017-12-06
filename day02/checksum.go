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
	for len(data) > 0 {
		for {
			i := bytes.IndexAny(data, "\t \n")
			if i < 1 {
				// if there are no bytes before the whitespace,
				// we must have reached the end of the line
				break
			}
			// parse column
			n, err := strconv.Atoi(string(data[0:i]))
			if err != nil {
				return 0, err
			}
			s.Column(n)
			// skip data to the next byte behind the whitespace
			for data[i] == '\t' || data[i] == ' ' {
				i++
			}
			data = data[i:]
		}
		sum += s.Row()
		// check for end of input
		i := bytes.IndexAny(data, "\n")
		if i != 0 {
			break
		}
		data = data[1:]
	}
	return sum, nil
}
