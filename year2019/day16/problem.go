package day16

import (
	"bufio"
	"fmt"
	aoc "github.com/SHyx0rmZ/advent-of-code"
	"io"
	"os"
	"runtime/pprof"
)

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	var v string
	for s.Scan() {
		v = s.Text()
		break
	}
	if err := s.Err(); err != nil{
		return "", err
	}
	pattern := []int{0, 1, 0, -1}
	for phase := 0; phase < 100; phase++ {
		var s string
		for i := range v {
			var c int
			for j := range v {
				x := ((j+1) / (i+1)) % 4
				c += pattern[x] * int(v[j]-'0')
			}
			if c < 0 {
				c = -c
			}
			s += string((c%10)+'0')
		}
		v = s
	}
	return v[0:8], nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	f, err := os.Create("pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	//C := time.After(10 *time.Second)
	C := make(chan struct{})
	s := bufio.NewScanner(r)
	var v string
	for s.Scan() {
		v = s.Text()
		break
	}
	if err := s.Err(); err != nil{
		return "", err
	}
	m := make([]byte, 0, len(v)*10000)
	for i := 0; i < 10000; i++ {
		m = append(m, v...)
	}
	v = string(m)
	//pattern := []int{0, 1, 0, -1}
	//type pair struct {
	//	J, I int
	//}
	//l := make(map[pair]int)
	for phase := 0; phase < 100; phase++ {
		s := make([]byte, len(v))

		for i := range v {
			var c int
			//for j := range v {
			//	x, ok := l[pair{j, i}]
			//	if !ok {
			//		x = ((j+1)/(i+1))&3
			//		l[pair{j,i}]=x
			//	}
			//	c += pattern[x] * int(v[j]-'0')
			//}
			c = len(v)%((i+1)*4)
			if c < 0 {
				c = -c
			}
			s[i] = byte((c%10)+'0')
			select {
			case <-C:
				return "", nil
			default:
			}
			fmt.Println(i)
		}
		v = string(s)
		fmt.Println("phase", phase)
		return "", nil
	}
	return v[0:8], nil
}
