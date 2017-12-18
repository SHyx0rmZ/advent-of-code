package day16

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code/lib"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) Dance(pr *program, moves []Move) {
	for _, move := range moves {
		move.Apply(pr)
	}
}

func (p problem) PartOne(data []byte) (string, error) {
	moves, err := p.Parse(data)
	if err != nil {
		return "", err
	}
	pr := Program()
	p.Dance(pr, moves)
	return pr.String(), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	moves, err := p.Parse(data)
	if err != nil {
		return "", err
	}
	pr := Program()
	s := lib.Set()
	d := lib.Dict()
	k := 1000000000
	for i := 0; i < k; i++ {
		fmt.Printf("\r%10.6f%%", float64(i*100)/float64(k))
		h := pr.String()
		if pr.offset == 0 && s.Contains(h) {
			fmt.Println(", detected cycle")
			v, _ := d.Get(k % i)
			return v.(string), nil
		}
		s.Add(h)
		d.Set(i, h)
		p.Dance(pr, moves)
	}
	fmt.Println()
	return pr.String(), nil
}

func (problem) Parse(data []byte) ([]Move, error) {
	var moves []Move
	for _, i := range bytes.Split(bytes.TrimSpace(data), []byte(",")) {
		if len(i) == 0 {
			continue
		}
		ps := bytes.Split(i[1:], []byte("/"))
		switch i[0] {
		case 's':
			x, err := strconv.Atoi(string(ps[0]))
			if err != nil {
				return nil, err
			}
			moves = append(moves, Spin{X: x})
		case 'x':
			a, err := strconv.Atoi(string(ps[0]))
			if err != nil {
				return nil, err
			}
			b, err := strconv.Atoi(string(ps[1]))
			if err != nil {
				return nil, err
			}
			moves = append(moves, Exchange{A: a, B: b})
		case 'p':
			moves = append(moves, Partner{A: rune(ps[0][0]), B: rune(ps[1][0])})
		default:
			panic("unknown dance move")
		}
	}
	return moves, nil
}
