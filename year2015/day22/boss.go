package day22

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type boss struct {
	unit
	Damage int
}

func NewBoss(r io.Reader) (boss, error) {
	br := bufio.NewReader(r)
	sh, err := br.ReadString('\n')
	if err != nil {
		return boss{}, err
	}
	sd, err := br.ReadString('\n')
	if err != nil {
		return boss{}, err
	}
	var b boss
	b.Health, err = strconv.Atoi(strings.Split(strings.TrimSpace(sh), ": ")[1])
	if err != nil {
		return boss{}, err
	}
	b.Damage, err = strconv.Atoi(strings.Split(strings.TrimSpace(sd), ": ")[1])
	if err != nil {
		return boss{}, err
	}
	return b, nil
}

func (b *boss) Attack(u *unit) {
	u.Health -= max(1, b.Damage-u.Armor)
}

func (b boss) String() string {
	return strings.TrimSpace(fmt.Sprintln("Boss has", b.Health, "hit points"))
}
