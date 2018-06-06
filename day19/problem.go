package day19

import (
	"fmt"
	"bufio"
	"bytes"
	"log"
)

type dir int

const (
	Down dir = iota
	Left
	Up
	Right
)

func (d dir) reverse() dir {
	switch d {
	case Down:
		return Up
	case Left:
		return Right
	case Up:
		return Down
	case Right:
		return Left
	}
	panic("invalid direction")
}

type problem struct {}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	bs, err := p.parse(data)
	if err != nil {
		return "", err
	}

	px := -1
	for i, b := range bs[0] {
		if b == '|' {
			px = i
			break
		}
	}
	if px == -1 {
		return "", fmt.Errorf("start not found")
	}

	var d dir
	py := 0
	var cs []byte

	for {
		switch d {
		case Down:
			py += 1
		case Left:
			px -= 1
		case Up:
			py -= 1
		case Right:
			px += 1
		}
		var ds []dir
		if bs[py][px] == '+' {
			if bs[py][px-1] != ' ' {
				ds = append(ds, Left)
			}
			if bs[py][px+1] != ' ' {
				ds = append(ds, Right)
			}
			if bs[py-1][px] != ' ' {
				ds = append(ds, Up)
			}
			if bs[py+1][px] != ' ' {
				ds = append(ds, Down)
			}
			for i, sd := range ds {
				if sd.reverse() == d {
					ds = append(ds[:i], ds[i+1:]...)
					break
				}
			}
			if len(ds) == 0 {
				return "", nil
			}
			if len(ds) != 1 {
				log.Fatalf("bs[%d][%d]: %+v\n", py, px, ds)
			}
			d = ds[0]
		}
		if bs[py][px] == ' ' {
			return string(cs), nil
		}
		if bs[py][px] >= 'A' && bs[py][px] <= 'Z' {
			cs = append(cs, bs[py][px])
		}
	}
}

func (p problem) PartTwo(data []byte) (string, error) {
	bs, err := p.parse(data)
	if err != nil {
		return "", err
	}

	px := -1
	for i, b := range bs[0] {
		if b == '|' {
			px = i
			break
		}
	}
	if px == -1 {
		return "", fmt.Errorf("start not found")
	}

	var d dir
	py := 0
	var cs []byte
	var s int

	for {
		s++
		switch d {
		case Down:
			py += 1
		case Left:
			px -= 1
		case Up:
			py -= 1
		case Right:
			px += 1
		}
		var ds []dir
		if bs[py][px] == '+' {
			if bs[py][px-1] != ' ' {
				ds = append(ds, Left)
			}
			if bs[py][px+1] != ' ' {
				ds = append(ds, Right)
			}
			if bs[py-1][px] != ' ' {
				ds = append(ds, Up)
			}
			if bs[py+1][px] != ' ' {
				ds = append(ds, Down)
			}
			for i, sd := range ds {
				if sd.reverse() == d {
					ds = append(ds[:i], ds[i+1:]...)
					break
				}
			}
			if len(ds) != 1 {
				log.Fatalf("bs[%d][%d]: %+v\n", py, px, ds)
			}
			d = ds[0]
		}
		if bs[py][px] == ' ' {
			return fmt.Sprintf("%d", s), nil
		}
		if bs[py][px] >= 'A' && bs[py][px] <= 'Z' {
			cs = append(cs, bs[py][px])
		}
	}
}

func (problem) parse(data []byte) ([][]byte, error) {
	var es [][]byte
	s := bufio.NewScanner(bytes.NewReader(data))
	for s.Scan() {
		var cs []byte
		ls := bufio.NewScanner(bytes.NewReader(s.Bytes()))
		ls.Split(bufio.ScanBytes)
		for ls.Scan() {
			cs = append(cs, ls.Bytes()[0])
			fmt.Printf("%s", ls.Text())
		}
		if err := ls.Err(); err != nil {
			return nil, err
		}
		fmt.Println()
		es = append(es, cs)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return es, nil
}
