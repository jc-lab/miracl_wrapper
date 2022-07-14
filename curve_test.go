package miracl_wrapper_test

import (
	"github.com/jc-lab/miracl_wrapper"
	"go.bryk.io/miracl/core/BLS48581"
	"testing"
)

func NewCurveReflectWithBLS48581() miracl_wrapper.CurveReflect {
	BLS48581.Init()
	return miracl_wrapper.NewCurveReflect(&miracl_wrapper.CurveFunctions{
		BGS:             BLS48581.BGS,
		BFS:             BLS48581.BFS,
		KeyPairGenerate: BLS48581.KeyPairGenerate,
		CoreSign:        BLS48581.Core_Sign,
		CoreVerify:      BLS48581.Core_Verify,
		ECP8Generator:   BLS48581.ECP8_generator,
		ECP8FromBytes:   BLS48581.ECP8_fromBytes,
		FromBytes:       BLS48581.FromBytes,
		G2mul:           BLS48581.G2mul,
		BIGCurveOrder: func() any {
			return BLS48581.NewBIGints(BLS48581.CURVE_Order)
		},
		Modmul: BLS48581.Modmul,
	})
}

func TestReflectedCurve_GetBGSWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	if r.GetBGS() != BLS48581.BGS {
		t.Errorf("expected %d, got %d", BLS48581.BGS, r.GetBGS())
	}
}

func TestReflectedCurve_GetBFSWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	if r.GetBFS() != BLS48581.BFS {
		t.Errorf("expected %d, got %d", BLS48581.BFS, r.GetBFS())
	}
}

func TestReflectedCurve_GetG1SWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	_ = r.GetG1S()
}

func TestReflectedCurve_GetG2SWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	_ = r.GetG2S()
}

func TestReflectedCurve_KeyPairGenerateWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	ikm := make([]byte, 16)
	S := make([]byte, r.GetBGS())
	W := make([]byte, r.GetG2S())

	result := r.KeyPairGenerate(ikm, S, W)
	if result != 0 {
		t.Errorf("KeyPairGenerate failed: %d", result)
	}
}

func TestReflectedCurve_ECP8GeneratorWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	r.ECP8Generator()
}

func TestReflectedCurve_ECP8ToBytesAndFromBytesWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	G := r.ECP8Generator()

	buffer := make([]byte, r.GetG2S())
	G.ToBytes(buffer, true)
	P := r.ECP8FromBytes(buffer)

	if P == nil {
		t.Errorf("returned nil")
	}
}

func TestReflectedCurve_BIGCurveOrderWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	r.BIGCurveOrder()
	if r.GetBGS() != BLS48581.BGS {
		t.Errorf("expected %d, got %d", BLS48581.BGS, r.GetBGS())
	}
}

func TestReflectedCurve_BIGFromBytesWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()
	a := r.BIGCurveOrder()

	buffer := make([]byte, r.GetBFS())
	a.ToBytes(buffer)

	b := r.FromBytes(buffer)

	if b == nil {
		t.Errorf("returned nil")
	}
}

func TestReflectedCurve_G2mulWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()

	G := r.ECP8Generator()
	a := r.BIGCurveOrder()

	b := r.G2mul(G, a)
	b.Affine()
}

func TestReflectedCurve_ModmulWithBLS48581(t *testing.T) {
	r := NewCurveReflectWithBLS48581()

	a := r.BIGCurveOrder()
	b := r.Modmul(a, a, a)
	b.Invmodp(b)
}

func TestReflectedCurve_CoreSignAndCoreVerify(t *testing.T) {
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

func TestReflectedCurve_CoreSignAndCoreVerify_WithDifferentMessage(t *testing.T) {
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
