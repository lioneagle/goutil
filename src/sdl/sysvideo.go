package sdl

type SDL_Window struct {
	Magic                 *byte
	Id                    uint32
	Title                 *byte
	Icon                  *SDL_Surface
	X                     int32
	Y                     int32
	W                     int32
	H                     int32
	Min_w                 int32
	Min_h                 int32
	Max_w                 int32
	Max_h                 int32
	Flags                 uint32
	Last_fullscreen_flags uint32

	/* Stored position and size for windowed mode */
	Windowed SDL_Rect

	Fullscreen_mode SDL_DisplayMode

	Opacity float32

	Brightness  float32
	Gamma       *uint16
	Saved_gamma *uint16 /* (just offset into gamma) */

	Surface       *SDL_Surface
	Surface_valid *SDL_Surface

	Is_hiding     int32 // SDL_bool
	Is_destroying int32
	Is_dropping   int32

	Shaper *SDL_WindowShaper

	Hit_test      uintptr //SDL_HitTest is a function ptr
	Hit_test_data *byte

	Data       *SDL_WindowUserData
	Driverdata *byte

	Prev *SDL_Window
	Next *SDL_Window
}

type SDL_WindowShaper struct {
	Window     *SDL_Window /* The window associated with the shaper */
	Userx      uint32      /* The user's specified coordinates for the window, for once we give it a shape. */
	Usery      uint32
	Mode       SDL_WindowShapeMode /* The parameters for shape calculation. */
	hasShape   int32
	Driverdata *byte
}

type SDL_WindowUserData struct {
	Name *byte
	Data *byte
	Next *SDL_WindowUserData
}
