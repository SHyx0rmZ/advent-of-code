package day21

import (
	"testing"
	"reflect"
	"fmt"
)

func TestMatrixFromBytes(t *testing.T) {
	tests := []struct{
		Name string
		Input []byte
		Matrix *Matrix
	}{
		{
			Name:   "4 succeeds",
			Input:  []byte{'#','.','/','.','#'},
			Matrix: &Matrix{1,0,0,1},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var m Matrix
			defer func() {
				if test.Matrix == nil {
					recover()
				}
				fmt.Printf("asdsd %#v %#v", m, *test.Matrix)
				if !reflect.DeepEqual(m, *test.Matrix) {
					t.Errorf("got %q, want %q", m, *test.Matrix)
				}
			}()
			m = MatrixFromBytes(test.Input)
		})
	}
}
