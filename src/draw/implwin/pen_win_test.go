package implwin

import (
	"core"
	"fmt"
	"testing"
	"win"
)

func TestGetPenStyleWin(t *testing.T) {
	testdata := []struct {
		style    core.PenStyle
		styleWin uint32
	}{
		{core.PenStyle{}, win.PS_COSMETIC | win.PS_SOLID | win.PS_ENDCAP_ROUND | win.PS_JOIN_ROUND},
		{core.PenStyle{core.PEN_TYPE_GEOMETRIC, core.PEN_DASH, core.PEN_CAP_SQUARE, core.PEN_JOIN_MITER}, win.PS_GEOMETRIC | win.PS_DASH | win.PS_ENDCAP_SQUARE | win.PS_JOIN_MITER},
		{core.PenStyle{core.PEN_TYPE_COSMETIC, core.PEN_DASH_DOT_DOT, core.PEN_CAP_FLAT, core.PEN_JOIN_BEVEL}, win.PS_COSMETIC | win.PS_DASHDOTDOT | win.PS_ENDCAP_FLAT | win.PS_JOIN_BEVEL},
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
