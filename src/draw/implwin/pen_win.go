package implwin

import (
	"github.com/lioneagle/goutil/src/draw"
	"github.com/lioneagle/goutil/src/win"
)

type PenWin interface {
	draw.IPen
	handle() win.HPEN
}

func getPenStyleWin(style draw.PenStyle) uint32 {
	winPenStyle := uint32(style.Style)

	switch style.Type {
	case draw.PEN_TYPE_COSMETIC:
		winPenStyle |= win.PS_COSMETIC
	case draw.PEN_TYPE_GEOMETRIC:
		winPenStyle |= win.PS_GEOMETRIC
	}

	switch style.CapStyle {
	case draw.PEN_CAP_ROUND:
		winPenStyle |= win.PS_ENDCAP_ROUND
	case draw.PEN_CAP_SQUARE:
		winPenStyle |= win.PS_ENDCAP_SQUARE
	case draw.PEN_CAP_FLAT:
		winPenStyle |= win.PS_ENDCAP_FLAT
	}

	switch style.JoinStyle {
	case draw.PEN_JOIN_ROUND:
		winPenStyle |= win.PS_JOIN_ROUND
	case draw.PEN_JOIN_BEVEL:
		winPenStyle |= win.PS_JOIN_BEVEL
	case draw.PEN_JOIN_MITER:
		winPenStyle |= win.PS_JOIN_MITER
	}
	return uint32(winPenStyle)
}
