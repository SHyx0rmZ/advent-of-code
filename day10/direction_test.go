package day10_test

import (
	"github.com/SHyx0rmZ/advent-of-code/day10"
	"testing"
)

func TestDirection_ForwardNext(t *testing.T) {
	marks := day10.NewMarks(3)
	dir := day10.Forward
	if dir.Next(&marks[1]) != &marks[2] {
		t.Errorf("got %p, want %p", dir.Next(&marks[1]), &marks[2])
	}
}

func TestDirection_ForwardPrev(t *testing.T) {
	marks := day10.NewMarks(3)
	dir := day10.Forward
	if dir.Prev(&marks[1]) != &marks[0] {
		t.Errorf("got %p, want %p", dir.Prev(&marks[1]), &marks[0])
	}
}

func TestDirection_BackwardNext(t *testing.T) {
	marks := day10.NewMarks(3)
	dir := day10.Backward
	if dir.Next(&marks[1]) != &marks[0] {
		t.Errorf("got %p, want %p", dir.Next(&marks[1]), &marks[0])
	}
}

func TestDirection_BackwardPrev(t *testing.T) {
	marks := day10.NewMarks(3)
	dir := day10.Backward
	if dir.Prev(&marks[1]) != &marks[2] {
		t.Errorf("got %p, want %p", dir.Prev(&marks[1]), &marks[2])
	}
}

func TestDirection_Toggle(t *testing.T) {
	dir := func(d day10.Direction) *day10.Direction {
		return &d
	}(day10.Forward)
	dir.Toggle()
	if *dir != day10.Backward {
		t.Errorf("got %s, want %s", *dir, day10.Backward)
	}
	dir.Toggle()
	if *dir != day10.Forward {
		t.Errorf("got %s, want %s", *dir, day10.Forward)
	}
}
