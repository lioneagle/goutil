package sdl

type SDL_BlitMap struct {
	Dst      *SDL_Surface
	Identity int32
	Blit     uintptr //SDL_Blit is a function ptr
	Data     *byte
	Info     SDL_BlitInfo
}

type SDL_BlitInfo struct {
	Src       *uint8
	Src_w     int32
	Src_h     int32
	Src_pitch int32
	Src_skip  int32
	Dst       *uint8
	Dst_w     int32
	Dst_h     int32
	Dst_pitch int32
	Dst_skip  int32
	Src_fmt   *SDL_PixelFormat
	Dst_fmt   *SDL_PixelFormat
	Table     *uint8
	Flags     int32
	Colorkey  uint32
	R         uint8
	G         uint8
	B         uint8
	A         uint8
}
