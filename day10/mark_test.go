package day10_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/day10"
)

type mark struct {
	Value int
	Next  int
	Prev  int
}

func TestNewMarks(t *testing.T) {
	tests := []struct {
		Name    string
		Lengths int
		Marks   []mark
	}{
		{
			Name:    "0",
			Lengths: 0,
			Marks:   nil,
		},
		{
			Name:    "1",
			Lengths: 1,
			Marks: []mark{
				{
					Value: 0,
					Next:  0,
					Prev:  0,
				},
			},
		},
		{
			Name:    "3",
			Lengths: 3,
			Marks: []mark{
				{
					Value: 0,
					Next:  1,
					Prev:  2,
				},
				{
					Value: 1,
					Next:  2,
					Prev:  0,
				},
				{
					Value: 2,
					Next:  0,
					Prev:  1,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			marks := day10.NewMarks(test.Lengths)
			if len(marks) != len(test.Marks) {
				t.Fatalf("got len of %d, want %d", len(marks), len(test.Marks))
			}
			for i := 0; i < len(test.Marks); i++ {
				if marks[i].Value != test.Marks[i].Value {
					t.Errorf("got Value of %d, want %d", marks[i].Value, test.Marks[i].Value)
				}
				if marks[i].PtrNext != &marks[test.Marks[i].Next] {
					t.Errorf("got PtrNext of %p, want %p", marks[i].PtrNext, &marks[test.Marks[i].Next])
				}
				if marks[i].PtrPrev != &marks[test.Marks[i].Prev] {
					t.Errorf("got PtrPrev of %p, want %p", marks[i].PtrPrev, &marks[test.Marks[i].Prev])
				}
			}
		})
	}
}

func TestMark_Next(t *testing.T) {
	tests := []struct {
		Name  string
		Marks []mark
		Mark  int
		Next  int
	}{
		{
			Name: "3-0",
			Marks: []mark{
				{
					Next: 1,
					Prev: 2,
				},
				{
					Next: 2,
					Prev: 0,
				},
				{
					Next: 0,
					Prev: 1,
				},
			},
			Mark: 0,
			Next: 1,
		},
		{
			Name: "3-2",
			Marks: []mark{
				{
					Next: 1,
					Prev: 2,
				},
				{
					Next: 2,
					Prev: 0,
				},
				{
					Next: 0,
					Prev: 1,
				},
			},
			Mark: 2,
			Next: 0,
		},
		{
			Name: "5-1",
			Marks: []mark{
				{
					Next: 1,
					Prev: 4,
				},
				{
					Next: 2,
					Prev: 0,
				},
				{
					Next: 3,
					Prev: 1,
				},
				{
					Next: 4,
					Prev: 2,
				},
				{
					Next: 0,
					Prev: 3,
				},
			},
			Mark: 1,
			Next: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			marks := day10.NewMarks(len(test.Marks))
			for i, mark := range test.Marks {
				marks[i].PtrNext = &marks[mark.Next]
				marks[i].PtrPrev = &marks[mark.Prev]
			}
			mark := marks[test.Mark]
			next := mark.Next()
			if next != &marks[test.Next] {
				t.Errorf("got %p, want %p", next, &marks[test.Next])
			}
		})
	}
}
