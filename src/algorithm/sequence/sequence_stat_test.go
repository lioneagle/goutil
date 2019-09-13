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
		{&SliceFloat64{[]float64{8, 6, 5, 9, 1}}, &SequenceStat{max: 9, min: 1, average: 5.8, stdev: 3.1144823, stdevp: 2.785677655}},
		{&SliceFloat64{[]float64{2, 3, 2, 7, 3, 4}}, &SequenceStat{max: 7, min: 2, average: 3.5, stdev: 1.870828693, stdevp: 1.707825128}},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			stat := NewSequenceStat()
			stat.Calc(v.data)

			test.EXPECT_TRUE(t, mathex.CompareFloat64Ex(stat.max, v.wanted.max, 0.0000001) == 0, "stat = %v, wanted = %v", stat, v.wanted)
		})
	}
}
