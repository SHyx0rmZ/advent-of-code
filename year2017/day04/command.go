package day04

import (
	"bytes"
	"strconv"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (problem) PartOne(data []byte) (string, error) {
	var sum int
	for _, passphrase := range bytes.Split(data, []byte("\n")) {
		if Valid(string(passphrase)) {
			sum++
		}
	}
	return strconv.Itoa(sum), nil
}

func (problem) PartTwo(data []byte) (string, error) {
	var sum int
	for _, passphrase := range bytes.Split(data, []byte("\n")) {
		if ValidEx(string(passphrase)) {
			sum++
		}
	}
	return strconv.Itoa(sum), nil
}
