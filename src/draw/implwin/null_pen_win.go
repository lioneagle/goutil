package implwin

import (
	"github.com/lioneagle/goutil/src/draw"
	"github.com/lioneagle/goutil/src/win"
)

type nullPenWin struct {
	hPen win.HPEN
}

func newNullPenWin() *nullPenWin {
	lb := &win.LOGBRUSH{LbStyle: win.BS_NULL}

	hPen := win.ExtCreatePen(win.PS_COSMETIC|win.PS_NULL, 1, lb, 0, nil)
	if hPen == 0 {
		panic("failed to create null brush")
	}

	return &nullPenWin{hPen: hPen}
}

func (this *nullPenWin) Dispose() {
	if this.hPen != 0 {
		win.DeleteObject(win.HGDIOBJ(this.hPen))
		this.hPen = 0
	}
}

func (this *nullPenWin) handle() win.HPEN {
	return this.hPen
}

func (this *nullPenWin) Style() draw.PenStyle {
	return draw.PenStyle{}
}

func (this *nullPenWin) Width() int {
	return 0
}

var nullPenWinSingleton *nullPenWin = newNullPenWin()

func NullPenWin() *nullPenWin {
	return nullPenWinSingleton
}
