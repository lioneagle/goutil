package model

type FileRange struct {
	constGroup *ConstGroup
	macros     *MacroList
	structs    *StructList
	functions  *FunctionList
}

func NewFileRange() *FileRange {
	ret := &FileRange{}
	ret.constGroup = NewConstGroup()
	ret.macros = NewMacroList()
	ret.structs = NewStructList()
	ret.functions = NewFunctionList()
	return ret
}

type File struct {
	name    string
	comment string

	rangePublic  *FileRange
	rangePrivate *FileRange
}

func NewFile() *File {
	ret := &File{}
	ret.rangePublic = NewFileRange()
	return ret
}

func (this *File) Accept(v CodeVisitor) {
	//v.VisitFile(this)
}

func (this *File) AppendConst(val *ConstList) {
	this.rangePublic.constGroup.AppendConstList(val)
}

type FileList struct {
	Files []*File
}

func NewFileList() *FileList {
	return &FileList{}
}

func (this *FileList) Accept(v CodeVisitor) {
	//v.VisitFileList(this)
}

func (this *FileList) Append(val *File) {
	this.Files = append(this.Files, val)
}
