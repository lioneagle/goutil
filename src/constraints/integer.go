package constraints

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type Unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr
}
