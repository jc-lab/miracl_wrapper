package miracl_wrapper

import "go.bryk.io/miracl/core/BLS48581"

func NewCurveReflectWithBLS48581() CurveReflect {
	BLS48581.Init()
	return NewCurveReflect(&CurveFunctions{
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
