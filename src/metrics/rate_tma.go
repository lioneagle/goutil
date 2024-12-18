package metrics

import (
	//"fmt"
	"time"
)

type Nower func() time.Time

type tmaNode struct {
	timestamp time.Time
	count     uint64
}

type RateTmaConfig struct {
	Name              string // 名称
	SampleInterval_ms uint64 // 采样间隔，单位：毫秒s
	MaxPeriod_ms      uint64 // 最大计算周期, 单位：毫秒
	Timer             Nower
}

/* RateTma以当前时间与第一个采样点时间差来计算平均速率
 * TMA = sum(X(1)...X(n)) / (Now - T(1))
 */
type RateTma struct {
	name              string    // 名称
	sampleInterval_ms uint64    // 采样间隔，单位：毫秒，用于定时填充空包
	totalSamples      uint64    // 采样总数
	total             uint64    // 当前计数总数
	maxPeriod_ms      uint64    // 最大计算周期, 单位：毫秒
	ring              []tmaNode // 用于计算速率的计数环
	timer             Nower     // 获取当前时间的函数接口
}

/* NewRateTma创建一个以当前时间与第一个采样点时间差来计算平均速率的metric
 *
 * 比如，MaxPeriod=100，则RateTma可以输出以当前时间与第一个采样点时间差来计算平均速率
 *（最多100个最近的采样点），这些采样点的时间间隔不一定相等，比如第一个采样点（数量为2）时
 * 间戳为1s，第二个采样点（数量为1）时间戳为1.2s，
 * 当前时间戳为2s，则当前平均速率为: (3+1)/ (2-1) = 4
 * 当前时间戳为2.5s，则当前平均速率为: (3+1)/ (2.5-1) = 4/1.5 = 8/3
 */
func NewRateTma(config *RateTmaConfig) *RateTma {
	capacity := config.MaxPeriod_ms/config.SampleInterval_ms + 1

	ret := &RateTma{
		name:              config.Name,
		sampleInterval_ms: config.SampleInterval_ms,
		maxPeriod_ms:      config.MaxPeriod_ms,
		ring:              make([]tmaNode, capacity, capacity),
		timer:             config.Timer,
	}

	if config.Timer == nil {
		ret.timer = time.Now
	}

	ret.Clear()
	return ret
}

func (this *RateTma) GetName() string {
	return this.name
}

func (this *RateTma) GetRateType() string {
	return "tma"
}

func (this *RateTma) GetSampleInterval() uint64 {
	return this.sampleInterval_ms
}

func (this *RateTma) GetTotal() uint64 {
	return this.total
}

func (this *RateTma) Len() int {
	return len(this.ring)
}

func (this *RateTma) Inc(num uint64) {
	index := this.totalSamples % uint64(this.Len())
	this.total += num
	this.ring[index].count = this.total
	this.ring[index].timestamp = this.timer()
	this.totalSamples++
}

func (this *RateTma) CanUpdate(millisecond uint64) bool {
	return (millisecond % this.sampleInterval_ms) == 0
}

func (this *RateTma) Update(millisecond uint64) {
	if !this.CanUpdate(millisecond) {
		return
	}

	this.Inc(0)
}

func (this *RateTma) Clear() {
	now := this.timer()
	for i := range this.ring {
		this.ring[i].count = this.total
		this.ring[i].timestamp = now
	}
}

func (this *RateTma) Period() uint64 {
	return this.maxPeriod_ms
}

func (this *RateTma) Calc() float64 {
	return this.CalcByPeriod(this.maxPeriod_ms)
}

func (this *RateTma) CalcByPeriod(period uint64) float64 {
	return this.calcFromPrevSample(period)
}

func (this *RateTma) calcFromPrevSample(period uint64) float64 {
	now := this.timer()
	ringLen := uint64(this.Len())
	curr := (this.totalSamples - 1 + ringLen) % ringLen
	prev := (this.totalSamples - 1 - period/this.sampleInterval_ms + ringLen) % ringLen

	return float64(this.ring[curr].count-this.ring[prev].count) /
		now.Sub(this.ring[prev].timestamp).Seconds()
}
