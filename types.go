package miracl_wrapper

type BIGInterface interface {
	Invmodp(p BIGInterface)
	ToBytes(b []byte)
}

type ECP8Interface interface {
	ToBytes(b []byte, compress bool)
	Affine()
}

type CurveReflect interface {
	GetBGS() int
	GetBFS() int
	GetG1S() int
	GetG2S() int
	KeyPairGenerate(IKM []byte, S []byte, W []byte) int
	CoreSign(SIG []byte, M []byte, S []byte) int
	CoreVerify(SIG []byte, M []byte, W []byte) int
	ECP8Generator() ECP8Interface
	ECP8FromBytes(b []byte) ECP8Interface
	FromBytes(b []byte) BIGInterface
	G2mul(P ECP8Interface, e BIGInterface) ECP8Interface
	BIGCurveOrder() BIGInterface
	Modmul(a1 BIGInterface, b1 BIGInterface, m BIGInterface) BIGInterface
}

type CurveFunctions struct {
	BGS             int
	BFS             int
	KeyPairGenerate func(IKM []byte, S []byte, W []byte) int
	CoreSign        func(SIG []byte, M []byte, S []byte) int
	CoreVerify      func(SIG []byte, M []byte, W []byte) int
	ECP8Generator   any
	ECP8FromBytes   any
	FromBytes       any
	G2mul           any
	BIGCurveOrder   func() any
	Modmul          any
}
