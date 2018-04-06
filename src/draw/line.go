package draw

type LineGraphic struct {
	from   Point
	to     Point
	text   *Text
	style  LineStyle
	parent IGraphic
}

type LineType uint32

type LineStyle struct {
	level int
	color Color
	arrow ArrowStyle
}

type ArrowStyle struct {
	hasArrow         bool
	inverseDirection bool
}

func (this *LineGraphic) Draw(canvas ICanvas) error {
	return nil
}

func (this *LineGraphic) Add(child IGraphic) error {
	return nil
}

func (this *LineGraphic) Remove(child IGraphic) error {
	return nil
}

func (this *LineGraphic) Parent() IGraphic {
	return this.parent
}

func (this *LineGraphic) Level(child IGraphic) int {
	return this.style.level
}
