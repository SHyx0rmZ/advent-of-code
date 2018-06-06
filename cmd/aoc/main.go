package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/day01"
	"github.com/SHyx0rmZ/advent-of-code/day02"
	"github.com/SHyx0rmZ/advent-of-code/day03"
	"github.com/SHyx0rmZ/advent-of-code/day04"
	"github.com/SHyx0rmZ/advent-of-code/day05"
	"github.com/SHyx0rmZ/advent-of-code/day06"
	"github.com/SHyx0rmZ/advent-of-code/day07"
	"github.com/SHyx0rmZ/advent-of-code/day08"
	"github.com/SHyx0rmZ/advent-of-code/day09"
	"github.com/SHyx0rmZ/advent-of-code/day10"
	"github.com/SHyx0rmZ/advent-of-code/day11"
	"github.com/SHyx0rmZ/advent-of-code/day12"
	"github.com/SHyx0rmZ/advent-of-code/day13"
	"github.com/SHyx0rmZ/advent-of-code/day14"
	"github.com/SHyx0rmZ/advent-of-code/day15"
	"github.com/SHyx0rmZ/advent-of-code/day16"
	"github.com/SHyx0rmZ/advent-of-code/day17"
	"github.com/SHyx0rmZ/advent-of-code/day18"
	"github.com/SHyx0rmZ/advent-of-code/day21"
	"github.com/SHyx0rmZ/advent-of-code/input"
	"github.com/SHyx0rmZ/advent-of-code/day19"
)

var problems = []aoc.Problem{
	day01.Problem(),
	day02.Problem(),
	day03.Problem(),
	day04.Problem(),
	day05.Problem(),
	day06.Problem(),
	day07.Problem(),
	day08.Problem(),
	day09.Problem(),
	day10.Problem(256),
	day11.Problem(),
	day12.Problem(),
	day13.Problem(),
	day14.Problem(),
	day15.Problem(591, 393),
	day16.Problem(),
	day17.Problem(),
	day18.Problem(),
	day19.Problem(),
	nil,
	day21.Problem(),
}

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "usage: %s <01-24> <a|b> <input>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	problem, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	problem--

	p := problems[problem]

	var answer string

	if rap, ok := p.(aoc.ReaderAwareProblem); ok {
		answer, err = solveReaderAwareProblem(rap)
	} else {
		answer, err = solveProblem(p)
	}

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", answer)
}

func solveProblem(problem aoc.Problem) (string, error) {
	data, err := input.ReadInput(os.Args[3])
	if err != nil {
		panic(err)
	}

	switch os.Args[2] {
	case "a":
		return problem.PartOne(data)
	case "b":
		return problem.PartTwo(data)
	}
	panic("expect either 'a' or 'b'")
}

func solveReaderAwareProblem(problem aoc.ReaderAwareProblem) (string, error) {
	r, err := input.OpenInputFile(os.Args[3])
	if err != nil {
		panic(err)
	}
	defer r.Close()

	switch os.Args[2] {
	case "a":
		return problem.PartOneWithReader(r)
	case "b":
		return problem.PartTwoWithReader(r)
	}
	panic("expect either 'a' or 'b'")
}
