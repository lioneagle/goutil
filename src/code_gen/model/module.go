package model

type ModuleImport struct {
	name    string
	alias   string
	comment string
}

func NewModuleImport(name string) *ModuleImport {
	return &ModuleImport{name: name}
}

func (this *ModuleImport) SetName(name string) {
	this.name = name
}

func (this *ModuleImport) GetName() string {
	return this.name
}

func (this *ModuleImport) SetAlias(alias string) {
	this.alias = alias
}

func (this *ModuleImport) GetAlias() string {
	return this.alias
}

func (this *ModuleImport) SetComment(comment string) {
	this.comment = comment
}

func (this *ModuleImport) GetComment() string {
	return this.comment
}

func (this *ModuleImport) Accept(visitor CodeVisitor) {
	visitor.VisitModuleImport(this)
}

type ModuleImportList struct {
	comment string
	modules []*ModuleImport
}

func NewModuleImportList() *ModuleImportList {
	return &ModuleImportList{
		modules: make([]*ModuleImport, 0),
	}
}

func (this *ModuleImportList) SetComment(comment string) {
	this.comment = comment
}

func (this *ModuleImportList) GetComment() string {
	return this.comment
}

func (this *ModuleImportList) AppendModule(val ...*ModuleImport) {
	this.modules = append(this.modules, val...)
}

func (this *ModuleImportList) Len() int {
	return len(this.modules)
}

func (this *ModuleImportList) Accept(visitor CodeVisitor) {
	if len(this.modules) <= 0 {
		return
	}

	visitor.VisitModuleImportListBegin(this)

	for _, v := range this.modules {
		v.Accept(visitor)
	}

	visitor.VisitModuleImportListEnd(this)
}
