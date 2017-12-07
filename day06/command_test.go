package day06_test

import (
	"github.com/SHyx0rmZ/advent-of-code/day06"
	"testing"
)

func TestBalanceState(t *testing.T) {
	r := day06.BalanceState([]int{0, 2, 7, 0})
	if r != 5 {
		t.Errorf("got %d, want %d", r, 5)
	}
}
