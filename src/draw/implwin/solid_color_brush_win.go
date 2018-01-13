package implwin

import (
	"core"
	"win"
)

type SolidColorBrushWin struct {
	hBrush win.HBRUSH
	color  core.Color
}

func NewSolidColorBrush(color core.Color) (*SolidColorBrushWin, error) {
	lb := &win.LOGBRUSH{LbStyle: win.BS_SOLID, LbColor: win.COLORREF(color)}

	hBrush := win.CreateBrushIndirect(lb)
	if hBrush == 0 {
		return nil, core.NewError("CreateBrushIndirect failed")
	}

	return &SolidColorBrushWin{hBrush: hBrush, color: color}, nil
}

func (this *SolidColorBrushWin) Color() core.Color {
	return this.color
}

func (this *SolidColorBrushWin) Dispose() {
	if this.hBrush != 0 {
		win.DeleteObject(win.HGDIOBJ(this.hBrush))

		this.hBrush = 0
	}
}

func (this *SolidColorBrushWin) handle() win.HBRUSH {
	return this.hBrush
}

func (this *SolidColorBrushWin) logbrush() *win.LOGBRUSH {
	return &win.LOGBRUSH{LbStyle: win.BS_SOLID, LbColor: win.COLORREF(this.color)}
}
