package sequence

import (
	//"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestSequenceDataFindMaxMin(t *testing.T) {
	slice := &SliceFloat64{[]float64{8, 6, 2, 9, 3}}

	max, maxPos := FindMax(slice, 0, slice.Len(), 1e-8)
	min, minPos := FindMin(slice, 0, slice.Len(), 1e-8)

	test.EXPECT_EQ(t, slice.Len(), 5, "")
	test.EXPECT_EQ(t, max, float64(9.0), "")
	test.EXPECT_EQ(t, maxPos, 3, "")
	test.EXPECT_EQ(t, min, float64(2.0), "")
	test.EXPECT_EQ(t, minPos, 2, "")
}

func TestSliceFloat64(t *testing.T) {
	slice := NewSliceFloat64(4)
	slice.Data[2] = 6

	test.EXPECT_EQ(t, slice.Len(), 4, "")
	test.EXPECT_EQ(t, slice.GetAt(2), float64(6.0), "")
}

func TestSliceInt(t *testing.T) {
	slice := NewSliceInt(7)
	slice.Data[3] = 17

	test.EXPECT_EQ(t, slice.Len(), 7, "")
	test.EXPECT_EQ(t, slice.GetAt(3), 17, "")
}
