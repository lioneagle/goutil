package model

type CodeVisitor interface {
	//VisitPacakge(val *Package)
	//VisitPacakgeList(val *PackageList)

	VisitStructBegin(val *Struct)
	VisitStructRangePublicBegin(val *Struct)
	VisitStructRangePublicEnd(val *Struct)
	VisitStructRangeProtectedBegin(val *Struct)
	VisitStructRangeProtectedEnd(val *Struct)
	VisitStructRangePrivateBegin(val *Struct)
	VisitStructRangePrivateEnd(val *Struct)
	VisitStructEnd(val *Struct)

	//VisitBlock(val *Block)
	VisitBlockBegin(val *Block)
	VisitBlockEnd(val *Block)

	//VisitFunction(val *Function)
	VisitSentence(val *Sentence)

	VisitConstsBegin(val *ConstList)
	VisitConst(val *Var)
	VisitConstsEnd(val *ConstList)

	VisitStructFieldVarListBegin(val *VarList)
	VisitStructFieldVar(val *Var)
	VisitStructFieldVarListEnd(val *VarList)

	VisitSingleChoiceBegin(val *SingleChoice)
	VisitSingleChoiceTrueBegin(val *SingleChoice)
	VisitSingleChoiceTrueEnd(val *SingleChoice)
	VisitSingleChoiceFalseBegin(val *SingleChoice)
	VisitSingleChoiceFalseEnd(val *SingleChoice)
	VisitSingleChoiceEnd(val *SingleChoice)

	VisitRepeatAsForBegin(val *Repeat)
	VisitRepeatAsForEnd(val *Repeat)

	VisitRepeatAsWhileBegin(val *Repeat)
	VisitRepeatAsWhileEnd(val *Repeat)

	VisitRepeatAsDoWhileBegin(val *Repeat)
	VisitRepeatAsDoWhileEnd(val *Repeat)
}

type NullCodeVisitor struct {
}

//func (this *NullCodeVisitor) VisitPacakge(val *Package)         {}
//func (this *NullCodeVisitor) VisitPacakgeList(val *PackageList) {}
//func (this *NullCodeVisitor) VisitVar(val *Var)                 {}
//func (this *NullCodeVisitor) VisitConst(val *Const)             {}
func (this *NullCodeVisitor) VisitBlockBegin(val *Block) {}

//func (this *NullCodeVisitor) VisitBlock(val *Block)             {}
func (this *NullCodeVisitor) VisitBlockEnd(val *Block) {}

//func (this *NullCodeVisitor) VisitFunction(val *Function)       {}
func (this *NullCodeVisitor) VisitSentence(val *Sentence) {}

func (this *NullCodeVisitor) VisitStructBegin(val *Struct)                  {}
func (this *NullCodeVisitor) VisitStructRangePublicBegin(val *Struct)       {}
func (this *NullCodeVisitor) VisitStructRangePublicEnd(val *Struct)         {}
func (this *NullCodeVisitor) VisitStructRangeProtectedBegin(val *Struct)    {}
func (this *NullCodeVisitor) VisitStructRangeProtectedEnd(val *Struct)      {}
func (this *NullCodeVisitor) VisitStructRangePrivateBegin(val *Struct)      {}
func (this *NullCodeVisitor) VisitStructRangePrivateEnd(val *Struct)        {}
func (this *NullCodeVisitor) VisitStructEnd(val *Struct)                    {}
func (this *NullCodeVisitor) VisitStructFieldVarListBegin(val *VarList)     {}
func (this *NullCodeVisitor) VisitStructFieldVar(val *Var)                  {}
func (this *NullCodeVisitor) VisitStructFieldVarListEnd(val *VarList)       {}
func (this *NullCodeVisitor) VisitSingleChoiceBegin(val *SingleChoice)      {}
func (this *NullCodeVisitor) VisitSingleChoiceTrueBegin(val *SingleChoice)  {}
func (this *NullCodeVisitor) VisitSingleChoiceTrueEnd(val *SingleChoice)    {}
func (this *NullCodeVisitor) VisitSingleChoiceFalseBegin(val *SingleChoice) {}
func (this *NullCodeVisitor) VisitSingleChoiceFalseEnd(val *SingleChoice)   {}
func (this *NullCodeVisitor) VisitSingleChoiceEnd(val *SingleChoice)        {}
func (this *NullCodeVisitor) VisitRepeatAsForBegin(val *Repeat)             {}
func (this *NullCodeVisitor) VisitRepeatAsForEnd(val *Repeat)               {}
func (this *NullCodeVisitor) VisitRepeatAsWhileBegin(val *Repeat)           {}
func (this *NullCodeVisitor) VisitRepeatAsWhileEnd(val *Repeat)             {}
func (this *NullCodeVisitor) VisitRepeatAsDoWhileBegin(val *Repeat)         {}
func (this *NullCodeVisitor) VisitRepeatAsDoWhileEnd(val *Repeat)           {}
