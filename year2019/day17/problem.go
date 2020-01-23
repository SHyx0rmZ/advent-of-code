package day17

import (
	"fmt"
	aoc "github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/pkg/lib"
	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
	"io"
	"strconv"
	"strings"
)

type point lib.Point

type board map[point]rune

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	output := make(chan int, 1)
	b := make(board)
	var j []point
	go prg.Run(input, output)
	var x, y int
	for ch := range output {
		switch ch {
		case '#', '^', '<', '>', 'v':
			b[point{x, y}] = '#'
			if x > 1 && b[point{x - 2, y}] == '#' && b[point{x - 1, y - 1}] == '#' && b[point{x - 1, y}] == '#' && b[point{x - 1, y + 1}] == '#' {
				j = append(j, point{x - 1, y})
			}
			if y > 1 && b[point{x - 1, y - 1}] == '#' && b[point{x, y - 2}] == '#' && b[point{x, y - 1}] == '#' && b[point{x + 1, y - 1}] == '#' {
				j = append(j, point{x, y - 1})
			}
		case '\n':
			x = -1
			y++
		}
		x++
	}
	var t int
	for _, p := range j {
		t += p.X * p.Y
	}
	return strconv.Itoa(t), nil
}

type writer struct {
	ch chan<- int
}

func (w writer) Write(p []byte) (n int, err error) {
	for _, b := range p {
		w.ch <- int(b)
		n++
	}
	return
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 100)
	output := make(chan int, 1)
	b := make(board)
	var j []point
	prg[0] = 2
	go prg.Run(input, output)
	var x, y int
	_, err = io.Copy(writer{input}, strings.NewReader(`A,B,A,C,B,A,B,C,C,B
L,12,L,12,R,4
R,10,R,6,R,4,R,4
R,6,L,12,L,12
n
`))
	if err != nil {
		return "", err
	}
	var l int
	for ch := range output {
		l = ch
		if ch <= 255 {
			fmt.Print(string(ch))
		}
		//time.Sleep(2 * time.Microsecond)
		switch ch {
		case '#', '^', '<', '>', 'v':
			b[point{x, y}] = '#'
			if x > 1 && b[point{x - 2, y}] == '#' && b[point{x - 1, y - 1}] == '#' && b[point{x - 1, y}] == '#' && b[point{x - 1, y + 1}] == '#' {
				j = append(j, point{x - 1, y})
			}
			if y > 1 && b[point{x - 1, y - 1}] == '#' && b[point{x, y - 2}] == '#' && b[point{x, y - 1}] == '#' && b[point{x + 1, y - 1}] == '#' {
				j = append(j, point{x, y - 1})
			}
		case '\n':
			x = -1
			y++
		}
		x++
	}
	return strconv.Itoa(l), nil
}
