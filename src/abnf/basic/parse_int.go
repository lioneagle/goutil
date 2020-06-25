package abnf

import (
	_ "fmt"

	"github.com/lioneagle/goutil/src/chars"
)

/* ParseUInt64 parse digit in src which can contain non-digit char,
   and convert to int64, return false if overflow
   ParseUInt64 is faster than strconv.ParseUint and strconv.Atoi
*/
func ParseUint64(src []byte, pos Pos) (digit uint64, newPos Pos, ok bool) {
	// Cutoff is the smallest number such that cutoff*base > maxUint64
	const maxUint64 = (1<<64 - 1)
	const cutoff = maxUint64/10 + 1

	len1 := Pos(len(src))
	charset := &chars.Charsets0
	if pos >= len1 || ((charset[src[pos]] & chars.MASK_DIGIT) == 0) {
		return 0, pos, false
	}

	var end1 Pos

	if len1-pos < 19 {
		end1 = len1
	} else {
		end1 = pos + 19
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

/* EncodeUint encode uint64 to byte buffer
   EncodeUint is faster than strconv.FormatUint
*/
func EncodeUint(buf *ByteBuffer, digit uint64) {
	if digit < 10 {
		buf.Write(g_digitsString[digit : digit+1])
	} else if digit < 100 {
		buf.Write(g_smallDigitsString[digit*2 : digit*2+2])
	} else {

		var val [32]byte
		num := 31
		for digit > 0 {
			mod := digit
			digit /= 10
			val[num] = '0' + byte(mod-digit*10)
			num--
		}

		buf.Write(val[num+1:])
	}
}

/* EncodeUintWithWidth encode uint64 to byte buffer with width,
   if length of digit is less than width, spaces are filled before digit.
   if length of digit is larger than width, no space is before digit.
   EncodeUintWithWidth is faster than fmt.Sprintf
*/
func EncodeUintWithWidth(buf *ByteBuffer, digit uint64, width int) {
	if digit < 10 {
		for i := 1; i < width; i++ {
			buf.WriteByte(' ')
		}
		buf.Write(g_digitsString[digit : digit+1])
	} else if digit < 100 {
		for i := 2; i < width; i++ {
			buf.WriteByte(' ')
		}
		buf.Write(g_smallDigitsString[digit*2 : digit*2+2])
	} else {
		var val [32]byte
		num := 31
		for digit > 0 {
			mod := digit
			digit /= 10
			val[num] = '0' + byte(mod-digit*10)
			num--
		}

		for i := width - 32 + num + 1; i > 0; i-- {
			buf.WriteByte(' ')
		}

		buf.Write(val[num+1:])
	}
}
