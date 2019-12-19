package day16

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"time"

	aoc "github.com/SHyx0rmZ/advent-of-code"
)

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	var v string
	for s.Scan() {
		v = s.Text()
		break
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	pattern := []int{0, 1, 0, -1}
	for phase := 0; phase < 100; phase++ {
		var s string
		for i := range v {
			var c int
			for j := range v {
				x := ((j + 1) / (i + 1)) % 4
				c += pattern[x] * int(v[j]-'0')
			}
			if c < 0 {
				c = -c
			}
			s += string((c % 10) + '0')
		}
		v = s
	}
	return v[0:8], nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	var v []byte
	for s.Scan() {
		v = []byte(s.Text())
		break
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	m := make([]byte, 0, len(v)*10000)
	for i := 0; i < 10000; i++ {
		m = append(m, v...)
	}
	v = m
	offset, err := strconv.Atoi(string(v[0:7]))
	if err != nil {
		return "", err
	}
	if offset < len(v)/2 {
		return "", fmt.Errorf("offset not far enough into data, will produce wrong result")
	}
	v = v[offset:]
	t := time.Now()
	n := len(v)
	for phase := 0; phase < 100; phase++ {
		s := make([]byte, len(v))
		var c int
		for i := range v {
			c += int(v[n-i-1] - '0')
			s[n-i-1] = byte((c % 10) + '0')
		}
		v = s
		fmt.Println("phase", phase, "took", time.Since(t))
	}
	return string(v[0:8]), nil
}
