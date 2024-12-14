package implwin

import (
	"github.com/lioneagle/goutil/src/draw"
	"github.com/lioneagle/goutil/src/win"
)

type CosmeticPenWin struct {
	HPen  win.HPEN
	style draw.PenStyle
	color draw.Color
}

func NewCosmeticPenWin(style draw.PenStyle, color draw.Color) (*CosmeticPenWin, error) {
	lb := &win.LOGBRUSH{LbStyle: win.BS_SOLID, LbColor: win.COLORREF(color)}

	style.Type = draw.PEN_TYPE_GEOMETRIC
	winPenStyle := getPenStyleWin(style)

	hPen := win.ExtCreatePen(winPenStyle, 1, lb, 0, nil)
	if hPen == 0 {
		return nil, draw.NewError("ExtCreatePen failed")
	}

	return &CosmeticPenWin{HPen: hPen, style: style, color: color}, nil
}

func (this *CosmeticPenWin) Dispose() {
	if this.HPen != 0 {
		win.DeleteObject(win.HGDIOBJ(this.HPen))
		this.HPen = 0
	}
}

func (this *CosmeticPenWin) handle() win.HPEN {
	return this.HPen
}

func (this *CosmeticPenWin) Style() draw.PenStyle {
	return this.style
}

func (this *CosmeticPenWin) Color() draw.Color {
	return this.color
}

func (this *CosmeticPenWin) Width() int {
	return 1
}
