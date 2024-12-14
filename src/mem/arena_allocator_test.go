package mem

import (
	//"fmt"
	"testing"

	//"github.com/lioneagle/goutil/src/buffer"
	"github.com/lioneagle/goutil/src/test"
)

/*func TestArenaAllocatorAllocOk(t *testing.T) {
	allocator := NewArenaAllocator(1024, 1)
	test.EXPECT_NE(t, allocator.Capacity(), 1024, "")
	test.EXPECT_NE(t, allocator.Left(), 1024, "")

	addr, allocSize := allocator.Alloc(1024)
	test.EXPECT_NE(t, addr, MEM_PTR_NIL, "")
	test.EXPECT_EQ(t, allocSize, uint32(1024), "")
	test.EXPECT_EQ(t, allocator.Used(), uint32(1024), "")
	test.EXPECT_EQ(t, allocator.Stat().AllocNum(), StatNumber(1), "")
	test.EXPECT_EQ(t, allocator.Stat().AllocNumOk(), StatNumber(1), "")

	allocator.FreeAll()
	addr, _ = allocator.AllocWithClear(8)
	test.EXPECT_NE(t, addr, MEM_PTR_NIL, "")
	test.EXPECT_EQ(t, allocator.Used(), uint32(8), "")
	test.EXPECT_EQ(t, allocator.mem[allocator.used-8:allocator.used], []byte{0, 0, 0, 0, 0, 0, 0, 0}, "")

	allocator.FreeAll()
	addr = allocator.AllocBytes([]byte("12346"))
	test.EXPECT_NE(t, addr, MEM_PTR_NIL, "")
	test.EXPECT_EQ(t, allocator.GetString(addr), "12346", "")
	test.EXPECT_EQ(t, allocator.mem[addr-2], byte(5), "")

	allocator.FreeAll()

	addr = allocator.AllocBytesBegin()
	test.EXPECT_NE(t, addr, MEM_PTR_NIL, "")
	ret := allocator.AppendBytes([]byte("12378"))
	test.EXPECT_TRUE(t, ret, "")
	allocator.AllocBytesEnd(addr)
	test.EXPECT_EQ(t, allocator.Strlen(addr), 5, "")
	test.EXPECT_EQ(t, allocator.GetString(addr), "12378", "")

	allocator.FreeAll()
	addr = allocator.AllocBytesBegin()
	test.EXPECT_NE(t, addr, MEM_PTR_NIL, "")
	ret = allocator.AppendBytes([]byte("12378"))
	test.EXPECT_TRUE(t, ret, "")
	ret = allocator.AppendByte('6')
	test.EXPECT_TRUE(t, ret, "")
	allocator.AppendByteNoCheck('a')
	allocator.AllocBytesEnd(addr)
	test.EXPECT_EQ(t, allocator.Strlen(addr), 7, "")
	test.EXPECT_EQ(t, allocator.GetString(addr), "123786a", "")
}*/

func TestArenaAllocatorAllocNOk(t *testing.T) {
	allocator := NewArenaAllocator(1024, 1)
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
	allocator := NewArenaAllocator(1024, 1)
	test.EXPECT_NE(t, allocator.Capacity(), 1024, "")
	test.EXPECT_NE(t, allocator.Left(), 1024, "")

	allocator.AllocWithClear(16)
	allocator.FreePart(8)
	test.EXPECT_EQ(t, allocator.Used(), uint32(8), "")

	allocator.FreeAll()
	test.EXPECT_EQ(t, allocator.Used(), uint32(0), "")
}

func TestArenaAllocatorFreeNOk(t *testing.T) {
	allocator := NewArenaAllocator(1024, 1)
	test.EXPECT_NE(t, allocator.Capacity(), 1024, "")
	test.EXPECT_NE(t, allocator.Left(), 1024, "")

	allocator.AllocWithClear(16)
	allocator.FreePart(28)
	test.EXPECT_EQ(t, allocator.Used(), uint32(16), "")
}

func TestArenaAllocatorGetString(t *testing.T) {
	allocator := NewArenaAllocator(1024, 1)

	addr := allocator.AllocBytes(nil)
	test.EXPECT_EQ(t, allocator.Used(), uint32(2), "")
	test.EXPECT_EQ(t, allocator.GetString(addr), "", "")
	test.EXPECT_EQ(t, allocator.GetString(MEM_PTR_NIL), "", "")

	allocator.FreeAll()
	addr = allocator.AllocBytes([]byte("abfg"))
	test.EXPECT_EQ(t, allocator.GetString(addr), "abfg", "")
}

func TestArenaAllocatorStrlen(t *testing.T) {
	allocator := NewArenaAllocator(1024, 1)

	addr := allocator.AllocBytes(nil)
	test.EXPECT_EQ(t, allocator.Used(), uint32(2), "")
	test.EXPECT_EQ(t, allocator.Strlen(addr), 0, "")
	test.EXPECT_EQ(t, allocator.Strlen(MEM_PTR_NIL), 0, "")

	allocator.FreeAll()
	addr = allocator.AllocBytes([]byte("abfg"))
	test.EXPECT_EQ(t, allocator.Strlen(addr), 4, "")
}

func TestArenaAllocatorZeroMem(t *testing.T) {
	allocator := NewArenaAllocator(1024, 1)

	addr := allocator.AllocBytes([]byte("1234"))
	allocator.ZeroMem(addr, 4)
	test.EXPECT_EQ(t, allocator.mem[allocator.used-4:allocator.used], []byte{0, 0, 0, 0}, "")
}

func TestArenaAllocatorClone(t *testing.T) {
	allocator := NewArenaAllocator(1024, 1)

	addr := allocator.AllocBytes([]byte("1234"))

	newAllocator := allocator.Clone()
	test.ASSERT_NE(t, newAllocator, nil, "")

	test.EXPECT_EQ(t, allocator.GetString(addr), "1234", "")
	test.EXPECT_EQ(t, newAllocator.GetString(addr), "1234", "")
}

/*
func TestArenaAllocatorString(t *testing.T) {
	allocator := NewArenaAllocator(1024, 1)
	allocator.AllocBytes([]byte("1234"))

	wanted := `-------------------------- ArenaAllocator show begin ----------------------------
00000008h: 04 00 31 32 33 34                                ; ..1234
---------------------------------------------------------------------------------
STAT:
alloc num: 1
alloc num ok: 1

MEMORY:
used     = 6
left     = 1018
capacity = 1024
-------------------------- ArenaAllocator show end   ----------------------------
`
	str := allocator.String()
	test.EXPECT_EQ(t, str, wanted, "")

	allocator.FreeAll()
	allocator.AllocBytes([]byte("12345678901234567890123456789012345678901234567890123456789012345"))

	wanted = `-------------------------- ArenaAllocator show begin ----------------------------
00000008h: 41 00 31 32 33 34 35 36  37 38 39 30 31 32 33 34 ; A.123456 78901234
00000018h: 35 36 37 38 39 30 31 32  33 34 35 36 37 38 39 30 ; 56789012 34567890
00000028h: 31 32 33 34 35 36 37 38  39 30 31 32 33 34 35 36 ; 12345678 90123456
00000038h: 37 38 39 30 31 32 33 34  35 36 37 38 39 30 31 32 ; 78901234 56789012
---------------------------------------------------------------------------------
STAT:
alloc num: 2
alloc num ok: 2
free all num: 1

MEMORY:
used     = 67
left     = 957
capacity = 1024
-------------------------- ArenaAllocator show end   ----------------------------
`

	str = allocator.String()
	test.EXPECT_EQ(t, str, wanted, "")
}

func TestArenaAllocatorPrintAll(t *testing.T) {
	allocator := NewArenaAllocator(16, 1)
	allocator.AllocBytes([]byte("1234"))

	wanted := `-------------------------- ArenaAllocator show begin ----------------------------
00000008h: 04 00 31 32 33 34 00 00  00 00 00 00 00 00 00 00 ; ..1234.. ........
---------------------------------------------------------------------------------
STAT:
alloc num: 1
alloc num ok: 1

MEMORY:
used     = 6
left     = 10
capacity = 16
-------------------------- ArenaAllocator show end   ----------------------------
`
	buf := buffer.NewByteBuffer(nil)
	allocator.PrintAll(buf)
	test.EXPECT_EQ(t, buf.String(), wanted, "")
}

func TestArenaAllocatorPrintUsed(t *testing.T) {
	allocator := NewArenaAllocator(16, 1)
	allocator.AllocBytes([]byte("1234"))

	wanted := `-------------------------- ArenaAllocator show begin ----------------------------
00000008h: 04 00 31 32 33 34                                ; ..1234
---------------------------------------------------------------------------------
STAT:
alloc num: 1
alloc num ok: 1

MEMORY:
used     = 6
left     = 10
capacity = 16
-------------------------- ArenaAllocator show end   ----------------------------
`
	buf := buffer.NewByteBuffer(nil)
	allocator.PrintUsed(buf)
	test.EXPECT_EQ(t, buf.String(), wanted, "")
}
*/
func TestZeroMem(t *testing.T) {
	allocator := NewArenaAllocator(1024, 1)

	addr := allocator.AllocBytes([]byte("1234"))
	ZeroMem(allocator.GetUintptr(addr), 4)
	test.EXPECT_EQ(t, allocator.mem[allocator.used-4:allocator.used], []byte{0, 0, 0, 0}, "")
}

func BenchmarkArenaAllocatorAlloc(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024*128, 1)
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
	allocator := NewArenaAllocator(1024*128, 1)
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
	allocator := NewArenaAllocator(1024*128, 1)
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
	allocator := NewArenaAllocator(1024*128, 1)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		addr := allocator.AllocBytesBegin()
		allocator.AppendBytes(data)
		allocator.AllocBytesEnd(addr)
	}
}

func BenchmarkArenaAllocatorAllocBytes3(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024*128, 1)
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

func BenchmarkArenaAllocatorAllocBytes4(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024*128, 1)
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

func BenchmarkArenaAllocatorAllocBytes5(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024*128, 1)
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

func f1(allocator *ArenaAllocator, data []byte) {
	_, buf := allocator.AllocBytesBeginEx()
	/*for i, v := range data {
		buf[i] = v
	}*/
	if len(buf) < len(data) {
		return
	}

	copy(buf, data)
	allocator.AllocBytesEndEx(uint32(len(data)))
}

func BenchmarkArenaAllocatorAllocBytes5Ex1(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024*128, 1)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		_, buf := allocator.AllocBytesBeginEx()
		/*for i, v := range data {
			buf[i] = v
		}*/

		copy(buf, data)
		allocator.AllocBytesEndEx(uint32(len(data)))
	}
}

func BenchmarkArenaAllocatorAllocBytes5Ex2(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024*128, 1)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		_, buf := allocator.AllocBytesBeginEx()
		num := 0
		for i, v := range data {
			buf[i] = v
			num++
		}
		allocator.AllocBytesEndEx(uint32(num))
	}
}

func BenchmarkArenaAllocatorAllocBytes5Ex3(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024*128, 1)
	data := []byte("01234567890123456789012345678901")
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.FreeAll()
		f1(allocator, data)
	}
}

func BenchmarkZeroMem1(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024*128, 1)
	addr, _ := allocator.Alloc(32)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ZeroMem(allocator.GetUintptr(addr), 32)
	}
}

func BenchmarkZeroMem2(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024*128, 1)
	addr, _ := allocator.Alloc(32)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.ZeroMem(addr, 32)
	}
}

func BenchmarkZeroMem3(b *testing.B) {
	b.StopTimer()
	allocator := NewArenaAllocator(1024*128, 1)
	data := []byte("01234567890123456789012345678901")
	addr := allocator.AllocBytes(data)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		allocator.ZeroBytes(addr)
	}
}
