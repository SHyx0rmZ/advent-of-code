package day16

import (
	"fmt"
	"unsafe"
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

//func (p *program) Exchange(a, b int) {
//	//pa := (a + p.offset) & 0xf
//	//pb := (b + p.offset) & 0xf
//	//pa := a & 0xf
//	//pb := b & 0xf
//	/*
//	pa := a
//	pb := b
//	va := p.data[pa]
//	vb := p.data[pb]
//	*/
//	ca := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[0])) + uintptr(a) * unsafe.Sizeof(int(0))))
//	cb := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[0])) + uintptr(b) * unsafe.Sizeof(int(0))))
//	va := *ca
//	vb := *cb
//	//p.data[(va + 16) & 0xf] = pb
//	//p.data[(vb + 16) & 0xf] = pa
//	*(*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[16])) + uintptr(va) * unsafe.Sizeof(int(0)))) = b
//	*(*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[16])) + uintptr(vb) * unsafe.Sizeof(int(0)))) = a
//	//*(*int)(unsafe.Pointer(&p.data[(va + 16)])) = pb
//	//*(*int)(unsafe.Pointer(&p.data[(vb + 16)])) = pa
//	/*
//	p.data[pa] = vb
//	p.data[pb] = va
//	*/
//	*ca = vb
//	*cb = va
//}

func (p *program) Exchange(a, b int) {
	p.do(a, b, 0, 16)
}

func (p *program) Partner(a, b int) {
	p.do(a, b, 16, 0)
}

func (p *program) do(a, b, c1, c2 int) {
	ca := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[c1])) + uintptr(a) * unsafe.Sizeof(int(0))))
	cb := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[c1])) + uintptr(b) * unsafe.Sizeof(int(0))))
	pa := *ca
	pb := *cb
	*(*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[c2])) + uintptr(pa) * unsafe.Sizeof(int(0)))) = b
	*(*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[c2])) + uintptr(pb) * unsafe.Sizeof(int(0)))) = a
	*ca = pb
	*cb = pa
}

//func (p *program) Partner(a, b int) {
//	//ca := &p.data[(a + 16)]
//	ca := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[16])) + uintptr(a) * unsafe.Sizeof(int(0))))
//	cb := (*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[16])) + uintptr(b) * unsafe.Sizeof(int(0))))
//	pa := *ca
//	pb := *cb
//	*(*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[0])) + uintptr(pa) * unsafe.Sizeof(int(0)))) = b
//	*(*int)(unsafe.Pointer((uintptr)(unsafe.Pointer(&p.data[0])) + uintptr(pb) * unsafe.Sizeof(int(0)))) = a
//	//p.data[pb & 0xf] = a
//	//p.data[pa & 0xf] = b
//	//cb := &p.data[(b + 16)]
//	*ca = pb
//	*cb = pa
//}

func (p *program) Spin(x int, _ int) {
	p.offset = (p.offset - x) & 0xf
}

func (p *program) String() string {
	s := "["
	for i := 0; i < 16; i++ {
		s += fmt.Sprintf("%c", p.data[(i+p.offset)&0xf]+'a')
	}
	return s + "]"
}

// Not used:
//func programSpin(control [16]byte, data *[16]byte)
//func programPartner(data *[16]byte, a, b byte) (pa, pb int)
