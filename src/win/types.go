package win

import (
	//"fmt"
	"reflect"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// From MSDN: Windows Data Types
// http://msdn.microsoft.com/en-us/library/s3f49ktz.aspx
// http://msdn.microsoft.com/en-us/library/windows/desktop/aa383751.aspx
// ATOM                  WORD
// BOOL                  int32
// BOOLEAN               byte
// BYTE                  byte
// CCHAR                 int8
// CHAR                  int8
// COLORREF              DWORD
// DWORD                 uint32
// DWORDLONG             ULONGLONG
// DWORD_PTR             ULONG_PTR
// DWORD32               uint32
// DWORD64               uint64
// FLOAT                 float32
// HACCEL                HANDLE
// HALF_PTR              struct{} // ???
// HANDLE                PVOID
// HBITMAP               HANDLE
// HBRUSH                HANDLE
// HCOLORSPACE           HANDLE
// HCONV                 HANDLE
// HCONVLIST             HANDLE
// HCURSOR               HANDLE
// HDC                   HANDLE
// HDDEDATA              HANDLE
// HDESK                 HANDLE
// HDROP                 HANDLE
// HDWP                  HANDLE
// HENHMETAFILE          HANDLE
// HFILE                 HANDLE
// HFONT                 HANDLE
// HGDIOBJ               HANDLE
// HGLOBAL               HANDLE
// HHOOK                 HANDLE
// HICON                 HANDLE
// HINSTANCE             HANDLE
// HKEY                  HANDLE
// HKL                   HANDLE
// HLOCAL                HANDLE
// HMENU                 HANDLE
// HMETAFILE             HANDLE
// HMODULE               HANDLE
// HPALETTE              HANDLE
// HPEN                  HANDLE
// HRESULT               int32
// HRGN                  HANDLE
// HSZ                   HANDLE
// HWINSTA               HANDLE
// HWND                  HANDLE
// INT                   int32
// INT_PTR               uintptr
// INT8                  int8
// INT16                 int16
// INT32                 int32
// INT64                 int64
// LANGID                WORD
// LCID                  DWORD
// LCTYPE                DWORD
// LGRPID                DWORD
// LONG                  int32
// LONGLONG              int64
// LONG_PTR              uintptr
// LONG32                int32
// LONG64                int64
// LPARAM                LONG_PTR
// LPBOOL                *BOOL
// LPBYTE                *BYTE
// LPCOLORREF            *COLORREF
// LPCSTR                *int8
// LPCTSTR               LPCWSTR
// LPCVOID               unsafe.Pointer
// LPCWSTR               *WCHAR
// LPDWORD               *DWORD
// LPHANDLE              *HANDLE
// LPINT                 *INT
// LPLONG                *LONG
// LPSTR                 *CHAR
// LPTSTR                LPWSTR
// LPVOID                unsafe.Pointer
// LPWORD                *WORD
// LPWSTR                *WCHAR
// LRESULT               LONG_PTR
// PBOOL                 *BOOL
// PBOOLEAN              *BOOLEAN
// PBYTE                 *BYTE
// PCHAR                 *CHAR
// PCSTR                 *CHAR
// PCTSTR                PCWSTR
// PCWSTR                *WCHAR
// PDWORD                *DWORD
// PDWORDLONG            *DWORDLONG
// PDWORD_PTR            *DWORD_PTR
// PDWORD32              *DWORD32
// PDWORD64              *DWORD64
// PFLOAT                *FLOAT
// PHALF_PTR             *HALF_PTR
// PHANDLE               *HANDLE
// PHKEY                 *HKEY
// PINT_PTR              *INT_PTR
// PINT8                 *INT8
// PINT16                *INT16
// PINT32                *INT32
// PINT64                *INT64
// PLCID                 *LCID
// PLONG                 *LONG
// PLONGLONG             *LONGLONG
// PLONG_PTR             *LONG_PTR
// PLONG32               *LONG32
// PLONG64               *LONG64
// POINTER_32            struct{} // ???
// POINTER_64            struct{} // ???
// POINTER_SIGNED        uintptr
// POINTER_UNSIGNED      uintptr
// PSHORT                *SHORT
// PSIZE_T               *SIZE_T
// PSSIZE_T              *SSIZE_T
// PSTR                  *CHAR
// PTBYTE                *TBYTE
// PTCHAR                *TCHAR
// PTSTR                 PWSTR
// PUCHAR                *UCHAR
// PUHALF_PTR            *UHALF_PTR
// PUINT                 *UINT
// PUINT_PTR             *UINT_PTR
// PUINT8                *UINT8
// PUINT16               *UINT16
// PUINT32               *UINT32
// PUINT64               *UINT64
// PULONG                *ULONG
// PULONGLONG            *ULONGLONG
// PULONG_PTR            *ULONG_PTR
// PULONG32              *ULONG32
// PULONG64              *ULONG64
// PUSHORT               *USHORT
// PVOID                 unsafe.Pointer
// PWCHAR                *WCHAR
// PWORD                 *WORD
// PWSTR                 *WCHAR
// QWORD                 uint64
// SC_HANDLE             HANDLE
// SC_LOCK               LPVOID
// SERVICE_STATUS_HANDLE HANDLE
// SHORT                 int16
// SIZE_T                ULONG_PTR
// SSIZE_T               LONG_PTR
// TBYTE                 WCHAR
// TCHAR                 WCHAR
// UCHAR                 uint8
// UHALF_PTR             struct{} // ???
// UINT                  uint32
// UINT_PTR              uintptr
// UINT8                 uint8
// UINT16                uint16
// UINT32                uint32
// UINT64                uint64
// ULONG                 uint32
// ULONGLONG             uint64
// ULONG_PTR             uintptr
// ULONG32               uint32
// ULONG64               uint64
// USHORT                uint16
// USN                   LONGLONG
// WCHAR                 uint16
// WORD                  uint16
// WPARAM                UINT_PTR
type (
	ATOM            uint16
	BOOL            int32
	COLORREF        uint32
	DWM_FRAME_COUNT uint64
	DWORD           uint32
	HACCEL          HANDLE
	HANDLE          uintptr
	HBITMAP         HANDLE
	HBRUSH          HANDLE
	HCURSOR         HANDLE
	HDC             HANDLE
	HDROP           HANDLE
	HDWP            HANDLE
	HENHMETAFILE    HANDLE
	HFONT           HANDLE
	HGDIOBJ         HANDLE
	HGLOBAL         HANDLE
	HGLRC           HANDLE
	HHOOK           HANDLE
	HICON           HANDLE
	HIMAGELIST      HANDLE
	HINSTANCE       HANDLE
	HKEY            HANDLE
	HKL             HANDLE
	HMENU           HANDLE
	HMODULE         HANDLE
	HMONITOR        HANDLE
	HPEN            HANDLE
	HRESULT         int32
	HRGN            HANDLE
	HRSRC           HANDLE
	HTHUMBNAIL      HANDLE
	HWND            HANDLE
	LPARAM          uintptr
	LPCVOID         unsafe.Pointer
	LRESULT         uintptr
	PVOID           unsafe.Pointer
	QPC_TIME        uint64
	ULONG_PTR       uintptr
	WPARAM          uintptr
	TRACEHANDLE     uintptr
	UINT            uint32
	WCHAR           uint16
	PWCHAR          *uint16
	ULONG           uint32
	HPALETTE        HANDLE
	LANGID          uint16
	WORD            uint16
)

func StringToWcharPtr(str string) *WCHAR {
	return (*WCHAR)(syscall.StringToUTF16Ptr(str))
}

func WcharPtrToString(p *WCHAR) string {
	p1 := (*uint16)(unsafe.Pointer(p))
	var length uintptr = 0
	for {
		if *p1 == 0 {
			break
		}
		p1 = (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(p1)) + 2))
		length++
	}

	var str []uint16
	((*reflect.SliceHeader)(unsafe.Pointer(&str))).Data = uintptr(unsafe.Pointer(p))
	((*reflect.SliceHeader)(unsafe.Pointer(&str))).Len = int(length)
	((*reflect.SliceHeader)(unsafe.Pointer(&str))).Cap = int(length)

	return string(utf16.Decode(str))
}

// GDI
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162805.aspx
type POINT struct {
	X, Y int32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162897.aspx
type RECT struct {
	Left, Top, Right, Bottom int32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd145035.aspx
type LOGBRUSH struct {
	LbStyle uint32
	LbColor COLORREF
	LbHatch uintptr
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd183565.aspx
type DEVMODE struct {
	DmDeviceName       [CCHDEVICENAME]uint16
	DmSpecVersion      uint16
	DmDriverVersion    uint16
	DmSize             uint16
	DmDriverExtra      uint16
	DmFields           uint32
	DmOrientation      int16
	DmPaperSize        int16
	DmPaperLength      int16
	DmPaperWidth       int16
	DmScale            int16
	DmCopies           int16
	DmDefaultSource    int16
	DmPrintQuality     int16
	DmColor            int16
	DmDuplex           int16
	DmYResolution      int16
	DmTTOption         int16
	DmCollate          int16
	DmFormName         [CCHFORMNAME]uint16
	DmLogPixels        uint16
	DmBitsPerPel       uint32
	DmPelsWidth        uint32
	DmPelsHeight       uint32
	DmDisplayFlags     uint32
	DmDisplayFrequency uint32
	DmICMMethod        uint32
	DmICMIntent        uint32
	DmMediaType        uint32
	DmDitherType       uint32
	DmReserved1        uint32
	DmReserved2        uint32
	DmPanningWidth     uint32
	DmPanningHeight    uint32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa373931.aspx
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type CLSID GUID

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd183376.aspx
type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162938.aspx
type RGBQUAD struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd183375.aspx
type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors *RGBQUAD
}

// Gdiplus
type GpStatus int32

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms534067.aspx
type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr
	SuppressBackgroundThread BOOL
	SuppressExternalCodecs   BOOL
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms534068.aspx
type GdiplusStartupOutput struct {
	NotificationHook   uintptr
	NotificationUnhook uintptr
}

type GpGraphics struct{}
type GpImage struct{}
type GpBitmap struct {
	GpImage
}

type ARGB uint32
type REAL float32

type ImageCodecInfo struct {
	Clsid             CLSID
	FormatID          GUID
	CodecName         *WCHAR
	DllName           *WCHAR
	FormatDescription *WCHAR
	FilenameExtension *WCHAR
	MimeType          *WCHAR
	Flags             uint32
	Version           uint32
	SigCount          uint32
	SigSize           uint32
	SigPattern        *byte
	SigMask           *byte
}

type EncoderParameter struct {
	Guid           GUID
	NumberOfValues ULONG
}

type EncoderParameters struct {
	Count     UINT
	Parameter EncoderParameter
}

type GpBrush struct{}
type GpSolidFill struct {
	GpBrush
}
type GpBrushType int32

type GpPoint struct {
	X REAL
	Y REAL
}

type GpWrapMode int32
type GpLineGradient struct {
	GpBrush
}

type TextRenderingHint int32
type QualityMode int32
type SmoothingMode int32

type GpFontCollection struct{}

type GpFontFamily struct{}
type GpFont struct{}

type GpUnit int32

type StringAlignment int32

type GpStringFormat struct{}
type StringFormatFlags int32

type RectF struct {
	X      REAL
	Y      REAL
	Width  REAL
	Hright REAL
}

type GpPen struct{}
