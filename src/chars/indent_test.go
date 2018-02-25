package chars

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/buffer"
	"github.com/lioneagle/goutil/src/test"
)

func TestIndent(t *testing.T) {
	testdata := []struct {
		initIndent int
		delta      int
		enterNum   int
		exitNum    int
		dest       string
	}{
		{0, 4, 1, 0, "    "},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			indent := NewIndent(v.initIndent, v.delta)
			for i := 0; i < v.enterNum; i++ {
				indent.Enter()
			}
			for i := 0; i < v.exitNum; i++ {
				indent.Exit()
			}

			buf := buffer.NewByteBuffer(nil)
			indent.Print(buf)

			test.EXPECT_EQ(t, buf.String(), v.dest, "")
		})
	}
}
