package day17_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2017/day17"
)

func TestProblem_PartOne(t *testing.T) {
	r, err := day17.Problem().PartOne([]byte(""))
	if r != "1912" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "1912", nil)
	}
}
