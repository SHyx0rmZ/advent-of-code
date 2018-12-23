package day15

type tile struct {
	terrain
	*unit
	extra *byte
}
