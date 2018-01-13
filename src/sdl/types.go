package sdl

import (
	"reflect"
	"unsafe"
)

func StringToBytePtr(str string) *byte {
	newStr := []byte(str + "\x00")
	return &newStr[0]
}

func CharPtrToString(str *byte) string {
	p1 := (*byte)(unsafe.Pointer(str))
	var length int = 0
	for {
		if *p1 == 0 {
			break
		}
		p1 = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p1)) + 1))
		length++
	}

	var str1 string
	((*reflect.StringHeader)(unsafe.Pointer(&str1))).Data = uintptr(unsafe.Pointer(str))
	((*reflect.StringHeader)(unsafe.Pointer(&str1))).Len = length

	return str1
}
