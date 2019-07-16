package model

type Package struct {
	name    string
	alias   string
	comment string

	Files *FileList
}

func NewPackage() *Package {
	return &Package{}
}

func (this *Package) Accept(v CodeVisitor) {
	//v.VisitPackage(this)
}

func (this *Package) AppendFile(val *File) {
	this.Files.Append(val)
}

type PackageList struct {
	Packages []*Package
}

func NewPakcageList() *PackageList {
	return &PackageList{}
}

func (this *PackageList) Accept(v CodeVisitor) {
	//v.VisitPackageList(this)
}

func (this *PackageList) AppendPackage(val *Package) {
	this.Packages = append(this.Packages, val)
}
