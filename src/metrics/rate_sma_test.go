package metrics

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestRateSmaCalcFromPrevSampleWithDataCountOnSecond(t *testing.T) {
	testdata := []struct {
		incNum uint64
		wanted []float64
	}{
		{
			incNum: 1,
			wanted: []float64{1.0, 1 / 2.0, 1 / 3.0, 1 / 4.0, 1 / 5.0},
		},

		{
			incNum: 5,
			wanted: []float64{5.0, 6 / 2.0, 6 / 3.0, 6 / 4.0, 6 / 5.0},
		},
	}

	rate := NewRateSma(
		&RateSmaConfig{
			Name:              "A",
			MaxPeriod_ms:      5000,
			SampleInterval_ms: 1000,
		},
	)
	caps := make([]float64, 5, 5)
	time := uint64(0)

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			//t.Parallel()

			time += 1000

			rate.Inc(v.incNum)
			rate.Update(time)

			for i := 0; i < len(caps); i++ {
				caps[i] = rate.CalcByPeriod(uint64(i+1) * 1000)
			}

			test.EXPECT_EQ(t, caps, v.wanted, "")
		})
	}
}

func TestRateSmaCalcFromPrevSampleWithDataCountLessThenOneSecond(t *testing.T) {
	testdata := []struct {
		incNum uint64
		wanted []float64
	}{
		{
			incNum: 1,
			wanted: []float64{0.0, 0.0, 0.0, 0.0, 0.0},
		},

		{
			incNum: 5,
			wanted: []float64{6.0, 6 / 2.0, 6 / 3.0, 6 / 4.0, 6 / 5.0},
		},

		{
			incNum: 3,
			wanted: []float64{6.0, 6 / 2.0, 6 / 3.0, 6 / 4.0, 6 / 5.0},
		},

		{
			incNum: 4,
			wanted: []float64{7.0, 13 / 2.0, 13 / 3.0, 13 / 4.0, 13 / 5.0},
		},
	}

	rate := NewRateSma(
		&RateSmaConfig{
			Name:              "A",
			MaxPeriod_ms:      5000,
			SampleInterval_ms: 1000,
		},
	)

	caps := make([]float64, 5, 5)
	time := uint64(0)

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			//t.Parallel()

			time += 500

			rate.Inc(v.incNum)
			rate.Update(time)

			for i := 0; i < len(caps); i++ {
				caps[i] = rate.CalcByPeriod(uint64(i+1) * 1000)
			}

			test.EXPECT_EQ(t, caps, v.wanted, "")
		})
	}
}
