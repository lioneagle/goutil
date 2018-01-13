package implwin

import (
	"core"
	"win"
)

type BrushWin interface {
	core.Brush
	handle() win.HBRUSH
	logbrush() *win.LOGBRUSH
}
