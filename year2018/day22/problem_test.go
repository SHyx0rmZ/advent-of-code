package day22

import (
	"reflect"
	"strconv"
	"testing"
)

func TestAStar(t *testing.T) {
	for i, tt := range []struct {
		M      map[point]*terrain
		Tx, Ty int
		P      []point
	}{
	  /*

	  ...
	  =|.
	  |=.

	   */
		{
			M: map[point]*terrain{
				point{0, 0}: &terrain{typ: rocky},
				point{0, 1}: &terrain{typ: wet},
				point{0, 2}: &terrain{typ: narrow},
				point{1, 0}: &terrain{typ: rocky},
				point{1, 1}: &terrain{typ: narrow},
				point{1, 2}: &terrain{typ: wet},
				point{2, 0}: &terrain{typ: rocky},
				point{2, 1}: &terrain{typ: rocky},
				point{2, 2}: &terrain{typ: rocky},
			},
			Tx: 2,
			Ty: 2,
			P: []point{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
				{2, 2},
			},
		},
		/*

		.=|....
		....=|.

		 */
		{
		  M: map[point]*terrain{

          },
          Tx:
        },
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			p := aStar(point{0, 0}, point{tt.Tx, tt.Ty}, tt.M)
			if !reflect.DeepEqual(tt.P, p) {
				t.Errorf("\n%v\n%v", tt.P, p)
			}
		})
	}
}

//
//func TestProblem_PartOneWithReader(t *testing.T) {
//  want := ""
//  r, err := day22.Problem().PartOneWithReader(strings.NewReader("3,4,1,5"))
//  if r != want || err != nil {
//    t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
//  }
//}
//
//func TestProblem_PartTwoWithReader(t *testing.T) {
//  want := ""
//  r, err := day22.Problem().PartTwoWithReader(strings.NewReader("3,4,1,5"))
//  if r != want || err != nil {
//    t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, want, nil)
//  }
//}
//
//func BenchmarkProblem_PartOneWithReader(b *testing.B) {
//  for i := 0; i < b.N; i++ {
//    day22.Problem().PartOneWithReader(strings.NewReader(""))
//  }
//}
//
//func BenchmarkProblem_PartTwoWithReader(b *testing.B) {
//  for i := 0; i < b.N; i++ {
//    day22.Problem().PartTwoWithReader(strings.NewReader(""))
//  }
//}
