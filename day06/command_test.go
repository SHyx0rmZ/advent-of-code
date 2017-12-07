package day06_test

import (
	"github.com/SHyx0rmZ/advent-of-code/day06"
	"testing"
)

func TestBalanceState(t *testing.T) {
	r1, r2 := day06.BalanceState([]int{0, 2, 7, 0})
	if r1 != 4 || r2 != 5 {
		t.Errorf("got (%d, %d), want (%d, %d)", r1, r2, 4, 5)
	}
}
