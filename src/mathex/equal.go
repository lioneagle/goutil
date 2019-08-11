package mathex

import (
	"math"
)

func EqualFloat32(x, y float32) bool {
	return EqualFloat32Ex(x, y, math.SmallestNonzeroFloat32)
}

func EqualFloat32Ex(x, y, epsilon float32) bool {
	return math.Abs(float64(x-y)) < float64(epsilon)
}

func EqualFloat64(x, y float64) bool {
	return EqualFloat64Ex(x, y, math.SmallestNonzeroFloat64)
}

func EqualFloat64Ex(x, y, epsilon float64) bool {
	return math.Abs(x-y) < epsilon
}

func EqualComplex64(x, y complex64) bool {
	return EqualComplex64Ex(x, y, math.SmallestNonzeroFloat32)
}

func EqualComplex64Ex(x, y complex64, epsilon float32) bool {
	return EqualFloat32Ex(real(x), real(y), epsilon) && EqualFloat32Ex(imag(x), imag(y), epsilon)
}

func EqualComplex128(x, y complex128) bool {
	return EqualComplex128Ex(x, y, math.SmallestNonzeroFloat64)
}

func EqualComplex128Ex(x, y complex128, epsilon float64) bool {
	return EqualFloat64Ex(real(x), real(y), epsilon) && EqualFloat64Ex(imag(x), imag(y), epsilon)
}
