package day10_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2017/day10"
)

func TestProblem_PartOne(t *testing.T) {
	r, err := day10.Problem(5).PartOne([]byte("3,4,1,5"))
	if r != "12" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "12", nil)
	}
}

func TestProblem_PartOne2(t *testing.T) {
	r, err := day10.Problem(5).PartOne2([]byte("3,4,1,5"), 5)
	if r != "12" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "12", nil)
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day10.Problem(5).PartTwo([]byte("3,4,1,5"))
	if r != "04" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "04", nil)
	}
}

//func TestProblem_PartThree(t *testing.T) {
//	r, err := day10.Problem(64).PartThree([]byte("3,4,1,5"))
//	if r != "04" || err != nil {
//		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "04", nil)
//	}
//}

func BenchmarkProblem_PartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.Problem(5).PartOne([]byte("1,2,3,4,5,6,7,8,9,10,11,12,13,14,15"))
	}
}

func BenchmarkProblem_PartOne2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.Problem(5).PartOne2([]byte("1,2,3,4,5,6,7,8,9,10,11,12,13,14,15"), 256)
	}
}

func BenchmarkProblem_PartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.Problem(256).PartTwo([]byte("1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16"))
	}
}

func BenchmarkProblem_PartThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.Problem(256).PartThree([]byte("1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16"))
	}
}
