package draw

type Font struct {
	Name            string
	FrontColor      Color
	BackgroundColor Color
	Style           int
	Size            int
}

const (
	FONT_STYLE_NORMAL     = 0x00
	FONT_STYLE_BOLD       = 0x01
	FONT_STYLE_ITALIC     = 0x02
	FONT_STYLE_UNDER_LINE = 0x04
	FONT_STYLE_STRIKE_OUT = 0x08
)

func NewFont(Name string, Style, Size int) *Font {
	f := &Font{Name: Name, Style: Style, FrontColor: ColorBlack, BackgroundColor: ColorWhite, Size: Size}
	return f
}

func (this *Font) GetDotName() string {
	s := this.Name
	if this.Style == FONT_STYLE_BOLD {
		s += " bold"
	}
	return s
}

func (this *Font) GetPlantumlName() string {
	return this.Name
}

func (this *Font) GetStyleName() string {
	switch this.Style {
	case FONT_STYLE_BOLD:
		return "bold"
	case FONT_STYLE_ITALIC:
		return "italic"
	case FONT_STYLE_UNDER_LINE:
		return "underline"
	}
	return ""
}
