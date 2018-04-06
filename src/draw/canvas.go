package draw

type ICanvas interface {
	DrawLine(pen IPen, from, to Point) error
	DrawRectangle(pen IPen, rect Rectangle) error
	FillRectangle(brush IBrush, rect Rectangle) error
	DrawEllipse(pen IPen, rect Rectangle) error
	FillEllipse(brush IBrush, rect Rectangle) error
	DrawCircle(pen IPen, center Point, radius int) error
	FillCircle(brush IBrush, center Point, radius int) error
	//DrawImage(image Image, location Point) error
	//PaintImage(image Image, f func() error) error
	Dispose()
}
