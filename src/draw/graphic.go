package draw

type IGraphic interface {
	Draw(canvas ICanvas) error
	Add(child IGraphic) error
	Remove(child IGraphic) error
	Parent() IGraphic
	Level() int
}
