package day04

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"

	aoc "github.com/SHyx0rmZ/advent-of-code"
)

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func parse(r io.Reader) (from, to int, err error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, 0, err
	}
	in := bytes.SplitN(bytes.TrimSpace(bs), []byte{'-'}, 2)
	from, err = strconv.Atoi(string(in[0]))
	if err != nil {
		return 0, 0, err
	}
	to, err = strconv.Atoi(string(in[1]))
	if err != nil {
		return 0, 0, err
	}
	return from, to, nil
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	from, to, err := parse(r)
	if err != nil {
		return "", err
	}
	var c int
numbers:
	for i := from; i < to; i++ {
		s := []byte(strconv.Itoa(i))
		l := s[0]
		var f bool
		for _, b := range s[1:] {
			if b == l {
				f = true
			}
			if b < l {
				continue numbers
			}
			l = b
		}
		if !f {
			continue
		}
		c++
	}
	return strconv.Itoa(c), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	from, to, err := parse(r)
	if err != nil {
		return "", err
	}
	var c int
numbers:
	for i := from; i < to; i++ {
		s := []byte(strconv.Itoa(i))
		l := s[0]
		var l2 byte
		var f [10]bool
		for _, b := range s[1:] {
			if b == l {
				f[b-'0'] = true
			}
			if b == l2 {
				f[b-'0'] = false
			}
			if b < l {
				continue numbers
			}
			l2 = l
			l = b
		}
		var tf bool
		for _, b := range f {
			tf = tf || b
		}
		if !tf {
			continue
		}
		c++
	}
	return strconv.Itoa(c), nil
}
