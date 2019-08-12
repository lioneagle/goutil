package mathex

import (
	"math"
)

const (
	MinAccuracyFloat32 = 1.1920928955078125e-7                 // 1 / 2**23
	MinAccuracyFloat64 = 2.2204460492503130808472633361816e-16 // 1 / 2**52
)

func EqualFloat32(x, y float32) bool {
	return EqualFloat32Ex(x, y, MinAccuracyFloat32)
}

func EqualFloat32Ex(x, y, epsilon float32) bool {
	return math.Abs(float64(x-y)) <= float64(epsilon)
}

func EqualFloat64(x, y float64) bool {
	return EqualFloat64Ex(x, y, MinAccuracyFloat64)
}

func EqualFloat64Ex(x, y, epsilon float64) bool {
	return math.Abs(x-y) <= epsilon
}

func EqualComplex64(x, y complex64) bool {
	return EqualComplex64Ex(x, y, MinAccuracyFloat32)
}

func EqualComplex64Ex(x, y complex64, epsilon float32) bool {
	return EqualFloat32Ex(real(x), real(y), epsilon) && EqualFloat32Ex(imag(x), imag(y), epsilon)
}

func EqualComplex128(x, y complex128) bool {
	return EqualComplex128Ex(x, y, MinAccuracyFloat64)
}

func EqualComplex128Ex(x, y complex128, epsilon float64) bool {
	return EqualFloat64Ex(real(x), real(y), epsilon) && EqualFloat64Ex(imag(x), imag(y), epsilon)
}
