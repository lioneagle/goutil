package draw

// Pen types
const (
	PEN_TYPE_COSMETIC  = 0
	PEN_TYPE_GEOMETRIC = 1
)

// Pen styles
const (
	PEN_SOLID        = 0
	PEN_DASH         = 1
	PEN_DOT          = 2
	PEN_DASH_DOT     = 3
	PEN_DASH_DOT_DOT = 4
	PEN_NULL         = 5
)

// Pen cap styles
const (
	PEN_CAP_ROUND  = 0
	PEN_CAP_SQUARE = 1
	PEN_CAP_FLAT   = 2
)

// Pen join styles
const (
	PEN_JOIN_ROUND = 0
	PEN_JOIN_BEVEL = 1
	PEN_JOIN_MITER = 2
)

type PenStyle struct {
	Type      byte
	Style     byte
	CapStyle  byte
	JoinStyle byte
}

type Pen interface {
	Dispose()
	Width() int
	Style() PenStyle
}
