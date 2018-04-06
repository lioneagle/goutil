package implwin

import (
	"draw"
	"win"
)

type BrushWin interface {
	draw.IBrush
	handle() win.HBRUSH
	logbrush() *win.LOGBRUSH
}
