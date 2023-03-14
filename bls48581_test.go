package miracl_wrapper

import (
	"go.bryk.io/miracl/core/BLS48581"
	"testing"
)

func TestReflectedCurve_BLS48581_GetBGS(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	if r.GetBGS() != BLS48581.BGS {
		t.Errorf("expected %d, got %d", BLS48581.BGS, r.GetBGS())
	}
}

func TestReflectedCurve_BLS48581_GetBFS(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	if r.GetBFS() != BLS48581.BFS {
		t.Errorf("expected %d, got %d", BLS48581.BFS, r.GetBFS())
	}
}

func TestReflectedCurve_BLS48581_GetG1S(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	_ = r.GetG1S()
}

func TestReflectedCurve_BLS48581_GetG2S(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	_ = r.GetG2S()
}

func TestReflectedCurve_BLS48581_KeyPairGenerate(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	ikm := make([]byte, 16)
	S := make([]byte, r.GetBGS())
	W := make([]byte, r.GetG2S())

	result := r.KeyPairGenerate(ikm, S, W)
	if result != 0 {
		t.Errorf("KeyPairGenerate failed: %d", result)
	}
}

func TestReflectedCurve_BLS48581_ECP8Generator(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	r.ECP8Generator()
}

func TestReflectedCurve_BLS48581_ECP8ToBytesAndFromBytes(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	G := r.ECP8Generator()

	buffer := make([]byte, r.GetG2S())
	G.ToBytes(buffer, true)
	P := r.ECP8FromBytes(buffer)

	if P == nil {
		t.Errorf("returned nil")
	}
}

func TestReflectedCurve_BLS48581_BIGCurveOrder(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	r.BIGCurveOrder()
	if r.GetBGS() != BLS48581.BGS {
		t.Errorf("expected %d, got %d", BLS48581.BGS, r.GetBGS())
	}
}

func TestReflectedCurve_BLS48581_BIGFromBytes(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	a := r.BIGCurveOrder()

	buffer := make([]byte, r.GetBFS())
	a.ToBytes(buffer)

	b := r.FromBytes(buffer)

	if b == nil {
		t.Errorf("returned nil")
	}
}

func TestReflectedCurve_BLS48581_G2mul(t *testing.T) {
	r := NewCurveReflectWithBLS48581()

	G := r.ECP8Generator()
	a := r.BIGCurveOrder()

	b := r.G2mulEcp8(G, a)
	b.Affine()
}

func TestReflectedCurve_BLS48581_Modmul(t *testing.T) {
	r := NewCurveReflectWithBLS48581()

	a := r.BIGCurveOrder()
	b := r.Modmul(a, a, a)
	b.Invmodp(b)
}

func TestReflectedCurve_BLS48581_CoreSignAndCoreVerify(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
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

func TestReflectedCurve_BLS48581_CoreSignAndCoreVerify_WithDifferentMessage(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
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
