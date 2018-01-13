package sdl

type SDL_JoystickID uint32
type SDL_JoystickPowerLevel int32

const (
	SDL_JOYSTICK_POWER_UNKNOWN SDL_JoystickPowerLevel = -1
	SDL_JOYSTICK_POWER_EMPTY   SDL_JoystickPowerLevel = 0
	SDL_JOYSTICK_POWER_LOW     SDL_JoystickPowerLevel = 1
	SDL_JOYSTICK_POWER_MEDIUM  SDL_JoystickPowerLevel = 2
	SDL_JOYSTICK_POWER_FULL    SDL_JoystickPowerLevel = 3
	SDL_JOYSTICK_POWER_WIRED   SDL_JoystickPowerLevel = 4
	SDL_JOYSTICK_POWER_MAX     SDL_JoystickPowerLevel = 5
)
