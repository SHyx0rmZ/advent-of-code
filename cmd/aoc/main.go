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
	"github.com/SHyx0rmZ/advent-of-code/input"
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

	data, err := input.ReadInput(os.Args[3])
	if err != nil {
		panic(err)
	}

	var answer string

	switch os.Args[2] {
	case "a":
		answer, err = problems[problem].PartOne(data)
	case "b":
		answer, err = problems[problem].PartTwo(data)
	default:
		panic("expect either 'a' or 'b'")
	}
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", answer)
}
