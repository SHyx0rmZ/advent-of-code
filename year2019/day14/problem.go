package day14

import (
	"fmt"
	aoc "github.com/SHyx0rmZ/advent-of-code"
	"io"
	"sort"
	"strconv"
)

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	rs, err := Parse(r)
	if err != nil {
		return "", err
	}
	m := make(map[string][]recipe)
	for _, r := range rs {
		sort.Slice(r.In, func(i, j int) bool {
			l, r := r.In[i], r.In[j]
			return l.Component < r.Component
		})
		m[r.Out.Component] = append(m[r.Out.Component], r)
		sort.Slice(m[r.Out.Component], func(i, j int) bool {
			l, r := m[r.Out.Component][i], m[r.Out.Component][j]
			return l.Out.Amount > r.Out.Amount
		})
	}
	var ks []string
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for _, k := range ks {
		fmt.Println(k)
		for _, i := range m[k] {
			fmt.Println("  ", i)
		}
	}
	have := map[string]int{
		"FUEL": -1,
	}
loop:
	for {
		for c, n := range have {
			if n < 0 && c != "ORE" {
				r := m[c][0]
				for have[c] < 0 {
					for _, i := range r.In {
						have[i.Component] -= i.Amount
					}
					have[r.Out.Component] += r.Out.Amount
				}
				continue loop
			}
		}
		return strconv.Itoa(-have["ORE"]), nil
	}
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	rs, err := Parse(r)
	if err != nil {
		return "", err
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].Out.Component < rs[j].Out.Component
	})
	for _, r := range rs {
		fmt.Println(r.Out, "=", r.In)
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		var ro []int
		used := func(r ingredient) bool {
			if r.Component == "FUEL" {
				return true
			}
			for _, ro := range rs {
				for _, ri := range ro.In {
					if ri.Component == r.Component {
						return true
					}
				}
			}
			return false
		}
		for l, r1 := range rs {
			for k, r2 := range rs {
				var ni []ingredient
				var ri []int
				for j, r3 := range r2.In {
					if r1.Out.Component == r3.Component {
						n := r3.Amount / r1.Out.Amount
						if n == 0 {
							continue
						}
						if r3.Amount%r1.Out.Amount == 0 {
							ri = append(ri, j)
						}
						r2.In[j].Amount -= n * r1.Out.Amount
						for _, r4 := range r1.In {
							ni = append(ni, ingredient{n * r4.Amount, r4.Component})
						}
					}
				}
				sort.Sort(sort.Reverse(sort.IntSlice(ri)))
				for _, j := range ri {
					r2.In = append(r2.In[:j], r2.In[j+1:]...)
				}
			update:
				for _, n := range ni {
					for j, r3 := range r2.In {
						if r3.Component == n.Component {
							r2.In[j].Amount += n.Amount
							continue update
						}
					}
					r2.In = append(r2.In, n)
				}
				rs[k] = r2
			}
			if !used(r1.Out) {
				ro = append(ro, l)
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(ro)))
		for _, ri := range ro {
			rs = append(rs[:ri], rs[ri+1:]...)
		}
	}
	for _, r := range rs {
		fmt.Println(r.Out, "=", r.In)
	}
	return "", nil
}
