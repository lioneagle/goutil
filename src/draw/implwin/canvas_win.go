package implwin

import (
	"core"
	//"fmt"
	"win"
)

type CanvasWin struct {
	hdc          win.HDC
	hwnd         win.HWND
	dpix         int
	dpiy         int
	doNotDispose bool
}

func (this *CanvasWin) HDC() win.HDC {
	return this.hdc
}

func NewCanvasWin() (*CanvasWin, error) {
	hdc := win.CreateCompatibleDC(0)
	if hdc == 0 {
		return nil, core.NewError("CreateCompatibleDC failed")
	}

	return (&CanvasWin{hdc: hdc}).init()
}

func NewCanvasWinFromHDC(hdc win.HDC) (*CanvasWin, error) {
	if hdc == 0 {
		return nil, core.NewError("invalid hdc")
	}

	return (&CanvasWin{hdc: hdc, doNotDispose: true}).init()
}

func (this *CanvasWin) init() (*CanvasWin, error) {
	this.dpix = int(win.GetDeviceCaps(this.hdc, win.LOGPIXELSX))
	this.dpiy = int(win.GetDeviceCaps(this.hdc, win.LOGPIXELSY))

	if win.SetBkMode(this.hdc, win.TRANSPARENT) == 0 {
		return nil, core.NewError("SetBkMode failed")
	}

	switch win.SetStretchBltMode(this.hdc, win.HALFTONE) {
	case 0, win.ERROR_INVALID_PARAMETER:
		return nil, core.NewError("SetStretchBltMode failed")
	}

	if !win.SetBrushOrgEx(this.hdc, 0, 0, nil) {
		return nil, core.NewError("SetBrushOrgEx failed")
	}

	return this, nil
}

func (this *CanvasWin) Dispose() {
	if !this.doNotDispose && this.hdc != 0 {
		if this.hwnd == 0 {
			win.DeleteDC(this.hdc)
		} else {
			win.ReleaseDC(this.hwnd, this.hdc)
		}
		this.hdc = 0
	}
}

func (this *CanvasWin) DrawLine(pen core.Pen, from, to core.Point) error {
	if !win.MoveToEx(this.hdc, int32(from.X), int32(from.Y), nil) {
		return core.NewError("MoveToEx failed")
	}
	win_pen, _ := pen.(PenWin)
	return this.withPen(win_pen, func() error {
		if !win.LineTo(this.hdc, int32(to.X), int32(to.Y)) {
			return core.NewError("LineTo failed")
		}
		return nil
	})
}

func (this *CanvasWin) DrawImage(image core.Image, location core.Point) error {
	return nil
}

func (this *CanvasWin) PaintImage(image core.Image, f func() error) error {
	return nil
}

func (this *CanvasWin) DrawRectangle(pen core.Pen, rect core.Rectangle) error {
	return this.drawRectangle(NullBrushWin(), pen, rect, 0)
}

func (this *CanvasWin) FillRectangle(brush core.Brush, rect core.Rectangle) error {
	return this.drawRectangle(brush, NullPenWin(), rect, 1)
}

func (this *CanvasWin) drawRectangle(brush core.Brush, pen core.Pen, rect core.Rectangle, sizeCorrection int) error {
	win_pen, _ := pen.(PenWin)
	win_brush, _ := brush.(BrushWin)
	return this.withBrushAndPen(win_brush, win_pen, func() error {
		if !win.Rectangle(
			this.hdc,
			int32(rect.X),
			int32(rect.Y),
			int32(rect.X+rect.Width+sizeCorrection),
			int32(rect.Y+rect.Height+sizeCorrection)) {
			return core.NewError("drawRectangle failed")
		}
		return nil
	})
}

func (this *CanvasWin) DrawEllipse(pen core.Pen, rect core.Rectangle) error {
	return this.drawEllipse(NullBrushWin(), pen, rect, 0)
}

func (this *CanvasWin) FillEllipse(brush core.Brush, rect core.Rectangle) error {
	return this.drawEllipse(brush, NullPenWin(), rect, 1)
}

func (this *CanvasWin) drawEllipse(brush core.Brush, pen core.Pen, rect core.Rectangle, sizeCorrection int) error {
	win_pen, _ := pen.(PenWin)
	win_brush, _ := brush.(BrushWin)
	return this.withBrushAndPen(win_brush, win_pen, func() error {
		if !win.Ellipse(
			this.hdc,
			int32(rect.X),
			int32(rect.Y),
			int32(rect.X+rect.Width+sizeCorrection),
			int32(rect.Y+rect.Height+sizeCorrection)) {
			return core.NewError("drawRectangle failed")
		}
		return nil
	})
}

func (this *CanvasWin) DrawCircle(pen core.Pen, center core.Point, radius int) error {
	return this.drawEllipse(NullBrushWin(), pen, core.Rectangle{center.X - radius, center.Y - radius, 2 * radius, 2 * radius}, 0)
}

func (this *CanvasWin) FillCircle(brush core.Brush, center core.Point, radius int) error {
	return this.drawEllipse(brush, NullPenWin(), core.Rectangle{center.X - radius, center.Y - radius, 2 * radius, 2 * radius}, 1)
}

func (this *CanvasWin) withGdiObj(handle win.HGDIOBJ, f func() error) error {
	oldHandle := win.SelectObject(this.hdc, handle)
	if oldHandle == 0 {
		return core.NewError("SelectObject failed")
	}
	defer win.SelectObject(this.hdc, oldHandle)

	return f()
}

func (this *CanvasWin) withPen(pen PenWin, f func() error) error {
	return this.withGdiObj(win.HGDIOBJ(pen.handle()), f)
}

func (this *CanvasWin) withBrush(brush BrushWin, f func() error) error {
	return this.withGdiObj(win.HGDIOBJ(brush.handle()), f)
}

func (this *CanvasWin) withBrushAndPen(brush BrushWin, pen PenWin, f func() error) error {
	return this.withBrush(brush, func() error {
		return this.withPen(pen, f)
	})
}
