package sdl

type SDL_TouchID int64
type SDL_FingerID int64

type SDL_Finger struct {
	Id       SDL_FingerID
	X        float32
	Y        float32
	Pressure float32
}
