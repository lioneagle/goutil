package timewheel

import (
	//"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestTimeWheelAllocatorAlloc1(t *testing.T) {
	allocator := NewTimeWheelAllocator(1)

	c := allocator.Alloc()
	test.ASSERT_NE(t, c, int32(-1), "should alloc ok")
	test.ASSERT_EQ(t, c, int32(0), "wrong id")
	test.ASSERT_EQ(t, allocator.freeHead, int32(-1), "wrong freeHead")
	test.ASSERT_EQ(t, allocator.size, int32(1), "wrong size")

	c = allocator.Alloc()
	test.ASSERT_EQ(t, c, int32(-1), "should alloc nok")
}

func TestTimeWheelAllocatorAlloc2(t *testing.T) {
	allocator := NewTimeWheelAllocator(2)

	c := allocator.Alloc()
	test.ASSERT_NE(t, c, int32(-1), "should alloc ok")
	test.ASSERT_EQ(t, c, int32(0), "wrong id")
	test.ASSERT_EQ(t, allocator.freeHead, int32(1), "wrong freeHead")
	test.ASSERT_EQ(t, allocator.size, int32(1), "wrong size")

	c = allocator.Alloc()
	test.ASSERT_NE(t, c, int32(-1), "should alloc ok")
	test.ASSERT_EQ(t, c, int32(1), "wrong id")
	test.ASSERT_EQ(t, allocator.freeHead, int32(-1), "wrong freeHead")
	test.ASSERT_EQ(t, allocator.size, int32(2), "wrong size")

	c = allocator.Alloc()
	test.ASSERT_EQ(t, c, int32(-1), "should alloc nok")
}

func TestTimeWheelAllocatorFree(t *testing.T) {
	allocator := NewTimeWheelAllocator(2)

	c1 := allocator.Alloc()
	test.ASSERT_NE(t, c1, int32(-1), "should alloc ok")

	c2 := allocator.Alloc()
	test.ASSERT_NE(t, c2, int32(-1), "should alloc ok")

	allocator.Free(c1)
	test.ASSERT_EQ(t, allocator.size, int32(1), "wrong size")
	test.ASSERT_EQ(t, allocator.freeHead, int32(0), "wrong freeHead")

	allocator.Free(c1)
	test.ASSERT_EQ(t, allocator.size, int32(1), "wrong size")
	test.ASSERT_EQ(t, allocator.freeHead, int32(0), "wrong freeHead")

	allocator.Free(c2)
	test.ASSERT_EQ(t, allocator.size, int32(0), "wrong size")
	test.ASSERT_EQ(t, allocator.freeHead, int32(0), "wrong freeHead")
}

func BenchmarkTimeWheelAllocatorAllocFree(b *testing.B) {
	b.StopTimer()
	allocator := NewTimeWheelAllocator(1000)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		c := allocator.Alloc()
		allocator.Free(c)
	}
}
