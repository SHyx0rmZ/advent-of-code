package day10_test

import "github.com/SHyx0rmZ/advent-of-code/year2017/day10"

func mark7() []day10.Mark {
	m := day10.NewMarks(7)
	ops := []struct {
		Ptr **day10.Mark
		Val *day10.Mark
	}{
		{&m[0].Next, &m[4]}, // 0
		{&m[1].Next, &m[2]},
		{&m[2].Next, &m[3]},
		{&m[3].Next, &m[4]},
		{&m[4].Next, &m[0]}, // 4
		{&m[5].Next, &m[6]},
		{&m[6].Next, &m[0]},

		{&m[0].Prev, &m[6]},
		{&m[1].Prev, &m[5]}, // 8
		{&m[2].Prev, &m[1]},
		{&m[3].Prev, &m[2]},
		{&m[4].Prev, &m[3]},
		{&m[5].Prev, &m[1]}, // 12
		{&m[6].Prev, &m[5]},
	}
	for _, op := range ops {
		*op.Ptr = op.Val
	}
	m[4].Fwd = 4
	m[1].Bkwd = 4
	return m
}
