package day21

import "math"

type Matrix []int

func MatrixFromBytes(p []byte) Matrix {
	var m Matrix
	var i, j int
	switch len(p) {
	case 5:
		m = make([]int, 4)
		i = 5
		j = 4
	case 11:
		m = make([]int, 9)
		i = 11
		j = 9
	case 19:
		m = make([]int, 16)
		i = 19
		j = 16
	default:
		panic("invalid grid")
	}
	_ = m[j-1]
	_ = p[i-1]
	for ; i > 0; i-- {
		if p[i-1] == '/' {
			continue
		}
		m[j-1] = bit(p[i-1])
		j--
	}
	return m
}

func (m Matrix) Transpose() Matrix {
	n := m.dimension() * m.dimension()
	switch n {
	case 4:
		return Matrix{
			m[0], m[2],
			m[1], m[3],
		}
	case 9:
		return Matrix{
			m[0], m[3], m[6],
			m[1], m[4], m[7],
			m[2], m[5], m[8],
		}
	}
	panic("invalid dimensions")
}

func (m Matrix) Flip() Matrix {
	n := m.dimension() * m.dimension()
	switch n {
	case 4:
		return Matrix{
			m[2], m[3],
			m[0], m[1],
		}
	case 9:
		return Matrix{
			m[6], m[7], m[8],
			m[3], m[4], m[5],
			m[0], m[1], m[2],
		}
	}
	panic("invalid dimensions")
}

func (m Matrix) Rotate() Matrix {
	return m.Transpose().Flip()
}

func (m Matrix) Enhance() Matrix {
	d := m.dimension()
	if d%2 == 0 {
		n := d / 2
		r := make(Matrix, n*3*n*3)
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				s := Matrix{
					m[offset(2, n, x, y, 0, 0)],
					m[offset(2, n, x, y, 1, 0)],
					m[offset(2, n, x, y, 0, 1)],
					m[offset(2, n, x, y, 1, 1)],
				}.Substitute()
				r[offset(3, n, x, y, 0, 0)] = s[0]
				r[offset(3, n, x, y, 1, 0)] = s[1]
				r[offset(3, n, x, y, 2, 0)] = s[2]
				r[offset(3, n, x, y, 0, 1)] = s[3]
				r[offset(3, n, x, y, 1, 1)] = s[4]
				r[offset(3, n, x, y, 2, 1)] = s[5]
				r[offset(3, n, x, y, 0, 2)] = s[6]
				r[offset(3, n, x, y, 1, 2)] = s[7]
				r[offset(3, n, x, y, 2, 2)] = s[8]
			}
		}
		return r
	}
	n := d / 3
	r := make(Matrix, n*4*n*4)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			s := Matrix{
				m[offset(3, n, x, y, 0, 0)],
				m[offset(3, n, x, y, 1, 0)],
				m[offset(3, n, x, y, 2, 0)],
				m[offset(3, n, x, y, 0, 1)],
				m[offset(3, n, x, y, 1, 1)],
				m[offset(3, n, x, y, 2, 1)],
				m[offset(3, n, x, y, 0, 2)],
				m[offset(3, n, x, y, 1, 2)],
				m[offset(3, n, x, y, 2, 2)],
			}.Substitute()
			r[offset(4, n, x, y, 0, 0)] = s[0]
			r[offset(4, n, x, y, 1, 0)] = s[1]
			r[offset(4, n, x, y, 2, 0)] = s[2]
			r[offset(4, n, x, y, 3, 0)] = s[3]
			r[offset(4, n, x, y, 0, 1)] = s[4]
			r[offset(4, n, x, y, 1, 1)] = s[5]
			r[offset(4, n, x, y, 2, 1)] = s[6]
			r[offset(4, n, x, y, 3, 1)] = s[7]
			r[offset(4, n, x, y, 0, 2)] = s[8]
			r[offset(4, n, x, y, 1, 2)] = s[9]
			r[offset(4, n, x, y, 2, 2)] = s[10]
			r[offset(4, n, x, y, 3, 2)] = s[11]
			r[offset(4, n, x, y, 0, 3)] = s[12]
			r[offset(4, n, x, y, 1, 3)] = s[13]
			r[offset(4, n, x, y, 2, 3)] = s[14]
			r[offset(4, n, x, y, 3, 3)] = s[15]
		}
	}
	return r
}

func (m Matrix) Equals(o Matrix) bool {
	if len(m) != len(o) {
		return false
	}
	for i := 0; i < len(m); i++ {
		if m[i] != o[i] {
			return false
		}
	}
	return true
}

var rules []EnhancementRule

func (m Matrix) Substitute() Matrix {
	for _, rule := range rules {
		r := rule.From
		if m.Equals(r) {
			return rule.To
		}
		r = r.Rotate()
		if m.Equals(r) {
			return rule.To
		}
		r = r.Rotate()
		if m.Equals(r) {
			return rule.To
		}
		r = r.Rotate()
		if m.Equals(r) {
			return rule.To
		}
		r = r.Flip()
		if m.Equals(r) {
			return rule.To
		}
		r = r.Rotate()
		if m.Equals(r) {
			return rule.To
		}
		r = r.Rotate()
		if m.Equals(r) {
			return rule.To
		}
		r = r.Rotate()
		if m.Equals(r) {
			return rule.To
		}
	}
	panic("invalid matrix")
}

func (m Matrix) Count() int {
	var s int
	for i := 0; i < len(m); i++ {
		s += m[i]
	}
	return s
}

func (m Matrix) dimension() int {
	return int(math.Sqrt(float64(len(m))))
}
