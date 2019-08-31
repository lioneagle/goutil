package chars

import (
	"bytes"
)

type IsInCharset func(ch byte) bool

func PercentUnescape(src []byte) (dst []byte) {
	if bytes.IndexByte(src, '%') == -1 {
		return src
	}

	len1 := len(src)

	for i := 0; i < len1; {
		if (src[i] == '%') && ((i + 2) < len1) && IsHex(src[i+1]) && IsHex(src[i+2]) {
			dst = append(dst, PercentUnescapeToByteEx(src[i+1], src[i+2]))
			i += 3
		} else {
			dst = append(dst, src[i])
			i++
		}
	}

	return dst
}

func PercentUnescapeToByte(src []byte) byte {
	return HexToByte(src[1])<<4 | HexToByte(src[2])
}

func PercentUnescapeToByteEx(v1, v2 byte) byte {
	return HexToByte(v1)<<4 | HexToByte(v2)
}

func PercentEscape(src []byte, inCharset IsInCharset) (dst []byte) {
	if !NeedEscape(src, inCharset) {
		return src
	}

	for _, v := range src {
		if inCharset(v) {
			dst = append(dst, v)
		} else {
			dst = append(dst, '%', ToUpperHex(v>>4), ToUpperHex(v))
		}
	}

	return dst
}

func NeedEscape(src []byte, inCharset IsInCharset) bool {
	for _, v := range src {
		if !inCharset(v) {
			return true
		}
	}
	return false
}

func PercentEscapeEx(src []byte, charset *[256]uint32, mask uint32) (dst []byte) {
	if !NeedEscapeEx(src, charset, mask) {
		return src
	}

	charset1 := charset

	for _, v := range src {
		if (charset1[v] & mask) != 0 {
			dst = append(dst, v)
		} else {
			dst = append(dst, '%', ToUpperHex(v>>4), ToUpperHex(v))
		}
	}

	return dst
}

func NeedEscapeEx(src []byte, charset *[256]uint32, mask uint32) bool {
	charset1 := charset
	for _, v := range src {
		if (charset1[v] & mask) == 0 {
			return true
		}
	}
	return false
}

func NeedEscapeNum(src []byte, charset *[256]uint32, mask uint32) int {
	charset1 := charset
	num := 0
	for _, v := range src {
		if (charset1[v] & mask) == 0 {
			num++
		}
	}
	return num
}
