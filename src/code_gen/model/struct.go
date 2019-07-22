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
	name      string
	isGeneric bool

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

func (this *Struct) AppendGenericType(val *Type) {
	this.genericTypes.Append(val)
}

func (this *Struct) AppendTypeDefinePublic(val *Sentence) {
	this.rangePublic.typeDefine.Append(val)
}

func (this *Struct) AppendFieldPublic(val *Var) {
	this.rangePublic.fields.Append(val)
}

func (this *Struct) AppendMethodPublic(val *Function) {
	this.rangePublic.methods.Append(val)
}

func (this *Struct) AppendTypeDefineProtected(val *Sentence) {
	this.getRangeProtected().typeDefine.Append(val)
}

func (this *Struct) AppendFieldProtected(val *Var) {
	this.getRangeProtected().fields.Append(val)
}

func (this *Struct) AppendMethodProtected(val *Function) {
	this.getRangeProtected().methods.Append(val)
}

func (this *Struct) getRangeProtected() *StructRange {
	if this.rangeProtected == nil {
		this.rangeProtected = NewStructRange()
	}
	return this.rangeProtected
}

func (this *Struct) AppendTypeDefinePrivate(val *Sentence) {
	this.getRangePrivate().typeDefine.Append(val)
}

func (this *Struct) AppendFieldPrivate(val *Var) {
	this.getRangePrivate().fields.Append(val)
}

func (this *Struct) AppendMethodPrivate(val *Function) {
	this.getRangePrivate().methods.Append(val)
}

func (this *Struct) getRangePrivate() *StructRange {
	if this.rangePrivate == nil {
		this.rangePrivate = NewStructRange()
	}
	return this.rangePrivate
}

type StructList struct {
	Structs []*Struct
}

func NewStructList() *StructList {
	return &StructList{}
}

func (this *StructList) Append(val *Struct) {
	this.Structs = append(this.Structs, val)
}
