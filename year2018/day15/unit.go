package day15

type unit struct {
	point
	faction
	hp int
}

type unitSlice []*unit

func (u unitSlice) Len() int {
	return len(u)
}

func (u unitSlice) Less(i, j int) bool {
	if u[i].Y == u[j].Y {
		return u[i].X < u[j].X
	}
	return u[i].Y < u[j].Y
}

func (u unitSlice) Swap(i, j int) {
	t := u[i]
	u[i] = u[j]
	u[j] = t
}

func (u unitSlice) DeadUnits() unitSlice {
	var dus unitSlice
	for _, c := range u {
		if c.hp <= 0 {
			dus = append(dus, c)
		}
	}
	return dus
}

func (u unitSlice) EnemiesOf(a *unit) unitSlice {
	var eus unitSlice
	for _, c := range u {
		if c.faction != a.faction {
			eus = append(eus, c)
		}
	}
	return eus
}

func (u unitSlice) UnitsAttackableBy(a *unit) unitSlice {
	var aus unitSlice
	for _, c := range u {
		if c.faction == a.faction {
			continue
		}
		switch {
		case c.X == a.X && c.Y == a.Y-1:
			fallthrough
		case c.X == a.X-1 && c.Y == a.Y:
			fallthrough
		case c.X == a.X+1 && c.Y == a.Y:
			fallthrough
		case c.X == a.X && c.Y == a.Y+1:
			aus = append(aus, c)
		}
	}
	return aus
}
