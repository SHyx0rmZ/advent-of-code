package day20_test

import (
  "strings"
  "testing"

  "github.com/SHyx0rmZ/advent-of-code/year2018/day20"
)

func TestProblem_PartOneWithReader(t *testing.T) {
  want := ""
  r, err := day20.Problem().PartOneWithReader(strings.NewReader("3,4,1,5"))
  if r != want || err != nil {
    t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
  }
}

func TestProblem_PartTwoWithReader(t *testing.T) {
  want := ""
  r, err := day20.Problem().PartTwoWithReader(strings.NewReader("3,4,1,5"))
  if r != want || err != nil {
    t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
  }
}

func BenchmarkProblem_PartOneWithReader(b *testing.B) {
  for i := 0; i < b.N; i++ {
    day20.Problem().PartOneWithReader(strings.NewReader(""))
  }
}

func BenchmarkProblem_PartTwoWithReader(b *testing.B) {
  for i := 0; i < b.N; i++ {
    day20.Problem().PartTwoWithReader(strings.NewReader(""))
  }
}
