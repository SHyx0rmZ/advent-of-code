package day07

import (
	"os"
	"github.com/SHyx0rmZ/advent-of-code/input"
	"strconv"
	"strings"
	"fmt"
)

type program struct {
	Name string
	Weight int
	Supports []string
}

func Command() error {
	if len(os.Args) < 3 {
		panic("not enough arguments")
	}

	c, err := input.ReadInput(os.Args[2])
	if err != nil {
		return err
	}

	ps := make(map[string]*program)
	sp := set{
		M: make(map[string]struct{}),
	}

	for len(c) > 0 {
		var e int
		var s int
		for c[e] != ' ' {
			e++
		}
		p := &program{}
		p.Name = string(c[s:e])
		s=e+2
		e+=2
		for c[e] != ')' {
			e++
		}
		p.Weight, err = strconv.Atoi(string(c[s:e]))
		if err != nil {
			return err
		}
		s=e+1
		e+=1
		if c[e] == ' ' && c[e+1] == '-' && c[e+2] == '>' && c[e+3] == ' ' {
			s=e+4
			e+=4
			for c[e] != '\n' {
				e++
			}
			p.Supports = strings.Split(string(c[s:e]), ", ")
		}
		fmt.Printf("%+v\n", p)
		c = c[(e+1):]
		ps[p.Name] = p
		sp.Add(p.Name)
	}
	for _, p := range ps {
		for _, s := range p.Supports {
			sp.Del(s)
		}
	}
	fmt.Printf("%+v\n", sp.Elements())
	return nil
}
