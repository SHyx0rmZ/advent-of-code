package day07

import (
	"os"
	"github.com/SHyx0rmZ/advent-of-code/input"
	"strconv"
	"strings"
	"fmt"
	"math"
)

type program struct {
	Name string
	Weight int
	Supports []string

	combined *int
}

func (p *program) Combined(m map[string]*program) int {
	if p.combined == nil {
		p.combined = new(int)
		*p.combined = p.Weight
		for _, s := range p.Supports {
			*p.combined += m[s].Combined(m)
		}
	}
	return *p.combined
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
	if len(sp.Elements()) != 1 {
		panic("invalid tree")
	}
	r := ps[sp.Elements()[0]]
	pr := r
	cc := 4
	d := 0.0
	p := 0
	for cc > 0 {
		if len(r.Supports) == 0 {
			panic("empty disc")
		}
		cws := make([]int, len(r.Supports))
		for i, s := range r.Supports {
			cws[i] = ps[s].Combined(ps)
		}
		if stddev(cws) == 0.0 {
			var o *program
			switch {
			case p < 1 && len(pr.Supports) > 1:
				o = ps[pr.Supports[p+1]]
			case p > 0:
				o = ps[pr.Supports[p-1]]
			default:
				panic("no neighbor")
			}
			fmt.Printf("%d\n", o.Combined(ps) - r.Combined(ps) + r.Weight)
		}
		d = 0.0
		p = 0
		for i, w := range cws {
			cd := dev(average(cws), w)
			if cd > d {
				d = cd
				p = i
			}
		}
		pr = r
		r = ps[r.Supports[p]]
		cc--
	}
	return nil
}

func sum(ws []int) float64 {
	var s float64
	for _, w := range ws {
		s += float64(w)
	}
	return s
}

func sumf(ws []float64) float64 {
	var s float64
	for _, w := range ws {
		s += w
	}
	return s
}

func average(ws []int) float64 {
	return sum(ws) / float64(len(ws))
}

func stddev(ws []int) float64 {
	avg := average(ws)
	var ds []float64
	for _, w := range ws {
		ds = append(ds, dev(avg, w) * dev(avg, w))
	}
	return math.Sqrt(sumf(ds)/float64(len(ds)))
}

func dev(avg float64, w int) float64 {
	return math.Abs(float64(w) - avg)
}
