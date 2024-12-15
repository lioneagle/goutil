package metrics

type GaugeNumber[T GaugeContraints] struct {
	name string
	val  T
}

func NewGaugeNumber[T GaugeContraints](name string) *GaugeNumber[T] {
	return &GaugeNumber[T]{
		name: name,
	}
}

func (this *GaugeNumber[T]) GetName() string {
	return this.name
}

func (this *GaugeNumber[T]) GetGaugeType() string {
	return "number"
}

func (this *GaugeNumber[T]) Clear() {
	this.val = 0
}

func (this *GaugeNumber[T]) Inc(val T) {
	this.val += val
}

func (this *GaugeNumber[T]) Dec(val T) {
	this.val -= val
}

func (this *GaugeNumber[T]) SettValue(val T) {
	this.val = val
}

func (this *GaugeNumber[T]) GetValue(milisecond uint64) T {
	return this.val
}
