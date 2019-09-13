package sequence

import (
	"github.com/lioneagle/goutil/src/mathex"
)

type SequenceData interface {
	GetAt(index int) float64
	Len() int
}

func FindMax(data SequenceData, from, to int, precision float64) (float64, int) {
	max := data.GetAt(from)
	maxPos := from
	for i := from + 1; i < to; i++ {
		val := data.GetAt(i)
		if mathex.CompareFloat64Ex(val, max, precision) > 0 {
			max = val
			maxPos = i
		}
	}

	return max, maxPos
}

func FindMin(data SequenceData, from, to int, precision float64) (float64, int) {
	min := data.GetAt(from)
	minPos := from
	for i := from + 1; i < to; i++ {
		val := data.GetAt(i)
		if mathex.CompareFloat64Ex(val, min, precision) < 0 {
			min = val
			minPos = i
		}
	}

	return min, minPos
}

type SliceFloat64 struct {
	Data []float64
}

func NewSliceFloat64(num int) *SliceFloat64 {
	return &SliceFloat64{Data: make([]float64, num)}
}

func (slice *SliceFloat64) Len() int {
	return len(slice.Data)
}

func (slice *SliceFloat64) GetAt(index int) float64 {
	return slice.Data[index]
}

type SliceInt struct {
	Data []int
}

func NewSliceInt(num int) *SliceInt {
	return &SliceInt{Data: make([]int, num)}
}

func (slice *SliceInt) Len() int {
	return len(slice.Data)
}

func (slice *SliceInt) GetAt(index int) int {
	return slice.Data[index]
}
