package day02_test

import (
  "strings"
  "testing"

  "github.com/SHyx0rmZ/advent-of-code/year2018/day02"
)

func TestProblem_PartOneWithReader(t *testing.T) {
  want := "12"
  r, err := day02.Problem().PartOneWithReader(strings.NewReader(`abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`))
  if r != want || err != nil {
    t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
  }
}

func TestProblem_PartTwoWithReader(t *testing.T) {
  want := "fgij"
  r, err := day02.Problem().PartTwoWithReader(strings.NewReader(`abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`))
  if r != want || err != nil {
    t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
  }
}

func BenchmarkProblem_PartOneWithReader(b *testing.B) {
  for i := 0; i < b.N; i++ {
    day02.Problem().PartOneWithReader(strings.NewReader(`abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`))
  }
}

func BenchmarkProblem_PartTwoWithReader(b *testing.B) {
  for i := 0; i < b.N; i++ {
    day02.Problem().PartTwoWithReader(strings.NewReader(`abcde
    fghij
    klmno
    pqrst
    fguij
    axcye
    wvxyz`))
  }
}
