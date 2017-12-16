package day15_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/day15"
)

func TestProblem_PartOne(t *testing.T) {
	r, err := day15.Problem(65, 8921).PartOne([]byte(""))
	if r != "588" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "588", nil)
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day15.Problem(65, 8921).PartTwo([]byte(""))
	if r != "309" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "309", nil)
	}
}
