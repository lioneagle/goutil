package implwin

import (
	"win"
)

type nullBrushWin struct {
	hBrush win.HBRUSH
}

func newNullBrushWin() *nullBrushWin {
	lb := &win.LOGBRUSH{LbStyle: win.BS_NULL}

	hBrush := win.CreateBrushIndirect(lb)
	if hBrush == 0 {
		panic("failed to create null brush")
	}

	return &nullBrushWin{hBrush: hBrush}
}

func (this *nullBrushWin) Dispose() {
	if this.hBrush != 0 {
		win.DeleteObject(win.HGDIOBJ(this.hBrush))
		this.hBrush = 0
	}
}

func (this *nullBrushWin) handle() win.HBRUSH {
	return this.hBrush
}

func (this *nullBrushWin) logbrush() *win.LOGBRUSH {
	return &win.LOGBRUSH{LbStyle: win.BS_NULL}
}

var nullBrushWinSingleton *nullBrushWin = newNullBrushWin()

func NullBrushWin() *nullBrushWin {
	return nullBrushWinSingleton
}
