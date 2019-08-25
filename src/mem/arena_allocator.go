package mem

import (
	//"fmt"
	//"io"
	"encoding/binary"
	"reflect"
	"unsafe"
	//"github.com/lioneagle/goutil/src/buffer"
)

const SLICE_HEADER_LEN = int32(unsafe.Sizeof(reflect.SliceHeader{}))
const ARENA_ALLOCATOR_ALIGN = uint32(1)

//const ABNF_MEM_LIGN_MASK = ^(ABNF_MEM_ALIGN - 1)
//const ABNF_MEM_LIGN_MASK2 = (ABNF_MEM_ALIGN - 1)
const ARENA_ALLOCATOR_PREFIX_LEN = 8

func RoundToAlign(x, align uint32) uint32 {
	return (x + align - 1) & ^(align - 1)
}

/* ArenaAllocator is a memory allocator for text/binary protocol processing,
 * it allocates memories and does not deallocate every memories which have
 * been allocated. It deallacator all memories by FreeAll.
 *
 * ArenaAllocator has no lock, so it is not gorouting safe, each gorouting
 * should use their own ArenaAllocator
 *
 * ArenaAllocator allocates memory fast and support alloctes memory when parsing.
 * ArenaAllocator can be reused when one protocol message is processed completed.
 */
type ArenaAllocator struct {
	used uint32
	stat ArenaAllocatorStat
	mem  []byte
}

func NewArenaAllocator(capacity uint32) *ArenaAllocator {
	ret := ArenaAllocator{}
	ret.Init(capacity)
	return &ret
}

func (this *ArenaAllocator) Init(capacity uint32) *ArenaAllocator {
	this.used = ARENA_ALLOCATOR_PREFIX_LEN
	this.mem = make([]byte, int(RoundToAlign(capacity+ARENA_ALLOCATOR_PREFIX_LEN, ARENA_ALLOCATOR_ALIGN)))
	this.stat.Init()
	return this
}

func (this *ArenaAllocator) Stat() *ArenaAllocatorStat {
	return &this.stat
}

func (this *ArenaAllocator) Used() uint32 {
	return this.used - ARENA_ALLOCATOR_PREFIX_LEN
}

func (this *ArenaAllocator) Capacity() uint32 {
	return uint32(cap(this.mem) - ARENA_ALLOCATOR_PREFIX_LEN)
}

func (this *ArenaAllocator) Left() uint32 {
	return uint32(cap(this.mem)) - this.used
}

func (this *ArenaAllocator) Alloc(size uint32) (addr MemPtr, allocSize uint32) {
	this.stat.allocNum++

	used := this.used

	this.used = (this.used + size + ARENA_ALLOCATOR_ALIGN - 1) & ^(ARENA_ALLOCATOR_ALIGN - 1)
	if this.used > uint32(cap(this.mem)) {
		return MEM_PTR_NIL, 0
	}

	this.stat.allocNumOk++

	return MemPtr(used), this.used - used
}

func (this *ArenaAllocator) AllocWithClear(size uint32) (addr MemPtr, allocSize uint32) {
	addr, num := this.Alloc(size)
	if addr != MEM_PTR_NIL {
		src := this.mem[addr : addr+MemPtr(num)]
		for i := range src {
			src[i] = 0
		}
	}

	return addr, num
}

func (this *ArenaAllocator) ZeroMem(addr MemPtr, num uint32) {
	src := this.mem[addr : uint32(addr)+num]
	for i := range src {
		src[i] = 0
	}
}

func (this *ArenaAllocator) AllocBytes(data []byte) MemPtr {
	addr, _ := this.Alloc(uint32(len(data)))
	if addr != MEM_PTR_NIL {
		copy(this.mem[addr:], data)
		binary.LittleEndian.PutUint16(this.mem[addr-2:], uint16(len(data)))
	}

	return addr
}

func (this *ArenaAllocator) AllocBytesBegin() MemPtr {
	this.stat.allocNum++

	this.used += 2

	if this.used > uint32(cap(this.mem)) {
		return MEM_PTR_NIL
	}

	return MemPtr(this.used)
}

func (this *ArenaAllocator) AppendBytes(data []byte) bool {
	newsize := this.used + uint32(len(data))
	if newsize > uint32(cap(this.mem)) {
		return false
	}
	copy(this.mem[this.used:], data)
	this.used = newsize
	return true
}

func (this *ArenaAllocator) AppendByte(data byte) bool {
	newsize := this.used + 1
	if newsize > uint32(cap(this.mem)) {
		return false
	}
	this.mem[this.used] = data
	this.used++
	return true
}

func (this *ArenaAllocator) AllocBytesEnd(addr MemPtr) {
	num := this.used - uint32(addr)
	this.used = (this.used + ARENA_ALLOCATOR_ALIGN - 1) & ^(ARENA_ALLOCATOR_ALIGN - 1)
	binary.LittleEndian.PutUint16(this.mem[addr-2:], uint16(num))
	this.stat.allocNumOk++
}

func (this *ArenaAllocator) BytesLen(addr MemPtr) int {
	if addr < 2 {
		return 0
	}
	return int(binary.LittleEndian.Uint16(this.mem[addr-2:]))
}

func ZeroMem(addr uintptr, size int) {
	h := reflect.SliceHeader{Data: addr, Len: size, Cap: size}
	x := *(*[]byte)(unsafe.Pointer(&h))

	for i := range x {
		x[i] = 0
	}
}

func (this *ArenaAllocator) FreeAll() {
	this.stat.freeAllNum++
	this.used = ARENA_ALLOCATOR_PREFIX_LEN
}

func (this *ArenaAllocator) FreePart(remain uint32) {
	this.stat.freePartNum++
	if remain >= this.used {
		return
	}
	this.used = remain + ARENA_ALLOCATOR_PREFIX_LEN
	if this.used < ARENA_ALLOCATOR_PREFIX_LEN {
		this.used = ARENA_ALLOCATOR_PREFIX_LEN
	}
}
