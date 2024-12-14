package implwin

import (
	"github.com/lioneagle/goutil/src/draw"
	"github.com/lioneagle/goutil/src/win"
)

type BrushWin interface {
	draw.IBrush
	handle() win.HBRUSH
	logbrush() *win.LOGBRUSH
}
