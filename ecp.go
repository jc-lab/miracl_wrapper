package miracl_wrapper

import "reflect"

type reflectedEcp struct {
	instance reflect.Value
}

func (r reflectedEcp) ToBytes(b []byte, compress bool) {
	r.instance.MethodByName("ToBytes").Call([]reflect.Value{
		reflect.ValueOf(b),
		reflect.ValueOf(compress),
	})
}

func (r reflectedEcp) Affine() {
	r.instance.MethodByName("Affine").Call([]reflect.Value{})
}
