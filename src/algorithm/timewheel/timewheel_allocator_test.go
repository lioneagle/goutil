package timewheel

import (
	//"fmt"
	"testing"
)

func TestTimeWheelAllocatorAlloc1(t *testing.T) {

	allocator := NewTimeWheelAllocator(1)
	funcName := "TestTimeWheelAllocatorAlloc1"

	c := allocator.Alloc()
	if c == -1 {
		t.Errorf("%s: should alloc ok", funcName)
		return
	}

	if c != 0 {
		t.Errorf("%s: wrong id = %d, wanted = 0", funcName, c)
		return
	}

	if allocator.freeHead != -1 {
		t.Errorf("%s: wrong freeHead = %d, wanted = -1", funcName, allocator.freeHead)
		return
	}

	if allocator.size != 1 {
		t.Errorf("%s: wrong size = %d, wanted = 1", funcName, allocator.size)
		return
	}

	c = allocator.Alloc()
	if c != -1 {
		t.Errorf("%s: should alloc nok", funcName)
		return
	}

}

func TestTimeWheelAllocatorAlloc2(t *testing.T) {

	allocator := NewTimeWheelAllocator(2)
	funcName := "TestTimeWheelAllocatorAlloc2"

	c := allocator.Alloc()
	if c == -1 {
		t.Errorf("%s: should alloc ok", funcName)
		return
	}

	if c != 0 {
		t.Errorf("%s: wrong id = %d, wanted = 0", funcName, c)
		return
	}

	if allocator.freeHead != 1 {
		t.Errorf("%s: wrong freeHead = %d, wanted = 1", funcName, allocator.freeHead)
		return
	}

	if allocator.size != 1 {
		t.Errorf("%s: wrong size = %d, wanted = 1", funcName, allocator.size)
		return
	}

	c = allocator.Alloc()
	if c == -1 {
		t.Errorf("%s: should alloc ok", funcName)
		return
	}

	if c != 1 {
		t.Errorf("%s: wrong id = %d, wanted = 1", funcName, c)
		return
	}

	if allocator.freeHead != -1 {
		t.Errorf("%s: wrong freeHead = %d, wanted = -1", funcName, allocator.freeHead)
		return
	}

	if allocator.size != 2 {
		t.Errorf("%s: wrong size = %d, wanted = 2", funcName, allocator.size)
		return
	}

	c = allocator.Alloc()
	if c != -1 {
		t.Errorf("%s: should alloc nok", funcName)
		return
	}

}

func TestTimeWheelAllocatorFree(t *testing.T) {
	allocator := NewTimeWheelAllocator(2)
	funcName := "TestTimeWheelAllocatorFree"

	c1 := allocator.Alloc()
	if c1 == -1 {
		t.Errorf("%s: should alloc ok", funcName)
		return
	}

	c2 := allocator.Alloc()
	if c2 == -1 {
		t.Errorf("%s: should alloc ok", funcName)
		return
	}

	allocator.Free(c1)

	if allocator.size != 1 {
		t.Errorf("%s: wrong size = %d, wanted = 1", funcName, allocator.size)
		return
	}

	if allocator.freeHead != 0 {
		t.Errorf("%s: wrong freeHead = %d, wanted = 0", funcName, allocator.freeHead)
		return
	}

	allocator.Free(c1)

	if allocator.size != 1 {
		t.Errorf("%s: wrong size = %d, wanted = 1", funcName, allocator.size)
		return
	}

	if allocator.freeHead != 0 {
		t.Errorf("%s: wrong freeHead = %d, wanted = 0", funcName, allocator.freeHead)
		return
	}

	allocator.Free(c2)

	if allocator.size != 0 {
		t.Errorf("%s: wrong size = %d, wanted = 0", funcName, allocator.size)
		return
	}

	if allocator.freeHead != 0 {
		t.Errorf("%s: wrong freeHead = %d, wanted = 0", funcName, allocator.freeHead)
		return
	}

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
