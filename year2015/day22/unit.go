package day22

type unit struct {
	Health int
	Armor  int
}

func (u unit) Alive() bool {
	return u.Health > 0
}

func (u *unit) Unit() *unit {
	return u
}
