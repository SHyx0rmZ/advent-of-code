package day10_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/day10"
)

func TestProblem_PartOne(t *testing.T) {
	r, err := day10.Problem(5).PartOne([]byte("3,4,1,5"))
	if r != "12" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "12", nil)
	}
}
