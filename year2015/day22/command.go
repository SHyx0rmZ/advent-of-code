package day22

import (
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code"
	"io"
	"strconv"
)

type item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

type spell struct {
	Name   string
	Cost   int
	Effect effect
}

type effect func(player *player, boss *boss, turns <-chan struct{})

var spells = [...]spell{
	{"Magic Missile", 53, func(player *player, boss *boss, turns <-chan struct{}) {
		boss.Health -= 4
	}},
	{"Drain", 73, func(player *player, boss *boss, turns <-chan struct{}) {
		boss.Health -= 2
		player.Health += 2
	}},
	{"Shield", 113, func(player *player, boss *boss, turns <-chan struct{}) {
		player.Armor += 7
		for i := 0; i < 6; i++ {
			<-turns
		}
		player.Armor -= 7
	}},
	{"Poison", 173, func(player *player, boss *boss, turns <-chan struct{}) {
		for i := 0; i < 6; i++ {
			<-turns
			boss.Health -= 3
		}
	}},
	{"Recharge", 229, func(player *player, boss *boss, turns <-chan struct{}) {
		for i := 0; i < 5; i++ {
			player.Mana += 101
		}
	}},
}

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cast(player *player, boss *boss, spell spell, turns <-chan struct{}) {
	go spell.Effect(player, boss, turns)

}
func battle(player player, boss boss) bool { //, spells <-chan spell) bool {
	var attacker, defender interface {
		Attack(u *unit)
		Unit() *unit
	}
	attacker, defender = &player, &boss
	for player.Alive() && boss.Alive() {
		if attacker == &player {
			fmt.Println("-- Player turn --")
		} else {
			fmt.Println("-- Boss turn --")
		}
		fmt.Println(player)
		fmt.Println(boss)

		attacker.Attack(defender.Unit())
		fmt.Println()

		attacker, defender = defender, attacker
	}
	return player.Alive()
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	b, err := NewBoss(r)
	if err != nil {
		return "", err
	}
	pl := player{
		unit: unit{
			Health: 100,
			Armor:  0,
		},
		Mana: 500,
	}
	pl = player{unit{10, 0}, 250}
	b = boss{unit{13, 0}, 8}
	min := 999
	s := make(chan spell)
	go func() {
		s <- spells[3]
		s <- spells[0]
		close(s)
	}()
	for {
		if !battle(pl, b) {
			continue
		}
		//if c < min {
		//	min = c
		//}
	}
	return strconv.Itoa(min), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	return "", nil
}
