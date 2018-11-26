package main

import (
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code/year2015/day02"
	"github.com/SHyx0rmZ/advent-of-code/year2015/day03"
	"github.com/SHyx0rmZ/advent-of-code/year2015/day04"

	"os"
	"path/filepath"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/pkg/input"
	"github.com/SHyx0rmZ/advent-of-code/year2015/day01"
)

var problems = []aoc.ReaderAwareProblem{
	day01.Problem(),
	day02.Problem(),
	day03.Problem(),
	day04.Problem(),
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
		path := fmt.Sprintf("year2015/day%02d/data/input.txt", problem)
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

	answer, err = solveReaderAwareProblem(p, args.Challenge, args.Input)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", answer)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <01-24> <a|b> <input>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
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
