package sequence

import (
	"math"
)

type SequenceStat struct {
	max     float64
	min     float64
	average float64
	stdev   float64 // 样本标准差
	stdevp  float64 // 总体标准差
}

func NewSequenceStat() *SequenceStat {
	ret := &SequenceStat{}
	ret.Clear()
	return ret
}

func (stat *SequenceStat) Clear() {
	stat.max = -math.SmallestNonzeroFloat64
	stat.min = math.MaxFloat64
	stat.average = 0.0
	stat.stdev = 0.0
}

func (stat *SequenceStat) Calc(data SequenceData) {
	stat.Clear()

	if data.Len() <= 0 {
		return
	}

	N := float64(data.Len())
	squareSum := 0.0
	average := 0.0

	for i := 0; i < data.Len(); i++ {
		v := data.GetAt(i)

		if v > stat.max {
			stat.max = v
		}

		if v < stat.min {
			stat.min = v
		}

		average += v
		squareSum += v * v
	}

	stat.average = average / N

	if N > 1 {
		stat.stdev = math.Sqrt((squareSum + N*average*average) / (N - 1))
		stat.stdevp = math.Sqrt((squareSum + N*average*average) / N)
	}
}
