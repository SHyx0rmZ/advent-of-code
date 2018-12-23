package day07

import (
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code"
	"io"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	return p.do(r, 1, 0)
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	return p.do(r, 5, 60)
}

func (p problem) do(r io.Reader, workers int, time int32) (string, error) {
	rs := ParseRules(r)
	dm := make(map[rune][]rune)
	as := make(map[rune]bool)
	for ru := range rs {
		fmt.Printf("Step %c depends on step %c\n", ru.ID, ru.Dependency)
		dm[ru.ID] = append(dm[ru.ID], ru.Dependency)
	}
	const max = 'Z'
	type workload func() workload
	var wf []workload
	var work func(rune) workload
	var done string
	findSchedulable := func() {
		for ru := 'A'; ru <= max; ru++ {
			if !as[ru] && len(dm[ru]) == 0 {
				//fmt.Printf("Step %c can begin\n", ru)
				as[ru] = true
				wf = append(wf, work(ru))
			}
		}
	}
	removeDependency := func(dep rune) {
		for ru := 'A'; ru <= max; ru++ {
			for i, d := range dm[ru] {
				if d == dep {
					dm[ru] = append(dm[ru][:i], dm[ru][i+1:]...)
					if len(dm[ru]) == 0 {
						delete(dm, ru)
					}
					break
				}
			}
		}
		done += string(dep)
		findSchedulable()
	}
	work = func(ru rune) workload {
		return func() workload {
			//fmt.Printf("Step %c began\n", ru)
			oru := ru
			var work2 func(rune) workload
			work2 = func(ru rune) workload {
				fmt.Printf("        %c  ", oru)
				t := ru - 'A' + time
				if t == 0 {
					return func() workload {
						removeDependency(oru)
						//fmt.Printf("Step %c finished\n", oru)
						return nil
					}
				}
				return func() workload {
					return work2(ru - 1)
				}
			}
			return work2(ru)
		}
	}
	run := func() {
		//fmt.Println("running")
		for w := 0; w < workers && w < len(wf); w++ {
			wf[w] = wf[w]()
			if wf[w] == nil {
				wf = append(wf[:w], wf[w+1:]...)
				w--
			}
		}
		for w := len(wf); w < workers; w++ {
			fmt.Printf("        %c  ", '.')
		}
	}
	fmt.Printf("Second")
	for w := 0; w < workers; w++ {
		fmt.Printf("   Worker %d", w)
	}
	fmt.Println("   Done")
	findSchedulable()
	var t int
	for t = 0; len(wf) > 0; t++ {
		fmt.Printf("  %2d", t)
		run()
		fmt.Printf("     %s\n", done)
	}
	return fmt.Sprintf("%d %s\n", t-1, done), nil
}

// ABLCFNSXZPRHVEGUYKDIMQTWJO
