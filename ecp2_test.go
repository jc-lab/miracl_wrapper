package miracl_wrapper

import "testing"

func TestEcp2_BLS12381_ToBytes(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	G := r.ECP2Generator()

	buffer := make([]byte, r.GetG2S())
	G.ToBytes(buffer, true)
}

func TestEcp2_BLS12381_Affine(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	G := r.ECP2Generator()

	G.Affine()
}
