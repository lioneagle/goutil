package implwin

import (
	"core"
	"win"
)

type CosmeticPenWin struct {
	HPen  win.HPEN
	style core.PenStyle
	color core.Color
}

func NewCosmeticPenWin(style core.PenStyle, color core.Color) (*CosmeticPenWin, error) {
	lb := &win.LOGBRUSH{LbStyle: win.BS_SOLID, LbColor: win.COLORREF(color)}

	style.Type = core.PEN_TYPE_GEOMETRIC
	winPenStyle := getPenStyleWin(style)

	hPen := win.ExtCreatePen(winPenStyle, 1, lb, 0, nil)
	if hPen == 0 {
		return nil, core.NewError("ExtCreatePen failed")
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

func (this *CosmeticPenWin) Style() core.PenStyle {
	return this.style
}

func (this *CosmeticPenWin) Color() core.Color {
	return this.color
}

func (this *CosmeticPenWin) Width() int {
	return 1
}
