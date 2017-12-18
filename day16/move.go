package day16

type Move interface {
	Apply(p *program)
}

type Spin struct {
	X int
}

func (m Spin) Apply(p *program) {
	p.Spin(m.X)
}

type Exchange struct {
	A int
	B int
}

func (m Exchange) Apply(p *program) {
	p.Exchange(m.A, m.B)
}

type Partner struct {
	A rune
	B rune
}

func (m Partner) Apply(p *program) {
	p.Partner(m.A, m.B)
}
