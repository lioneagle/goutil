package draw

type RectangleGraphic struct {
	rect   Rectangle
	Style  RectangStyle
	childs []IGraphic
	parent IGraphic
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

func (this *RectangleGraphic) Draw(canvas ICanvas) error {
	return nil
}

func (this *RectangleGraphic) Add(child IGraphic) error {
	return nil
}

func (this *RectangleGraphic) Remove(child IGraphic) error {
	return nil
}

func (this *RectangleGraphic) Parent() IGraphic {
	return this.parent
}

func (this *RectangleGraphic) Level(child IGraphic) int {
	return this.Style.Level
}
