package mathex

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

func CompareSliceFloat64(x, y []float64) int {
	return CompareSliceFloat64Ex(x, y, MinAccuracyFloat64)
}

func CompareSliceFloat64Ex(x, y []float64, epsilon float64) int {
	if len(x) < len(y) {
		return -1
	}

	if len(x) > len(y) {
		return 1
	}

	for i := 0; i < len(x); i++ {
		ret := CompareFloat64Ex(x[i], y[i], epsilon)
		if ret != 0 {
			return ret
		}
	}

	return 0
}

func CompareSliceFloat32(x, y []float32) int {
	return CompareSliceFloat32Ex(x, y, MinAccuracyFloat32)
}

func CompareSliceFloat32Ex(x, y []float32, epsilon float32) int {
	if len(x) < len(y) {
		return -1
	}

	if len(x) > len(y) {
		return 1
	}

	for i := 0; i < len(x); i++ {
		ret := CompareFloat32Ex(x[i], y[i], epsilon)
		if ret != 0 {
			return ret
		}
	}

	return 0
}
