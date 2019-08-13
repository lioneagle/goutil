package mathex

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestCompareFloat32(t *testing.T) {
	testdata := []struct {
		x      float32
		y      float32
		wanted int
	}{
		{0.1, 0.1 + MinAccuracyFloat32, 0},
		{0.1, 0.1 - MinAccuracyFloat32, 0},

		{0.1, 0.1 - 1.1*MinAccuracyFloat32, 1},
		{0.1, 0.1 - 2.0*MinAccuracyFloat32, 1},

		{0.1, 0.1 + 1.1*MinAccuracyFloat32, -1},
		{0.1, 0.1 + 2.0*MinAccuracyFloat32, -1},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			test.EXPECT_EQ(t, CompareFloat32(v.x, v.y), v.wanted, "")
		})
	}
}

func TestCompareFloat64(t *testing.T) {
	testdata := []struct {
		x      float64
		y      float64
		wanted int
	}{
		{0.1, 0.1 + MinAccuracyFloat64, 0},
		{0.1, 0.1 - MinAccuracyFloat64, 0},

		{0.1, 0.1 - 1.1*MinAccuracyFloat64, 1},
		{0.1, 0.1 - 2.0*MinAccuracyFloat64, 1},

		{0.1, 0.1 + 1.1*MinAccuracyFloat64, -1},
		{0.1, 0.1 + 2.0*MinAccuracyFloat64, -1},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			test.EXPECT_EQ(t, CompareFloat64(v.x, v.y), v.wanted, "")
		})
	}
}

func TestCompareFloat64Ex(t *testing.T) {
	testdata := []struct {
		x       float64
		y       float64
		epsilon float64
		wanted  int
	}{
		{0.1, 0.1 + MinAccuracyFloat32, MinAccuracyFloat32, 0},
		{0.1, 0.1 - MinAccuracyFloat32, MinAccuracyFloat32, 0},

		{0.1, 0.1 - 1.1*MinAccuracyFloat32, MinAccuracyFloat32, 1},
		{0.1, 0.1 - 2.0*MinAccuracyFloat32, MinAccuracyFloat32, 1},

		{0.1, 0.1 + 1.1*MinAccuracyFloat32, MinAccuracyFloat32, -1},
		{0.1, 0.1 + 2.0*MinAccuracyFloat32, MinAccuracyFloat32, -1},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			test.EXPECT_EQ(t, CompareFloat64Ex(v.x, v.y, v.epsilon), v.wanted, "")
		})
	}
}
