package sequence

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestDrawdownsCalcData(t *testing.T) {
	testdata := []struct {
		data   *SliceFloat64
		wanted *Drawdowns
	}{
		{&SliceFloat64{[]float64{8, 6, 2, 9, 3}}, &Drawdowns{9, 3, []*Drawdown{&Drawdown{0, 2, 3, 0.75, true}, &Drawdown{3, 4, 4, 2.0 / 3.0, false}}}},
		{&SliceFloat64{[]float64{8, 6, 2, 9, 3, 4}}, &Drawdowns{9, 3, []*Drawdown{&Drawdown{0, 2, 3, 0.75, true}, &Drawdown{3, 4, 5, 2.0 / 3.0, false}}}},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			d := NewDrawdowns()
			d.CalcData(v.data, 1e-8)

			test.EXPECT_EQ(t, d, v.wanted, "")
		})
	}
}
