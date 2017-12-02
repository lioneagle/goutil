package chars

import (
	"strings"
)

func Filter(src []string, filter []string) []string {
	var ret []string

	for i := 0; i < len(src); i++ {
		if Contains(src[i], filter) {
			ret = append(ret, src[i])
		}
	}
	return ret
}

func FilterReverse(src []string, filter []string) []string {
	var ret []string

	for i := 0; i < len(src); i++ {
		if !Contains(src[i], filter) {
			ret = append(ret, src[i])
		}
	}
	return ret
}

func Contains(src string, filter []string) bool {
	for i := 0; i < len(filter); i++ {
		if strings.Contains(src, filter[i]) {
			return true
		}
	}
	return false
}

func StringsEqual(lhs, rhs []string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}
