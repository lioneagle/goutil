package implwin

import (
	"fmt"

	"draw"
	"win"
)

type BitmapWin struct {
	hBmp       win.HBITMAP
	hPackedDIB win.HGLOBAL
	size       draw.Size
	oldHandle  win.HGDIOBJ
	hdc        win.HDC
}

func NewBitmap(size draw.Size) (bmp *BitmapWin, err error) {
	bmp = &BitmapWin{}
	bmp.hBmp = win.CreateBitmap(int32(size.Width), int32(size.Height), 1, 32, nil)
	if bmp.hBmp == 0 {
		return nil, draw.NewError("CreateBitmap failed")
	}

	return bmp, nil
}

func (this *BitmapWin) BeginPaint(canvas draw.Canvas) error {
	if this.hBmp == 0 {
		return draw.NewError("hBmp is invalid")
	}
	canvas_win, _ := canvas.(*CanvasWin)
	this.oldHandle = win.SelectObject(canvas_win.HDC(), win.HGDIOBJ(this.hBmp))
	if this.oldHandle == 0 {
		return draw.NewError("SelectObject failed")
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

func (this *BitmapWin) Size() draw.Size {
	return this.size
}

func (this *BitmapWin) Draw(canvas draw.Canvas) error {
	return nil
}

func (this *BitmapWin) SaveToFile(filename, format string) error {
	var bitmap *win.GpBitmap

	err := win.GdipCreateBitmapFromHBITMAP(this.hBmp, 0, &bitmap)
	if err != nil {
		return draw.NewError(fmt.Sprintf("GdipCreateBitmapFromHBITMAP failed, err =", err.Error()))
	}
	defer win.GdipDisposeImage(&bitmap.GpImage)

	clsid, _ := win.GetEncoderClsid(fmt.Sprintf("image/%s", format))
	if clsid == nil {
		return draw.NewError(fmt.Sprintf("Do not support %s", format))
	}

	err = win.GdipSaveImageToFile(&bitmap.GpImage, filename, clsid, nil)
	if err != nil {
		return draw.NewError(fmt.Sprintf("GdipSaveImageToFile failed, err =", err.Error()))
	}

	return nil
}
