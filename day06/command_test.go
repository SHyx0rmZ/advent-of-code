package day06_test

import (
	"github.com/SHyx0rmZ/advent-of-code/day06"
	"testing"
)

func TestProblem_PartOne(t *testing.T) {
	r, err := day06.Problem().PartOne([]byte("0\t2\t7\t0\n"))
	if r != "5" || err != nil {
		t.Errorf("got (%s, %+v), want (%d, %+v)", r, err, 5, nil)
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day06.Problem().PartTwo([]byte("0\t2\t7\t0\n"))
	if r != "4" || err != nil {
		t.Errorf("got (%s, %+v), want (%d, %+v)", r, err, 4, nil)
	}
}
