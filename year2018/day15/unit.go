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
