package day02

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	aoc "github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
)

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func parse(r io.Reader) ([]int, error) {
	var ns []int
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	for _, d := range bytes.Split(bytes.TrimSpace(bs), []byte{','}) {
		n, err := strconv.Atoi(string(d))
		if err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	ns, err := intcode.NewProgram(r)
	fmt.Println(ns)
	//ns, err := parse(r)
	if err != nil {
		return "", err
	}
	ns[1] = 12
	ns[2] = 2
	for i := 0; i < len(ns) && ns[i] != 99; i++ {
		switch ns[i] {
		case 1:
			args := ns[i+1 : i+4]
			ns[args[2]] = ns[args[0]] + ns[args[1]]
			i += 3
		case 2:
			args := ns[i+1 : i+4]
			ns[args[2]] = ns[args[0]] * ns[args[1]]
			i += 3
		}
	}
	return strconv.Itoa(ns[0]), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	ns, err := parse(r)
	if err != nil {
		return "", err
	}
	bu := make([]int, len(ns))
	copy(bu, ns)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(ns, bu)
			ns[1] = noun
			ns[2] = verb
			for i := 0; i < len(ns) && ns[i] != 99; i++ {
				switch ns[i] {
				case 1:
					args := ns[i+1 : i+4]
					ns[args[2]] = ns[args[0]] + ns[args[1]]
					i += 3
				case 2:
					args := ns[i+1 : i+4]
					ns[args[2]] = ns[args[0]] * ns[args[1]]
					i += 3
				}
			}
			if ns[0] == 19690720 {
				return strconv.Itoa(noun*100 + verb), nil
			}
		}
	}
	panic("no solution")
}
