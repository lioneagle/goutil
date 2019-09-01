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
   using goto for performance
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

	var prevPos Pos

	for {
		prevPos = newPos
		for ; newPos < len1; newPos++ {
			v = src[newPos]
			if ((charset[v]) & mask) == 0 {
				if v == '%' {
					goto unescape
				}
				goto end
			}
		}
		goto end
	unescape:
		if (newPos + 2) >= len1 {
			return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "reach end after '%'")
		}

		v1 := src[newPos+1]
		v2 := src[newPos+2]

		if !chars.IsHex(v1) || !chars.IsHex(v2) {
			return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "not HEX after '%'")
		}

		v = chars.PercentUnescapeToByteEx(v1, v2)

		if (prevPos < newPos) && !allocator.AppendBytes(src[prevPos:newPos]) {
			return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "no mem")
		}

		if !allocator.AppendByte(v) {
			return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "no mem")
		}

		newPos += 3
	}
end:
	if (prevPos < newPos) && !allocator.AppendBytes(src[prevPos:newPos]) {
		return mem.MEM_PTR_NIL, newPos, NewError(src, newPos, "no mem")
	}

	allocator.AllocBytesEnd(addr)
	return addr, newPos, nil
}

/* ParseUInt64 parse digit in src which can contain non-digit char,
   and convert to int64, return false if overflow
   ParseUInt64 is faster than strconv.ParseUint and strconv.Atoi
*/
func ParseUint64(src []byte, pos Pos) (digit uint64, newPos Pos, ok bool) {
	const cutoff = (1<<64-1)/10 + 1

	len1 := Pos(len(src))
	charset := &chars.Charsets0
	if pos >= len1 || ((charset[src[pos]] & chars.MASK_DIGIT) == 0) {
		return 0, pos, false
	}

	var end1 Pos

	if len1 < 19 {
		end1 = newPos + len1
	} else {
		end1 = newPos + 19
	}

	digit = uint64(0)
	for newPos = pos; newPos < end1 && ((charset[src[newPos]] & chars.MASK_DIGIT) != 0); newPos++ {
		digit = digit*10 + uint64(src[newPos]) - '0'
	}

	if digit >= cutoff {
		return 0, newPos, false
	}

	if newPos >= len1 || (charset[src[newPos]]&chars.MASK_DIGIT) == 0 {
		return digit, newPos, true
	}

	digit1 := digit * 10
	digit = digit1 + uint64(src[newPos]) - '0'
	if digit < digit1 {
		return 0, newPos, false
	}

	newPos++

	if newPos >= len1 {
		return digit, newPos, true
	}

	if (charset[src[newPos]] & chars.MASK_DIGIT) != 0 {
		return digit, newPos, false
	}

	return digit, newPos, true
}

/* ParseUInt32 parse digit in src which can contain non-digit char,
   and convert to int64, return false if overflow
   ParseUInt32 is faster than strconv.ParseUint and strconv.Atoi
*/
func ParseUint32(src []byte, pos Pos) (digit uint32, newPos Pos, ok bool) {
	len1 := Pos(len(src))
	charset := &chars.Charsets0
	if pos >= len1 || ((charset[src[pos]] & chars.MASK_DIGIT) == 0) {
		return 0, pos, false
	}

	digit1 := uint64(0)

	for newPos = pos; newPos < len1 && ((charset[src[newPos]] & chars.MASK_DIGIT) != 0); newPos++ {
		digit1 = digit1*10 + uint64(src[newPos]) - '0'
	}

	if (newPos-pos) > 10 || (digit1 > 0xffffffff) {
		return 0, newPos, false
	}

	return uint32(digit1), newPos, true
}

/* ParseUInt16 parse digit in src which can contain non-digit char,
   and convert to int64, return false if overflow
   ParseUInt16 is faster than strconv.ParseUint and strconv.Atoi
*/
func ParseUint16(src []byte, pos Pos) (digit uint16, newPos Pos, ok bool) {
	len1 := Pos(len(src))
	charset := &chars.Charsets0
	if pos >= len1 || ((charset[src[pos]] & chars.MASK_DIGIT) == 0) {
		return 0, pos, false
	}

	digit1 := uint32(0)

	for newPos = pos; newPos < len1 && ((charset[src[newPos]] & chars.MASK_DIGIT) != 0); newPos++ {
		digit1 = digit1*10 + uint32(src[newPos]) - '0'
	}

	if (newPos-pos) > 5 || (digit1 > 0xffff) {
		return 0, newPos, false
	}

	return uint16(digit1), newPos, true
}

/* ParseUInt8 parse digit in src which can contain non-digit char,
   and convert to int64, return false if overflow
   ParseUInt16 is faster than strconv.ParseUint and strconv.Atoi
*/
func ParseUint8(src []byte, pos Pos) (digit uint8, newPos Pos, ok bool) {
	len1 := Pos(len(src))
	charset := &chars.Charsets0
	if pos >= len1 || ((charset[src[pos]] & chars.MASK_DIGIT) == 0) {
		return 0, pos, false
	}

	digit1 := uint16(0)

	for newPos = pos; newPos < len1 && ((charset[src[newPos]] & chars.MASK_DIGIT) != 0); newPos++ {
		digit1 = digit1*10 + uint16(src[newPos]) - '0'
	}

	if (newPos-pos) > 3 || (digit1 > 0xff) {
		return 0, newPos, false
	}

	return uint8(digit1), newPos, true
}
