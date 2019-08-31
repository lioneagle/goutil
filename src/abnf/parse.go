package abnf

import (
	"github.com/lioneagle/goutil/src/mem"
)

/* parse input when char is in charset and not empty, and allocte memory for output
 */
func ParseInCharsetAndAlloc(allocator *mem.ArenaAllocator, src []byte, pos int, charset *[256]uint32, mask uint32) (addr mem.MemPtr, newPos int) {
	newPos = ParseInCharset(src, pos, charset, mask)

	if newPos <= pos {
		return mem.MEM_PTR_NIL, newPos
	}

	addr = allocator.AllocBytes(src[pos:newPos])
	return addr, newPos
}

func ParseInCharset(src []byte, pos int, charset *[256]uint32, mask uint32) (newPos int) {
	for newPos = pos; newPos < len(src); newPos++ {
		if (charset[src[newPos]] & mask) == 0 {
			break
		}
	}

	return newPos
}

func ParseInCharsetPercentEscapable(src []byte, pos int, charset *[256]uint32, mask uint32) (newPos int) {
	for newPos = pos; newPos < len(src); newPos++ {
		if (charset[src[newPos]] & mask) == 0 {
			break
		}
	}

	return newPos
}
