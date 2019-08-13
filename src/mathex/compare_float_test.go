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
	test.EXPECT_FALSE(t, EqualFloat64Ex(x, y+2*MinAccuracyFloat32, MinAccuracyFloat32), "")
	test.EXPECT_FALSE(t, EqualFloat64Ex(x, y-2*MinAccuracyFloat32, MinAccuracyFloat32), "")
}

func TestGreateFloat32(t *testing.T) {
	var x, y float32

	x, y = 0.1, 0.1

	test.EXPECT_TRUE(t, GreateFloat32(x+1.1*MinAccuracyFloat32, y), "")
	test.EXPECT_TRUE(t, GreateFloat32(x+2*MinAccuracyFloat32, y), "")

	test.EXPECT_FALSE(t, GreateFloat32(x+MinAccuracyFloat32, y), "")
	test.EXPECT_FALSE(t, GreateFloat32(x+0.99*MinAccuracyFloat32, y), "")
}

func TestGreateFloat64(t *testing.T) {
	x, y := 0.1, 0.1
	test.EXPECT_TRUE(t, GreateFloat64(x+1.2*MinAccuracyFloat64, y), "")
	test.EXPECT_TRUE(t, GreateFloat64(x+2*MinAccuracyFloat64, y), "")
	test.EXPECT_TRUE(t, GreateFloat64Ex(x+1.01*MinAccuracyFloat32, y, MinAccuracyFloat32), "")
	test.EXPECT_TRUE(t, GreateFloat64Ex(x+2*MinAccuracyFloat32, y, MinAccuracyFloat32), "")

	test.EXPECT_FALSE(t, GreateFloat64(x+MinAccuracyFloat64, y), "")
	test.EXPECT_FALSE(t, GreateFloat64(x+0.99*MinAccuracyFloat64, y), "")
	test.EXPECT_FALSE(t, GreateFloat64Ex(x+MinAccuracyFloat32, y, MinAccuracyFloat32), "")
	test.EXPECT_FALSE(t, GreateFloat64Ex(x+0.99*MinAccuracyFloat32, y, MinAccuracyFloat32), "")
}

func TestGreateThanFloat32(t *testing.T) {
	var x, y float32

	x, y = 0.1, 0.1
	test.EXPECT_TRUE(t, GreateThanFloat32(x+MinAccuracyFloat32, y), "")
	test.EXPECT_TRUE(t, GreateThanFloat32(x+1.01*+MinAccuracyFloat32, y), "")

	test.EXPECT_FALSE(t, GreateThanFloat32(x+0.9*MinAccuracyFloat32, y), "")
}

func TestGreateThanFloat64(t *testing.T) {
	x, y := 0.1, 0.1
	test.EXPECT_TRUE(t, GreateThanFloat64(x+MinAccuracyFloat64, y), "")
	test.EXPECT_TRUE(t, GreateThanFloat64(x+1.01*MinAccuracyFloat64, y), "")
	test.EXPECT_TRUE(t, GreateThanFloat64Ex(x+MinAccuracyFloat32, y, MinAccuracyFloat32), "")
	test.EXPECT_TRUE(t, GreateThanFloat64Ex(x+1.01*MinAccuracyFloat32, y, MinAccuracyFloat32), "")

	test.EXPECT_FALSE(t, GreateThanFloat64(x+0.9*MinAccuracyFloat64, y), "")
	test.EXPECT_FALSE(t, GreateThanFloat64Ex(x+0.99*MinAccuracyFloat32, y, MinAccuracyFloat32), "")
}

func TestLessFloat32(t *testing.T) {
	var x, y float32

	x, y = 0.1, 0.1
	test.EXPECT_TRUE(t, LessFloat32(x-1.01*MinAccuracyFloat32, y), "")
	test.EXPECT_TRUE(t, LessFloat32(x-2*MinAccuracyFloat32, y), "")

	test.EXPECT_FALSE(t, LessFloat32(x-MinAccuracyFloat32, y), "")
	test.EXPECT_FALSE(t, LessFloat32(x-0.99*MinAccuracyFloat32, y), "")
}

func TestLessFloat64(t *testing.T) {
	x, y := 0.1, 0.1
	test.EXPECT_TRUE(t, LessFloat64(x-1.01*MinAccuracyFloat64, y), "")
	test.EXPECT_TRUE(t, LessFloat64(x-2*MinAccuracyFloat64, y), "")
	test.EXPECT_TRUE(t, LessFloat64Ex(x-1.01*MinAccuracyFloat32, y, MinAccuracyFloat32), "")
	test.EXPECT_TRUE(t, LessFloat64Ex(x-2*MinAccuracyFloat32, y, MinAccuracyFloat32), "")

	test.EXPECT_FALSE(t, LessFloat64(x-MinAccuracyFloat64, y), "")
	test.EXPECT_FALSE(t, LessFloat64(x-0.99*MinAccuracyFloat64, y), "")
	test.EXPECT_FALSE(t, LessFloat64Ex(x-MinAccuracyFloat32, y, MinAccuracyFloat32), "")
	test.EXPECT_FALSE(t, LessFloat64Ex(x-0.99*MinAccuracyFloat32, y, MinAccuracyFloat32), "")
}

func TestLessThanFloat32(t *testing.T) {
	var x, y float32

	x, y = 0.1, 0.1
	test.EXPECT_TRUE(t, LessThanFloat32(x-MinAccuracyFloat32, y), "")
	test.EXPECT_TRUE(t, LessThanFloat32(x-1.01*+MinAccuracyFloat32, y), "")

	test.EXPECT_FALSE(t, LessThanFloat32(x-0.99*MinAccuracyFloat32, y), "")
}

func TestLessThanFloat64(t *testing.T) {
	x, y := 0.1, 0.1
	test.EXPECT_TRUE(t, LessThanFloat64(x-MinAccuracyFloat64, y), "")
	test.EXPECT_TRUE(t, LessThanFloat64(x-1.01*MinAccuracyFloat64, y), "")
	test.EXPECT_TRUE(t, LessThanFloat64Ex(x-MinAccuracyFloat32, y, MinAccuracyFloat32), "")
	test.EXPECT_TRUE(t, LessThanFloat64Ex(x-1.01*MinAccuracyFloat32, y, MinAccuracyFloat32), "")

	test.EXPECT_FALSE(t, LessThanFloat64(x-0.99*MinAccuracyFloat64, y), "")
	test.EXPECT_FALSE(t, LessThanFloat64Ex(x-0.99*MinAccuracyFloat32, y, MinAccuracyFloat32), "")
}
