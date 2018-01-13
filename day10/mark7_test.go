package day10_test

import "github.com/SHyx0rmZ/advent-of-code/day10"

func mark7() []day10.Mark {
	m := day10.NewMarks(7)
	ops := []struct {
		Ptr **day10.Mark
		Val *day10.Mark
	}{
		{&m[0].PtrNext, &m[4]}, // 0
		{&m[1].PtrNext, &m[2]},
		{&m[2].PtrNext, &m[3]},
		{&m[3].PtrNext, &m[4]},
		{&m[4].PtrNext, &m[0]}, // 4
		{&m[5].PtrNext, &m[6]},
		{&m[6].PtrNext, &m[0]},
		{&m[0].PtrPrev, &m[6]},
		{&m[1].PtrPrev, &m[5]}, // 8
		{&m[2].PtrPrev, &m[1]},
		{&m[3].PtrPrev, &m[2]},
		{&m[4].PtrPrev, &m[3]},
		{&m[5].PtrPrev, &m[1]}, // 12
		{&m[6].PtrPrev, &m[5]},
	}
	for _, op := range ops {
		*op.Ptr = op.Val
	}
	m[4].TNFS = true
	m[5].TNFE = true
	m[5].TNBS = true
	m[4].TNBE = true
	m[0].ToggleBackward2 = true
	m[1].ToggleBackward = true
	return m
}
