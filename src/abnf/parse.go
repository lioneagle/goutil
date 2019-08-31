package abnf

import (
	//"fmt"

	"github.com/lioneagle/goutil/src/chars"
	"github.com/lioneagle/goutil/src/mem"
)

/* parse input when char is in charset and not empty, and allocte memory for output
 */
func ParseInCharsetAndAlloc(allocator *mem.ArenaAllocator, src []byte, pos Pos, charset *[256]uint32, mask uint32) (addr mem.MemPtr, newPos Pos) {
	newPos = ParseInCharset(src, pos, charset, mask)

	if newPos <= pos {
		return mem.MEM_PTR_NIL, newPos
	}

	addr = allocator.AllocBytes(src[pos:newPos])
	return addr, newPos
}

/* parse input when char is in charset
 */
func ParseInCharset(src []byte, pos Pos, charset *[256]uint32, mask uint32) (newPos Pos) {
	for newPos = pos; newPos < Pos(len(src)); newPos++ {
		if (charset[src[newPos]] & mask) == 0 {
			break
		}
	}

	return newPos
}

/* parse input when char is in charset and char may be percent escaped, such as "%61"
 */
func ParseInCharsetPercentEscapable(allocator *mem.ArenaAllocator, src []byte, pos Pos, charset *[256]uint32, mask uint32) (addr mem.MemPtr, newPos Pos, err error) {
	newPos = pos
	len1 := Pos(len(src))

	v := src[newPos]
	if ((charset[v] & mask) == 0) && (v != '%') {
		return mem.MEM_PTR_NIL, newPos, nil
	}

	addr = allocator.AllocBytesBegin()
	if addr == mem.MEM_PTR_NIL {
		return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "no mem")
	}

	prevPos := newPos
	for newPos < len1 {
		v = src[newPos]
		if ((charset[v]) & mask) == 0 {

			if v != '%' {
				break
			}

			if (newPos + 2) >= len1 {
				return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "reach end after '%'")
			}

			v1 := src[newPos+1]
			v2 := src[newPos+2]

			if !chars.IsHex(v1) || !chars.IsHex(v2) {
				return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "not HEX after '%'")
			}

			v = chars.UnescapeToByteEx(v1, v2)
			if ((charset[v]) & mask) == 0 {
				break
			}

			if (prevPos < newPos) && !allocator.AppendBytes(src[prevPos:newPos]) {
				return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "no mem")
			}

			if !allocator.AppendByte(v) {
				return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "no mem")
			}

			newPos += 3
			prevPos = newPos
		} else {
			newPos++
		}
	}

	if (prevPos < newPos) && !allocator.AppendBytes(src[prevPos:newPos]) {
		return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "no mem")
	}

	allocator.AllocBytesEnd(addr)
	return addr, newPos, nil
}
