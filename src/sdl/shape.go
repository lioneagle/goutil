package sdl

type WindowShapeMode int32

type SDL_WindowShapeMode struct {
	Mode       WindowShapeMode
	Parameters SDL_WindowShapeParams
}

type SDL_WindowShapeParams struct {
	SDL_Color
}
