package implwin

import (
	"core"
	"fmt"
	"win"
)

type BitmapWin struct {
	hBmp       win.HBITMAP
	hPackedDIB win.HGLOBAL
	size       core.Size
	oldHandle  win.HGDIOBJ
	hdc        win.HDC
}

func NewBitmap(size core.Size) (bmp *BitmapWin, err error) {
	bmp = &BitmapWin{}
	bmp.hBmp = win.CreateBitmap(int32(size.Width), int32(size.Height), 1, 32, nil)
	if bmp.hBmp == 0 {
		return nil, core.NewError("CreateBitmap failed")
	}

	return bmp, nil
}

func (this *BitmapWin) BeginPaint(canvas core.Canvas) error {
	if this.hBmp == 0 {
		return core.NewError("hBmp is invalid")
	}
	canvas_win, _ := canvas.(*CanvasWin)
	this.oldHandle = win.SelectObject(canvas_win.HDC(), win.HGDIOBJ(this.hBmp))
	if this.oldHandle == 0 {
		return core.NewError("SelectObject failed")
	}

	this.hdc = canvas_win.HDC()
	return nil
}

func (this *BitmapWin) EndPaint() {
	if this.oldHandle != 0 {
		win.SelectObject(this.hdc, this.oldHandle)
		this.hdc = 0
		this.oldHandle = 0
	}
}

func (this *BitmapWin) Dispose() {
	if this.hBmp != 0 {
		win.DeleteObject(win.HGDIOBJ(this.hBmp))
		this.hBmp = 0
	}
}

func (this *BitmapWin) Size() core.Size {
	return this.size
}

func (this *BitmapWin) Draw(canvas core.Canvas) error {
	return nil
}

func (this *BitmapWin) SaveToFile(filename, format string) error {
	var bitmap *win.GpBitmap

	err := win.GdipCreateBitmapFromHBITMAP(this.hBmp, 0, &bitmap)
	if err != nil {
		return core.NewError(fmt.Sprintf("GdipCreateBitmapFromHBITMAP failed, err =", err.Error()))
	}
	defer win.GdipDisposeImage(&bitmap.GpImage)

	clsid, _ := win.GetEncoderClsid(fmt.Sprintf("image/%s", format))
	if clsid == nil {
		return core.NewError(fmt.Sprintf("Do not support %s", format))
	}

	err = win.GdipSaveImageToFile(&bitmap.GpImage, filename, clsid, nil)
	if err != nil {
		return core.NewError(fmt.Sprintf("GdipSaveImageToFile failed, err =", err.Error()))
	}

	return nil
}
