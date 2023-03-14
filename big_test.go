package miracl_wrapper

import (
	"go.bryk.io/miracl/core/BLS48581"
	"testing"
)

func TestBLS48581_Invmodp(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	p := r.BIGCurveOrder()
	p.Invmodp(p)
}

func TestBLS48581_ToBytes(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	p := r.BIGCurveOrder()

	buffer := make([]byte, r.GetG2S())
	p.ToBytes(buffer)
}

func TestBLS48581_Mod(t *testing.T) {
	origP_add_1 := BLS48581.NewBIGints(BLS48581.CURVE_Order)
	origP_add_1.Plus(BLS48581.NewBIGint(1))

	r := NewCurveReflectWithBLS48581()
	p := r.BIGCurveOrder()

	buffer := make([]byte, r.GetG2S())
	origP_add_1.ToBytes(buffer)
	a := r.FromBytes(buffer)
	a.Mod(p)

	if "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" != a.ToString() {
		t.Fail()
	}
}
