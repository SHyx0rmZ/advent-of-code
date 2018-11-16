package day07

import (
	"errors"
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code/pkg/lib"
	"math"
	"strconv"
	"strings"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	ps, err := p.parse(data)
	if err != nil {
		return "", err
	}
	sp := lib.Set()
	for _, p := range ps {
		sp.Add(p.Name)
	}
	for _, p := range ps {
		for _, s := range p.Supports {
			sp.Delete(s)
		}
	}
	if len(sp.Elements()) != 1 {
		return "", errors.New("invalid tree")
	}
	return sp.Elements()[0].(string), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	ps, err := p.parse(data)
	if err != nil {
		return "", err
	}
	sp := lib.Set()
	for _, p := range ps {
		sp.Add(p.Name)
	}
	for _, p := range ps {
		for _, s := range p.Supports {
			sp.Delete(s)
		}
	}
	if len(sp.Elements()) != 1 {
		return "", errors.New("invalid tree")
	}
	r := ps[sp.Elements()[0].(string)]
	pr := r
	cc := 4
	d := 0.0
	pos := 0
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
			case pos < 1 && len(pr.Supports) > 1:
				o = ps[pr.Supports[pos+1]]
			case pos > 0:
				o = ps[pr.Supports[pos-1]]
			default:
				panic("no neighbor")
			}
			return fmt.Sprintf("%d", o.Combined(ps)-r.Combined(ps)+r.Weight), nil
		}
		d = 0.0
		pos = 0
		for i, w := range cws {
			cd := dev(average(cws), w)
			if cd > d {
				d = cd
				pos = i
			}
		}
		pr = r
		r = ps[r.Supports[pos]]
		cc--
	}
	return "", errors.New("never reached")
}

func (problem) parse(data []byte) (map[string]*program, error) {
	ps := make(map[string]*program)
	var err error
	for len(data) > 0 {
		var e int
		var s int
		for data[e] != ' ' {
			e++
		}
		p := &program{}
		p.Name = string(data[s:e])
		s = e + 2
		e += 2
		for data[e] != ')' {
			e++
		}
		p.Weight, err = strconv.Atoi(string(data[s:e]))
		if err != nil {
			return nil, err
		}
		s = e + 1
		e += 1
		if data[e] == ' ' && data[e+1] == '-' && data[e+2] == '>' && data[e+3] == ' ' {
			s = e + 4
			e += 4
			for data[e] != '\n' {
				e++
			}
			p.Supports = strings.Split(string(data[s:e]), ", ")
		}
		data = data[(e + 1):]
		ps[p.Name] = p
	}
	return ps, nil
}

type program struct {
	Name     string
	Weight   int
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
		ds = append(ds, dev(avg, w)*dev(avg, w))
	}
	return math.Sqrt(sumf(ds) / float64(len(ds)))
}

func dev(avg float64, w int) float64 {
	return math.Abs(float64(w) - avg)
}
