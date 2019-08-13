package mathex

// Equal for complex

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
