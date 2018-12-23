package day23

import (
  "io"

  "github.com/SHyx0rmZ/advent-of-code"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
  return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
  return "", nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
  return "", nil
}
