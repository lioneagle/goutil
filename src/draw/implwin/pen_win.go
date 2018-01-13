package implwin

import (
	"core"
	"win"
)

type PenWin interface {
	core.Pen
	handle() win.HPEN
}

func getPenStyleWin(style core.PenStyle) uint32 {
	winPenStyle := uint32(style.Style)

	switch style.Type {
	case core.PEN_TYPE_COSMETIC:
		winPenStyle |= win.PS_COSMETIC
	case core.PEN_TYPE_GEOMETRIC:
		winPenStyle |= win.PS_GEOMETRIC
	}

	switch style.CapStyle {
	case core.PEN_CAP_ROUND:
		winPenStyle |= win.PS_ENDCAP_ROUND
	case core.PEN_CAP_SQUARE:
		winPenStyle |= win.PS_ENDCAP_SQUARE
	case core.PEN_CAP_FLAT:
		winPenStyle |= win.PS_ENDCAP_FLAT
	}

	switch style.JoinStyle {
	case core.PEN_JOIN_ROUND:
		winPenStyle |= win.PS_JOIN_ROUND
	case core.PEN_JOIN_BEVEL:
		winPenStyle |= win.PS_JOIN_BEVEL
	case core.PEN_JOIN_MITER:
		winPenStyle |= win.PS_JOIN_MITER
	}
	return uint32(winPenStyle)
}
