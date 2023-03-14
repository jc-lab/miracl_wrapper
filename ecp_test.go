package miracl_wrapper

import "testing"

func TestEcp_BLS12381_ToBytes(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	G := r.ECPGenerator()

	buffer := make([]byte, r.GetG1S())
	G.ToBytes(buffer, true)
}

func TestEcp_BLS12381_Affine(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	G := r.ECPGenerator()

	G.Affine()
}
