package metrics

type GaugeFunctional[T GaugeContraints] struct {
	name string
	val  func(milisecond uint64) T
}

func NewGaugeFunctional[T GaugeContraints](name string, f func(milisecond uint64) T) *GaugeFunctional[T] {
	return &GaugeFunctional[T]{
		name: name,
		val:  f,
	}
}

func (this *GaugeFunctional[T]) GetName() string {
	return this.name
}

func (this *GaugeFunctional[T]) GetGaugeType() string {
	return "functional"
}

func (this *GaugeFunctional[T]) Clear()          {}
func (this *GaugeFunctional[T]) Inc(val T)       {}
func (this *GaugeFunctional[T]) Dec(val T)       {}
func (this *GaugeFunctional[T]) SettValue(val T) {}
func (this *GaugeFunctional[T]) GetValue(milisecond uint64) T {
	return this.val(milisecond)
}
