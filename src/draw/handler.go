package draw

type IHandler interface {
	SetObject(obj IObject)
	GetControlById(id int) IControlBase
	GetControlByName(name string) IControlBase
	//GetControlDialog(id int) IDialogBase

	SetVisible(name string, isVisible bool)
	SetDisable(name string, isDisable bool)

	SetTitle(name string, title string)
	GetTitle(name string) string
}

type Handler struct {
	obj IObject
}

func (this *Handler) SetObject(obj IObject) {
	this.obj = obj
}

func (this *Handler) SetVisible(controlName string, visible bool) {
	//@@TODO

	/*control := this.impl.GetControlByName(controlName)
	if control != nil {
		control.SetVisible(visible)
	}*/
}

func (this *Handler) GetControlById(id int) IControlBase {
	//@@TODO
	return nil
}

func (this *Handler) GetControlByName(name string) IControlBase {
	//@@TODO
	return nil
}
