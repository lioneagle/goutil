package mathex

import (
	"math"
)

const (
	MinAccuracyFloat32 = 1.1920928955078125e-7                 // 1 / 2**23
	MinAccuracyFloat64 = 2.2204460492503130808472633361816e-16 // 1 / 2**52
)

// Equal for float

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

// Greate for float

func GreateFloat32(x, y float32) bool {
	return GreateFloat32Ex(x, y, MinAccuracyFloat32)
}

func GreateFloat32Ex(x, y, epsilon float32) bool {
	return x > (y + epsilon)
}

func GreateFloat64(x, y float64) bool {
	return GreateFloat64Ex(x, y, MinAccuracyFloat64)
}

func GreateFloat64Ex(x, y, epsilon float64) bool {
	return x > (y + epsilon)
}

// GreateThan for float

func GreateThanFloat32(x, y float32) bool {
	return GreateThanFloat32Ex(x, y, MinAccuracyFloat32)
}

func GreateThanFloat32Ex(x, y, epsilon float32) bool {
	return x >= (y + epsilon)
}

func GreateThanFloat64(x, y float64) bool {
	return GreateThanFloat64Ex(x, y, MinAccuracyFloat64)
}

func GreateThanFloat64Ex(x, y, epsilon float64) bool {
	return x >= (y + epsilon)
}

// Less for float

func LessFloat32(x, y float32) bool {
	return LessFloat32Ex(x, y, MinAccuracyFloat32)
}

func LessFloat32Ex(x, y, epsilon float32) bool {
	return x < (y - epsilon)
}

func LessFloat64(x, y float64) bool {
	return LessFloat64Ex(x, y, MinAccuracyFloat64)
}

func LessFloat64Ex(x, y, epsilon float64) bool {
	return x < (y - epsilon)
}

// LessThan for float

func LessThanFloat32(x, y float32) bool {
	return LessThanFloat32Ex(x, y, MinAccuracyFloat32)
}

func LessThanFloat32Ex(x, y, epsilon float32) bool {
	return x <= (y - epsilon)
}

func LessThanFloat64(x, y float64) bool {
	return GreateThanFloat64Ex(x, y, MinAccuracyFloat64)
}

func LessThanFloat64Ex(x, y, epsilon float64) bool {
	return x <= (y - epsilon)
}
