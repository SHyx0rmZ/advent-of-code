package day16_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

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
		switch move.T {
		case day16.MoveSpin:
			movesSpin = append(movesSpin, move)
		case day16.MoveExchange:
			movesExchange = append(movesExchange, move)
		default:
			movesPartner = append(movesPartner, move)
		}
	}
}

func BenchmarkProblem_Dance(b *testing.B) {
	pr := day16.Program()
	p := day16.Problem()
	n := b.N
	b.ResetTimer()
	//for i := 0; i < n; i++ {
	//	p.Dance(pr, moves)
	//}
	p.DanceTest(n, pr, moves)
}

func BenchmarkProblem_DanceSpin(b *testing.B) {
	pr := day16.Program()
	p := day16.Problem()
	n := b.N
	b.ResetTimer()
	for i := 0; i < n; i++ {
		p.Dance(pr, movesSpin)
	}
}

func BenchmarkProblem_DanceExchange(b *testing.B) {
	pr := day16.Program()
	p := day16.Problem()
	n := b.N
	b.ResetTimer()
	for i := 0; i < n; i++ {
		p.Dance(pr, movesExchange)
	}
}

func BenchmarkProblem_DancePartner(b *testing.B) {
	pr := day16.Program()
	p := day16.Problem()
	n := b.N
	b.ResetTimer()
	for i := 0; i < n; i++ {
		p.Dance(pr, movesPartner)
	}
}
