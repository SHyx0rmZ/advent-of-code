package day16

import "fmt"

type program struct {
	data [16]byte
	cache [16]int
	offset int
}

func Program() *program {
	p := &program{}
	for i := 0; i < 16; i++ {
		b := byte('a' + i)
		p.data[i] = b
		p.cache[i] = i
	}
	return p
}

func (p *program) Exchange(a, b int) {
	am := (a + p.offset) & 0xf
	bm := (b + p.offset) & 0xf
	ba := p.data[am]
	bb := p.data[bm]
	p.data[bm] = ba
	p.data[am] = bb
	p.cache[ba-'a'] = bm
	p.cache[bb-'a'] = am
}

func (p *program) Partner(a, b rune) {
	ao := a-'a'
	bo := b-'a'
	pa := p.cache[ao]
	pb := p.cache[bo]
	p.data[pa], p.data[pb] = p.data[pb], p.data[pa]
	p.cache[ao], p.cache[bo] = p.cache[bo], p.cache[ao]
}

func (p *program) Spin(x int) {
	p.offset = (p.offset + 16 - x) & 0xf
}

func (p *program) String() string {
	s := "["
	for i := 0; i < 16; i++ {
		s += fmt.Sprintf("%c", p.data[(i + p.offset)%16])
	}
	return s + "]"
}

// Not used:
func programSpin(control [16]byte, data *[16]byte)
func programPartner(data *[16]byte, a, b byte) (pa, pb int)
