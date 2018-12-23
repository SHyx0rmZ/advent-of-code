package day14

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	n, err := strconv.Atoi(string(bytes.TrimSpace(bs)))
	if err != nil {
		return "", err
	}
	var recipes = []int{3, 7}
	elf1 := 0
	elf2 := 1
	//fmt.Println("(3)[7]")
	for len(recipes) < n {
		s := recipes[elf1] + recipes[elf2]
		ds := strconv.Itoa(s)
		for _, d := range ds {
			recipes = append(recipes, int(d-'0'))
		}
		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}
	var score []int
	if len(recipes) > n {
		score = append(score, recipes[len(recipes)-1])
	}
	for len(score) < 10 {
		s := recipes[elf1] + recipes[elf2]
		ds := strconv.Itoa(s)
		for _, d := range ds {
			recipes = append(recipes, int(d-'0'))
			score = append(score, int(d-'0'))
		}
		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}
	var s string
	for i := 0; i < 10; i++ {
		s += strconv.Itoa(score[i])
	}
	return s, nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	var match []int
	for _, d := range bytes.TrimSpace(bs) {
		match = append(match, int(d-'0'))
	}
	var recipes = []int{3, 7}
	elf1 := 0
	elf2 := 1
	for {
		s := recipes[elf1] + recipes[elf2]
		ds := strconv.Itoa(s)
		for _, d := range ds {
			recipes = append(recipes, int(d-'0'))
		}
		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
		if len(ds) == 2 {
			var wrong bool
			if len(recipes) > len(match)+1 {
				for i := range match {
					if match[i] != recipes[len(recipes)-len(match)-2+i] {
						wrong = true
					}
				}
			} else {
				wrong = true
			}
			if !wrong {
				return strconv.Itoa(len(recipes) - len(match) - 2), nil
			}
		}
		var wrong bool
		if len(recipes) > len(match) {
			for i := range match {
				if match[i] != recipes[len(recipes)-len(match)-1+i] {
					wrong = true
				}
			}
		} else {
			wrong = true
		}
		if !wrong {
			return strconv.Itoa(len(recipes) - len(match) - 1), nil
		}
	}
	return "", nil
}

func print(recipes []int, elf1, elf2 int) {
	for i := range recipes {
		switch {
		case elf1 == i:
			fmt.Printf("(%d)", recipes[i])
		case elf2 == i:
			fmt.Printf("[%d]", recipes[i])
		default:
			fmt.Printf(" %d ", recipes[i])
		}
	}
	fmt.Println()
}
