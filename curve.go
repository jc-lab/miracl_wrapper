package miracl_wrapper

import "reflect"

type reflectedCurve struct {
	input *CurveFunctions
}

func NewCurveReflect(input *CurveFunctions) CurveReflect {
	return &reflectedCurve{
		input: input,
	}
}

func (r reflectedCurve) GetBGS() int {
	return r.input.BGS
}

func (r reflectedCurve) GetBFS() int {
	return r.input.BFS
}

func (r reflectedCurve) GetG1S() int {
	return r.input.BFS + 1
}

func (r reflectedCurve) GetG2S() int {
	return 8*r.input.BFS + 1
}

func (r reflectedCurve) KeyPairGenerate(IKM []byte, S []byte, W []byte) int {
	return r.input.KeyPairGenerate(IKM, S, W)
}

func (r reflectedCurve) CoreSign(SIG []byte, M []byte, S []byte) int {
	return r.input.CoreSign(SIG, M, S)
}

func (r reflectedCurve) CoreVerify(SIG []byte, M []byte, W []byte) int {
	return r.input.CoreVerify(SIG, M, W)
}

func (r reflectedCurve) ECPGenerator() ECPInterface {
	instance := reflect.ValueOf(r.input.ECPGenerator).Call([]reflect.Value{})
	return &reflectedEcp{
		instance: instance[0],
	}
}

func (r reflectedCurve) ECPFromBytes(b []byte) ECPInterface {
	ret := reflect.ValueOf(r.input.ECPFromBytes).Call([]reflect.Value{
		reflect.ValueOf(b),
	})
	return &reflectedEcp8{
		instance: ret[0],
	}
}

func (r reflectedCurve) HasEcp2() bool {
	return r.input.ECP2Generator != nil
}

func (r reflectedCurve) ECP2Generator() ECP2Interface {
	instance := reflect.ValueOf(r.input.ECP2Generator).Call([]reflect.Value{})
	return &reflectedEcp2{
		instance: instance[0],
	}
}

func (r reflectedCurve) ECP2FromBytes(b []byte) ECP2Interface {
	ret := reflect.ValueOf(r.input.ECP2FromBytes).Call([]reflect.Value{
		reflect.ValueOf(b),
	})
	return &reflectedEcp8{
		instance: ret[0],
	}
}

func (r reflectedCurve) HasEcp8() bool {
	return r.input.ECP8Generator != nil
}

func (r reflectedCurve) ECP8Generator() ECP8Interface {
	instance := reflect.ValueOf(r.input.ECP8Generator).Call([]reflect.Value{})
	return &reflectedEcp8{
		instance: instance[0],
	}
}

func (r reflectedCurve) ECP8FromBytes(b []byte) ECP8Interface {
	ret := reflect.ValueOf(r.input.ECP8FromBytes).Call([]reflect.Value{
		reflect.ValueOf(b),
	})
	return &reflectedEcp8{
		instance: ret[0],
	}
}

func (r reflectedCurve) FromBytes(b []byte) BIGInterface {
	ret := reflect.ValueOf(r.input.FromBytes).Call([]reflect.Value{
		reflect.ValueOf(b),
	})
	return &reflectedBig{
		instance: ret[0],
	}
}

func (r reflectedCurve) G1mul(P ECPInterface, e BIGInterface) ECPInterface {
	ret := reflect.ValueOf(r.input.G1mul).Call([]reflect.Value{
		P.(*reflectedEcp).instance,
		e.(*reflectedBig).instance,
	})
	return &reflectedEcp{
		instance: ret[0],
	}
}

func (r reflectedCurve) G2mulEcp2(P ECP2Interface, e BIGInterface) ECP2Interface {
	ret := reflect.ValueOf(r.input.G2mul).Call([]reflect.Value{
		P.(*reflectedEcp2).instance,
		e.(*reflectedBig).instance,
	})
	return &reflectedEcp2{
		instance: ret[0],
	}
}

func (r reflectedCurve) G2mulEcp8(P ECP8Interface, e BIGInterface) ECP8Interface {
	ret := reflect.ValueOf(r.input.G2mul).Call([]reflect.Value{
		P.(*reflectedEcp8).instance,
		e.(*reflectedBig).instance,
	})
	return &reflectedEcp8{
		instance: ret[0],
	}
}

func (r reflectedCurve) BIGCurveOrder() BIGInterface {
	instance := reflect.ValueOf(r.input.BIGCurveOrder).Call([]reflect.Value{})
	return &reflectedBig{
		instance: reflect.ValueOf(instance[0].Interface()),
	}
}

func (r reflectedCurve) Modmul(a1 BIGInterface, b1 BIGInterface, m BIGInterface) BIGInterface {
	instance := reflect.ValueOf(r.input.Modmul).Call([]reflect.Value{
		a1.(*reflectedBig).instance,
		b1.(*reflectedBig).instance,
		m.(*reflectedBig).instance,
	})
	return &reflectedBig{
		instance: reflect.ValueOf(instance[0].Interface()),
	}
}
