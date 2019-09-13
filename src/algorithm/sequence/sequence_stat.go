package sequence

import (
	"math"
)

type SequenceStat struct {
	Size   int
	Max     float64
	Min     float64
	Average float64
	Stdev   float64 // 样本标准差
	Stdevp  float64 // 总体标准差
}

func NewSequenceStat() *SequenceStat {
	ret := &SequenceStat{}
	ret.Clear()
	return ret
}

func (stat *SequenceStat) Clear() {
	stat.Max = -math.SmallestNonzeroFloat64
	stat.Min = math.MaxFloat64
	stat.Average = 0.0
	stat.Stdev = 0.0
	stat.Size = 0
}

func (stat *SequenceStat) Calc(data SequenceData) {
	stat.Clear()

	if data.Len() <= 0 {
		return
	}

	N := float64(data.Len())
	stat.Size = data.Len()
	squareSum := 0.0
	average := 0.0

	for i := 0; i < data.Len(); i++ {
		v := data.GetAt(i)

		if v > stat.Max {
			stat.Max = v
		}

		if v < stat.Min {
			stat.Min = v
		}

		average += v
		squareSum += v * v
	}

	stat.Average = average / N

	if N > 1 {
		stat.Stdev = math.Sqrt((squareSum + N*average*average) / (N - 1))
		stat.Stdevp = math.Sqrt((squareSum + N*average*average) / N)
	}
}
