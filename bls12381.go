package miracl_wrapper

import "go.bryk.io/miracl/core/BLS12381"

func NewCurveReflectWithBLS12381() CurveReflect {
	BLS12381.Init()
	return NewCurveReflect(&CurveFunctions{
		BGS:             BLS12381.BGS,
		BFS:             BLS12381.BFS,
		KeyPairGenerate: BLS12381.KeyPairGenerate,
		CoreSign:        BLS12381.Core_Sign,
		CoreVerify:      BLS12381.Core_Verify,
		ECPGenerator:    BLS12381.ECP_generator,
		ECPFromBytes:    BLS12381.ECP_fromBytes,
		ECP2Generator:   BLS12381.ECP2_generator,
		ECP2FromBytes:   BLS12381.ECP2_fromBytes,
		FromBytes:       BLS12381.FromBytes,
		G1mul:           BLS12381.G1mul,
		G2mul:           BLS12381.G2mul,
		BIGCurveOrder: func() any {
			return BLS12381.NewBIGints(BLS12381.CURVE_Order)
		},
		Modmul: BLS12381.Modmul,
	})
}
