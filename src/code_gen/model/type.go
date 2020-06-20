package model

type Type struct {
	name string
}

func NewType() *Type {
	return &Type{}
}

type TypeList struct {
	types []*Type
}

func NewTypeList() *TypeList {
	return &TypeList{}
}

func (this *TypeList) Append(val ...*Type) *TypeList {
	this.types = append(this.types, val...)
	return this
}

func (this *TypeList) GetTypeNameLen() int {
	len1 := 0
	for _, v := range this.types {
		len1 += len(v.name)
	}

	if len(this.types) < 2 {
		return len1
	}

	return len1 + (len(this.types))*2 // 2 -- len(", ") and "()", such as "(int, bool)"
}
