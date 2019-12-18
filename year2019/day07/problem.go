package day07

import (
	"fmt"
	"io"

	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func amp(prg intcode.Program, phases []int) int {
	input := make(chan int, 1)
	output1 := make(chan int, 1)
	output2 := make(chan int, 1)
	output3 := make(chan int, 1)
	output4 := make(chan int, 1)
	output5 := make(chan int, 1)

	input <- phases[0]
	output1 <- phases[1]
	output2 <- phases[2]
	output3 <- phases[3]
	output4 <- phases[4]

	go prg.Run(input, output1)
	go prg.Run(output1, output2)
	go prg.Run(output2, output3)
	go prg.Run(output3, output4)
	go prg.Run(output4, output5)

	input <- 0

	defer close(input)

	return <-output5
}

func loop(prg intcode.Program, phases []int) int {
	output1 := make(chan int, 1)
	output2 := make(chan int, 1)
	output3 := make(chan int, 1)
	output4 := make(chan int, 1)
	output5 := make(chan int, 1)
	result := make(chan int, 1)

	result <- phases[0]
	output1 <- phases[1]
	output2 <- phases[2]
	output3 <- phases[3]
	output4 <- phases[4]

	go prg.Run(result, output1)
	go prg.Run(output1, output2)
	go prg.Run(output2, output3)
	go prg.Run(output3, output4)
	go prg.Run(output4, output5)

	result <- 0

	defer close(result)

	var r int
	for {
		select {
		case n, ok := <-output5:
			if !ok {
				return r
			}
			r = n
			result <- r
		}
	}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}

	var max struct {
		Value  int
		Phases []int
	}
	for phases := range permute([]int{0, 1, 2, 3, 4}, nil) {
		n := amp(prg, phases)
		if n > max.Value {
			max.Value = n
			max.Phases = phases
		}
	}

	return fmt.Sprintf("%v", max.Value), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}

	var max struct {
		Value  int
		Phases []int
	}
	for phases := range permute([]int{5, 6, 7, 8, 9}, nil) {
		n := loop(prg, phases)
		if n > max.Value {
			max.Value = n
			max.Phases = phases
		}
	}

	return fmt.Sprintf("%v", max.Value), nil
}

func permute(ns, ps []int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer func() {
			close(c)
		}()
		if len(ns) == 0 {
			c <- ps
		}
		for i := range ns {
			for p := range permute(append(ns[:i:i], ns[i+1:]...), append(ps, ns[i])) {
				c <- p
			}
		}
	}()
	return c
}
