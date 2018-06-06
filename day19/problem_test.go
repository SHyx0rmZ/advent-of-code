package day19_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/day19"
)

func TestProblem_PartOne(t *testing.T) {
	//r, err := day19.Problem().PartOne([]byte(
	//	"     |          "  +
	//	"	    |  +--+    " +
	//	"	    A  |  C    " +
	//	"	F---|----E|--+ " +
	//	"	    |  |  |  D " +
	//	"	    +B-+  +--+ "))

	r, err := day19.Problem().PartOne([]byte(
`      |          
      |  +--+    
      A  |  C    
  F---|----E|--+ 
      |  |  |  D 
      +B-+  +--+ 
                 `))
	if r != "ABCDEF" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "ABCDEF", nil)
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day19.Problem().PartTwo([]byte(
	`      |          
      |  +--+    
      A  |  C    
  F---|----E|--+ 
      |  |  |  D 
      +B-+  +--+ 
                 `))
	if r != "38" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "38", nil)
	}
}
