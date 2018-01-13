package sdl

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	// Library
	libsdl32 = syscall.NewLazyDLL("sdl2_64.dll")
	//libmsimg32 = syscall.NewLazyDLL("msimg32.dll")

	// Functions
	procSDL_Init          = libsdl32.NewProc("SDL_Init")
	procSDL_Quit          = libsdl32.NewProc("SDL_Quit")
	procSDL_WasInit       = libsdl32.NewProc("SDL_WasInit")
	procSDL_InitSubSystem = libsdl32.NewProc("SDL_InitSubSystem")
	procSDL_QuitSubSystem = libsdl32.NewProc("SDL_QuitSubSystem")
	procSDL_GetError      = libsdl32.NewProc("SDL_GetError")

	procSDL_CreateWindow     = libsdl32.NewProc("SDL_CreateWindow")
	procSDL_DestroyWindow    = libsdl32.NewProc("SDL_DestroyWindow")
	procSDL_PollEvent        = libsdl32.NewProc("SDL_PollEvent")
	procSDL_CreateRGBSurface = libsdl32.NewProc("SDL_CreateRGBSurface")
	procSDL_FreeSurface      = libsdl32.NewProc("SDL_FreeSurface")
	procSDL_FillRect         = libsdl32.NewProc("SDL_FillRect")

	procSDL_CreateRenderer         = libsdl32.NewProc("SDL_CreateRenderer")
	procSDL_CreateSoftwareRenderer = libsdl32.NewProc("SDL_CreateSoftwareRenderer")
	procSDL_DestroyRenderer        = libsdl32.NewProc("SDL_DestroyRenderer")
	procSDL_RenderDrawLine         = libsdl32.NewProc("SDL_RenderDrawLine")
	procSDL_RenderDrawRect         = libsdl32.NewProc("SDL_RenderDrawRect")
	procSDL_RenderFillRect         = libsdl32.NewProc("SDL_RenderFillRect")
)

func SDL_Init(flags uint32) error {
	ret, _, _ := procSDL_Init.Call(
		uintptr(flags))

	if int(ret) < 0 {
		return errors.New(fmt.Sprintf("SDL_Init failed: %s", SDL_GetError()))
	}

	return nil
}

func SDL_Quit() {
	procSDL_Quit.Call()
}

func SDL_WasInit(flags uint32) error {
	ret, _, _ := procSDL_WasInit.Call(
		uintptr(flags))

	if ret == 0 {
		return errors.New(fmt.Sprintf("SDL_WasInit failed: %s", SDL_GetError()))
	}

	return nil
}

func SDL_InitSubSystem(flags uint32) error {
	ret, _, _ := procSDL_InitSubSystem.Call(uintptr(flags))

	if int(ret) < 0 {
		return errors.New(fmt.Sprintf("SDL_InitSubSystem failed: %s", SDL_GetError()))
	}

	return nil
}

func SDL_QuitSubSystem(flags uint32) {
	procSDL_QuitSubSystem.Call(uintptr(flags))
}

func SDL_GetError() string {
	ret, _, _ := procSDL_GetError.Call()
	return CharPtrToString((*byte)(unsafe.Pointer(ret)))
}

func SDL_CreateWindow(title string, x, y, w, h int32, flags uint32) (*SDL_Window, error) {
	ret, _, _ := procSDL_CreateWindow.Call(
		uintptr(unsafe.Pointer(StringToBytePtr(title))),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h),
		uintptr(flags),
	)

	if int(ret) == 0 {
		return nil, errors.New(fmt.Sprintf("SDL_CreateWindow failed: %s", SDL_GetError()))
	}

	return (*SDL_Window)(unsafe.Pointer(ret)), nil
}

func SDL_DestroyWindow(window *SDL_Window) {
	procSDL_DestroyWindow.Call(uintptr(unsafe.Pointer(window)))
}

func SDL_FillRect(dst *SDL_Surface, rect *SDL_Rect, color uint32) error {
	ret, _, _ := procSDL_FillRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(color),
	)

	if int(ret) < 0 {
		return errors.New(fmt.Sprintf("SDL_CreateWindow failed: %s", SDL_GetError()))
	}

	return nil
}

func SDL_PollEvent(event *SDL_Event) bool {
	ret, _, _ := procSDL_PollEvent.Call(
		uintptr(unsafe.Pointer(event)),
	)

	if int(ret) == 0 {
		return false
	}

	return true
}

func SDL_CreateRGBSurface(flags uint32, width, height, depth int32, Rmask, Gmask, Bmask, Amask uint32) (*SDL_Surface, error) {
	ret, _, _ := procSDL_CreateRGBSurface.Call(
		uintptr(flags),
		uintptr(width),
		uintptr(height),
		uintptr(depth),
		uintptr(Rmask),
		uintptr(Gmask),
		uintptr(Bmask),
		uintptr(Amask),
	)

	if int(ret) < 0 {
		return nil, errors.New(fmt.Sprintf("SDL_CreateRGBSurface failed: %s", SDL_GetError()))
	}

	return (*SDL_Surface)(unsafe.Pointer(ret)), nil
}

func SDL_FreeSurface(surface *SDL_Surface) {
	procSDL_FreeSurface.Call(uintptr(unsafe.Pointer(surface)))
}

func SDL_CreateRenderer(window *SDL_Window, index int32, flags uint32) (*SDL_Renderer, error) {
	ret, _, _ := procSDL_CreateRenderer.Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(index),
		uintptr(flags),
	)

	if int(ret) == 0 {
		return nil, errors.New(fmt.Sprintf("SDL_CreateRenderer failed: %s", SDL_GetError()))
	}

	return (*SDL_Renderer)(unsafe.Pointer(ret)), nil
}

func SDL_CreateSoftwareRenderer(surface *SDL_Surface) (*SDL_Renderer, error) {
	ret, _, _ := procSDL_CreateSoftwareRenderer.Call(
		uintptr(unsafe.Pointer(surface)),
	)

	if int(ret) == 0 {
		return nil, errors.New(fmt.Sprintf("SDL_CreateSoftwareRenderer failed: %s", SDL_GetError()))
	}

	return (*SDL_Renderer)(unsafe.Pointer(ret)), nil
}

func SDL_DestroyRenderer(renderer *SDL_Renderer) {
	procSDL_DestroyRenderer.Call(uintptr(unsafe.Pointer(renderer)))
}

func SDL_RenderDrawLine(renderer *SDL_Renderer, x1, y1, x2, y2 int32) error {
	ret, _, _ := procSDL_RenderDrawLine.Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
	)

	if int(ret) < 0 {
		return errors.New(fmt.Sprintf("SDL_RenderDrawLine failed: %s", SDL_GetError()))
	}

	return nil
}

func SDL_RenderDrawRect(renderer *SDL_Renderer, rect *SDL_Rect) error {
	ret, _, _ := procSDL_RenderDrawRect.Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
	)

	if int(ret) < 0 {
		return errors.New(fmt.Sprintf("SDL_RenderDrawRect failed: %s", SDL_GetError()))
	}

	return nil
}

func SDL_RenderFillRect(renderer *SDL_Renderer, rect *SDL_Rect) error {
	ret, _, _ := procSDL_RenderFillRect.Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
	)

	if int(ret) < 0 {
		return errors.New(fmt.Sprintf("SDL_RenderFillRect failed: %s", SDL_GetError()))
	}

	return nil
}
