package day04

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/SHyx0rmZ/advent-of-code"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	h := md5.New()
	s = bytes.TrimSpace(s)
	var i int
	for {
		i++
		h.Reset()
		h.Write(s)
		_, err = io.WriteString(h, fmt.Sprintf("%d", i))
		if err != nil {
			return "", err
		}
		m := fmt.Sprintf("%x", h.Sum(nil))
		if strings.HasPrefix(m, "00000") {
			return fmt.Sprintf("%d", i), nil
		}
	}
	panic("unreachable")
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	h := md5.New()
	s = bytes.TrimSpace(s)
	var i int
	for {
		i++
		h.Reset()
		h.Write(s)
		_, err = io.WriteString(h, fmt.Sprintf("%d", i))
		if err != nil {
			return "", err
		}
		m := fmt.Sprintf("%x", h.Sum(nil))
		if strings.HasPrefix(m, "000000") {
			return fmt.Sprintf("%d", i), nil
		}
	}
	panic("unreachable")
}
