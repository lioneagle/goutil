package draw

type IImage interface {
	BeginPaint(canvas ICanvas) error
	EndPaint()
	Draw(canvas ICanvas) error
	Dispose()
	Size() Size
	SaveToFile(filename, format string) error
}
