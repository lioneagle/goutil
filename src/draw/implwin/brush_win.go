package implwin

import (
	"draw"
	"win"
)

type BrushWin interface {
	draw.Brush
	handle() win.HBRUSH
	logbrush() *win.LOGBRUSH
}
