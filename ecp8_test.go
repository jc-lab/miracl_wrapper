package miracl_wrapper_test

import "testing"

func TestToBytesBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	G := r.ECP8Generator()

	buffer := make([]byte, r.GetG2S())
	G.ToBytes(buffer, true)
}

func TestAffineWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	G := r.ECP8Generator()

	G.Affine()
}
