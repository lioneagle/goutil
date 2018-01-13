package win

import (
	//"reflect"
	"syscall"
	"unsafe"
)

var (
	// Library
	kernel32 = syscall.NewLazyDLL("kernel.dll")

	// Functions
	procMultiByteToWideChar = kernel32.NewProc("MultiByteToWideChar")
	procWideCharToMultiByte = kernel32.NewProc("WideCharToMultiByte")
)

func MultiByteToWideChar(CodePage UINT, dwFlags uint32, lpMultiByteStr *byte, cchMultiByte int32,
	lpWideCharStr *byte, cchWideChar int32) int32 {
	ret, _, _ := procMultiByteToWideChar.Call(
		uintptr(CodePage),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpMultiByteStr)),
		uintptr(cchMultiByte),
		uintptr(unsafe.Pointer(lpWideCharStr)),
		uintptr(cchWideChar))

	return int32(ret)
}

func WideCharToMultiByte(CodePage UINT, dwFlags uint32, lpWideCharStr *WCHAR, cchWideChar int32,
	lpMultiByteStr *byte, cchMultiByte int32, lpDefaultChar *byte, pfUsedDefaultChar *bool) int32 {
	ret, _, _ := procWideCharToMultiByte.Call(
		uintptr(CodePage),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpWideCharStr)),
		uintptr(cchWideChar),
		uintptr(unsafe.Pointer(lpMultiByteStr)),
		uintptr(cchMultiByte),
		uintptr(unsafe.Pointer(lpDefaultChar)),
		uintptr(unsafe.Pointer(pfUsedDefaultChar)))

	return int32(ret)
}
