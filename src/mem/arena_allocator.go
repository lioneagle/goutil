package mem

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"unsafe"

	"github.com/lioneagle/goutil/src/buffer"
)

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
	ArenaAllocatorStat
	used  uint32
	align uint32
	mem   []byte
}

/* align must be power of 2
 */
func NewArenaAllocator(capacity, align uint32) *ArenaAllocator {
	ret := ArenaAllocator{align: align}
	ret.Init(capacity)
	return &ret
}

func (this *ArenaAllocator) Init(capacity uint32) *ArenaAllocator {
	this.used = ARENA_ALLOCATOR_PREFIX_LEN
	this.mem = make([]byte, int(RoundToAlign(capacity+ARENA_ALLOCATOR_PREFIX_LEN, this.align)))
	this.ArenaAllocatorStat.Init()
	return this
}

func (this *ArenaAllocator) Stat() *ArenaAllocatorStat {
	return &this.ArenaAllocatorStat
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
	this.allocNum++
	used := this.used

	this.used = RoundToAlign(this.used+size, this.align)
	if this.used > uint32(cap(this.mem)) {
		return MEM_PTR_NIL, 0
	}

	this.allocNumOk++

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
	addr, _ := this.Alloc(uint32(len(data) + 2))
	if addr != MEM_PTR_NIL {
		copy(this.mem[addr+2:], data)
		binary.LittleEndian.PutUint16(this.mem[addr:], uint16(len(data)))
	}

	return addr + 2
}

func (this *ArenaAllocator) AllocBytesBegin() MemPtr {
	this.allocNum++
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

func (this *ArenaAllocator) AppendByteNoCheck(data byte) {
	this.mem[this.used] = data
	this.used++
}

func (this *ArenaAllocator) AllocBytesEnd(addr MemPtr) {
	num := this.used - uint32(addr)
	this.used = RoundToAlign(this.used, this.align)
	binary.LittleEndian.PutUint16(this.mem[addr-2:], uint16(num))
	this.allocNumOk++
}

func (this *ArenaAllocator) Strlen(addr MemPtr) int {
	if addr < 2 {
		return 0
	}
	return int(binary.LittleEndian.Uint16(this.mem[addr-2:]))
}

func (this *ArenaAllocator) FreeAll() {
	this.freeAllNum++
	this.used = ARENA_ALLOCATOR_PREFIX_LEN
}

func (this *ArenaAllocator) FreePart(remain uint32) {
	this.freePartNum++
	if remain >= this.used {
		return
	}
	this.used = remain + ARENA_ALLOCATOR_PREFIX_LEN
}

func (this *ArenaAllocator) GetString(addr MemPtr) string {
	if addr == MEM_PTR_NIL {
		return ""
	}

	header := reflect.StringHeader{Data: this.GetUintptr(addr), Len: this.Strlen(addr)}
	return *(*string)(unsafe.Pointer(&header))
}

func (this *ArenaAllocator) Clone() *ArenaAllocator {
	newAllocator := NewArenaAllocator(this.Capacity(), this.align)
	newAllocator.used = this.used
	copy(newAllocator.mem[:this.used], this.mem[:this.used])
	return newAllocator
}

func (this *ArenaAllocator) GetUintptr(addr MemPtr) uintptr {
	return (uintptr)(unsafe.Pointer(&this.mem[addr]))
}

func (this *ArenaAllocator) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.PrintBrief(buf)
	return buf.String()
}

func (this *ArenaAllocator) PrintAll(w io.Writer) {
	this.Print(w, 0, int(this.Capacity()))
}

func (this *ArenaAllocator) PrintUsed(w io.Writer) {
	this.Print(w, 0, int(this.Used()))
}

func (this *ArenaAllocator) PrintBrief(w io.Writer) {
	size := int(this.Used())
	if size > 64 {
		size = 64
	}
	this.Print(w, 0, size)
}

func (this *ArenaAllocator) Print(w io.Writer, memBegin, memEnd int) {
	fmt.Fprintln(w, "-------------------------- ArenaAllocator show begin ----------------------------")
	buffer.PrintAsHex(w, this.mem, memBegin+ARENA_ALLOCATOR_PREFIX_LEN, memEnd+ARENA_ALLOCATOR_PREFIX_LEN)
	fmt.Fprintln(w, "---------------------------------------------------------------------------------")
	fmt.Fprintln(w, "STAT:")
	this.ArenaAllocatorStat.Print(w)
	fmt.Fprintln(w, "\nMEMORY:")
	fmt.Fprintf(w, "used     = %d\n", this.Used())
	fmt.Fprintf(w, "left     = %d\n", this.Left())
	fmt.Fprintf(w, "capacity = %d\n", this.Capacity())
	fmt.Fprintln(w, "-------------------------- ArenaAllocator show end   ----------------------------")
}

func ZeroMem(addr uintptr, size int) {
	h := reflect.SliceHeader{Data: addr, Len: size, Cap: size}
	x := *(*[]byte)(unsafe.Pointer(&h))

	for i := range x {
		x[i] = 0
	}
}

func ParseCharsetAndAlloc(allocator *ArenaAllocator, src []byte, pos int, charset *[256]uint32, mask uint32) (addr MemPtr, newPos int) {
	newPos = ParseCharset(src, pos, charset, mask)

	if newPos <= pos {
		return MEM_PTR_NIL, newPos
	}

	addr = allocator.AllocBytes(src[pos:newPos])
	return addr, newPos
}

func ParseCharsetAndAllocEnableEmpty(allocator *ArenaAllocator, src []byte, pos int, charset *[256]uint32, mask uint32) (addr MemPtr, newPos int) {
	newPos = ParseCharset(src, pos, charset, mask)
	addr = allocator.AllocBytes(src[pos:newPos])
	return addr, newPos
}

func ParseCharset(src []byte, pos int, charset *[256]uint32, mask uint32) (newPos int) {
	for newPos = pos; newPos < len(src); newPos++ {
		if (charset[src[newPos]] & mask) == 0 {
			break
		}
	}

	return newPos
}
