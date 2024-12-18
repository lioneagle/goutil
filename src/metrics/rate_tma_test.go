package metrics

import (
	"fmt"
	"testing"
	"time"

	"github.com/lioneagle/goutil/src/test"
)

type nower struct {
	top   int
	queue []time.Time
}

func (this *nower) Push(t ...time.Time) {
	for _, v := range t {
		this.queue = append(this.queue, v)
	}
}

func (this *nower) Now() time.Time {
	if len(this.queue) <= 0 {
		return time.Time{}
	}
	v := this.queue[0]
	this.queue = this.queue[1:]
	return v
}

func TestRateTmaCalc(t *testing.T) {
	testdata := []struct {
		incNum   uint64
		interval time.Duration
		wanted   float64
	}{
		{incNum: 1, interval: time.Millisecond * 20, wanted: 50.0},
		{incNum: 2, interval: time.Millisecond * 100, wanted: 25.0},
		{incNum: 1, interval: time.Millisecond * 280, wanted: 10.0},
		{incNum: 1, interval: time.Millisecond * 600, wanted: 5.0},
		{incNum: 1, interval: time.Millisecond * 1000, wanted: 3.0},
		{incNum: 1, interval: time.Millisecond * 1020, wanted: 2.0},
	}

	nower := &nower{}

	rate := NewRateTma("A", 1000, 5, nower.Now)

	time1 := nower.Now()
	nower.Push(time1)

	total := time.Duration(0)

	for _, v := range testdata {
		total += v.interval
		time2 := time1.Add(total)
		nower.Push(time2)
		nower.Push(time2)
	}

	rate.Inc(0)
	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			//t.Parallel()

			rate.Inc(v.incNum)

			caps := rate.Calc()

			test.EXPECT_EQ(t, caps, v.wanted, "")
		})
	}
}
