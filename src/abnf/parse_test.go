package abnf

import (
	//"fmt"
	"testing"

	//"github.com/lioneagle/goutil/src/buffer"
	"github.com/lioneagle/goutil/src/chars"
	"github.com/lioneagle/goutil/src/mem"
	//"github.com/lioneagle/goutil/src/test"
)

func BenchmarkParseCharsetAndAlloc(b *testing.B) {
	b.StopTimer()
	allocator := mem.NewArenaAllocator(1024, 1)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		addr, newPos := ParseCharsetAndAlloc(allocator, data, 0, &chars.Charset0, chars.MASK_DIGIT)
		if addr == mem.MEM_PTR_NIL || newPos != len(data) {
			return
		}
	}
}
