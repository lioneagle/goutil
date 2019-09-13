package basic

import (
	"github.com/lioneagle/goutil/src/mathex"
)

type Drawdown struct {
	Begin       int     // 回撤开始下标
	End         int     // 回撤结束下标
	Recover     int     // 回撤恢复下标
	Rate        float32 // 回撤百分比
	IsRecovered bool    // 是否结束回撤
}

func NewDrawdown() *Drawdown {
	return &Drawdown{}
}

type Drawdowns struct {
	beginValue float64     // 起始值
	endValue   float64     // 最低值
	Data       []*Drawdown // 所有的回撤记录
}

func NewDrawdowns() *Drawdowns {
	ret := &Drawdowns{}
	ret.Data = append(ret.Data, NewDrawdown())
	return ret
}

func (d *Drawdowns) CalcData(data mathex.SequenceData, precision float64) {
	for i := 0; i < data.Len(); i++ {
		d.Calc(i, data.GetAt(i), precision)
	}
	d.CalcEnd()
}

/* drawdown算法原理：先找到最高点，如果开始下降，则纪录最低点，当从最低点上升到超过原有的最高点时，
   产生一次回撤，回撤最高点为前最高点，回撤最低点为最低点
*/
func (d *Drawdowns) Calc(index int, val float64, precision float64) {
	drawdown := d.Data[len(d.Data)-1]

	if index == 0 {
		if len(d.Data) > 0 {
			d.Data = d.Data[:1]
		}

		d.calcFirst(drawdown, index, val)
		return
	}

	if mathex.CompareFloat64Ex(val, d.beginValue, precision) >= 0 {
		if mathex.CompareFloat64Ex(d.beginValue, d.endValue, precision) > 0 {
			drawdown.IsRecovered = true
			drawdown.Recover = index
			drawdown.Rate = float32(1.0 - d.endValue/d.beginValue)
			drawdown = NewDrawdown()
			d.Data = append(d.Data, drawdown)
		}
		d.calcFirst(drawdown, index, val)
	} else if mathex.CompareFloat64Ex(val, d.endValue, precision) < 0 {
		d.endValue = val
		drawdown.End = index
		drawdown.Recover = index
		drawdown.Rate = float32(1.0 - d.endValue/d.beginValue)
	} else {
		drawdown.Recover = index
	}
}

func (d *Drawdowns) CalcEnd() {
	drawdown := d.Data[len(d.Data)-1]
	drawdown.Rate = float32(1.0 - d.endValue/d.beginValue)
}

func (d *Drawdowns) calcFirst(drawdown *Drawdown, index int, val float64) {
	d.beginValue = val
	d.endValue = val
	drawdown.IsRecovered = false
	drawdown.Begin = index
	drawdown.End = index
	drawdown.Recover = index
	drawdown.Rate = 0.0
}
