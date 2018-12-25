package day24

import (
	"bufio"
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code"
	"io"
	"sort"
	"strconv"
	"strings"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

type team int

const (
	immuneSystem team = iota
	infection
)

type group struct {
	Units      int
	HitPoints  int
	Attack     int
	AttackType typ
	Weak       []typ
	Immune     []typ
	Initiative int
	Team       team
	Target     *group
	Loss       int
}

func (g *group) EffectivePower() int {
	if g == nil {
		return 0
	}
	p := g.Units * g.Attack
	if p < 0 {
		//p = 999999999
	}
	return p
}

func parse(numbers <-chan int, keywords <-chan keyword, types <-chan typ, team team, gs *[]*group, quit <-chan struct{}, done chan<- struct{}) {
	var addTypes func(g *group) []typ
	addTypes = func(g *group) []typ {
		var ts []typ
		for {
			select {
			case t := <-types:
				ts = append(ts, t)
			case k := <-keywords:
				switch k {
				case weak:
					g.Weak = addTypes(g)
					return ts
				case immune:
					g.Immune = addTypes(g)
					return ts
				}
			case n := <-numbers:
				g.Attack = n
				g.AttackType = <-types
				g.Initiative = <-numbers
				return ts
			}
		}
	}
	for {
		select {
		case <-quit:
			done <- struct{}{}
			return
		case n := <-numbers:
			g := new(group)
			g.Team = team
			g.Units = n
			g.HitPoints = <-numbers
			select {
			case k := <-keywords:
				switch k {
				case weak:
					g.Weak = addTypes(g)
				case immune:
					g.Immune = addTypes(g)
				}
			case n := <-numbers:
				g.Attack = n
				g.AttackType = <-types
				g.Initiative = <-numbers
			}
			*gs = append(*gs, g)
		}
	}
}

type keyword int

const (
	weak keyword = iota
	immune
)

type typ int

const (
	bludgeoning typ = iota
	slashing
	fire
	cold
	radiation
)

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	//	r = strings.NewReader(`Immune System:
	//17 units each with 5390 hit points (weak to radiation, bludgeoning) with
	// an attack that does 4507 fire damage at initiative 2
	//989 units each with 1274 hit points (immune to fire; weak to bludgeoning,
	// slashing) with an attack that does 25 slashing damage at initiative 3
	//
	//Infection:
	//801 units each with 4706 hit points (weak to radiation) with an attack
	// that does 116 bludgeoning damage at initiative 1
	//4485 units each with 2961 hit points (immune to radiation; weak to fire,
	// cold) with an attack that does 12 slashing damage at initiative 4
	//
	//`)
	quit1 := make(chan struct{})
	done1 := make(chan struct{})
	quit2 := make(chan struct{})
	done2 := make(chan struct{})
	numbers := make(chan int)
	keywords := make(chan keyword)
	types := make(chan typ)
	//var immuneSystem []group
	//var infection []group
	var groups []*group
	s1 := bufio.NewScanner(r)
	for s1.Scan() {
		switch s1.Text() {
		case "Immune System:":
			go parse(numbers, keywords, types, immuneSystem, &groups, quit1, done1)
			continue
		case "Infection:":
			close(quit1)
			<-done1
			go parse(numbers, keywords, types, infection, &groups, quit2, done2)
			continue
		case "":
			continue
		}
		s2 := bufio.NewScanner(strings.NewReader(s1.Text()))
		s2.Split(bufio.ScanWords)
		for s2.Scan() {
			switch s2.Text() {
			case "units", "each", "with", "hit", "points", "to", "an", "attack", "that", "does", "damage", "at", "initiative":
			case "(weak", "weak":
				keywords <- weak
			case "(immune", "immune":
				keywords <- immune
			case "bludgeoning;", "bludgeoning,", "bludgeoning)", "bludgeoning":
				types <- bludgeoning
			case "slashing;", "slashing,", "slashing)", "slashing":
				types <- slashing
			case "fire;", "fire,", "fire)", "fire":
				types <- fire
			case "cold;", "cold,", "cold)", "cold":
				types <- cold
			case "radiation;", "radiation,", "radiation)", "radiation":
				types <- radiation
			default:
				n, err := strconv.Atoi(s2.Text())
				if err != nil {
					panic(err)
				}
				numbers <- n
			}
		}
	}
	close(quit2)
	<-done2
	//fmt.Println(immuneSystem)

	for {
		var a, d int
		for _, g := range groups {
			if g.Units >= 0 {
				switch g.Team {
				case immuneSystem:
					d += g.Units
				case infection:
					a += g.Units
				}
			}
		}
		fight(groups, true)
		if a == 0 || d == 0 {
			return fmt.Sprintf("immu %d, infe %d", d, a), nil
		}
	}

}

func print(groups []*group) {
	fmt.Println("Immune System:")
	for _, g := range groups {
		if g.Team == immuneSystem && g.Units > 0 {
			fmt.Printf("%d units each with %d hit points with an attack that does %d damage\n", g.Units, g.HitPoints, g.Attack)
		}
	}
	fmt.Println()
	fmt.Println("Infection:")
	for _, g := range groups {
		if g.Team == infection && g.Units > 0 {
			fmt.Printf("%d units each with %d hit points with an attack that does %d damage\n", g.Units, g.HitPoints, g.Attack)
		}
	}
	fmt.Println()
}

func (g *group) DamageFrom(o *group) int {
	if g == nil {
		return 0
	}

	for _, i := range g.Immune {
		if i == o.AttackType {
			return 0
		}
	}

	for _, w := range g.Weak {
		if w == o.AttackType {
			p := o.Units * o.Attack * 2
			if p < 0 {
				//p = 999999999
			}
			return p
		}
	}
	p := o.Units * o.Attack
	if p < 0 {
		//p = 999999999
	}
	return p
}

func fight(groups []*group, verbose bool) bool {

	if verbose {
		print(groups)
	}
	sort.Slice(groups, func(i, j int) bool {
		if groups[i].EffectivePower() == groups[j].EffectivePower() {
			return groups[i].Initiative > groups[j].Initiative
		}
		return groups[i].EffectivePower() > groups[j].EffectivePower()
	})
	for _, a := range groups {
		a.Target = nil

		if a.Units <= 0 {
			continue
		}

	defending:
		for _, d := range groups {
			if d.Team == a.Team {
				continue
			}

			if d.Units <= 0 {
				continue
			}

			for _, g := range groups {
				if g.Target == d {
					continue defending
				}
			}

			ct := a.Target.DamageFrom(a)
			pt := d.DamageFrom(a)

			if pt == 0 {
				continue
			}

			if ct == pt {
				ce := a.Target.EffectivePower()
				pe := d.EffectivePower()

				if ce == pe {
					if a.Target == nil || a.Target.Initiative < d.Initiative {
						a.Target = d
						continue
					}
				}

				if ce < pe {
					a.Target = d
					continue
				}
			}

			if ct < pt {
				a.Target = d
			}
		}
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Initiative > groups[j].Initiative
	})
	var attacked bool
	for i, g := range groups {
		if g.Target == nil {
			continue
		}

		loss := g.Target.DamageFrom(g) / g.Target.HitPoints
		if loss > 0 {
			attacked = true
		}
		g.Target.Units -= loss
		g.Target = nil
		groups[i] = g
	}
	return attacked
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	quit1 := make(chan struct{})
	done1 := make(chan struct{})
	quit2 := make(chan struct{})
	done2 := make(chan struct{})
	numbers := make(chan int)
	keywords := make(chan keyword)
	types := make(chan typ)
	var groups []*group
	s1 := bufio.NewScanner(r)
	for s1.Scan() {
		switch s1.Text() {
		case "Immune System:":
			go parse(numbers, keywords, types, immuneSystem, &groups, quit1, done1)
			continue
		case "Infection:":
			close(quit1)
			<-done1
			go parse(numbers, keywords, types, infection, &groups, quit2, done2)
			continue
		case "":
			continue
		}
		s2 := bufio.NewScanner(strings.NewReader(s1.Text()))
		s2.Split(bufio.ScanWords)
		for s2.Scan() {
			switch s2.Text() {
			case "units", "each", "with", "hit", "points", "to", "an", "attack", "that", "does", "damage", "at", "initiative":
			case "(weak", "weak":
				keywords <- weak
			case "(immune", "immune":
				keywords <- immune
			case "bludgeoning;", "bludgeoning,", "bludgeoning)", "bludgeoning":
				types <- bludgeoning
			case "slashing;", "slashing,", "slashing)", "slashing":
				types <- slashing
			case "fire;", "fire,", "fire)", "fire":
				types <- fire
			case "cold;", "cold,", "cold)", "cold":
				types <- cold
			case "radiation;", "radiation,", "radiation)", "radiation":
				types <- radiation
			default:
				n, err := strconv.Atoi(s2.Text())
				if err != nil {
					panic(err)
				}
				numbers <- n
			}
		}
	}
	close(quit2)
	<-done2

	var boost int
	var verbose bool

	for {
		var clones []*group

		//AttackType typ
		//Weak       []typ
		//Immune     []typ
		//Initiative int
		//Team       team
		//Target     *group
		//Loss       int

		for _, g := range groups {
			c := new(group)
			c.Units = g.Units
			c.Initiative = g.Initiative
			c.HitPoints = g.HitPoints
			c.Attack = g.Attack
			c.AttackType = g.AttackType
			c.Weak = g.Weak
			c.Immune = g.Immune
			c.Initiative = g.Initiative
			c.Team = g.Team
			if c.Team == immuneSystem {
				c.Attack += boost
			}
			clones = append(clones, c)
		}

		for {
			f := fight(clones, verbose)
			var a, d int
			for _, g := range clones {
				if g.Units >= 0 {
					switch g.Team {
					case immuneSystem:
						d += g.Units
					case infection:
						a += g.Units
					}
				}
			}
			if a == 0 || d == 0 || !f {
				fmt.Printf("immu %d, infe %d, boos %d\n", d, a, boost)
				if a == 0 && d > 0 && boost != 28 {
					return fmt.Sprintf("%d", boost), nil
				}
				break
			}
		}

		boost++
		//if boost == 13 {
		//	verbose = true
		//}
	}
}
