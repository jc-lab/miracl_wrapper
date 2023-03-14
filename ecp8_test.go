package miracl_wrapper

import "testing"

func TestEcp8_BLS48581_ToBytes(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	G := r.ECP8Generator()

	buffer := make([]byte, r.GetG2S())
	G.ToBytes(buffer, true)
}

func TestEcp8_BLS48581_Affine(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	G := r.ECP8Generator()

	G.Affine()
}
