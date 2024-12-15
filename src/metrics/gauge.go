package metrics

type GaugeContraints interface {
	int64 | float64
}

type Gauge[T GaugeContraints] interface {
	Metric
	GetGaugeType() string
	Clear()
	Inc(val T)
	Dec(val T)
	SettValue(val T)
	GetValue(milisecond uint64) T
}
