package miracl_wrapper

import (
	"go.bryk.io/miracl/core/BLS12381"
	"testing"
)

func TestReflectedCurve_BLS12381_GetBGS(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	if r.GetBGS() != BLS12381.BGS {
		t.Errorf("expected %d, got %d", BLS12381.BGS, r.GetBGS())
	}
}

func TestReflectedCurve_BLS12381_GetBFS(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	if r.GetBFS() != BLS12381.BFS {
		t.Errorf("expected %d, got %d", BLS12381.BFS, r.GetBFS())
	}
}

func TestReflectedCurve_BLS12381_GetG1S(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	_ = r.GetG1S()
}

func TestReflectedCurve_BLS12381_GetG2S(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	_ = r.GetG2S()
}

func TestReflectedCurve_BLS12381_KeyPairGenerate(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	ikm := make([]byte, 16)
	S := make([]byte, r.GetBGS())
	W := make([]byte, r.GetG2S())

	result := r.KeyPairGenerate(ikm, S, W)
	if result != 0 {
		t.Errorf("KeyPairGenerate failed: %d", result)
	}
}

func TestReflectedCurve_BLS12381_ECPGenerator(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	r.ECPGenerator()
}

func TestReflectedCurve_BLS12381_ECP2Generator(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	r.ECP2Generator()
}

func TestReflectedCurve_BLS12381_ECPToBytesAndFromBytes(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	G := r.ECPGenerator()

	buffer := make([]byte, r.GetG1S())
	G.ToBytes(buffer, true)
	P := r.ECPFromBytes(buffer)

	if P == nil {
		t.Errorf("returned nil")
	}
}

func TestReflectedCurve_BLS12381_ECP2ToBytesAndFromBytes(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	G := r.ECP2Generator()

	buffer := make([]byte, r.GetG2S())
	G.ToBytes(buffer, true)
	P := r.ECP2FromBytes(buffer)

	if P == nil {
		t.Errorf("returned nil")
	}
}

func TestReflectedCurve_BLS12381_BIGCurveOrder(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	r.BIGCurveOrder()
	if r.GetBGS() != BLS12381.BGS {
		t.Errorf("expected %d, got %d", BLS12381.BGS, r.GetBGS())
	}
}

func TestReflectedCurve_BLS12381_BIGFromBytes(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	a := r.BIGCurveOrder()

	buffer := make([]byte, r.GetBFS())
	a.ToBytes(buffer)

	b := r.FromBytes(buffer)

	if b == nil {
		t.Errorf("returned nil")
	}
}

func TestReflectedCurve_BLS12381_G2mul(t *testing.T) {
	r := NewCurveReflectWithBLS12381()

	G := r.ECP2Generator()
	a := r.BIGCurveOrder()

	b := r.G2mulEcp2(G, a)
	b.Affine()
}

func TestReflectedCurve_BLS12381_Modmul(t *testing.T) {
	r := NewCurveReflectWithBLS12381()

	a := r.BIGCurveOrder()
	b := r.Modmul(a, a, a)
	b.Invmodp(b)
}

func TestReflectedCurve_BLS12381_CoreSignAndCoreVerify(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	ikm := make([]byte, 16)
	S := make([]byte, r.GetBGS())
	W := make([]byte, r.GetG2S())

	result := r.KeyPairGenerate(ikm, S, W)
	if result != 0 {
		t.Errorf("KeyPairGenerate failed: %d", result)
	}

	msg := []byte("HELLO")

	SIG := make([]byte, r.GetG1S())
	ret := r.CoreSign(SIG, msg, S)
	if ret != 0 {
		t.Errorf("sign failed: %d", ret)
	}

	ret = r.CoreVerify(SIG, msg, W)
	if ret != 0 {
		t.Errorf("verify failed: %d", ret)
	}
}

func TestReflectedCurve_BLS12381_CoreSignAndCoreVerify_WithDifferentMessage(t *testing.T) {
	r := NewCurveReflectWithBLS12381()
	ikm := make([]byte, 16)
	S := make([]byte, r.GetBGS())
	W := make([]byte, r.GetG2S())

	result := r.KeyPairGenerate(ikm, S, W)
	if result != 0 {
		t.Errorf("KeyPairGenerate failed: %d", result)
	}

	msg := []byte("HELLO1")

	SIG := make([]byte, r.GetG1S())
	ret := r.CoreSign(SIG, msg, S)
	if ret != 0 {
		t.Errorf("sign failed: %d", ret)
	}

	msg = []byte("HELLO2")

	ret = r.CoreVerify(SIG, msg, W)
	t.Logf("CoreVerify result: %d", ret)
	if ret == 0 {
		t.Errorf("incorrect verify: %d", ret)
	}
}
