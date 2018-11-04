package day08

import (
	"io/ioutil"
	"testing"
)

var d []byte

func init() {
	var err error
	d, err = ioutil.ReadFile("data/input.txt")
	if err != nil {
		panic(err)
	}
}

func BenchmarkProblem_Parse1(b *testing.B) {
	b.SetBytes(int64(len(d)))
	p := problem{}
	for i := 0; i < b.N; i++ {
		_, err := p.parse(d)
		if err != nil {
			b.Error(err)
		}
	}
}

//func BenchmarkProblem_Parse2(b *testing.B) {
//	b.SetBytes(int64(len(d)))
//	p := problem{}
//	for i := 0; i < b.N; i++ {
//		_, err := p.parse2(d)
//		if err != nil {
//			b.Error(err)
//		}
//	}
//}
//
//func BenchmarkProblem_Parse3(b *testing.B) {
//	b.SetBytes(int64(len(d)))
//	p := problem{}
//	for i := 0; i < b.N; i++ {
//		_, err := p.parse3(d)
//		if err != nil {
//			b.Error(err)
//		}
//	}
//}
