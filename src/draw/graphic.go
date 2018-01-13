package core

type Graphic interface {
	Draw(canvas Canvas) error
	Add(child Graphic) error
	Remove(child Graphic) error
	Parent() Graphic
	Level() int
}
