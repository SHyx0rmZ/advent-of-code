package day16_test

import (
	"testing"
	"io/ioutil"
	"path/filepath"

	"github.com/SHyx0rmZ/advent-of-code/day16"
)

var moves []day16.Move
var movesSpin []day16.Move
var movesExchange []day16.Move
var movesPartner []day16.Move

func init() {
	bs, err := ioutil.ReadFile(filepath.Join("data", "input.txt"))
	if err != nil {
		panic(err)
	}
	moves, err = day16.Problem().Parse(bs)
	if err != nil {
		panic(err)
	}
	movesSpin = make([]day16.Move, 0)
	movesExchange = make([]day16.Move, 0)
	movesPartner = make([]day16.Move, 0)
	for _, move := range moves {
		switch m := move.(type) {
		case day16.Spin:
			movesSpin = append(movesSpin, m)
		case day16.Exchange:
			movesExchange = append(movesExchange, m)
		case day16.Partner:
			movesPartner = append(movesPartner, m)
		}
	}
}

func BenchmarkProblem_Dance(b *testing.B) {
	pr := day16.Program()
	for i := 0; i < b.N; i++ {
		day16.Problem().Dance(pr, moves)
	}
}

func BenchmarkProblem_DanceSpin(b *testing.B) {
	pr := day16.Program()
	for i := 0; i < b.N; i++ {
		day16.Problem().Dance(pr, movesSpin)
	}
}

func BenchmarkProblem_DanceExchange(b *testing.B) {
	pr := day16.Program()
	for i := 0; i < b.N; i++ {
		day16.Problem().Dance(pr, movesExchange)
	}
}

func BenchmarkProblem_DancePartner(b *testing.B) {
	pr := day16.Program()
	for i := 0; i < b.N; i++ {
		day16.Problem().Dance(pr, movesPartner)
	}
}
