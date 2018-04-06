package draw

import (
	"encoding/xml"
	"strconv"
)

type IObject interface {
	GetId() int
	GetName() string
	SetName(name string)

	IsClass(name string) bool
	IsThisObject(id int, name string) bool

	GetObjectClass() string
	BaseObjectClassName() string

	RegisterHandler(handler IHandler) bool
	SetHandler(handler IHandler)
	GetHandler() IHandler

	OnBaseMessage(id, msg int, wParam, lParam int) int
	OnControlUpdate(rect *Rectangle, isUpdate bool, controlBase IControlBase) int

	SetAttribute(name string, value string, isLoading bool) int
	LoadAttributesInfo() bool
	LoadFromXml(decoder xml.Decoder, start *xml.StartElement, loadSubControl bool) bool

	OnInit() bool

	SetRect(rect *Rectangle)
	GetRect() Rectangle
}

type Object struct {
	id      int
	name    string
	rect    Rectangle
	handler IHandler
	attr    map[string]*ObjectAttributeInfo
}

func (this *Object) GetId() int {
	return this.id
}

func (this *Object) GetName() string {
	return this.name
}

func (this *Object) SetName(name string) {
	this.name = name
}

func (this *Object) IsClass(name string) bool {
	return false
}

func (this *Object) IsThisObject(id int, name string) bool {
	return false
}

func (this *Object) GetObjectClass() string {
	return ""
}

func (this *Object) BaseObjectClassName() string {
	return ""
}

func (this *Object) RegisterHandler(handler IHandler) bool {
	if handler == nil {
		return false
	}
	this.handler = handler
	this.handler.SetObject(this)
	return true
}

func (this *Object) SetHandler(handler IHandler) {
	this.handler = handler
}

func (this *Object) GetHandler() IHandler {
	return this.handler
}

func (this *Object) OnBaseMessage(id, msg int, wParam, lParam int) int {
	return 0
}

func (this *Object) OnControlUpdate(rect *Rectangle, isUpdate bool, controlBase IControlBase) int {
	return 0
}

func (this *Object) SetAttribute(name string, value string, isLoading bool) int {
	if name == "id" {
		var err error
		this.id, err = strconv.Atoi(name)
		if err != nil {
			return -1
		}
	} else if name == "name" {
		this.name = name
	} else {
		return -1
	}
	return 0
}

func (this *Object) LoadAttributesInfo() bool {
	id := &ObjectAttributeInfo{}
	id.name = "id"

	name := &ObjectAttributeInfo{}

	this.attr["id"] = id
	this.attr["name"] = name

	return true
}

func (this *Object) LoadFromXml(decoder xml.Decoder, start *xml.StartElement, loadSubControl bool) bool {
	//@@TODO
	return true
}

func (this *Object) OnInit() bool {
	return true
}

func (this *Object) SetRect(rect *Rectangle) {
	this.rect = *rect
}

func (this *Object) GetRect() Rectangle {
	return this.rect
}

const (
	OBJ_ATTR_TYPE_INT  = 0
	OBJ_ATTR_TYPE_BOOL = 1
	OBJ_ATTR_TYPE_UINT = 2
)

type ObjectAttributeInfo struct {
	name      string
	valueType int
	value     interface{}
	setFunc   interface{}
	getFunc   interface{}
}
