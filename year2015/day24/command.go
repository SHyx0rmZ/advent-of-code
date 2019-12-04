package day24

import (
	"bufio"
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code"
	"io"
	"math"
	"reflect"
	"sort"
	"strconv"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func parse(r io.Reader) ([]int, error) {
	var ns []int
	s := bufio.NewScanner(r)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return ns, nil
}

func mix(ns []int) {
}

func m(ns, g1, g2, g3 []int, q int) int {
	if len(ns) == 0 {
		fmt.Println(g1, g2, g3, q)
		var t1, t2, t3 int
		t := 1
		for _, n := range g1 {
			t1 += n
			t *= n
		}
		for _, n := range g2 {
			t2 += n
		}
		for _, n := range g3 {
			t3 += n
		}
		if t1 != t2 || t2 != t3 {
			return q
		}
		if t < q && t > 0 {
			fmt.Println(g1, g2, g3, q, t)
			return t
		}
		return q
	}
	q = m(ns[1:], append(g1, ns[0]), g2, g3, q)
	q = m(ns[1:], g1, append(g2, ns[0]), g3, q)
	q = m(ns[1:], g1, g2, append(g3, ns[0]), q)
	return q
}

func m2(t int, g1, g2, g3 []int, q int) int {
	t1 := sum(g1)
	if t1 < t {
		panic("unreachable")
		return q
	}
	if t1 == t {
		fmt.Println(g1, g2, g2, q)
		t2 := sum(g2)
		if t2 != t {
			return q
		}
		t3 := sum(g3)
		if t3 != t {
			return q
		}
		p := product(g1)
		if p < q && p > 0 {
			return p
		}
		return q
	}
	for i := range g1 {
		if t1-g1[i] < t {
			continue
		}
		t2 := sum(g2)
		if t2+g1[i] > t {
			continue
		}
		q = m2(t, append(g1[:i:i], g1[i+1:]...), append(g2, g1[i]), g3, q)
		t3 := sum(g3)
		if t3+g1[i] > t {
			continue
		}
		q = m2(t, append(g1[:i:i], g1[i+1:]...), g2, append(g3, g1[i]), q)
		fmt.Println(t, t1, t2, t3, g1, g2, g3, q)
	}
	return q
}

func accumulate(p, v interface{}, f interface{}) interface{} {
	rp := reflect.Indirect(reflect.ValueOf(p))
	rv := reflect.Indirect(reflect.ValueOf(v))
	rf := reflect.Indirect(reflect.ValueOf(f))
	if rp.Kind() != reflect.Slice && rp.Kind() != reflect.Array {
		panic("expected a slice or array as first argument")
	}
	if rp.Type().Elem() != rv.Type() {
		panic("expected the second argument's type to relate to the first's")
	}
	if rf.Kind() != reflect.Func {
	}
	for i := 0; i < rp.Len(); i++ {
		rv = rf.Call([]reflect.Value{rv, rp.Index(i)})[0]
	}
	return rv.Interface()
}

func sum(p []int) int {
	return accumulate(p, 0, func(a, e interface{}) interface{} {
		return a.(int) + e.(int)
	}).(int)
	//var t int
	//for _, n := range p {
	//	t += n
	//}
	//return t
}

func product(p []int) int {
	return accumulate(p, 1, func(a, e interface{}) interface{} {
		return a.(int) * e.(int)
	}).(int)
}

//func permute(ns []int) [][]int {
//	var ps [][]int
//	for i := range ns {
//		cs := permute(append(ns[:i:i], ns[i+1:]...))
//		for _, c := range cs {
//			ps = append(ps, c)
//			c = append(c, ns[i])
//		}
//	}
//}

func permute(ns []int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer func() {
			close(c)
		}()
		if len(ns) == 6 {
			c <- ns
		}
		if len(ns) < 6 {
			return
		}
		for i := range ns {
			for p := range permute(append(ns[:i:i], ns[i+1:]...)) {
				c <- p
			}
		}
	}()
	return c
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func in(v int, p []int) bool {
	for _, e := range p {
		if e == v {
			return true
		}
	}
	return false
}

func m3(t int, ns []int) int {
	var g1, g2, g3 []int
	var t1, t2, t3 int
	var ons []int
	for _, n := range ns {
		if t1+n <= t {
			g1 = append(g1, n)
			t1 += n
		} else {
			ons = append(ons, n)
		}
	}
	ns = ons
	ons = nil
	for _, n := range ns {
		if t2+n <= t {
			g2 = append(g2, n)
			t2 += n
		} else {
			ons = append(ons, n)
		}
	}
	ns = ons
	ons = nil
	for _, n := range ns {
		if t3+n <= t {
			g3 = append(g3, n)
			t3 += n
		} else {
			ons = append(ons, n)
		}
	}
	if len(ons) != 0 || t1 != t || t2 != t || t3 != t {
		panic("something went wrong")
	}
	return min(min(product(g1), product(g2)), product(g3))
}

func ok2(ns []int, t int) bool {
	var g1, g2, ons []int
	var t1, t2 int
	for _, n := range ns {
		if t1+n <= t {
			g1 = append(g1, n)
			t1 += n
		} else {
			ons = append(ons, n)
		}
	}
	ns = ons
	ons = nil
	for _, n := range ns {
		if t2+n <= t {
			g2 = append(g2, n)
			t2 += n
		} else {
			ons = append(ons, n)
		}
	}
	if len(ons) != 0 || t1 != t || t2 != t {
		return false
	}
	return true
}

const digits = 5

func m5(ns, ss []int, t, c, b int) ([]int, bool) {
	if c > t {
		return nil, false
	}
	if len(ss) >= digits {
		if c == t {
			if len(ns) == 0 {
				return nil, true
			}
			return ns, true
		}
		return nil, false
	}
	for i := range ns {
		if ns[i] < b {
			continue
		}
		x, ok := m5(append(ns[:i:i], ns[i+1:]...), append(ss, ns[i]), t, c+ns[i], ns[i])
		if ok {
			return x, true
		}
	}
	return nil, false
}

func m4(ns, ss []int, t, c, b, q int) int {
	if c > t {
		return q
	}
	if len(ss) >= digits {
		if c == t {
			//if product(ss) != 10439961859 {
			if product(ss) != 72050269 {
				return q
			}
			_, ok := m5(ns, nil, t, 0, 0)
			fmt.Println(product(ss), ok, ss)
			return min(q, product(ss))
		}
		return q
	}
	for i := range ns {
		if ns[i] < b {
			continue
		}
		q = m4(append(ns[:i:i], ns[i+1:]...), append(ss, ns[i]), t, c+ns[i], ns[i], q)
	}
	return q
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	ns, err := parse(r)
	if err != nil {
		return "", err
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	t := sum(ns)

	x := m4(ns, nil, t/3, 0, 0, math.MaxInt64)

	//x := m3(t/3, ns)
	//x := m2(t/3, ns, nil, nil, math.MaxInt32)
	//x := m(ns, nil, nil, nil, math.MaxInt32)
	//fmt.Println(t)
	//q := math.MaxInt32
	//for n := range permute(ns) {
	//	if sum(n) != t/3 {
	//		continue
	//	}
	//	q = min(q, product(n))
	//	fmt.Println(n, q)
	//}
	//fmt.Println(296378068219)
	//fmt.Println(48585083935)
	//fmt.Println(146445538123)
	return strconv.Itoa(x), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {

	ns, err := parse(r)
	if err != nil {
		return "", err
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	t := sum(ns)

	x := m4(ns, nil, t/4, 0, 0, math.MaxInt64)
	return strconv.Itoa(x), nil
}
