package container

import (
	//"fmt"
	"testing"
)

func TestChunkAllocatorAlloc1(t *testing.T) {

	allocator := NewChunkAllocator(1)

	c := allocator.Alloc()
	if c == -1 {
		t.Errorf("TestChunkAllocatorAlloc: should alloc ok")
		return
	}

	if c != 0 {
		t.Errorf("TestChunkAllocatorAlloc: wrong id = %d, wanted = 0", c)
		return
	}

	/*if allocator.busyHead != 0 {
		t.Errorf("TestChunkAllocatorAlloc: wrong busyHead = %d, wanted = 0", allocator.busyHead)
		return
	}*/

	if allocator.freeHead != -1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong freeHead = %d, wanted = -1", allocator.freeHead)
		return
	}

	if allocator.size != 1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong size = %d, wanted = 1", allocator.size)
		return
	}

	c = allocator.Alloc()
	if c != -1 {
		t.Errorf("TestChunkAllocatorAlloc: should alloc nok")
		return
	}

}

func TestChunkAllocatorAlloc2(t *testing.T) {

	allocator := NewChunkAllocator(2)

	c := allocator.Alloc()
	if c == -1 {
		t.Errorf("TestChunkAllocatorAlloc: should alloc ok")
		return
	}

	if c != 0 {
		t.Errorf("TestChunkAllocatorAlloc: wrong id = %d, wanted = 0", c)
		return
	}

	/*if allocator.busyHead != 0 {
		t.Errorf("TestChunkAllocatorAlloc: wrong busyHead = %d, wanted = 0", allocator.busyHead)
		return
	}*/

	if allocator.freeHead != 1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong freeHead = %d, wanted = 1", allocator.freeHead)
		return
	}

	if allocator.size != 1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong size = %d, wanted = 1", allocator.size)
		return
	}

	c = allocator.Alloc()
	if c == -1 {
		t.Errorf("TestChunkAllocatorAlloc: should alloc ok")
		return
	}

	if c != 1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong id = %d, wanted = 1", c)
		return
	}

	/*if allocator.busyHead != 0 {
		t.Errorf("TestChunkAllocatorAlloc: wrong busyHead = %d, wanted = 1", allocator.busyHead)
		return
	}*/

	if allocator.freeHead != -1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong freeHead = %d, wanted = -1", allocator.freeHead)
		return
	}

	if allocator.size != 2 {
		t.Errorf("TestChunkAllocatorAlloc: wrong size = %d, wanted = 2", allocator.size)
		return
	}

	c = allocator.Alloc()
	if c != -1 {
		t.Errorf("TestChunkAllocatorAlloc: should alloc nok")
		return
	}

}

func TestChunkAllocatorFree(t *testing.T) {
	allocator := NewChunkAllocator(2)

	c1 := allocator.Alloc()
	if c1 == -1 {
		t.Errorf("TestChunkAllocatorAlloc: should alloc ok")
		return
	}

	c2 := allocator.Alloc()
	if c2 == -1 {
		t.Errorf("TestChunkAllocatorAlloc: should alloc ok")
		return
	}

	allocator.Free(c1)

	if allocator.size != 1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong size = %d, wanted = 1", allocator.size)
		return
	}

	/*if allocator.busyHead != 1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong busyHead = %d, wanted = 0", allocator.busyHead)
		return
	}*/

	if allocator.freeHead != 0 {
		t.Errorf("TestChunkAllocatorAlloc: wrong freeHead = %d, wanted = 0", allocator.freeHead)
		return
	}

	allocator.Free(c1)

	if allocator.size != 1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong size = %d, wanted = 1", allocator.size)
		return
	}

	/*if allocator.busyHead != 1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong busyHead = %d, wanted = 0", allocator.busyHead)
		return
	}*/

	if allocator.freeHead != 0 {
		t.Errorf("TestChunkAllocatorAlloc: wrong freeHead = %d, wanted = 0", allocator.freeHead)
		return
	}

	allocator.Free(c2)

	if allocator.size != 0 {
		t.Errorf("TestChunkAllocatorAlloc: wrong size = %d, wanted = 0", allocator.size)
		return
	}

	/*if allocator.busyHead != -1 {
		t.Errorf("TestChunkAllocatorAlloc: wrong busyHead = %d, wanted = -1", allocator.busyHead)
		return
	}*/

	if allocator.freeHead != 0 {
		t.Errorf("TestChunkAllocatorAlloc: wrong freeHead = %d, wanted = 0", allocator.freeHead)
		return
	}

}

func BenchmarkChunkAllocatorAllocFree(b *testing.B) {
	b.StopTimer()
	allocator := NewChunkAllocator(1000)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		c := allocator.AllocEx()
		allocator.Free(c.id)
	}
}

/*func BenchmarkChunkAllocatorAllocFreeEx(b *testing.B) {
	b.StopTimer()
	allocator := NewChunkAllocator(1000)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		c := allocator.AllocEx()
		allocator.FreeEx(c)
	}
}*/
