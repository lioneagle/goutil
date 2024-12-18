package metrics

import (
	_ "fmt"
)

type RateSmaConfig struct {
	Name              string // 名称
	MaxPeriod_ms      uint64 // 最大计算周期, 单位：毫秒
	SampleInterval_ms uint64 // 采样间隔，单位：毫秒
}

/* RateSma用于计算SMA（Simple Moving Average）速率的metric
 * SMA = sum(X(1)...X(n)) / period
 */
type RateSma struct {
	name              string   // 名称
	sampleInterval_ms uint64   // 采样间隔，单位：毫秒
	totalSamples      uint64   // 采样总数
	total             uint64   // 计数总数
	maxPeriod_ms      uint64   // 最大计算周期, 单位：毫秒
	ring              []uint64 // 用于计算速率的计数环
}

/* NewRateSma创建一个计算SMA（Simple Moving Average）速率的metric
 *
 * 比如，MaxPeriod=10000，SampleInterval_ms=1000，则RateSma可以输出按1秒到按10秒计算的SMA速率
 *
 * 目前可以计算1-N（N = maxPeriod_ms/sampleInterval_ms）之间任意一个周期的SMA，如果只考虑最大
 * 周期数，可以在Update时计算好，Calc时之间返回计算好的值
 */
func NewRateSma(config *RateSmaConfig) *RateSma {
	capacity := config.MaxPeriod_ms/config.SampleInterval_ms + 1

	return &RateSma{
		name:              config.Name,
		sampleInterval_ms: config.SampleInterval_ms,
		ring:              make([]uint64, capacity, capacity),
	}
}

func (this *RateSma) GetName() string {
	return this.name
}

func (this *RateSma) GetRateType() string {
	return "sma"
}

func (this *RateSma) GetSampleInterval() uint64 {
	return this.sampleInterval_ms
}

func (this *RateSma) GetTotal() uint64 {
	return this.total
}

func (this *RateSma) Len() int {
	return len(this.ring)
}

func (this *RateSma) Inc(num uint64) {
	this.total += num
}

func (this *RateSma) CanUpdate(millisecond uint64) bool {
	return (millisecond % this.sampleInterval_ms) == 0
}

func (this *RateSma) Update(millisecond uint64) {
	if !this.CanUpdate(millisecond) {
		return
	}

	/*if this.totalSamples == 0 {
		for i := 0; i < len(this.ring); i++ {
			this.ring[i] = this.total
		}
	} else*/{
		index := this.totalSamples % uint64(this.Len())
		this.ring[index] = this.total
	}

	this.totalSamples++
}

func (this *RateSma) Clear() {
	for i := range this.ring {
		this.ring[i] = this.total
	}
}

func (this *RateSma) Period() uint64 {
	return this.maxPeriod_ms
}

func (this *RateSma) Calc() float64 {
	return this.CalcByPeriod(this.maxPeriod_ms)
}

func (this *RateSma) CalcByPeriod(period uint64) float64 {
	return this.calcFromPrevSample(period)
}

/* calcFromPrevSample从前一采样周期开始计算，不包含当前秒
 * period: 计算SMA的周期（毫秒）
 *
 * 比如，period=1000，则从前一个采样周期开始计算1秒的速率，
 * 假设period=1000，this.sampleTime=1000
 * 假设当前时刻为1.5秒，0秒时刻计数为10，1秒时刻计数为20，当前计数为25，则函数返回速率为10.0
 * 假设当前时刻为2.3秒，0秒时刻计数为10，1秒时刻计数为20，2秒时刻计数为26，当前计数为29，则函数返回速率为6.0
 */
func (this *RateSma) calcFromPrevSample(period uint64) float64 {
	/* @@NOTE:
	 * 该算法利用最开始的初始值都为0秒时的值，来避开开始几个计算时的if判断，正常的算法应该如下：
	 * if this.totalSamples < uint64(this.maxPeriod_ms/this.sampleInterval_ms) {
	 *	 prev = 0
	 * }
	 *
	 */
	ringLen := uint64(this.Len())
	curr := (this.totalSamples - 1 + ringLen) % ringLen
	prev := (this.totalSamples - 1 - uint64(period/this.sampleInterval_ms) + ringLen) % ringLen

	/*fmt.Printf("ring=%v, curr=%d, prev=%d, samples = %d\n",
	this.ring, curr, prev, this.totalSamples)*/
	return float64(this.ring[curr]-this.ring[prev]) * 1000.0 / float64(period)
}
