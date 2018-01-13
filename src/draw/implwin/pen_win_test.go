package implwin

import (
	"fmt"
	"testing"

	"draw"
	"win"
)

func TestGetPenStyleWin(t *testing.T) {
	testdata := []struct {
		style    draw.PenStyle
		styleWin uint32
	}{
		{draw.PenStyle{}, win.PS_COSMETIC | win.PS_SOLID | win.PS_ENDCAP_ROUND | win.PS_JOIN_ROUND},
		{draw.PenStyle{draw.PEN_TYPE_GEOMETRIC, draw.PEN_DASH, draw.PEN_CAP_SQUARE, draw.PEN_JOIN_MITER}, win.PS_GEOMETRIC | win.PS_DASH | win.PS_ENDCAP_SQUARE | win.PS_JOIN_MITER},
		{draw.PenStyle{draw.PEN_TYPE_COSMETIC, draw.PEN_DASH_DOT_DOT, draw.PEN_CAP_FLAT, draw.PEN_JOIN_BEVEL}, win.PS_COSMETIC | win.PS_DASHDOTDOT | win.PS_ENDCAP_FLAT | win.PS_JOIN_BEVEL},
	}

	prefix := "TestGetPenStyleWin"

	for i, v := range testdata {
		val := getPenStyleWin(v.style)

		if val != v.styleWin {
			fmt.Println("style =", v.style)
			t.Errorf("%s[%d] failed: winStyle = 0x%08x, wanted = 0x%08x\n", prefix, i, val, v.styleWin)
			continue
		}
	}

}
