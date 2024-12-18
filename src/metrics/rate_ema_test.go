package metrics

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/lioneagle/goutil/src/test"
)

func TestRateEmaCalc(t *testing.T) {
	period := 9

	rate1 := NewRateEma(&RateEmaConfig{
		Name:              "A",
		SampleInterval_ms: 1,
		Period_ms:         uint64(period),
	})
	capacity := 100

	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	num := make([]uint64, capacity, capacity)
	for i := 0; i < capacity; i++ {
		num[i] = r1.Uint64() % 10000
		//num[i] = uint64((i + 1) * (i + 1))
		rate1.Inc(num[i])
		ma := 0.0
		if i < period {
			for j := 0; j <= i; j++ {
				ma += float64(num[j])
			}
			ma /= float64(i + 1)
		} else {
			for j := i - period + 1; j <= i; j++ {
				ma += float64(num[j])
			}
			ma /= float64(period)
		}
		rate1.Update(uint64(i))
		ema := rate1.Calc() / 1000.0
		fmt.Printf("[%d]: rate = %.2f, ma = %.2f, ema = %.2f\n",
			i, float64(num[i]), ma, ema)
	}

	testdata := []struct {
		incNum uint64
		wanted float64
	}{
		{incNum: 1, wanted: 1.0},
		{incNum: 4, wanted: 2.0},
	}

	rate := NewRateEma(&RateEmaConfig{
		Name:              "A",
		SampleInterval_ms: 1000,
		Period_ms:         5000,
	})

	time := uint64(0)

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			//t.Parallel()

			time += 1000

			rate.Inc(v.incNum)
			rate.Update(time)

			caps := rate.Calc()

			test.EXPECT_EQ(t, caps, v.wanted, "")
		})
	}

}
