package win

import (
	"errors"
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
)

func (s GpStatus) String() string {
	switch s {
	case Ok:
		return "Ok"

	case GenericError:
		return "GenericError"

	case InvalidParameter:
		return "InvalidParameter"

	case OutOfMemory:
		return "OutOfMemory"

	case ObjectBusy:
		return "ObjectBusy"

	case InsufficientBuffer:
		return "InsufficientBuffer"

	case NotImplemented:
		return "NotImplemented"

	case Win32Error:
		return "Win32Error"

	case WrongState:
		return "WrongState"

	case Aborted:
		return "Aborted"

	case FileNotFound:
		return "FileNotFound"

	case ValueOverflow:
		return "ValueOverflow"

	case AccessDenied:
		return "AccessDenied"

	case UnknownImageFormat:
		return "UnknownImageFormat"

	case FontFamilyNotFound:
		return "FontFamilyNotFound"

	case FontStyleNotFound:
		return "FontStyleNotFound"

	case NotTrueTypeFont:
		return "NotTrueTypeFont"

	case UnsupportedGdiplusVersion:
		return "UnsupportedGdiplusVersion"

	case GdiplusNotInitialized:
		return "GdiplusNotInitialized"

	case PropertyNotFound:
		return "PropertyNotFound"

	case PropertyNotSupported:
		return "PropertyNotSupported"

	case ProfileNotFound:
		return "ProfileNotFound"
	}

	return "Unknown Status Value"
}

var (
	token uintptr

	// Library
	libgdiplus = syscall.NewLazyDLL("gdiplus.dll")

	// Functions
	procGdiplusStartup  = libgdiplus.NewProc("GdiplusStartup")
	procGdiplusShutdown = libgdiplus.NewProc("GdiplusShutdown")

	// Grahics Functions
	procGdipCreateFromHDC        = libgdiplus.NewProc("GdipCreateFromHDC")
	procGdipDeleteGraphics       = libgdiplus.NewProc("GdipDeleteGraphics")
	procGdipGetDC                = libgdiplus.NewProc("GdipGetDC")
	procGdipReleaseDC            = libgdiplus.NewProc("GdipReleaseDC")
	procGdipSetTextRenderingHint = libgdiplus.NewProc("GdipSetTextRenderingHint")
	procGdipSetSmoothingMode     = libgdiplus.NewProc("GdipSetSmoothingMode")
	procGdipDrawRectangle        = libgdiplus.NewProc("GdipDrawRectangle")
	procGdipFillRectangle        = libgdiplus.NewProc("GdipFillRectangle")

	// Bitmap Functions
	procGdipCreateBitmapFromHBITMAP = libgdiplus.NewProc("GdipCreateBitmapFromHBITMAP")
	procCreateHBITMAPFromBitmap     = libgdiplus.NewProc("GdipCreateHBITMAPFromBitmap")

	// Image Functions
	procGdipGetImageGraphicsContext = libgdiplus.NewProc("GdipGetImageGraphicsContext")
	procGdipGetImageEncodersSize    = libgdiplus.NewProc("GdipGetImageEncodersSize")
	procGdipGetImageEncoders        = libgdiplus.NewProc("GdipGetImageEncoders")
	procGdipSaveImageToFile         = libgdiplus.NewProc("GdipSaveImageToFile")
	procGdipDisposeImage            = libgdiplus.NewProc("GdipDisposeImage")

	// SolidBrush Functions
	procGdipCloneBrush        = libgdiplus.NewProc("GdipCloneBrush")
	procGdipDeleteBrush       = libgdiplus.NewProc("GdipDeleteBrush")
	procGdipGetBrushType      = libgdiplus.NewProc("GdipGetBrushType")
	procGdipCreateSolidFill   = libgdiplus.NewProc("GdipCreateSolidFill")
	procGdipSetSolidFillColor = libgdiplus.NewProc("GdipSetSolidFillColor")
	procGdipGetSolidFillColor = libgdiplus.NewProc("GdipGetSolidFillColor")

	// LinearGradientBrush Functions
	procGdipCreateLineBrush = libgdiplus.NewProc("GdipCreateLineBrush")

	// FontFamily Functions
	procGdipCreateFontFamilyFromName = libgdiplus.NewProc("GdipCreateFontFamilyFromName")
	procGdipDeleteFontFamily         = libgdiplus.NewProc("GdipDeleteFontFamily")

	// Font functions
	procGdipCreateFont = libgdiplus.NewProc("GdipCreateFont")
	procGdipDeleteFont = libgdiplus.NewProc("GdipDeleteFont")

	// String Format Functions
	procGdipCreateStringFormat   = libgdiplus.NewProc("GdipCreateStringFormat")
	procGdipDeleteStringFormat   = libgdiplus.NewProc("GdipDeleteStringFormat")
	procGdipSetStringFormatAlign = libgdiplus.NewProc("GdipSetStringFormatAlign")

	// Text Functions
	procGdipDrawString = libgdiplus.NewProc("GdipDrawString")

	// Pen Functions
	procGdipCreatePen1 = libgdiplus.NewProc("GdipCreatePen1")
	procGdipDeletePen  = libgdiplus.NewProc("GdipDeletePen")
)

func GdiplusStartup(input *GdiplusStartupInput, output *GdiplusStartupOutput) error {
	ret, _, _ := procGdiplusStartup.Call(
		uintptr(unsafe.Pointer(&token)),
		uintptr(unsafe.Pointer(input)),
		uintptr(unsafe.Pointer(output)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdiplusStartup failed with status '%s'", GpStatus(ret)))
	}

	return nil
}

func GdiplusShutdown() {
	procGdiplusShutdown.Call(token)
}

func GdipCreateFromHDC(hdc HDC, graphics **GpGraphics) error {
	ret, _, _ := procGdipCreateFromHDC.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(graphics)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipCreateFromHDC failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipDeleteGraphics(graphics *GpGraphics) error {
	ret, _, _ := procGdipDeleteGraphics.Call(
		uintptr(unsafe.Pointer(graphics)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipDeleteGraphics failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipGetDC(graphics *GpGraphics, hdc *HDC) error {
	ret, _, _ := procGdipGetDC.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(hdc)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipGetDC failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipReleaseDC(graphics *GpGraphics, hdc HDC) error {
	ret, _, _ := procGdipReleaseDC.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(hdc))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipReleaseDC failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipSetTextRenderingHint(graphics *GpGraphics, mode TextRenderingHint) error {
	ret, _, _ := procGdipSetTextRenderingHint.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(mode))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipSetTextRenderingHint failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipSetSmoothingMode(graphics *GpGraphics, smoothingMode SmoothingMode) error {
	ret, _, _ := procGdipSetSmoothingMode.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(smoothingMode))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipSetSmoothingMode failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipDrawRectangle(graphics *GpGraphics, pen *GpPen, x, y, width, height REAL) error {
	ret, _, _ := procGdipDrawRectangle.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(pen)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipDrawRectangle failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipFillRectangle(graphics *GpGraphics, brush *GpBrush, x, y, width, height REAL) error {
	ret, _, _ := procGdipFillRectangle.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipFillRectangle failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipCreateBitmapFromHBITMAP(hbm HBITMAP, hpal HPALETTE, bitmap **GpBitmap) error {
	ret, _, _ := procGdipCreateBitmapFromHBITMAP.Call(
		uintptr(hbm),
		uintptr(hpal),
		uintptr(unsafe.Pointer(bitmap)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipCreateBitmapFromHBITMAP failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipCreateHBITMAPFromBitmap(bitmap *GpBitmap, hbmReturn *HBITMAP, background ARGB) error {
	ret, _, _ := procCreateHBITMAPFromBitmap.Call(
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(hbmReturn)),
		uintptr(background))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipCreateHBITMAPFromBitmap failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms534041(v=vs.85).aspx
func GdipGetImageGraphicsContext(image *GpImage, graphics **GpGraphics) error {
	ret, _, _ := procGdipGetImageGraphicsContext.Call(
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(graphics)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipGetImageGraphicsContext failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipGetImageEncodersSize(numEncoders *UINT, size *UINT) error {
	ret, _, _ := procGdipGetImageEncodersSize.Call(
		uintptr(unsafe.Pointer(numEncoders)),
		uintptr(unsafe.Pointer(size)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipGetImageEncodersSize failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipGetImageEncoders(numEncoders UINT, size UINT, decoders *ImageCodecInfo) error {
	ret, _, _ := procGdipGetImageEncoders.Call(
		uintptr(numEncoders),
		uintptr(size),
		uintptr(unsafe.Pointer(decoders)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipGetImageEncoders failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipSaveImageToFile(image *GpImage, filename string, clsidEncoder *CLSID, encoderParams *EncoderParameters) error {
	ret, _, _ := procGdipSaveImageToFile.Call(
		uintptr(unsafe.Pointer(image)),
		uintptr(unsafe.Pointer(StringToWcharPtr(filename))),
		uintptr(unsafe.Pointer(clsidEncoder)),
		uintptr(unsafe.Pointer(encoderParams)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipSaveImageToFile failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipDisposeImage(image *GpImage) error {
	ret, _, _ := procGdipDisposeImage.Call(
		uintptr(unsafe.Pointer(image)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipDisposeImage failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipCloneBrush(brush *GpBrush, cloneBrush **GpBrush) error {
	ret, _, _ := procGdipCloneBrush.Call(
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(cloneBrush)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipCloneBrush failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipDeleteBrush(brush *GpBrush) error {
	ret, _, _ := procGdipDeleteBrush.Call(
		uintptr(unsafe.Pointer(brush)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipDeleteBrush failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipGetBrushType(brush *GpBrush, brushType *GpBrushType) error {
	ret, _, _ := procGdipGetBrushType.Call(
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(brushType)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipGetBrushType failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipCreateSolidFill(color ARGB, brush **GpSolidFill) error {
	ret, _, _ := procGdipCreateSolidFill.Call(
		uintptr(color),
		uintptr(unsafe.Pointer(brush)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipCreateSolidFill failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipSetSolidFillColor(brush **GpSolidFill, color ARGB) error {
	ret, _, _ := procGdipSetSolidFillColor.Call(
		uintptr(unsafe.Pointer(brush)),
		uintptr(color))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipSetSolidFillColor failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipGetSolidFillColor(brush **GpSolidFill, color *ARGB) error {
	ret, _, _ := procGdipGetSolidFillColor.Call(
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(color)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipGetSolidFillColor failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipCreateLineBrush(point1, point2 *GpPoint, color1, color2 ARGB, wrapMode GpWrapMode, lineGradient **GpLineGradient) error {
	ret, _, _ := procGdipCreateLineBrush.Call(
		uintptr(unsafe.Pointer(point1)),
		uintptr(unsafe.Pointer(point2)),
		uintptr(color1),
		uintptr(color2),
		uintptr(wrapMode),
		uintptr(unsafe.Pointer(lineGradient)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipCreateLineBrush failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipCreateFontFamilyFromName(name string, fontCollection *GpFontCollection, fontFamily **GpFontFamily) error {
	ret, _, _ := procGdipCreateFontFamilyFromName.Call(
		uintptr(unsafe.Pointer(StringToWcharPtr(name))),
		uintptr(unsafe.Pointer(fontCollection)),
		uintptr(unsafe.Pointer(fontFamily)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipCreateFontFamilyFromName failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipDeleteFontFamily(fontFamily *GpFontFamily) error {
	ret, _, _ := procGdipDeleteFontFamily.Call(
		uintptr(unsafe.Pointer(fontFamily)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipDeleteFontFamily failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipCreateFont(fontFamily *GpFontFamily, emSize REAL, style int32, unit GpUnit, font **GpFont) error {
	ret, _, _ := procGdipCreateFont.Call(
		uintptr(unsafe.Pointer(fontFamily)),
		uintptr(emSize),
		uintptr(style),
		uintptr(unit),
		uintptr(unsafe.Pointer(font)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipCreateFont failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipDeleteFont(font *GpFont) error {
	ret, _, _ := procGdipDeleteFont.Call(
		uintptr(unsafe.Pointer(font)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipDeleteFont failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipCreateStringFormat(formatAttributes int32, language LANGID, format **GpStringFormat) error {
	ret, _, _ := procGdipCreateStringFormat.Call(
		uintptr(formatAttributes),
		uintptr(language),
		uintptr(unsafe.Pointer(format)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipCreateStringFormat failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipDeleteStringFormat(format *GpStringFormat) error {
	ret, _, _ := procGdipDeleteStringFormat.Call(
		uintptr(unsafe.Pointer(format)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipDeleteStringFormat failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipSetStringFormatAlign(format *GpStringFormat, align StringAlignment) error {
	ret, _, _ := procGdipSetStringFormatAlign.Call(
		uintptr(unsafe.Pointer(format)),
		uintptr(align))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipSetStringFormatAlign failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipDrawString(graphics *GpGraphics, str string, length int32, font *GpFont, layoutRect *RectF,
	stringFormat *GpStringFormat, brush *GpBrush) error {
	ret, _, _ := procGdipDrawString.Call(
		uintptr(unsafe.Pointer(graphics)),
		uintptr(unsafe.Pointer(StringToWcharPtr(str))),
		uintptr(length),
		uintptr(unsafe.Pointer(font)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(stringFormat)),
		uintptr(unsafe.Pointer(brush)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipDrawString failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipCreatePen1(color ARGB, width REAL, unit GpUnit, pen **GpPen) error {
	ret, _, _ := procGdipCreatePen1.Call(
		uintptr(color),
		uintptr(width),
		uintptr(unit),
		uintptr(unsafe.Pointer(pen)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipCreatePen1 failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GdipDeletePen(pen *GpPen) error {
	ret, _, _ := procGdipDeletePen.Call(
		uintptr(unsafe.Pointer(pen)))

	if GpStatus(ret) != Ok {
		return errors.New(fmt.Sprintf("GdipDeletePen failed with status '%s'", GpStatus(ret)))
	}
	return nil
}

func GetEncoderClsid(format string) (clsid *CLSID, index int) {
	var num UINT = 0
	var size UINT = 0

	clsid = &CLSID{}

	err := GdipGetImageEncodersSize(&num, &size)
	if err != nil {
		fmt.Println("GetEncoderClsid:: call GdipGetImageEncodersSize failed, err = %s", err.Error())
		return nil, -1
	}

	buf := make([]byte, size)
	err = GdipGetImageEncoders(num, size, (*ImageCodecInfo)(unsafe.Pointer(&buf[0])))
	if err != nil {
		fmt.Println("GetEncoderClsid:: call GdipGetImageEncoders failed, err = %s", err.Error())
		return nil, -1
	}

	var imageCodecInfo []ImageCodecInfo

	((*reflect.SliceHeader)(unsafe.Pointer(&imageCodecInfo))).Data = uintptr(unsafe.Pointer(&buf[0]))
	((*reflect.SliceHeader)(unsafe.Pointer(&imageCodecInfo))).Len = int(num)
	((*reflect.SliceHeader)(unsafe.Pointer(&imageCodecInfo))).Cap = int(num)

	for i := UINT(0); i < num; i++ {
		str := WcharPtrToString(imageCodecInfo[i].MimeType)
		//fmt.Println("type =", str)
		if str == format {
			*clsid = imageCodecInfo[i].Clsid
			return clsid, int(i)
		}
	}
	return nil, -1
}
