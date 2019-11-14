package day21

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/SHyx0rmZ/advent-of-code"
)

type unit struct {
	Health int
	Damage int
	Armor  int
}

type item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

var weapons = [...]item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

var armors = [...]item{
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings = [...]item{
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

var null = item{"", 0, 0, 0}

func equip(u unit, is []item) (unit, int) {
	var c int
	for _, i := range is {
		c += i.Cost
		u.Damage += i.Damage
		u.Armor += i.Armor
	}
	return u, c
}

func items() [][]item {
	var is [][]item
	weapon(func(w item) {
		armor(func(a item) {
			ring(func(rs ...item) {
				set := []item{w, a}
				for _, r := range rs {
					set = append(set, r)
				}
				is = append(is, set)
			})
		})
	})
	return is
}

func weapon(f func(item)) {
	for _, w := range weapons {
		f(w)
	}
}

func armor(f func(item)) {
	f(null)
	for _, a := range armors {
		f(a)
	}
}

func ring(f func(...item)) {
	f()
	for _, r1 := range rings {
		f(r1)
		for _, r2 := range rings {
			if r1 == r2 {
				continue
			}
			f(r1, r2)
		}
	}
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

func battle(player, boss unit) bool {
	attacker, defender := &player, &boss
	for player.Health > 0 && boss.Health > 0 {
		defender.Health -= max(1, attacker.Damage-defender.Armor)

		attacker, defender = defender, attacker
	}
	return player.Health > 0
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	br := bufio.NewReader(r)
	sh, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}
	sd, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}
	sa, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}
	var boss unit
	boss.Health, err = strconv.Atoi(strings.Split(strings.TrimSpace(sh), ": ")[1])
	if err != nil {
		return "", err
	}
	boss.Damage, err = strconv.Atoi(strings.Split(strings.TrimSpace(sd), ": ")[1])
	if err != nil {
		return "", err
	}
	boss.Armor, err = strconv.Atoi(strings.Split(strings.TrimSpace(sa), ": ")[1])
	if err != nil {
		return "", err
	}
	player := unit{
		Health: 100,
		Damage: 0,
		Armor:  0,
	}
	min := 999
	for _, is := range items() {
		p, c := equip(player, is)
		if !battle(p, boss) {
			continue
		}
		if c < min {
			min = c
		}
	}
	return strconv.Itoa(min), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	br := bufio.NewReader(r)
	sh, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}
	sd, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}
	sa, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}
	var boss unit
	boss.Health, err = strconv.Atoi(strings.Split(strings.TrimSpace(sh), ": ")[1])
	if err != nil {
		return "", err
	}
	boss.Damage, err = strconv.Atoi(strings.Split(strings.TrimSpace(sd), ": ")[1])
	if err != nil {
		return "", err
	}
	boss.Armor, err = strconv.Atoi(strings.Split(strings.TrimSpace(sa), ": ")[1])
	if err != nil {
		return "", err
	}
	player := unit{
		Health: 100,
		Damage: 0,
		Armor:  0,
	}
	max := 0
	for _, is := range items() {
		p, c := equip(player, is)
		if battle(p, boss) {
			continue
		}
		if c > max {
			max = c
		}
	}
	return strconv.Itoa(max), nil
}
