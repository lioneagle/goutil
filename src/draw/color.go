package core

import (
	"fmt"
)

type Color uint32

func RGB(red, green, blue byte) Color {
	return Color(uint32(red) | uint32(green)<<8 | uint32(blue)<<16)
}

func ARGB(alpha, red, green, blue byte) Color {
	return Color(uint32(red) | uint32(green)<<8 | uint32(blue)<<16 | uint32(alpha)<<24)
}

func (c Color) RGBString() string {
	return fmt.Sprintf("#%02x%02x%02x", c.Red(), c.Green(), c.Blue())
}

func (c Color) ARGBString() string {
	return fmt.Sprintf("#%02x%02x%02x%02x", c.Alpha(), c.Red(), c.Green(), c.Blue())
}

func (c Color) Red() byte {
	return byte(c & 0xff)
}

func (c Color) Green() byte {
	return byte((c >> 8) & 0xff)
}

func (c Color) Blue() byte {
	return byte((c >> 16) & 0xff)
}

func (c Color) Alpha() byte {
	return byte(c >> 24)
}

type ColorPool struct {
	data map[string]Color
}

func NewColorPool() *ColorPool {
	return &ColorPool{}
}

func (this *ColorPool) AddColor(name string, color Color) {
	this.data[name] = color
}

func (this *ColorPool) GetColor(name string) (color Color, ok bool) {
	color, ok = this.data[name]
	return color, ok
}

var (
	ColorBlack Color = RGB(0, 0, 0)
	ColorWhite Color = RGB(255, 255, 255)
	ColorRed   Color = RGB(255, 0, 0)
	ColorGreen Color = RGB(0, 255, 0)
	ColorBlue  Color = RGB(0, 0, 255)
	ColorAqua  Color = RGB(0, 255, 255)
	ColorGold  Color = RGB(255, 201, 14)
)
