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

func TesttCompareSliceFloat64Ex(t *testing.T) {
	testdata := []struct {
		x       []float64
		y       []float64
		epsilon float64
		wanted  int
	}{
		{[]float64{}, []float64{}, 0.00001, 0},
		{[]float64{1}, []float64{}, 0.00001, 1},
		{[]float64{1}, []float64{2, 3}, 0.00001, -1},
		{[]float64{1, 3}, []float64{2, 3}, 0.00001, -1},
		{[]float64{2, 3}, []float64{2, 3}, 0.00001, 0},
		{[]float64{2, 4}, []float64{2, 3}, 0.00001, 1},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			test.EXPECT_EQ(t, CompareSliceFloat64Ex(v.x, v.y, v.epsilon), v.wanted, "")
		})
	}
}

func TestCompareSliceFloat64(t *testing.T) {
	testdata := []struct {
		x      []float64
		y      []float64
		wanted int
	}{
		{[]float64{}, []float64{}, 0},
		{[]float64{1}, []float64{}, 1},
		{[]float64{1}, []float64{2, 3}, -1},
		{[]float64{1, 3}, []float64{2, 3}, -1},
		{[]float64{2, 3}, []float64{2, 3}, 0},
		{[]float64{2, 4}, []float64{2, 3}, 1},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			test.EXPECT_EQ(t, CompareSliceFloat64(v.x, v.y), v.wanted, "")
		})
	}
}

func TestCompareSliceFloat32Ex(t *testing.T) {
	testdata := []struct {
		x       []float32
		y       []float32
		epsilon float32
		wanted  int
	}{
		{[]float32{}, []float32{}, 0.00001, 0},
		{[]float32{1}, []float32{}, 0.00001, 1},
		{[]float32{1}, []float32{2, 3}, 0.00001, -1},
		{[]float32{1, 3}, []float32{2, 3}, 0.00001, -1},
		{[]float32{2, 3}, []float32{2, 3}, 0.00001, 0},
		{[]float32{2, 4}, []float32{2, 3}, 0.00001, 1},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			test.EXPECT_EQ(t, CompareSliceFloat32Ex(v.x, v.y, v.epsilon), v.wanted, "")
		})
	}
}

func TestEqualSliceFloat32(t *testing.T) {
	testdata := []struct {
		x      []float32
		y      []float32
		wanted int
	}{
		{[]float32{}, []float32{}, 0},
		{[]float32{1}, []float32{}, 1},
		{[]float32{1}, []float32{2, 3}, -1},
		{[]float32{1, 3}, []float32{2, 3}, -1},
		{[]float32{2, 3}, []float32{2, 3}, 0},
		{[]float32{2, 4}, []float32{2, 3}, 1},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			test.EXPECT_EQ(t, CompareSliceFloat32(v.x, v.y), v.wanted, "")
		})
	}
}
