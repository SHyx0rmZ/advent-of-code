package day02_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2019/day02"
)

func TestProblem_PartOneWithReader(t *testing.T) {
	want := "3500"
	r, err := day02.Problem().PartOneWithReader(strings.NewReader(`1,9,10,3,2,3,11,0,99,30,40,50`))
	if r != want || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
	}
}

func TestProblem_PartTwoWithReader(t *testing.T) {
	want := ""
	r, err := day02.Problem().PartTwoWithReader(strings.NewReader(``))
	if r != want || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
	}
}
