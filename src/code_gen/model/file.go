package model

type File struct {
	path    string
	name    string
	comment string
	codes   *Codes
}

func NewFile(name string) *File {
	return &File{
		name:  name,
		codes: NewCodes(),
	}
}

func (this *File) SetName(name string) {
	this.name = name
}

func (this *File) GetName() string {
	return this.name
}

func (this *File) SetPath(path string) {
	this.path = path
}

func (this *File) GetPath() string {
	return this.path
}

func (this *File) SetComment(comment string) {
	this.comment = comment
}

func (this *File) GetComment() string {
	return this.comment
}

func (this *File) GetCodes() *Codes {
	return this.codes
}

func (this *File) SetCodes(codes *Codes) {
	this.codes = codes
}

func (this *File) Accept(visitor CodeVisitor) {
	visitor.VisitFileBegin(this)
	this.codes.Accept(visitor)
	visitor.VisitFileEnd(this)
}

/*
type FileList struct {
	files []*File
}

func NewFileList() *FileList {
	return &FileList{
		files: make([]*File, 0),
}
}

func (this *FileList) Append(val ...*File) *FileList {
	this.files = append(this.files, val...)
	return this
}

func (this *FileList) Accept(visitor CodeVisitor) {
	visitor.VisitFileList(this)
}*/
