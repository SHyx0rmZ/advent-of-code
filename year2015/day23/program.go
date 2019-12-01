package day23

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type program []in

func parse(r io.Reader) (program, error) {
	var p program
	s := bufio.NewScanner(r)
lines:
	for s.Scan() {
		ps := strings.SplitN(s.Text(), " ", 2)
		for _, i := range ins {
			if i.Mnemonic == ps[0] {
				ps = strings.Split(ps[1], ", ")
				switch i.X.typ {
				case reg:
					i.X.arg = register(ps[0])
				case off:
					n, err := strconv.Atoi(ps[0])
					if err != nil {
						return nil, err
					}
					i.X.arg = offset(n)
				}
				switch i.Y.typ {
				case reg:
					i.Y.arg = register(ps[1])
				case off:
					n, err := strconv.Atoi(ps[1])
					if err != nil {
						return nil, err
					}
					i.Y.arg = offset(n)
				}
				p = append(p, i)
				continue lines
			}
		}
		return nil, fmt.Errorf("unknown instruction: %s", ps[0])
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return p, nil
}
