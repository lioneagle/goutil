package mem

import (
	//"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestArenaAllocatorAllocOk(t *testing.T) {
	allocator := NewArenaAllocator(1024)
	test.EXPECT_NE(t, allocator.Capacity(), 1024, "")
	test.EXPECT_NE(t, allocator.Left(), 1024, "")

	addr, allocSize := allocator.Alloc(1024)
	test.EXPECT_NE(t, addr, MEM_PTR_NIL, "")
	test.EXPECT_EQ(t, allocSize, uint32(1024), "")
	test.EXPECT_EQ(t, allocator.Used(), uint32(1024), "")
	test.EXPECT_EQ(t, allocator.Stat().AllocNum(), uint64(1), "")
	test.EXPECT_EQ(t, allocator.Stat().AllocNumOk(), uint64(1), "")

	allocator.FreeAll()
	test.EXPECT_EQ(t, allocator.Used(), uint32(0), "")

	addr, _ = allocator.AllocWithClear(8)
	test.EXPECT_NE(t, addr, MEM_PTR_NIL, "")
	test.EXPECT_EQ(t, allocator.Used(), uint32(8), "")
	test.EXPECT_EQ(t, allocator.mem[allocator.used-8:allocator.used], []byte{0, 0, 0, 0, 0, 0, 0, 0}, "")

	allocator.FreeAll()
	test.EXPECT_EQ(t, allocator.Used(), uint32(0), "")

	addr = allocator.AllocBytes([]byte("12346"))
	test.EXPECT_NE(t, addr, MEM_PTR_NIL, "")
	test.EXPECT_EQ(t, allocator.GetString(addr), "12346", "")

	allocator.FreeAll()
	test.EXPECT_EQ(t, allocator.Used(), uint32(0), "")

	addr = allocator.AllocBytesBegin()
	test.EXPECT_NE(t, addr, MEM_PTR_NIL, "")
	ret := allocator.AppendBytes([]byte("12378"))
	test.EXPECT_TRUE(t, ret, "")
	allocator.AllocBytesEnd(addr)
	test.EXPECT_EQ(t, allocator.GetString(addr), "12378", "")

	allocator.FreeAll()
	test.EXPECT_EQ(t, allocator.Used(), uint32(0), "")
	addr = allocator.AllocBytesBegin()
	test.EXPECT_NE(t, addr, MEM_PTR_NIL, "")
	ret = allocator.AppendBytes([]byte("12378"))
	test.EXPECT_TRUE(t, ret, "")
	ret = allocator.AppendByte('6')
	test.EXPECT_TRUE(t, ret, "")
	allocator.AppendByteNoCheck('a')
	allocator.AllocBytesEnd(addr)
	test.EXPECT_EQ(t, allocator.GetString(addr), "123786a", "")
}

func TestArenaAllocatorAllocNOk(t *testing.T) {
	allocator := NewArenaAllocator(1024)
	test.EXPECT_NE(t, allocator.Capacity(), 1024, "")
	test.EXPECT_NE(t, allocator.Left(), 1024, "")

	addr, _ := allocator.Alloc(1025)
	test.EXPECT_EQ(t, addr, MEM_PTR_NIL, "")

	allocator.Alloc(1023)
	addr = allocator.AllocBytesBegin()
	test.EXPECT_EQ(t, addr, MEM_PTR_NIL, "")

	allocator.FreeAll()
	allocator.Alloc(1020)
	addr = allocator.AllocBytesBegin()
	ret := allocator.AppendBytes([]byte{0, 1, 2})
	test.EXPECT_FALSE(t, ret, "")

	allocator.AppendBytes([]byte{0, 1})
	ret = allocator.AppendByte('a')
	test.EXPECT_FALSE(t, ret, "")
}

func TestArenaAllocatorFreeOk(t *testing.T) {
	allocator := NewArenaAllocator(1024)
	test.EXPECT_NE(t, allocator.Capacity(), 1024, "")
	test.EXPECT_NE(t, allocator.Left(), 1024, "")

	allocator.AllocWithClear(16)
	allocator.FreePart(8)
	test.EXPECT_EQ(t, allocator.Used(), uint32(8), "")

	allocator.FreeAll()
	test.EXPECT_EQ(t, allocator.Used(), uint32(0), "")
}

func TestArenaAllocatorFreeNOk(t *testing.T) {
	allocator := NewArenaAllocator(1024)
	test.EXPECT_NE(t, allocator.Capacity(), 1024, "")
	test.EXPECT_NE(t, allocator.Left(), 1024, "")

	allocator.AllocWithClear(16)
	allocator.FreePart(28)
	test.EXPECT_EQ(t, allocator.Used(), uint32(16), "")
}

func TestArenaAllocatorGetString(t *testing.T) {
	allocator := NewArenaAllocator(1024)

	addr := allocator.AllocBytes(nil)
	test.EXPECT_EQ(t, allocator.Used(), uint32(2), "")
	test.EXPECT_EQ(t, allocator.GetString(addr), "", "")
	test.EXPECT_EQ(t, allocator.GetString(MEM_PTR_NIL), "", "")

	allocator.FreeAll()
	addr = allocator.AllocBytes([]byte("abfg"))
	test.EXPECT_EQ(t, allocator.GetString(addr), "abfg", "")
}

func TestArenaAllocatorStrlen(t *testing.T) {
	allocator := NewArenaAllocator(1024)

	addr := allocator.AllocBytes(nil)
	test.EXPECT_EQ(t, allocator.Used(), uint32(2), "")
	test.EXPECT_EQ(t, allocator.Strlen(addr), 0, "")
	test.EXPECT_EQ(t, allocator.Strlen(MEM_PTR_NIL), 0, "")

	allocator.FreeAll()
	addr = allocator.AllocBytes([]byte("abfg"))
	test.EXPECT_EQ(t, allocator.Strlen(addr), 4, "")
}

func TestArenaAllocatorZeroMem(t *testing.T) {
	allocator := NewArenaAllocator(1024)

	addr := allocator.AllocBytes([]byte("1234"))
	allocator.ZeroMem(addr, 4)
	test.EXPECT_EQ(t, allocator.mem[allocator.used-4:allocator.used], []byte{0, 0, 0, 0}, "")
}

func TestArenaAllocatorClone(t *testing.T) {
	allocator := NewArenaAllocator(1024)

	addr := allocator.AllocBytes([]byte("1234"))

	newAllocator := allocator.Clone()
	test.ASSERT_NE(t, newAllocator, nil, "")

	test.EXPECT_EQ(t, allocator.GetString(addr), "1234", "")
	test.EXPECT_EQ(t, newAllocator.GetString(addr), "1234", "")
}

func TestZeroMem(t *testing.T) {
	allocator := NewArenaAllocator(1024)

	addr := allocator.AllocBytes([]byte("1234"))
	ZeroMem(allocator.GetUintptr(addr), 4)
	test.EXPECT_EQ(t, allocator.mem[allocator.used-4:allocator.used], []byte{0, 0, 0, 0}, "")
}

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

func BenchmarkArenaAllocatorAllocBytes1(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024 * 128)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		allocator.AllocBytes(data)
	}
}

func BenchmarkArenaAllocatorAllocBytes2(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024 * 128)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		addr := allocator.AllocBytesBegin()
		allocator.AppendBytes(data[0:16])
		allocator.AppendBytes(data[16:])
		allocator.AllocBytesEnd(addr)
	}
}

func BenchmarkArenaAllocatorAllocBytes3(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024 * 128)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		addr := allocator.AllocBytesBegin()
		for _, v := range data {
			allocator.AppendByte(v)
		}
		allocator.AllocBytesEnd(addr)
	}
}

func BenchmarkArenaAllocatorAllocBytes4(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024 * 128)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		addr := allocator.AllocBytesBegin()
		for _, v := range data {
			allocator.AppendByteNoCheck(v)
		}
		allocator.AllocBytesEnd(addr)
	}
}
