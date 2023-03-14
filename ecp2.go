package miracl_wrapper

import "reflect"

type reflectedEcp2 struct {
	instance reflect.Value
}

func (r reflectedEcp2) ToBytes(b []byte, compress bool) {
	r.instance.MethodByName("ToBytes").Call([]reflect.Value{
		reflect.ValueOf(b),
		reflect.ValueOf(compress),
	})
}

func (r reflectedEcp2) Affine() {
	r.instance.MethodByName("Affine").Call([]reflect.Value{})
}
