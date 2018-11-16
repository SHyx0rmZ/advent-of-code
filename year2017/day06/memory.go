package day06

import (
	"container/ring"
	"fmt"
	"strings"
)

type memory struct {
	banks []*ring.Ring
}

func (m *memory) Balance() {
	var pos int
	var v = m.banks[0].Value.(int)
	for i, b := range m.banks {
		if b.Value.(int) > v {
			pos = i
			v = b.Value.(int)
		}
	}
	r := m.banks[pos]
	r.Value = 0
	for v > 0 {
		r = r.Next()
		r.Value = r.Value.(int) + 1
		v--
	}
}

func (m *memory) Reload(banks []int) {
	m.banks = make([]*ring.Ring, len(banks))
	r := ring.New(len(banks))

	for i, b := range banks {
		r.Value = b
		m.banks[i] = r
		r = r.Next()
	}
}

func (m *memory) String() string {
	var banks []string
	for _, b := range m.banks {
		banks = append(banks, fmt.Sprintf("%2d", b.Value.(int)))
	}
	return strings.Join(banks, " ")
}
