package day16

import (
	"fmt"
)

type program struct {
	data   [32]int
	offset int
}

func Program() *program {
	p := &program{}
	for i := 0; i < 16; i++ {
		p.data[i] = i
		p.data[i+16] = i
	}
	return p
}

func (p *program) Exchange(a, b int) {
	//d := &p.data
	pa := (a + p.offset) & 0xf
	pb := (b + p.offset) & 0xf
	va := p.data[pa]
	vb := p.data[pb]
	p.data[(va + 16)] = pb
	p.data[(vb + 16)] = pa
	p.data[pa] = vb
	p.data[pb] = va
}

func (p *program) Partner(a, b int) {
	ca := &p.data[(a + 16)]
	pa := *ca
	cb := &p.data[(b + 16)]
	pb := *cb
	*ca = pb
	*cb = pa
	p.data[pa] = b
	p.data[pb] = a
}

func (p *program) Spin(x int, _ int) {
	p.offset = (p.offset - x) & 0xf
}

func (p *program) String() string {
	s := "["
	for i := 0; i < 16; i++ {
		s += fmt.Sprintf("%c", p.data[(i+p.offset)%16]+'a')
	}
	return s + "]"
}

// Not used:
func programSpin(control [16]byte, data *[16]byte)
func programPartner(data *[16]byte, a, b byte) (pa, pb int)
