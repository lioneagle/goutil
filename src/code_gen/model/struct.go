package model

type StructRange struct {
	typeDefine *SentenceList
	fields     *VarList
	methods    *FunctionList
}

func NewStructRange() *StructRange {
	ret := &StructRange{}
	ret.typeDefine = NewSentenceList()
	ret.fields = NewVarList()
	ret.methods = NewFunctionList()
	return ret
}

type Struct struct {
	isGeneric bool
	name      string

	genericTypes *TypeList

	rangePublic    *StructRange
	rangeProtected *StructRange
	rangePrivate   *StructRange
}

func NewStruct() *Struct {
	ret := &Struct{}
	ret.rangePublic = NewStructRange()
	return ret
}

func (this *Struct) Accept(visitor CodeVisitor) {
	visitor.VisitStructBegin(this)

	visitor.VisitStructRangePublicBegin(this)
	this.acceptVisitRange(visitor, this.rangePublic)
	visitor.VisitStructRangePublicEnd(this)

	if this.rangeProtected != nil {
		visitor.VisitStructRangeProtectedBegin(this)
		this.acceptVisitRange(visitor, this.rangeProtected)
		visitor.VisitStructRangeProtectedEnd(this)
	}

	if this.rangePrivate != nil {
		visitor.VisitStructRangePrivateBegin(this)
		this.acceptVisitRange(visitor, this.rangePrivate)
		visitor.VisitStructRangePrivateEnd(this)
	}

	visitor.VisitStructEnd(this)
}

func (this *Struct) acceptVisitRange(visitor CodeVisitor, structRange *StructRange) {
	if structRange == nil {
		return
	}

	structRange.fields.AcceptAsStructField(visitor)

	//TODO: typedef and methods will be supported in future
}

func (this *Struct) GetName() string {
	return this.name
}

func (this *Struct) SetName(name string) {
	this.name = name
}

func (this *Struct) AppendGenericType(val ...*Type) *Struct {
	this.genericTypes.Append(val...)
	return this
}

func (this *Struct) AppendTypeDefinePublic(val ...*Sentence) *Struct {
	this.rangePublic.typeDefine.Append(val...)
	return this
}

func (this *Struct) AppendFieldPublic(val ...*Var) *Struct {
	this.rangePublic.fields.Append(val...)
	return this
}

func (this *Struct) AppendMethodPublic(val ...*Function) *Struct {
	this.rangePublic.methods.Append(val...)
	return this
}

func (this *Struct) AppendTypeDefineProtected(val ...*Sentence) *Struct {
	this.getRangeProtected().typeDefine.Append(val...)
	return this
}

func (this *Struct) AppendFieldProtected(val ...*Var) *Struct {
	this.getRangeProtected().fields.Append(val...)
	return this
}

func (this *Struct) AppendMethodProtected(val ...*Function) *Struct {
	this.getRangeProtected().methods.Append(val...)
	return this
}

func (this *Struct) getRangeProtected() *StructRange {
	if this.rangeProtected == nil {
		this.rangeProtected = NewStructRange()
	}
	return this.rangeProtected
}

func (this *Struct) AppendTypeDefinePrivate(val ...*Sentence) *Struct {
	this.getRangePrivate().typeDefine.Append(val...)
	return this
}

func (this *Struct) AppendFieldPrivate(val ...*Var) *Struct {
	this.getRangePrivate().fields.Append(val...)
	return this
}

func (this *Struct) AppendMethodPrivate(val ...*Function) *Struct {
	this.getRangePrivate().methods.Append(val...)
	return this
}

func (this *Struct) getRangePrivate() *StructRange {
	if this.rangePrivate == nil {
		this.rangePrivate = NewStructRange()
	}
	return this.rangePrivate
}

type StructList struct {
	structs []*Struct
}

func NewStructList() *StructList {
	return &StructList{}
}

func (this *StructList) Append(val ...*Struct) *StructList {
	this.structs = append(this.structs, val...)
	return this
}
