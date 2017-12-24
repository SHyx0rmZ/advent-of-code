package day16

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code/lib"
	"runtime/pprof"
	"os"
	"unsafe"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func programExchange(a, b int, d *[32]int)

func (p problem) Dance(pr *program, moves []Move) {
	n := len(moves)
	for i := 0; i < n; i++ {
		m := moves[i]
		switch {
		//case m.S:
		//	pr.Spin(m.A, m.B)
		//case m.E:
		//	pr.Exchange(m.A, m.B)
		//default:
		//	pr.Partner(m.A, m.B)
		case m.E:
			//pr.Exchange(m.A, m.B)
			//pr.do(m.A, m.B, 0, 16)
			ca := (*int)(unsafe.Pointer(&pr.data[0]))
			cb := (*int)(unsafe.Pointer(&pr.data[0]))
			ca = (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(ca)) + uintptr(m.A) * unsafe.Sizeof(int(0))))
			cb = (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(cb)) + uintptr(m.B) * unsafe.Sizeof(int(0))))
			pa := *ca
			pb := *cb
			oa := (*int)(unsafe.Pointer(&pr.data[16]))
			ob := (*int)(unsafe.Pointer(&pr.data[16]))
			oa = (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(oa)) + uintptr(pa) * unsafe.Sizeof(int(0))))
			ob = (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(ob)) + uintptr(pb) * unsafe.Sizeof(int(0))))
			*oa = m.B
			*ob = m.A
			*ca = pb
			*cb = pa
			//programExchange(m.A, m.B, &pr.data)
		case !m.S:
			ca := (*int)(unsafe.Pointer(&pr.data[16]))
			cb := (*int)(unsafe.Pointer(&pr.data[16]))
			ca = (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(ca)) + uintptr(m.A) * unsafe.Sizeof(int(0))))
			cb = (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(cb)) + uintptr(m.B) * unsafe.Sizeof(int(0))))
			pa := *ca
			pb := *cb
			oa := (*int)(unsafe.Pointer(&pr.data[0]))
			ob := (*int)(unsafe.Pointer(&pr.data[0]))
			oa = (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(oa)) + uintptr(pa) * unsafe.Sizeof(int(0))))
			ob = (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(ob)) + uintptr(pb) * unsafe.Sizeof(int(0))))
			*oa = m.B
			*ob = m.A
			*ca = pb
			*cb = pa
			//pr.Partner(m.A, m.B)
			//pr.do(m.A, m.B, 16, 0)
		default:
			pr.Spin(m.A, m.B)
		}
	}
}

//func (p problem) Dance(pr *program, moves []Move) {
//	n := len(moves)
//	var c1, c2 int
//	for i := 0; i < n; i++ {
//		m := moves[i]
//		switch {
//		//case m.S:
//		//	pr.Spin(m.A, m.B)
//		//case m.E:
//		//	pr.Exchange(m.A, m.B)
//		//default:
//		//	pr.Partner(m.A, m.B)
//		case m.E:
//			//pr.Exchange(m.A, m.B)
//			//pr.do(m.A, m.B, 0, 16)
//			c1 = 0
//			c2 = 16
//		case !m.S:
//			c1 = 16
//			c2 = 0
//			//pr.Partner(m.A, m.B)
//			//pr.do(m.A, m.B, 16, 0)
//		default:
//			pr.Spin(m.A, m.B)
//			continue
//		}
//		ca := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&pr.data[0])) + uintptr(m.A | c1) * unsafe.Sizeof(int(0))))
//		cb := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&pr.data[0])) + uintptr(m.B | c1) * unsafe.Sizeof(int(0))))
//		pa := *ca
//		pb := *cb
//		*(*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&pr.data[0])) + uintptr(pa | c2) * unsafe.Sizeof(int(0)))) = m.B
//		*(*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&pr.data[0])) + uintptr(pb | c2) * unsafe.Sizeof(int(0)))) = m.A
//		*ca = pb
//		*cb = pa
//	}
//}

func (p problem) PartOne(data []byte) (string, error) {
	moves, err := p.Parse(data)
	if err != nil {
		return "", err
	}
	pr := Program()
	p.Dance(pr, moves)
	return pr.String(), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	pprof.StartCPUProfile(os.Stderr)
	defer pprof.StopCPUProfile()

	moves, err := p.Parse(data)
	if err != nil {
		return "", err
	}
	pr := Program()
	s := lib.Set()
	d := lib.Dict()
	k := 1000000000
	for i := 0; i < k; i++ {
		fmt.Printf("\r%10.6f%%", float64(i*100)/float64(k))
		h := pr.String()
		if i > 10000 {
			break
		}
		//if pr.offset == 0 && s.Contains(h) {
		//	fmt.Println(", detected cycle")
		//	v, _ := d.Get(k % i)
		//	return v.(string), nil
		//}
		s.Add(h)
		d.Set(i, h)
		p.Dance(pr, moves)
	}
	fmt.Println()
	return pr.String(), nil
}

func (problem) Parse(data []byte) ([]Move, error) {
	var moves []Move
	var offset int
	last := &Move{A:-1, B:-1}
	lastI := -1
	for _, i := range bytes.Split(bytes.TrimSpace(data), []byte(",")) {
		if len(i) == 0 {
			continue
		}
		ps := bytes.Split(i[1:], []byte("/"))
		switch i[0] {
		case 's':
			x, err := strconv.Atoi(string(ps[0]))
			if err != nil {
				return nil, err
			}
			//moves = append(moves, Move{S: true, A: x})
			offset = (offset - x) & 0xf
		case 'x':
			a, err := strconv.Atoi(string(ps[0]))
			if err != nil {
				return nil, err
			}
			b, err := strconv.Atoi(string(ps[1]))
			if err != nil {
				return nil, err
			}
			m := Move{E: true, A: (a+offset)&0xf, B: (b+offset)&0xf}
			la, lb := last.A, last.B
			ma, mb := m.A, m.B
			switch {
			case (la == ma && lb == mb) || (la == mb && lb == ma):
				//fmt.Printf("=")
				moves = append(moves[:lastI], moves[lastI+1:]...)
				last = &Move{A:-1, B:-1}
				lastI = -1
			//case la == ma:
			//	fmt.Printf("a")
			//	fmt.Printf("\n%#v %#v\n", *last, m)
			//	panic("asd")
			//	//moves = append(moves, m)
			//	//last = &moves[len(moves) - 1]
			//	//lastI = len(moves) - 1
			//	last.A = m.B
			//case la == mb:
			//	fmt.Printf("A")
			//	moves = append(moves, m)
			//	last = &moves[len(moves) - 1]
			//	lastI = len(moves) - 1
			//case lb == mb:
			//	fmt.Printf("b")
			//	moves = append(moves, m)
			//	last = &moves[len(moves) - 1]
			//	lastI = len(moves) - 1
			//case lb == ma:
			//	fmt.Printf("B")
			//	moves = append(moves, m)
			//	last = &moves[len(moves) - 1]
			//	lastI = len(moves) - 1
			default:
				//fmt.Printf(".")
				moves = append(moves, m)
				last = &moves[len(moves) - 1]
				lastI = len(moves) - 1
			}
			//moves = append(moves, Move{E: true, A: a, B: b})
		case 'p':
			moves = append(moves, Move{A: int(ps[0][0] - 'a'), B: int(ps[1][0] - 'a')})
		default:
			panic("unknown dance move")
		}
	}
	moves = append(moves, Move{S: true, A: -offset})
	return moves, nil
}
