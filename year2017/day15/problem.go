package day15

import "fmt"

type problem struct {
	A generator
	B generator
}

func Problem(a, b uint) *problem {
	return &problem{
		A: generator{
			State:  a,
			Stride: 16807,
			Check:  1,
		},
		B: generator{
			State:  b,
			Stride: 48271,
			Check:  1,
		},
	}
}

func (p problem) PartOne(data []byte) (string, error) {
	var s int
	for i := 0; i < 40000000; i++ {
		if judge(p.A.Step(), p.B.Step()) {
			s++
		}
	}
	return fmt.Sprintf("%d", s), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	p.A.Check = 4
	p.B.Check = 8
	var s int
	for i := 0; i < 5000000; i++ {
		if judge(p.A.Step(), p.B.Step()) {
			s++
		}
	}
	return fmt.Sprintf("%d", s), nil
}
