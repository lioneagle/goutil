package chars

var g_tolower_table [256]byte
var g_toupper_table [256]byte
var g_return Return

func init() {
	for i := 0; i < 256; i++ {
		g_tolower_table[i] = toLower(byte(i))
		g_toupper_table[i] = toUpper(byte(i))
	}

	g_return.Init()

}

func toLower(ch byte) byte {
	//if IsUpper(ch) {
	if (Charsets0[ch] & MASK_UPPER) != 0 {
		//return ch - 'A' + 'a'
		return ch | 0x20
	}
	return ch
}

func toUpper(ch byte) byte {
	//if IsLower(ch) {
	if (Charsets0[ch] & MASK_LOWER) != 0 {
		//return ch - 'a' + 'A'
		return ch & 0xDF
	}
	return ch
}
