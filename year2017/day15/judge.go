package day15

func judge(a uint, b uint) bool {
	return (a & 0xffff) == (b & 0xffff)
}
