package day07_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/day07"
)

func TestProblem_PartOne(t *testing.T) {
	r, err := day07.Problem().PartOneWithReader(strings.NewReader(""))
	if r != "" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "", nil)
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day07.Problem().PartTwoWithReader(strings.NewReader(""))
	if r != "" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "", nil)
	}
}
