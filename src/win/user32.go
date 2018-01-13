package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libuser32 = syscall.NewLazyDLL("user32.dll")

	// Functions
	procReleaseDC = libuser32.NewProc("ReleaseDC")
	procFillRect  = libuser32.NewProc("FillRect")
)

func ReleaseDC(hwnd HWND, hDC HDC) bool {
	ret, _, _ := procReleaseDC.Call(
		uintptr(hwnd),
		uintptr(hDC))

	return ret != 0
}

// https://msdn.microsoft.com/en-us/library/windows/desktop/dd162719(v=vs.85).aspx
func FillRect(hdc HDC, rect *RECT, brush HBRUSH) bool {
	ret, _, _ := procFillRect.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(rect)),
		uintptr(brush))

	return ret != 0
}
