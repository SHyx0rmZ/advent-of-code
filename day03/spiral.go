package day03

import (
	"fmt"
	"math"
)

func Command() error {
	n := 368078.0
	sqrt := math.Sqrt(n)
	sqrt = math.Floor(sqrt)
	if math.Mod(sqrt, 2.0) == 0.0 {
		sqrt -= 1.0
	}
	sqrt += 2.0
	var ds []float64
	for i := sqrt-2.0; i > math.Floor((sqrt - 1.0)/ 2.0); i-- {
		ds = append(ds, i)
	}
	for i := math.Floor(sqrt / 2.0) + 0.0; i <= sqrt - 1.0 ; i++ {
		ds = append(ds, i)
	}
	m2 := (sqrt - 2.0) * (sqrt - 2.0)
	rt := sqrt * sqrt
	d := rt - m2

	fmt.Printf("%#v\n", ds[int(math.Mod(n-m2-1.0, d/4.0))])

	return nil
}
