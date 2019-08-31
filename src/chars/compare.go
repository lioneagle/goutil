package chars

import (
	"unsafe"
)

func ToUpperHex(ch byte) byte {
	return "0123456789ABCDEF"[ch&0x0F]
}

func ToLowerHex(ch byte) byte {
	return "0123456789abcdef"[ch&0x0F]
}

func ToUpper(ch byte) byte {
	return g_toupper_table[ch]
}

func ToLower(ch byte) byte {
	return g_tolower_table[ch]
}

func HexToByte(ch byte) byte {
	//if IsDigit(ch) {
	if (Charsets0[ch] & MASK_DIGIT) != 0 {
		//return ch - '0'
		return ch & 0x0F
	}
	//if IsLowerHexAlpha(ch) {
	if (Charsets0[ch] & MASK_LOWER_HEX_ALPHA) != 0 {
		return ch - 'a' + 10
	}
	return ch - 'A' + 10
}

func CompareNoCase(s1, s2 []byte) int {
	len1 := len(s1)
	if len1 != len(s2) {
		return len1 - len(s2)
	}

	for i := 0; i < len1; i++ {
		if s1[i] != s2[i] {
			ch1 := ToLower(s1[i])
			ch2 := ToLower(s2[i])
			if ch1 != ch2 {
				return int(ch1) - int(ch2)
			}
		}
	}

	return 0
}

/*func EqualNoCase0(s1, s2 []byte) bool {
	len1 := len(s1)
	if len1 != len(s2) {
		return false
	}

	if ToLower(s1[0]) != ToLower(s2[0]) {
		return false
	}

	for i := 1; i < len1; i++ {
		if s1[i] != s2[i] {
			if ToLower(s1[i]) != ToLower(s2[i]) {
				return false
			}
		}
	}

	return true
}*/

func EqualNoCase1(s1, s2 []byte) bool {
	len1 := len(s1)
	if len1 != len(s2) {
		return false
	}

	for i := 0; i < len1; i++ {
		if ToLower(s1[i]) != ToLower(s2[i]) {
			return false
		}
	}

	return true
}

func EqualNoCase(s1, s2 []byte) bool {
	len1 := len(s1)
	if len1 != len(s2) {
		return false
	}

	p1 := uintptr(unsafe.Pointer(&s1[0]))
	p2 := uintptr(unsafe.Pointer(&s2[0]))
	end := p1 + uintptr(len1)
	end1 := p1 + uintptr((len1>>3)<<3)

	for p1 < end1 {
		if *((*int64)(unsafe.Pointer(p1))) != *((*int64)(unsafe.Pointer(p2))) {
			break
		}
		p1 += 8
		p2 += 8
	}

	for p1 < end {
		if *((*byte)(unsafe.Pointer(p1))) != *((*byte)(unsafe.Pointer(p2))) {
			break
		}
		p1++
		p2++
	}
	for p1 < end {
		if ToLower(*((*byte)(unsafe.Pointer(p1)))) != ToLower(*((*byte)(unsafe.Pointer(p2)))) {
			return false
		}
		p1++
		p2++
	}

	return true
}
