package day01_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2019/day01"
)

func TestProblem_PartOneWithReader(t *testing.T) {
	want := ""
	r, err := day01.Problem().PartOneWithReader(strings.NewReader(``))
	if r != want || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
	}
}

func TestProblem_PartTwoWithReader(t *testing.T) {
	want := "50346"
	r, err := day01.Problem().PartTwoWithReader(strings.NewReader(`100756`))
	if r != want || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
	}
}
