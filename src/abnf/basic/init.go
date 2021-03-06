package abnf

var g_smallDigitsString []byte
var g_digitsString []byte

func init() {
	initDigits()
}

func initDigits() {
	g_digitsString = []byte("0123456789abcdefghijklmnopqrstuvwxyz")

	g_smallDigitsString = []byte("00010203040506070809" +
		"10111213141516171819" +
		"20212223242526272829" +
		"30313233343536373839" +
		"40414243444546474849" +
		"50515253545556575859" +
		"60616263646566676869" +
		"70717273747576777879" +
		"80818283848586878889" +
		"90919293949596979899")
}
