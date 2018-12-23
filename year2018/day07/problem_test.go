package day07_test

import (
  "strings"
  "testing"

  "github.com/SHyx0rmZ/advent-of-code/year2018/day07"
)

const one = `Step C must be finished before rule A can begin.
Step C must be finished before rule F can begin.
Step A must be finished before rule B can begin.
Step A must be finished before rule D can begin.
Step B must be finished before rule E can begin.
Step D must be finished before rule E can begin.
Step F must be finished before rule E can begin.`

func TestProblem_PartOneWithReader(t *testing.T) {
  want := "CABDFE"
  r, err := day07.Problem().PartOneWithReader(strings.NewReader(one))
  if r != want || err != nil {
    t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
  }
}

func TestProblem_PartTwoWithReader(t *testing.T) {
  want := ""
  r, err := day07.Problem().PartTwoWithReader(strings.NewReader("3,4,1,5"))
  if r != want || err != nil {
    t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
  }
}

func BenchmarkProblem_PartOneWithReader(b *testing.B) {
  for i := 0; i < b.N; i++ {
    day07.Problem().PartOneWithReader(strings.NewReader(""))
  }
}

func BenchmarkProblem_PartTwoWithReader(b *testing.B) {
  for i := 0; i < b.N; i++ {
    day07.Problem().PartTwoWithReader(strings.NewReader(""))
  }
}
