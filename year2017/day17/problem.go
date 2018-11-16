package day17

import (
	"container/ring"
	"fmt"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	r := new(ring.Ring)
	r.Value = 0
	for i := 1; i < 2018; i++ {
		r = r.Move(355)
		r = r.Link(&ring.Ring{
			Value: i,
		}).Prev()
	}
	return fmt.Sprintf("%d", r.Next().Value), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	r := new(ring.Ring)
	r.Value = 0
	r0 := r
	for i := 1; i < 50000000; i++ {
		fmt.Printf("\r%10.6f%%", float64(i*100)/float64(50000000))
		r = r.Move(355)
		r = r.Link(&ring.Ring{
			Value: i,
		}).Prev()
	}
	fmt.Println()
	return fmt.Sprintf("%d", r0.Next().Value), nil
}
