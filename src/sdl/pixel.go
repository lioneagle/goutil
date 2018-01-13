package sdl

/** Pixel type. */
const (
	SDL_PIXELTYPE_UNKNOWN  = 0
	SDL_PIXELTYPE_INDEX1   = 1
	SDL_PIXELTYPE_INDEX4   = 2
	SDL_PIXELTYPE_INDEX8   = 3
	SDL_PIXELTYPE_PACKED8  = 4
	SDL_PIXELTYPE_PACKED16 = 5
	SDL_PIXELTYPE_PACKED32 = 6
	SDL_PIXELTYPE_ARRAYU8  = 7
	SDL_PIXELTYPE_ARRAYU16 = 8
	SDL_PIXELTYPE_ARRAYU32 = 9
	SDL_PIXELTYPE_ARRAYF16 = 10
	SDL_PIXELTYPE_ARRAYF32 = 11
)

/** Bitmap pixel order, high bit -> low bit. */
const (
	SDL_BITMAPORDER_NONE = 0
	SDL_BITMAPORDER_4321 = 1
	SDL_BITMAPORDER_1234 = 2
)

/** Packed component order, high bit -> low bit. */
const (
	SDL_PACKEDORDER_NONE = 0
	SDL_PACKEDORDER_XRGB = 1
	SDL_PACKEDORDER_RGBX = 2
	SDL_PACKEDORDER_ARGB = 3
	SDL_PACKEDORDER_RGBA = 4
	SDL_PACKEDORDER_XBGR = 5
	SDL_PACKEDORDER_BGRX = 6
	SDL_PACKEDORDER_ABGR = 7
	SDL_PACKEDORDER_BGRA = 8
)

/** Array component order, low byte -> high byte. */
/* !!! FIXME: in 2.1, make these not overlap differently with
   !!! FIXME:  SDL_PACKEDORDER_*, so we can simplify SDL_ISPIXELFORMAT_ALPHA */
const (
	SDL_ARRAYORDER_NONE = 0
	SDL_ARRAYORDER_RGB  = 1
	SDL_ARRAYORDER_RGBA = 2
	SDL_ARRAYORDER_ARGB = 3
	SDL_ARRAYORDER_BGR  = 4
	SDL_ARRAYORDER_BGRA = 5
	SDL_ARRAYORDER_ABGR = 6
)

/** Packed component layout. */
const (
	SDL_PACKEDLAYOUT_NONE    = 0
	SDL_PACKEDLAYOUT_332     = 1
	SDL_PACKEDLAYOUT_4444    = 2
	SDL_PACKEDLAYOUT_1555    = 3
	SDL_PACKEDLAYOUT_5551    = 4
	SDL_PACKEDLAYOUT_565     = 5
	SDL_PACKEDLAYOUT_8888    = 6
	SDL_PACKEDLAYOUT_2101010 = 7
	SDL_PACKEDLAYOUT_1010102 = 8
)

type SDL_Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type SDL_Palette struct {
	Ncolors  int32
	Colors   *SDL_Color
	Version  uint32
	Refcount int32
}

/**
 *  \note Everything in the pixel format structure is read-only.
 */
type SDL_PixelFormat struct {
	Format        uint32
	Palette       *SDL_Palette
	BitsPerPixel  uint8
	BytesPerPixel uint8
	padding       [2]uint8
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	Refcount      int
	Next          *SDL_PixelFormat
}
