package miracl_wrapper

import (
	"reflect"
)

type reflectedBig struct {
	instance reflect.Value
}

func (r *reflectedBig) Nbits() int {
	return int(r.instance.MethodByName("Nbits").Call([]reflect.Value{})[0].Int())
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

func (r *reflectedBig) ToString() string {
	return r.instance.MethodByName("ToString").Call([]reflect.Value{})[0].String()
}

func (r *reflectedBig) Mod(p BIGInterface) {
	reflectedP, _ := p.(*reflectedBig)
	m := r.instance.MethodByName("Mod")
	m.Call([]reflect.Value{
		reflectedP.instance,
	})
}

func (r *reflectedBig) Plus(rhs BIGInterface) BIGInterface {
	reflectedRhs, _ := rhs.(*reflectedBig)
	m := r.instance.MethodByName("Plus")
	ret := m.Call([]reflect.Value{
		reflectedRhs.instance,
	})
	return &reflectedBig{
		instance: ret[0],
	}
}

func (r *reflectedBig) Minus(rhs BIGInterface) BIGInterface {
	reflectedRhs, _ := rhs.(*reflectedBig)
	m := r.instance.MethodByName("Minus")
	ret := m.Call([]reflect.Value{
		reflectedRhs.instance,
	})
	return &reflectedBig{
		instance: ret[0],
	}
}
