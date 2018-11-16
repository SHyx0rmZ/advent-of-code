package day10_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/pkg/lib"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day10"
)

func TestIter_Next(t *testing.T) {
	m := mark7()

	tests := []struct {
		Name  string
		Dir   day10.Direction
		Set   func() lib.GenSet
		Start int
		Vals  []int
	}{
		{
			Name:  "Forward0",
			Dir:   day10.Forward,
			Start: 0,
			Vals:  []int{0, 4, 3, 2, 1, 5, 6, 0},
		},
		{
			Name:  "Forward5",
			Dir:   day10.Forward,
			Start: 5,
			Vals:  []int{5, 6, 0, 4, 3, 2, 1, 5},
		},
		//{
		//	Name:  "Backward0",
		//	Dir:   day10.Backward,
		//	Start: 0,
		//	Vals:  []int{0, 6, 5, 1, 9, 3, 4, 0},
		//},
		//{
		//	Name:  "Backward1",
		//	Dir:   day10.Backward,
		//	Start: 1,
		//	Vals:  []int{1, 5, 6, 0, 4, 3, 2, 1},
		//	Set: func() lib.GenSet {
		//		var set lib.GenSet
		//		c := day10.Counter(2)
		//		set.Add(c)
		//		go c.Run()
		//		return set
		//	},
		//},
		//{
		//
		//	Name:  "Backward1",
		//	Dir:   day10.Backward,
		//	Start: 1,
		//	Vals:  []int{1, 5, 6, 0, 4, 3, 2, 1},
		//},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			c := &m[test.Start]
			it := day10.Iter{
				Direction: test.Dir,
			}
			if test.Set != nil {
				it.Set = test.Set()
			}
			for i := 0; i < len(test.Vals); i++ {
				if c.Value != test.Vals[i] {
					t.Errorf("Values[%d]: got %d, want %d", i, c.Value, test.Vals[i])
				}
				c = it.Next(c)
			}
		})
	}
}

func TestIter_Prev(t *testing.T) {
	m := mark7()

	tests := []struct {
		Name  string
		Dir   day10.Direction
		Start int
		Vals  []int
	}{
		//{
		//	Name:  "Backward0",
		//	Dir:   day10.Backward,
		//	Start: 0,
		//	Vals:  []int{0, 4, 3, 2, 1, 5, 6, 0},
		//},
		{
			Name:  "Forward0",
			Dir:   day10.Forward,
			Start: 0,
			Vals:  []int{0, 6, 5, 1, 2, 3, 4, 0},
		},
		{
			Name:  "Forward1",
			Dir:   day10.Forward,
			Start: 1,
			Vals:  []int{1, 2, 3, 4, 0, 6, 5, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			c := &m[test.Start]
			it := day10.Iter{
				Direction: test.Dir,
			}
			for i := 0; i < len(test.Vals); i++ {
				if c.Value != test.Vals[i] {
					t.Errorf("Values[%d]: got %d, want %d", i, c.Value, test.Vals[i])
				}
				c = it.Prev(c)
			}
		})
	}
}
