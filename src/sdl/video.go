package sdl

type SDL_DisplayMode struct {
	Format       uint32
	W            int32
	H            int32
	Refresh_rate int32
	Driverdata   *byte
}

/**
 *  \brief The flags on a window
 *
 *  \sa SDL_GetWindowFlags()
 */
const (
	/* !!! FIXME: change this to name = (1<<x). */
	SDL_WINDOW_FULLSCREEN         = 0x00000001 /**< fullscreen window */
	SDL_WINDOW_OPENGL             = 0x00000002 /**< window usable with OpenGL context */
	SDL_WINDOW_SHOWN              = 0x00000004 /**< window is visible */
	SDL_WINDOW_HIDDEN             = 0x00000008 /**< window is not visible */
	SDL_WINDOW_BORDERLESS         = 0x00000010 /**< no window decoration */
	SDL_WINDOW_RESIZABLE          = 0x00000020 /**< window can be resized */
	SDL_WINDOW_MINIMIZED          = 0x00000040 /**< window is minimized */
	SDL_WINDOW_MAXIMIZED          = 0x00000080 /**< window is maximized */
	SDL_WINDOW_INPUT_GRABBED      = 0x00000100 /**< window has grabbed input focus */
	SDL_WINDOW_INPUT_FOCUS        = 0x00000200 /**< window has input focus */
	SDL_WINDOW_MOUSE_FOCUS        = 0x00000400 /**< window has mouse focus */
	SDL_WINDOW_FULLSCREEN_DESKTOP = (SDL_WINDOW_FULLSCREEN | 0x00001000)
	SDL_WINDOW_FOREIGN            = 0x00000800 /**< window not created by SDL */
	SDL_WINDOW_ALLOW_HIGHDPI      = 0x00002000 /**< window should be created in high-DPI mode if supported */
	SDL_WINDOW_MOUSE_CAPTURE      = 0x00004000 /**< window has mouse captured (unrelated to INPUT_GRABBED) */
	SDL_WINDOW_ALWAYS_ON_TOP      = 0x00008000 /**< window should always be above others */
	SDL_WINDOW_SKIP_TASKBAR       = 0x00010000 /**< window should not be added to the taskbar */
	SDL_WINDOW_UTILITY            = 0x00020000 /**< window should be treated as a utility window */
	SDL_WINDOW_TOOLTIP            = 0x00040000 /**< window should be treated as a tooltip */
	SDL_WINDOW_POPUP_MENU         = 0x00080000 /**< window should be treated as a popup menu */
)

const (
	SDL_WINDOWPOS_UNDEFINED_MASK = 0x1FFF0000
	SDL_WINDOWPOS_UNDEFINED      = SDL_WINDOWPOS_UNDEFINED_MASK
	SDL_WINDOWPOS_CENTERED_MASK  = 0x2FFF0000
	SDL_WINDOWPOS_CENTERED       = SDL_WINDOWPOS_CENTERED_MASK
)
