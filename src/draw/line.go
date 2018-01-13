package core

type LineGraphic struct {
	from   Point
	to     Point
	text   *Text
	style  LineStyle
	parent Graphic
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

func (this *LineGraphic) Draw(canvas Canvas) error {
	return nil
}

func (this *LineGraphic) Add(child Graphic) error {
	return nil
}

func (this *LineGraphic) Remove(child Graphic) error {
	return nil
}

func (this *LineGraphic) Parent() Graphic {
	return this.parent
}

func (this *LineGraphic) Level(child Graphic) int {
	return this.style.level
}
