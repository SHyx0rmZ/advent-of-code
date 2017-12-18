package lib

import "reflect"

type dict struct {
	m map[interface{}]interface{}
	k reflect.Kind
	v reflect.Kind
}

func Dict() *dict {
	return &dict{
		m: make(map[interface{}]interface{}),
	}
}

func (d dict) Get(k interface{}) (interface{}, bool) {
	v, ok := d.m[k]
	return v, ok
}

func (d dict) Keys() []interface{} {
	var ks []interface{}
	for k := range d.m {
		ks = append(ks, k)
	}
	sort(ks, d.k)
	return ks
}

func (d *dict) Set(k interface{}, v interface{}) {
	kt := reflect.ValueOf(k).Kind()
	vt := reflect.ValueOf(v).Kind()
	if d.k == reflect.Invalid {
		d.k = kt
	}
	if d.v == reflect.Invalid {
		d.v = vt
	}
	if d.k != kt {
		panic("key of type " + kt.String() + " in dict with keys of type " + d.k.String())
	}
	if d.v != vt {
		panic("value of type " + vt.String() + " in dict with values of type " + d.v.String())
	}
	d.m[k] = v
}
