package mathex

import (
	"math"
)

const (
	MinAccuracyFloat32 = 1.1920928955078125e-7                 // 1 / 2**23
	MinAccuracyFloat64 = 2.2204460492503130808472633361816e-16 // 1 / 2**52
)

func CompareFloat32(x, y float32) int {
	return CompareFloat32Ex(x, y, MinAccuracyFloat32)
}

func CompareFloat32Ex(x, y, epsilon float32) int {
	d := x - y

	if d < -epsilon {
		return -1
	} else if d > epsilon {
		return 1
	}

	return 0
}

func CompareFloat64(x, y float64) int {
	return CompareFloat64Ex(x, y, MinAccuracyFloat64)
}

func CompareFloat64Ex(x, y, epsilon float64) int {
	d := x - y

	if d < -epsilon {
		return -1
	} else if d > epsilon {
		return 1
	}

	return 0
}
