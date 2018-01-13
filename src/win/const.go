package win

const (
	CCHDEVICENAME = 32
	CCHFORMNAME   = 32
)

// kernel32

// Error codes
const (
	ERROR_SUCCESS             = 0
	ERROR_INVALID_FUNCTION    = 1
	ERROR_FILE_NOT_FOUND      = 2
	ERROR_INVALID_PARAMETER   = 87
	ERROR_INSUFFICIENT_BUFFER = 122
	ERROR_MORE_DATA           = 234
)

// gdi

// Pen types
const (
	PS_COSMETIC  = 0x00000000
	PS_GEOMETRIC = 0x00010000
	PS_TYPE_MASK = 0x000F0000
)

// Pen styles
const (
	PS_SOLID       = 0
	PS_DASH        = 1
	PS_DOT         = 2
	PS_DASHDOT     = 3
	PS_DASHDOTDOT  = 4
	PS_NULL        = 5
	PS_INSIDEFRAME = 6
	PS_USERSTYLE   = 7
	PS_ALTERNATE   = 8
	PS_STYLE_MASK  = 0x0000000F
)

// Pen cap types
const (
	PS_ENDCAP_ROUND  = 0x00000000
	PS_ENDCAP_SQUARE = 0x00000100
	PS_ENDCAP_FLAT   = 0x00000200
	PS_ENDCAP_MASK   = 0x00000F00
)

// Pen join types
const (
	PS_JOIN_ROUND = 0x00000000
	PS_JOIN_BEVEL = 0x00001000
	PS_JOIN_MITER = 0x00002000
	PS_JOIN_MASK  = 0x0000F000
)

// Hatch styles
const (
	HS_HORIZONTAL = 0
	HS_VERTICAL   = 1
	HS_FDIAGONAL  = 2
	HS_BDIAGONAL  = 3
	HS_CROSS      = 4
	HS_DIAGCROSS  = 5
)

// Stock Logical Objects
const (
	WHITE_BRUSH         = 0
	LTGRAY_BRUSH        = 1
	GRAY_BRUSH          = 2
	DKGRAY_BRUSH        = 3
	BLACK_BRUSH         = 4
	NULL_BRUSH          = 5
	HOLLOW_BRUSH        = NULL_BRUSH
	WHITE_PEN           = 6
	BLACK_PEN           = 7
	NULL_PEN            = 8
	OEM_FIXED_FONT      = 10
	ANSI_FIXED_FONT     = 11
	ANSI_VAR_FONT       = 12
	SYSTEM_FONT         = 13
	DEVICE_DEFAULT_FONT = 14
	DEFAULT_PALETTE     = 15
	SYSTEM_FIXED_FONT   = 16
	DEFAULT_GUI_FONT    = 17
	DC_BRUSH            = 18
	DC_PEN              = 19
)

// Brush styles
const (
	BS_SOLID         = 0
	BS_NULL          = 1
	BS_HOLLOW        = BS_NULL
	BS_HATCHED       = 2
	BS_PATTERN       = 3
	BS_INDEXED       = 4
	BS_DIBPATTERN    = 5
	BS_DIBPATTERNPT  = 6
	BS_PATTERN8X8    = 7
	BS_DIBPATTERN8X8 = 8
	BS_MONOPATTERN   = 9
)

// GetDeviceCaps index constants
const (
	DRIVERVERSION   = 0
	TECHNOLOGY      = 2
	HORZSIZE        = 4
	VERTSIZE        = 6
	HORZRES         = 8
	VERTRES         = 10
	LOGPIXELSX      = 88
	LOGPIXELSY      = 90
	BITSPIXEL       = 12
	PLANES          = 14
	NUMBRUSHES      = 16
	NUMPENS         = 18
	NUMFONTS        = 22
	NUMCOLORS       = 24
	NUMMARKERS      = 20
	ASPECTX         = 40
	ASPECTY         = 42
	ASPECTXY        = 44
	PDEVICESIZE     = 26
	CLIPCAPS        = 36
	SIZEPALETTE     = 104
	NUMRESERVED     = 106
	COLORRES        = 108
	PHYSICALWIDTH   = 110
	PHYSICALHEIGHT  = 111
	PHYSICALOFFSETX = 112
	PHYSICALOFFSETY = 113
	SCALINGFACTORX  = 114
	SCALINGFACTORY  = 115
	VREFRESH        = 116
	DESKTOPHORZRES  = 118
	DESKTOPVERTRES  = 117
	BLTALIGNMENT    = 119
	SHADEBLENDCAPS  = 120
	COLORMGMTCAPS   = 121
	RASTERCAPS      = 38
	CURVECAPS       = 28
	LINECAPS        = 30
	POLYGONALCAPS   = 32
	TEXTCAPS        = 34
)

// GetDeviceCaps TECHNOLOGY constants
const (
	DT_PLOTTER    = 0
	DT_RASDISPLAY = 1
	DT_RASPRINTER = 2
	DT_RASCAMERA  = 3
	DT_CHARSTREAM = 4
	DT_METAFILE   = 5
	DT_DISPFILE   = 6
)

// GetDeviceCaps SHADEBLENDCAPS constants
const (
	SB_NONE          = 0x00
	SB_CONST_ALPHA   = 0x01
	SB_PIXEL_ALPHA   = 0x02
	SB_PREMULT_ALPHA = 0x04
	SB_GRAD_RECT     = 0x10
	SB_GRAD_TRI      = 0x20
)

// GetDeviceCaps COLORMGMTCAPS constants
const (
	CM_NONE       = 0x00
	CM_DEVICE_ICM = 0x01
	CM_GAMMA_RAMP = 0x02
	CM_CMYK_COLOR = 0x04
)

// GetDeviceCaps RASTERCAPS constants
const (
	RC_BANDING      = 2
	RC_BITBLT       = 1
	RC_BITMAP64     = 8
	RC_DI_BITMAP    = 128
	RC_DIBTODEV     = 512
	RC_FLOODFILL    = 4096
	RC_GDI20_OUTPUT = 16
	RC_PALETTE      = 256
	RC_SCALING      = 4
	RC_STRETCHBLT   = 2048
	RC_STRETCHDIB   = 8192
	RC_DEVBITS      = 0x8000
	RC_OP_DX_OUTPUT = 0x4000
)

// GetDeviceCaps CURVECAPS constants
const (
	CC_NONE       = 0
	CC_CIRCLES    = 1
	CC_PIE        = 2
	CC_CHORD      = 4
	CC_ELLIPSES   = 8
	CC_WIDE       = 16
	CC_STYLED     = 32
	CC_WIDESTYLED = 64
	CC_INTERIORS  = 128
	CC_ROUNDRECT  = 256
)

// GetDeviceCaps LINECAPS constants
const (
	LC_NONE       = 0
	LC_POLYLINE   = 2
	LC_MARKER     = 4
	LC_POLYMARKER = 8
	LC_WIDE       = 16
	LC_STYLED     = 32
	LC_WIDESTYLED = 64
	LC_INTERIORS  = 128
)

// GetDeviceCaps POLYGONALCAPS constants
const (
	PC_NONE        = 0
	PC_POLYGON     = 1
	PC_POLYPOLYGON = 256
	PC_PATHS       = 512
	PC_RECTANGLE   = 2
	PC_WINDPOLYGON = 4
	PC_SCANLINE    = 8
	PC_TRAPEZOID   = 4
	PC_WIDE        = 16
	PC_STYLED      = 32
	PC_WIDESTYLED  = 64
	PC_INTERIORS   = 128
)

// GetDeviceCaps TEXTCAPS constants
const (
	TC_OP_CHARACTER = 1
	TC_OP_STROKE    = 2
	TC_CP_STROKE    = 4
	TC_CR_90        = 8
	TC_CR_ANY       = 16
	TC_SF_X_YINDEP  = 32
	TC_SA_DOUBLE    = 64
	TC_SA_INTEGER   = 128
	TC_SA_CONTIN    = 256
	TC_EA_DOUBLE    = 512
	TC_IA_ABLE      = 1024
	TC_UA_ABLE      = 2048
	TC_SO_ABLE      = 4096
	TC_RA_ABLE      = 8192
	TC_VA_ABLE      = 16384
	TC_RESERVED     = 32768
	TC_SCROLLBLT    = 65536
)

// Tooltip notifications
const (
	TTN_FIRST       = -520
	TTN_LAST        = -549
	TTN_GETDISPINFO = (TTN_FIRST - 10)
	TTN_SHOW        = (TTN_FIRST - 1)
	TTN_POP         = (TTN_FIRST - 2)
	TTN_LINKCLICK   = (TTN_FIRST - 3)
	TTN_NEEDTEXT    = TTN_GETDISPINFO
)

const (
	TTF_IDISHWND    = 0x0001
	TTF_CENTERTIP   = 0x0002
	TTF_RTLREADING  = 0x0004
	TTF_SUBCLASS    = 0x0010
	TTF_TRACK       = 0x0020
	TTF_ABSOLUTE    = 0x0080
	TTF_TRANSPARENT = 0x0100
	TTF_PARSELINKS  = 0x1000
	TTF_DI_SETITEM  = 0x8000
)

const (
	SWP_NOSIZE         = 0x0001
	SWP_NOMOVE         = 0x0002
	SWP_NOZORDER       = 0x0004
	SWP_NOREDRAW       = 0x0008
	SWP_NOACTIVATE     = 0x0010
	SWP_FRAMECHANGED   = 0x0020
	SWP_SHOWWINDOW     = 0x0040
	SWP_HIDEWINDOW     = 0x0080
	SWP_NOCOPYBITS     = 0x0100
	SWP_NOOWNERZORDER  = 0x0200
	SWP_NOSENDCHANGING = 0x0400
	SWP_DRAWFRAME      = SWP_FRAMECHANGED
	SWP_NOREPOSITION   = SWP_NOOWNERZORDER
	SWP_DEFERERASE     = 0x2000
	SWP_ASYNCWINDOWPOS = 0x4000
)

// Background Modes
const (
	TRANSPARENT = 1
	OPAQUE      = 2
	BKMODE_LAST = 2
)

// StretchBlt modes
const (
	BLACKONWHITE        = 1
	WHITEONBLACK        = 2
	COLORONCOLOR        = 3
	HALFTONE            = 4
	MAXSTRETCHBLTMODE   = 4
	STRETCH_ANDSCANS    = BLACKONWHITE
	STRETCH_ORSCANS     = WHITEONBLACK
	STRETCH_DELETESCANS = COLORONCOLOR
	STRETCH_HALFTONE    = HALFTONE
)

// gdiplus
const (
	Ok                        GpStatus = 0
	GenericError              GpStatus = 1
	InvalidParameter          GpStatus = 2
	OutOfMemory               GpStatus = 3
	ObjectBusy                GpStatus = 4
	InsufficientBuffer        GpStatus = 5
	NotImplemented            GpStatus = 6
	Win32Error                GpStatus = 7
	WrongState                GpStatus = 8
	Aborted                   GpStatus = 9
	FileNotFound              GpStatus = 10
	ValueOverflow             GpStatus = 11
	AccessDenied              GpStatus = 12
	UnknownImageFormat        GpStatus = 13
	FontFamilyNotFound        GpStatus = 14
	FontStyleNotFound         GpStatus = 15
	NotTrueTypeFont           GpStatus = 16
	UnsupportedGdiplusVersion GpStatus = 17
	GdiplusNotInitialized     GpStatus = 18
	PropertyNotFound          GpStatus = 19
	PropertyNotSupported      GpStatus = 20
	ProfileNotFound           GpStatus = 21
)

const (
	BrushTypeSolidColor     GpBrushType = 0
	BrushTypeHatchFill      GpBrushType = 1
	BrushTypeTextureFill    GpBrushType = 2
	BrushTypePathGradient   GpBrushType = 3
	BrushTypeLinearGradient GpBrushType = 4
)

const (
	WrapModeTile       GpWrapMode = 0
	WrapModeTileFlipX  GpWrapMode = 1
	WrapModeTileFlipY  GpWrapMode = 2
	WrapModeTileFlipXY GpWrapMode = 3
	WrapModeClamp      GpWrapMode = 4
)

const (
	TextRenderingHintSystemDefault            TextRenderingHint = 0
	TextRenderingHintSingleBitPerPixelGridFit TextRenderingHint = 1
	TextRenderingHintSingleBitPerPixel        TextRenderingHint = 2
	TextRenderingHintAntiAliasGridFit         TextRenderingHint = 3
	TextRenderingHintAntiAlias                TextRenderingHint = 4
	TextRenderingHintClearTypeGridFit         TextRenderingHint = 5
)

const (
	QualityModeInvalid QualityMode = -1
	QualityModeDefault QualityMode = 0
	QualityModeLow     QualityMode = 1
	QualityModeHigh    QualityMode = 2
)

const (
	SmoothingModeInvalid      SmoothingMode = -1 //SmoothingMode(QualityModeInvalid)
	SmoothingModeDefault      SmoothingMode = 0  // SmoothingMode(QualityModeDefault)
	SmoothingModeHighSpeed    SmoothingMode = 1  // SmoothingMode(QualityModeLow)
	SmoothingModeHighQuality  SmoothingMode = 2  // SmoothingMode(QualityModeHigh)
	SmoothingModeNone         SmoothingMode = 3  // SmoothingMode(QualityModeHigh + 1)
	SmoothingModeAntiAlias8x4 SmoothingMode = 4  // SmoothingMode(QualityModeHigh + 2)
	SmoothingModeAntiAlias    SmoothingMode = 4  // SmoothingModeAntiAlias8x4
	SmoothingModeAntiAlias8x8 SmoothingMode = 5  // SmoothingModeAntiAlias + 1
)

const (
	StringAlignmentNear   StringAlignment = 0
	StringAlignmentCenter StringAlignment = 1
	StringAlignmentFar    StringAlignment = 2
)

const (
	StringFormatFlagsDirectionRightToLeft  StringFormatFlags = 0x00000001
	StringFormatFlagsDirectionVertical     StringFormatFlags = 0x00000002
	StringFormatFlagsNoFitBlackBox         StringFormatFlags = 0x00000004
	StringFormatFlagsDisplayFormatControl  StringFormatFlags = 0x00000020
	StringFormatFlagsNoFontFallback        StringFormatFlags = 0x00000400
	StringFormatFlagsMeasureTrailingSpaces StringFormatFlags = 0x00000800
	StringFormatFlagsNoWrap                StringFormatFlags = 0x00001000
	StringFormatFlagsLineLimit             StringFormatFlags = 0x00002000
	StringFormatFlagsNoClip                StringFormatFlags = 0x00004000
)

const (
	LANG_NEUTRAL LANGID = 0
)

const (
	UnitWorld      GpUnit = 0
	UnitDisplay    GpUnit = 1
	UnitPixel      GpUnit = 2
	UnitPoint      GpUnit = 3
	UnitInch       GpUnit = 4
	UnitDocument   GpUnit = 5
	UnitMillimeter GpUnit = 6
)
