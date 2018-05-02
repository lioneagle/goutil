package chars

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

func Exists(src []string, name string) bool {
	for i := 0; i < len(src); i++ {
		if name == src[i] {
			return true
		}
	}
	return false
}

func Remove(src []string, name string) []string {
	for i := 0; i < len(src); i++ {
		if name == src[i] {
			src[i] = src[len(src)-1]
			src = src[:len(src)-1]
			break
		}
	}
	return src
}

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

func ByteSlicePackSpace(src []byte) []byte {
	var ret []byte
	src = bytes.TrimSpace(src)
	currentIsSpace := false
	for i := 0; i < len(src); i++ {
		if IsWspChar(src[i]) {
			if currentIsSpace {
				continue
			}
			currentIsSpace = true
			ret = append(ret, ' ')
		} else {
			currentIsSpace = false
			ret = append(ret, src[i])
		}
	}
	return ret
}

func StringPackSpace(src string) string {
	return ByteSliceToString(ByteSlicePackSpace(StringToByteSlice(src)))
}

func StringToByteSlice(str string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	retHeader := reflect.SliceHeader{Data: strHeader.Data, Len: strHeader.Len, Cap: strHeader.Len}
	return *(*[]byte)(unsafe.Pointer(&retHeader))
}

func StringToByteSlice2(str string) *[]byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	retHeader := reflect.SliceHeader{Data: strHeader.Data, Len: strHeader.Len, Cap: strHeader.Len}
	return (*[]byte)(unsafe.Pointer(&retHeader))
}

func ByteSliceToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func IsPrintAscii(ch byte) bool {
	return ch < 128 && strconv.IsPrint(rune(ch))
}
