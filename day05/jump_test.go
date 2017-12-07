package day05_test

import (
	"github.com/SHyx0rmZ/advent-of-code/day05"
	"testing"
)

func TestJump(t *testing.T) {
	r, err := day05.Jump("0\n3\n0\n1\n-3")
	if r != 5 || err != nil {
		t.Errorf("got (%d, %+v), want (%d, %+v)", r, err, 5, nil)
	}
}
