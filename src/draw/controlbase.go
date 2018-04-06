package draw

type IControlBase interface {
	IObject
}

type ControlBase struct {
	impl IControlBase

	visible bool
}

func (this *ControlBase) SetVisible(visible bool) {
	this.visible = visible
}
