package core

type Canvas interface {
	DrawLine(pen Pen, from, to Point) error
	DrawRectangle(pen Pen, rect Rectangle) error
	FillRectangle(brush Brush, rect Rectangle) error
	DrawEllipse(pen Pen, rect Rectangle) error
	FillEllipse(brush Brush, rect Rectangle) error
	DrawCircle(pen Pen, center Point, radius int) error
	FillCircle(brush Brush, center Point, radius int) error
	//DrawImage(image Image, location Point) error
	//PaintImage(image Image, f func() error) error
	Dispose()
}
