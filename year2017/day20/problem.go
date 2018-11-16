package day20

import (
	"fmt"
	"bytes"
	"sort"
	"github.com/SHyx0rmZ/advent-of-code/pkg/lib"
)

type problem struct {}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	ps, err := p.parse(data)
	if err != nil {
		return "", err
	}
	cp := 0
	z := vec3{}
	for i, p := range ps {
		if p.A == z {
			cp = i
			break
		}
	}

	return fmt.Sprintf("%d", cp), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	ps, err := p.parse(data)
	if err != nil {
		return "", err
	}
	//cl := len(ps)
	r := 0
	//ps = []particle{
	//	{vec3{-6,0,0}, vec3{3,0,0}, vec3{}},
	//	{vec3{-4,0,0}, vec3{2,0,0}, vec3{}},
	//	{vec3{-2,0,0}, vec3{1,0,0}, vec3{}},
	//	{vec3{3,0,0}, vec3{-1,0,0}, vec3{}},
	//}
	for {
		sort.Sort(particles(ps))
		//for i := range ps {
		//	fmt.Printf(
		//		"( %6d / %6d / %6d ) ( %6d / %6d / %6d ) ( %6d / %6d / %6d )\n",
		//		ps[i].P.X, ps[i].P.Y, ps[i].P.Z,
		//		ps[i].V.X, ps[i].V.Y, ps[i].V.Z,
		//		ps[i].A.X, ps[i].A.Y, ps[i].A.Z,
		//	)
		//	if i >= 5 {
		//		break
		//	}
		//}
		pp := ps[len(ps) - 1]
		td := lib.Set()
		if len(ps) > 1 {
			for i := range ps {
				if ps[i].P == pp.P {
					td.Add(i - 1)
					td.Add(i)
				}
				pp = ps[i]
			}
		}
		var nps []particle
		cp := 0
		fmt.Printf("- %d\n", len(td.Elements()))
		//fmt.Printf("%#v\n", td.Elements())
		for _, i := range td.Elements() {
			ii := i.(int)
			//fmt.Printf("%4d[%5d:%5d]\n", len(ps), cp, ii)
			if ii != 0 {
				nps = append(nps, ps[cp:ii]...)
			}
			cp = ii+1
		}
		ps = append(nps, ps[cp:]...)
		//for i := range ps {
		//	fmt.Printf(
		//		"( %6d / %6d / %6d ) ( %6d / %6d / %6d ) ( %6d / %6d / %6d )\n",
		//		ps[i].P.X, ps[i].P.Y, ps[i].P.Z,
		//		ps[i].V.X, ps[i].V.Y, ps[i].V.Z,
		//		ps[i].A.X, ps[i].A.Y, ps[i].A.Z,
		//	)
		//	if i >= 5 {
		//		break
		//	}
		//}
		//fmt.Println("--------------------------------------------------------------------------------------")
		//l := len(ps)
		//if len(td) == 0 && l == cl {
		//	break
		//}
		//fmt.Printf("%5d: %4d\n", r, l)
		r++
		if r > 10000 {
			break
		}
		//cl = l
		for i := range ps {
			ps[i].V = ps[i].V.Add(ps[i].A)
			ps[i].P = ps[i].P.Add(ps[i].V)
		}
	}

	return fmt.Sprintf("%d", len(ps)), nil
}

type particles []particle

func (p particles) Len() int { return len(p) }
func (p particles) Less(i, j int) bool { return p[i].Less(p[j]) }
func (p particles) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type particle struct {
	P vec3
	V vec3
	A vec3
}

func (p particle) Less(o particle) bool {
	if p.P != o.P {
		return p.P.Less(o.P)
	}
	if p.V != o.V {
		return p.V.Less(o.V)
	}
	return p.A.Less(o.A)
}

type vec3 struct {
	X int
	Y int
	Z int
}

func (v vec3) Add(o vec3) vec3 {
	return vec3{
		X: v.X + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}

func (v vec3) Less(o vec3) bool {
	if v.X != o.X {
		return v.X < o.X
	}
	if v.Y != o.Y {
		return v.Y < o.Y
	}
	return v.Z < o.Z
}

func (problem) parse(data []byte) ([]particle, error) {
	var ps []particle
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		line = bytes.Replace(line, []byte(","), []byte(" "), -1)
		line = bytes.Replace(line, []byte("="), []byte(" "), -1)
		var p particle
		_, err := fmt.Sscanf(
			string(line),
			"p <%d %d %d> v <%d %d %d> a <%d %d %d>",
			&p.P.X, &p.P.Y, &p.P.Z,
			&p.V.X, &p.V.Y, &p.V.Z,
			&p.A.X, &p.A.Y, &p.A.Z,
		)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}
