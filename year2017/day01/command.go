package day01

import (
	"bytes"
	"strconv"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (problem) PartOne(data []byte) (string, error) {
	return strconv.Itoa(CaptchaNext(string(bytes.TrimSpace(data)))), nil
}

func (problem) PartTwo(data []byte) (string, error) {
	return strconv.Itoa(CaptchaHalf(string(bytes.TrimSpace(data)))), nil
}
