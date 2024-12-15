package algorithm

import (
	"github.com/lioneagle/goutil/src/constraints"
)

// calculate Greate Common Divisor
func Gcd[T constraints.Integer](a, b T) T {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
