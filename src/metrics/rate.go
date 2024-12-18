package metrics

import (
	"strings"
	_ "time"

	"github.com/pkg/errors"
)

type Rate interface {
	Metric
	GetRateType() string
	GetSampleInterval() uint64 /* 获取采样间隔 */
	GetTotal() uint64          /* 获取计数总数 */
	Inc(num uint64)            /* 增加数据 */
	Period() uint64            /* 计算周期 */
	Calc() float64             /* 当前计算得到的速率 */
	Update(millisecond uint64) /* 更新计算数据 */
}

func NewRate(rateType, name string, sampleInterval_ms, period_ms uint64, timer Nower) (Rate, error) {
	switch strings.ToLower(rateType) {
	case "sma":
		return NewRateSma(&RateSmaConfig{
			Name:              name,
			MaxPeriod_ms:      period_ms,
			SampleInterval_ms: sampleInterval_ms,
		}), nil
	case "ema":
		return NewRateEma(name, sampleInterval_ms, period_ms), nil
	case "tma":
		return NewRateTma(name, sampleInterval_ms, period_ms, timer), nil
	default:
		return nil, errors.Errorf("unknown metric rate type: %s", rateType)
	}
}
