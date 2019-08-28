package abnf

import (
	"github.com/lioneagle/goutil/src/mem"
)

func ParseCharsetAndAlloc(allocator *mem.ArenaAllocator, src []byte, pos int, charset *[256]uint32, mask uint32) (addr mem.MemPtr, newPos int) {
	newPos = ParseCharset(src, pos, charset, mask)

	if newPos <= pos {
		return mem.MEM_PTR_NIL, newPos
	}

	addr = allocator.AllocBytes(src[pos:newPos])
	return addr, newPos
}

func ParseCharsetAndAllocEnableEmpty(allocator *mem.ArenaAllocator, src []byte, pos int, charset *[256]uint32, mask uint32) (addr mem.MemPtr, newPos int) {
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
