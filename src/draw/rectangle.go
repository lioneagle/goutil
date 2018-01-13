package core

type RectangleGraphic struct {
	rect   Rectangle
	Style  RectangStyle
	childs []Graphic
	parent Graphic
	text   *Text
}

type RectangStyle struct {
	Level       int
	DrawBorder  bool
	DoFill      bool
	FillColor   Color
	BorderColor Color
	BorderWidth int
}

func (this *RectangleGraphic) Draw(canvas Canvas) error {
	return nil
}

func (this *RectangleGraphic) Add(child Graphic) error {
	return nil
}

func (this *RectangleGraphic) Remove(child Graphic) error {
	return nil
}

func (this *RectangleGraphic) Parent() Graphic {
	return this.parent
}

func (this *RectangleGraphic) Level(child Graphic) int {
	return this.Style.Level
}
