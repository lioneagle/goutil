package sequence

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/mathex"
	"github.com/lioneagle/goutil/src/test"
)

func TestSequenceStatCalc(t *testing.T) {
	testdata := []struct {
		data   *SliceFloat64
		wanted *SequenceStat
	}{
		{&SliceFloat64{[]float64{}}, NewSequenceStat()},
		{&SliceFloat64{[]float64{8, 6, 5, 9, 1}}, &SequenceStat{Max: 9, Min: 1, Average: 5.8, Stdev: 3.1144823, Stdevp: 2.785677655}},
		{&SliceFloat64{[]float64{2, 3, 2, 7, 3, 4}}, &SequenceStat{Max: 7, Min: 2, Average: 3.5, Stdev: 1.870828693, Stdevp: 1.707825128}},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			stat := NewSequenceStat()
			stat.Calc(v.data)

			test.EXPECT_TRUE(t, mathex.CompareFloat64Ex(stat.Max, v.wanted.Max, 0.0000001) == 0, "stat = %v, wanted = %v", stat, v.wanted)
		})
	}
}
