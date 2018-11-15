package lib

import (
	"reflect"
	std "sort"
)

func Sort(vs []interface{}) {

}

func sort(vs []interface{}, t reflect.Kind) {
	switch t {
	case reflect.Int:
		std.Sort(IntSlice(vs))
	case reflect.Int32:
		std.Sort(RuneSlice(vs))
	case reflect.String:
		std.Sort(StringSlice(vs))
	}
}

type IntSlice []interface{}

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i].(int) < p[j].(int) }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type RuneSlice []interface{}

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i].(rune) < p[j].(rune) }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type StringSlice []interface{}

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i].(string) < p[j].(string) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
