package day13

import (
	"bytes"
	"container/ring"
	"fmt"
	"strconv"
)

type layer struct {
	Depth int
	Range int
	S     *ring.Ring
	R     *ring.Ring
}

func (l *layer) step() {
	l.S = l.S.Next()
}

func (l *layer) top() bool {
	return l != nil && l.S == l.R
}

func (l *layer) draw() {
	if l == nil {
		fmt.Printf("...")
		return
	}
	for i := 0; i < l.R.Len(); i++ {
		if l.S == l.R.Move(i) {
			fmt.Printf(" [S]")
		} else {
			fmt.Printf(" [ ]")
		}
	}
}

func (l *layer) reset() {
	if l == nil {
		return
	}

	l.S = l.R
}

type firewall struct {
	F []*layer
}

func (f *firewall) step() {
	for _, l := range f.F {
		if l != nil {
			l.step()
		}
	}
}

func (f firewall) draw(p int) {
	for i, l := range f.F {
		if i == p {
			fmt.Printf("%2d:*", i)
		} else {
			fmt.Printf("%2d: ", i)
		}
		l.draw()
		fmt.Println()
	}
}

func (f firewall) reset() {
	for _, l := range f.F {
		l.reset()
	}
}

type problem struct{}

func (problem) PartOne(data []byte) (string, error) {
	var f firewall
	var err error
	var m int
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		parts := bytes.Split(line, []byte(": "))
		l := &layer{}
		l.Depth, err = strconv.Atoi(string(parts[0]))
		if err != nil {
			return "", err
		}
		l.Range, err = strconv.Atoi(string(parts[1]))
		if err != nil {
			return "", err
		}
		l.R = ring.New((l.Range - 1) * 2)
		l.S = l.R
		for m < l.Depth {
			f.F = append(f.F, nil)
			m++
		}
		f.F = append(f.F, l)
		m++
	}
	var s int
	for i := 0; i < len(f.F)*2; i++ {
		if i%2 == 0 {
			if f.F[i/2].top() {
				s += f.F[i/2].Depth * f.F[i/2].Range
			}
		} else {
			f.step()
		}
	}
	return fmt.Sprintf("%d", s), nil
}

func (problem) PartTwo(data []byte) (string, error) {
	var f firewall
	var err error
	var m int
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		parts := bytes.Split(line, []byte(": "))
		l := &layer{}
		l.Depth, err = strconv.Atoi(string(parts[0]))
		if err != nil {
			return "", err
		}
		l.Range, err = strconv.Atoi(string(parts[1]))
		if err != nil {
			return "", err
		}
		l.R = ring.New((l.Range - 1) * 2)
		l.S = l.R
		for m < l.Depth {
			f.F = append(f.F, nil)
			m++
		}
		f.F = append(f.F, l)
		m++
	}
	var d int
	for {
		for i, l := range f.F {
			if l == nil {
				continue
			}

			if (i+d)%((l.Range-1)*2) == 0 {
				goto meh2
			}
		}
		goto yay
	meh2:
		d++
	}
yay:
	return fmt.Sprintf("%d", d), nil
}

func Problem() *problem {
	return &problem{}
}
