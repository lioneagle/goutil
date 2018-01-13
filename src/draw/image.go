package core

type Image interface {
	BeginPaint(canvas Canvas) error
	EndPaint()
	Draw(canvas Canvas) error
	Dispose()
	Size() Size
	SaveToFile(filename, format string) error
}
