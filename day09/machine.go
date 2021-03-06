package day09

import (
	"errors"
	"fmt"
)

type stateFn func([]byte) ([]byte, stateFn)

type machine struct {
	F       stateFn
	err     chan error
	score   int
	current int
	ignored int
}

func Machine() *machine {
	m := &machine{
		err: make(chan error, 1),
	}
	m.F = m.group
	return m
}

func (m machine) Ignored() int {
	return m.ignored
}

func (m *machine) Run(data []byte) error {
	for {
		data, m.F = m.F(data)
		if m.F == nil {
			select {
			case err := <-m.err:
				return err
			default:
				if data == nil {
					return nil
				}
				return errors.New("still have data left")
			}
		}
	}
}

func (m machine) Score() int {
	return m.score
}

func (m *machine) errorf(format string, a ...interface{}) stateFn {
	m.err <- fmt.Errorf(format, a...)
	return nil
}

func (m *machine) garbage(data []byte) ([]byte, stateFn) {
	for len(data) > 0 {
		switch data[0] {
		case '!':
			if len(data) < 2 {
				return data, m.errorf("expected data after '!'")
			}
			data = data[2:]
		case '>':
			return data[1:], m.group
		default:
			m.ignored++
			data = data[1:]
		}
	}
	return nil, nil
}

func (m *machine) group(data []byte) ([]byte, stateFn) {
	for len(data) > 0 {
		switch data[0] {
		case '{':
			m.current++
			return data[1:], m.group
		case '<':
			return data[1:], m.garbage
		case '}':
			m.score += m.current
			m.current--
			return data[1:], m.group
		case ',':
			if m.current < 1 {
				return data, m.errorf("found ',' outside group")
			}
			fallthrough
		default:
			data = data[1:]
		}
	}
	return nil, nil
}
