package mem

import (
	"encoding/binary"
	//"fmt"
	//"reflect"
	//"strconv"
	"unsafe"
	//"github.com/lioneagle/goutil/src/chars"
)

/* MemPtr save a integer value or a address shift of memory.
 * If highest bit is 1, it save a integer value, integer value range is between
 * 0 ~ 2**sizeof(MemPtr)-1
 */
type MemPtr uint16

const (
	MEM_PTR_BITS = unsafe.Sizeof(MemPtr(0)) * 8
	MEM_PTR_BIT  = MemPtr(1 << (MEM_PTR_BITS - 1))
	MEM_PTR_MASK = MemPtr(^MEM_PTR_BIT)
)

const MEM_PTR_NIL = MemPtr(0)

func (p MemPtr) GetUint() uint {
	return uint(p & MEM_PTR_MASK)
}

func MemPtrSetUint(value MemPtr) MemPtr {
	return value | MEM_PTR_BIT
}

func (p MemPtr) GetMemAddr(allocator *ArenaAllocator) *byte {
	return (*byte)(unsafe.Pointer(&allocator.mem[p]))
}

func (p MemPtr) GetMemPointer(allocator *ArenaAllocator) unsafe.Pointer {
	return unsafe.Pointer(&allocator.mem[p])
}

func (p MemPtr) GetUintptr(allocator *ArenaAllocator) uintptr {
	return (uintptr)(unsafe.Pointer(&allocator.mem[p]))
}
