package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/pkg/input"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day01"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day02"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day03"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day04"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day05"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day06"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day07"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day08"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day09"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day10"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day11"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day12"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day13"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day14"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day15"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day16"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day17"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day18"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day20"
	"github.com/SHyx0rmZ/advent-of-code/year2017/day21"
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
	nil,
	day20.Problem(),
	day21.Problem(),
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	args := struct {
		Day       string
		Challenge string
		Input     string
	}{
		Day:       os.Args[1],
		Challenge: os.Args[2],
	}

	if len(os.Args) > 3 {
		args.Input = os.Args[3]
	}

	problem, err := strconv.Atoi(args.Day)
	if err != nil {
		panic(err)
	}

	if args.Input == "" {
		path := fmt.Sprintf("day%02d/data/input.txt", problem)
		f, err := os.Open(path)
		if err != nil {
			usage()
		}
		f.Close()
		args.Input = path
	}

	problem--

	p := problems[problem]

	var answer string

	if rap, ok := p.(aoc.ReaderAwareProblem); ok {
		answer, err = solveReaderAwareProblem(rap, args.Challenge, args.Input)
	} else {
		answer, err = solveProblem(p, args.Challenge, args.Input)
	}

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", answer)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <01-24> <a|b> <input>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func solveProblem(problem aoc.Problem, challenge, path string) (string, error) {
	data, err := input.ReadInput(path)
	if err != nil {
		panic(err)
	}

	switch challenge {
	case "a":
		return problem.PartOne(data)
	case "b":
		return problem.PartTwo(data)
	}
	panic("expect either 'a' or 'b'")
}

func solveReaderAwareProblem(problem aoc.ReaderAwareProblem, challenge, path string) (string, error) {
	r, err := input.OpenInputFile(path)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	switch challenge {
	case "a":
		return problem.PartOneWithReader(r)
	case "b":
		return problem.PartTwoWithReader(r)
	}
	panic("expect either 'a' or 'b'")
}
