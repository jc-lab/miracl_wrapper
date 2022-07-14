package miracl_wrapper

import "reflect"

type reflectedEcp8 struct {
	instance reflect.Value
}

func (r reflectedEcp8) ToBytes(b []byte, compress bool) {
	r.instance.MethodByName("ToBytes").Call([]reflect.Value{
		reflect.ValueOf(b),
		reflect.ValueOf(compress),
	})
}

func (r reflectedEcp8) Affine() {
	r.instance.MethodByName("Affine").Call([]reflect.Value{})
}
