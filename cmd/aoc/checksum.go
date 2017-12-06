package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

type checksumState interface {
	Column(n int)
	Row() int
}

type checksumStateDivision struct {
	Numbers []int
}

func (s *checksumStateDivision) Column(n int) {
	s.Numbers = append(s.Numbers, n)
}

func (s *checksumStateDivision) Row() int {
	for i, lp := range s.Numbers {
		if i+1 >= len(s.Numbers) {
			continue
		}
		for _, rp := range s.Numbers[(i + 1):] {
			if lp%rp == 0 {
				s.Numbers = make([]int, 0)
				return lp / rp
			}
			if rp%lp == 0 {
				s.Numbers = make([]int, 0)
				return rp / lp
			}
		}
	}
	s.Numbers = make([]int, 0)
	return 0
}

type checksumStateMinMax struct {
	High int
	Low  int
}

func (s *checksumStateMinMax) Column(n int) {
	if s != nil {
		if n > s.High {
			s.High = n
		}
		if n < s.Low {
			s.Low = n
		}
	}
}

func (s *checksumStateMinMax) Row() int {
	diff := s.High - s.Low
	s.High = math.MinInt64
	s.Low = math.MaxInt64
	return diff
}

func Checksum(data []byte) (int, error) {
	return checksum(data, &checksumStateMinMax{
		High: math.MinInt64,
		Low:  math.MaxInt64,
	})
}

func ChecksumDivision(data []byte) (int, error) {
	return checksum(data, &checksumStateDivision{
		Numbers: make([]int, 0),
	})
}

func checksum(data []byte, s checksumState) (int, error) {
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

func checksumCommand() error {
	if len(os.Args) < 4 {
		panic("not enough arguments")
	}

	var f *os.File
	var err error

	if os.Args[3] == "-" {
		f = os.Stdin
	} else {
		f, err = os.OpenFile(os.Args[3], os.O_RDONLY, 0755)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	c = bytes.TrimSpace(c)

	var sum int

	switch os.Args[2] {
	case "minmax":
		sum, err = Checksum(c)
	case "division":
		sum, err = ChecksumDivision(c)
	default:
		panic("unknown sub-command: " + os.Args[2])
	}

	if err != nil {
		return err
	}

	_, err = fmt.Printf("%d\n", sum)

	return err
}
