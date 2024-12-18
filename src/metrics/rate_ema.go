package metrics

/* RateEma用于计算EMA（Exponential Moving Average）速率的metric
 * EMA(n) = EMA(n-1) + alpha * [X(n) - EMA(n-1)]
 *
 */
type RateEma struct {
	name              string  // 名称
	sampleInterval_ms uint64  // 采样间隔，单位：毫秒
	alpha             float64 // 加权系数，范围：(0,1)
	totalSamples      uint64  // 采样总数
	total             uint64  // 计数总数
	totalLast         uint64  // 上个采样周期结束时计数总数
	period_ms         uint64  // 计算周期, 单位：毫秒
	rate              float64 // 上个周期的计算结果
}

/* NewRateEma创建一个计算EMA（Exponential Moving Average）速率的metric
 *
 * name: 名称
 * sampleInterval_ms: 采样间隔，单位：毫秒
 * period_ms: 输出的周期，单位：毫秒
 * N = period_ms / sampleInterval_ms
 * alpha取2 / (N + 1)时，对于直线，EMA的计算值能与N个周期的SMA的值基本是一致的
 * alpha缺省取 2 / (N + 1)
 *
 */
func NewRateEma(name string, sampleInterval_ms, period_ms uint64) *RateEma {
	return &RateEma{
		name:              name,
		sampleInterval_ms: sampleInterval_ms,
		period_ms:         period_ms,
		alpha:             2.0 / (float64(period_ms)/float64(sampleInterval_ms) + 1.0),
	}
}

func (this *RateEma) GetName() string {
	return this.name
}

func (this *RateEma) GetRateType() string {
	return "ema"
}

func (this *RateEma) GetSampleInterval() uint64 {
	return this.sampleInterval_ms
}

func (this *RateEma) GetTotal() uint64 {
	return this.total
}

func (this *RateEma) Inc(num uint64) {
	this.total += num
}

func (this *RateEma) Period() uint64 {
	return this.period_ms
}

func (this *RateEma) CanUpdate(millisecond uint64) bool {
	return (millisecond % this.sampleInterval_ms) == 0
}

func (this *RateEma) Update(millisecond uint64) {
	if !this.CanUpdate(millisecond) {
		return
	}

	rate := float64(this.total-this.totalLast) / float64(this.sampleInterval_ms) * 1000.0

	if this.totalSamples == 0 {
		this.rate = rate
	} else {
		this.rate += this.alpha * (rate - this.rate)
	}

	this.totalSamples++
	this.totalLast = this.total
}

func (this *RateEma) Calc() float64 {
	return this.rate
}
