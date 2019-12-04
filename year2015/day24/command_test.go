package day24_test

import (
	"github.com/SHyx0rmZ/advent-of-code/year2015/day24"
	"strings"
	"testing"
)

func TestProblem_PartOneWithReader(t *testing.T) {
	want := "99"
	r, err := day24.Problem().PartOneWithReader(strings.NewReader("1\n2\n3\n4\n5\n7\n8\n9\n10\n11"))
	if r != want || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
	}
}
