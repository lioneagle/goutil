package mathex

import (
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestEqualFloat32(t *testing.T) {
	var x, y float32

	x, y = 0.1, 0.1
	test.EXPECT_TRUE(t, EqualFloat32(x, y+MinAccuracyFloat32), "")
	test.EXPECT_TRUE(t, EqualFloat32(x, y-MinAccuracyFloat32), "")

	test.EXPECT_FALSE(t, EqualFloat32(x, y+2*MinAccuracyFloat32), "")
	test.EXPECT_FALSE(t, EqualFloat32(x, y-2*MinAccuracyFloat32), "")
}

func TestEqualFloat64(t *testing.T) {
	x, y := 0.1, 0.1
	test.EXPECT_TRUE(t, EqualFloat64(x, y+MinAccuracyFloat64), "")
	test.EXPECT_TRUE(t, EqualFloat64(x, y-MinAccuracyFloat64), "")
	test.EXPECT_TRUE(t, EqualFloat64Ex(x, y+MinAccuracyFloat32, MinAccuracyFloat32), "")
	test.EXPECT_TRUE(t, EqualFloat64Ex(x, y-MinAccuracyFloat32, MinAccuracyFloat32), "")

	test.EXPECT_FALSE(t, EqualFloat64(x, y+2*MinAccuracyFloat64), "")
	test.EXPECT_FALSE(t, EqualFloat64(x, y-2*MinAccuracyFloat64), "")
}

func TestEqualComplex64(t *testing.T) {
	var x, y, delta complex64

	x = 0.1 + 0.2i
	y = 0.1 + 0.2i
	delta = MinAccuracyFloat32 * (-1 + 1i)

	test.EXPECT_TRUE(t, EqualComplex64(x, y+delta), "")
	test.EXPECT_TRUE(t, EqualComplex64(x, y-delta), "")

	test.EXPECT_FALSE(t, EqualComplex64(x, y+2*delta), "")
	test.EXPECT_FALSE(t, EqualComplex64(x, y-2*delta), "")
}

func TestEqualComplex128(t *testing.T) {
	x := 0.1 + 0.2i
	y := 0.1 + 0.2i
	delta1 := MinAccuracyFloat64 * (-1 + 1i)
	delta2 := MinAccuracyFloat32 * (-1 + 1i)

	test.EXPECT_TRUE(t, EqualComplex128(x, y+delta1), "")
	test.EXPECT_TRUE(t, EqualComplex128(x, y-delta1), "")

	test.EXPECT_TRUE(t, EqualComplex128Ex(x, y+delta2, MinAccuracyFloat32), "")
	test.EXPECT_TRUE(t, EqualComplex128Ex(x, y-delta2, MinAccuracyFloat32), "")

	test.EXPECT_FALSE(t, EqualComplex128(x, y+2*delta1), "")
	test.EXPECT_FALSE(t, EqualComplex128(x, y-2*delta1), "")
}
