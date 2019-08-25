package mem

import (
	//"fmt"
	"testing"
	//"github.com/lioneagle/goutil/src/test"
)

func BenchmarkArenaAllocatorAlloc(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024 * 128)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		_, _ = allocator.Alloc(1024)
	}
}

func BenchmarkArenaAllocatorAllocWithClear1(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024 * 128)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		_, _ = allocator.AllocWithClear(32)
	}
}
