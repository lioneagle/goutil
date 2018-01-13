package implwin

import (
	"core"
	"win"
)

type GeometricPenWin struct {
	hPen  win.HPEN
	style core.PenStyle
	brush core.Brush
	width int
}

func NewGeometricPenWin(style core.PenStyle, width int, brush BrushWin) (*GeometricPenWin, error) {
	if brush == nil {
		return nil, core.NewError("brush cannot be nil")
	}
	style.Type = core.PEN_TYPE_GEOMETRIC
	winPenStyle := getPenStyleWin(style)

	hPen := win.ExtCreatePen(winPenStyle, uint32(width), brush.logbrush(), 0, nil)
	if hPen == 0 {
		return nil, core.NewError("ExtCreatePen failed")
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

func (this *GeometricPenWin) Style() core.PenStyle {
	return this.style
}

func (this *GeometricPenWin) Width() int {
	return this.width
}

func (this *GeometricPenWin) Brush() core.Brush {
	return this.brush
}
