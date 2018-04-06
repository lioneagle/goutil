package implwin

import (
	"draw"
	"win"
)

type GeometricPenWin struct {
	hPen  win.HPEN
	style draw.PenStyle
	brush draw.IBrush
	width int
}

func NewGeometricPenWin(style draw.PenStyle, width int, brush BrushWin) (*GeometricPenWin, error) {
	if brush == nil {
		return nil, draw.NewError("brush cannot be nil")
	}
	style.Type = draw.PEN_TYPE_GEOMETRIC
	winPenStyle := getPenStyleWin(style)

	hPen := win.ExtCreatePen(winPenStyle, uint32(width), brush.logbrush(), 0, nil)
	if hPen == 0 {
		return nil, draw.NewError("ExtCreatePen failed")
	}

	return &GeometricPenWin{
		hPen:  hPen,
		style: style,
		width: width,
		brush: brush,
	}, nil
}

func (this *GeometricPenWin) Dispose() {
	if this.hPen != 0 {
		win.DeleteObject(win.HGDIOBJ(this.hPen))
		this.hPen = 0
	}
}

func (this *GeometricPenWin) handle() win.HPEN {
	return this.hPen
}

func (this *GeometricPenWin) Style() draw.PenStyle {
	return this.style
}

func (this *GeometricPenWin) Width() int {
	return this.width
}

func (this *GeometricPenWin) Brush() draw.IBrush {
	return this.brush
}
