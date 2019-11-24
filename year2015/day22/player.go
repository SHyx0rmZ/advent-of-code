package day22

import (
	"fmt"
	"strings"
)

type player struct {
	unit
	Mana int
}

func (p player) Alive() bool {
	return p.unit.Alive() && p.Mana > 0
}

func (p player) Attack(u *unit) {}

func (p player) String() string {
	return strings.TrimSpace(fmt.Sprintln("Player has", p.Health, "hit points,", p.Armor, "armor,", p.Mana, "mana"))
}
