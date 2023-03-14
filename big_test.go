package miracl_wrapper

import "testing"

func TestInvmodpBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	p := r.BIGCurveOrder()
	p.Invmodp(p)

}

func TestToBytesWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	p := r.BIGCurveOrder()

	buffer := make([]byte, r.GetG2S())
	p.ToBytes(buffer)
}
