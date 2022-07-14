package miracl_wrapper

import "reflect"

type reflectedBig struct {
	instance reflect.Value
}

func (r *reflectedBig) Invmodp(p BIGInterface) {
	reflectedP, _ := p.(*reflectedBig)
	m := r.instance.MethodByName("Invmodp")
	m.Call([]reflect.Value{
		reflectedP.instance,
	})
}

func (r *reflectedBig) ToBytes(b []byte) {
	r.instance.MethodByName("ToBytes").Call([]reflect.Value{
		reflect.ValueOf(b),
	})
}
