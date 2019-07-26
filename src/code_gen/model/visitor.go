package model

type CodeVisitor interface {
	//VisitPacakge(val *Package)
	//VisitPacakgeList(val *PackageList)

	VisitComment(val *Comment)

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

	VisitMultiChoiceBegin(val *MultiChoice)
	VisitChoiceFirstBegin(val *Choice)
	VisitChoiceFirstEnd(val *Choice)
	VisitChoiceNonFirstBegin(val *Choice)
	VisitChoiceNonFirstEnd(val *Choice)
	VisitMultiChoiceLastCode(val Code)
	VisitMultiChoiceEnd(val *MultiChoice)

	VisitRepeatAsForBegin(val *Repeat)
	VisitRepeatAsForEnd(val *Repeat)

	VisitRepeatAsWhileBegin(val *Repeat)
	VisitRepeatAsWhileEnd(val *Repeat)

	VisitRepeatAsDoWhileBegin(val *Repeat)
	VisitRepeatAsDoWhileEnd(val *Repeat)

	VisitFuncParamVarFirst(val *Var)
	VisitFuncParamVarNonFirstBegin()
	VisitFuncParamVarNonFirst(val *Var)
	VisitFuncParamVarNonFirstEnd()

	VisitFuncDeclare(val *Function)
	VisitFuncDefine(val *Function)

	VisitFuncNoReturn()
	VisitFuncReturnFirst(val *Var)
	VisitFuncReturnNonFirst(val *Var)

	VisitMacroParamVarFirst(val *Var)
	VisitMacroParamVarNonFirstBegin()
	VisitMacroParamVarNonFirst(val *Var)
	VisitMacroParamVarNonFirstEnd()
	VisitMacroDefine(val *MacroDefine)
}

type NullCodeVisitor struct {
}

func (this *NullCodeVisitor) VisitComment(val *Comment) {}

//func (this *NullCodeVisitor) VisitPacakge(val *Package)         {}
//func (this *NullCodeVisitor) VisitPacakgeList(val *PackageList) {}
//func (this *NullCodeVisitor) VisitVar(val *Var)                 {}
//func (this *NullCodeVisitor) VisitConst(val *Const)             {}

func (this *NullCodeVisitor) VisitBlockBegin(val *Block) {}

//func (this *NullCodeVisitor) VisitBlock(val *Block)             {}
func (this *NullCodeVisitor) VisitBlockEnd(val *Block) {}

//func (this *NullCodeVisitor) VisitFunction(val *Function)       {}
func (this *NullCodeVisitor) VisitSentence(val *Sentence) {}

func (this *NullCodeVisitor) VisitStructBegin(val *Struct)               {}
func (this *NullCodeVisitor) VisitStructRangePublicBegin(val *Struct)    {}
func (this *NullCodeVisitor) VisitStructRangePublicEnd(val *Struct)      {}
func (this *NullCodeVisitor) VisitStructRangeProtectedBegin(val *Struct) {}
func (this *NullCodeVisitor) VisitStructRangeProtectedEnd(val *Struct)   {}
func (this *NullCodeVisitor) VisitStructRangePrivateBegin(val *Struct)   {}
func (this *NullCodeVisitor) VisitStructRangePrivateEnd(val *Struct)     {}
func (this *NullCodeVisitor) VisitStructEnd(val *Struct)                 {}
func (this *NullCodeVisitor) VisitStructFieldVarListBegin(val *VarList)  {}
func (this *NullCodeVisitor) VisitStructFieldVar(val *Var)               {}
func (this *NullCodeVisitor) VisitStructFieldVarListEnd(val *VarList)    {}
func (this *NullCodeVisitor) VisitMultiChoiceBegin(val *MultiChoice)     {}
func (this *NullCodeVisitor) VisitChoiceFirstBegin(val *Choice)          {}
func (this *NullCodeVisitor) VisitChoiceFirstEnd(val *Choice)            {}
func (this *NullCodeVisitor) VisitChoiceNonFirstBegin(val *Choice)       {}
func (this *NullCodeVisitor) VisitChoiceNonFirstEnd(val *Choice)         {}
func (this *NullCodeVisitor) VisitMultiChoiceLastCode(val Code)          {}
func (this *NullCodeVisitor) VisitMultiChoiceEnd(val *MultiChoice)       {}
func (this *NullCodeVisitor) VisitRepeatAsForBegin(val *Repeat)          {}
func (this *NullCodeVisitor) VisitRepeatAsForEnd(val *Repeat)            {}
func (this *NullCodeVisitor) VisitRepeatAsWhileBegin(val *Repeat)        {}
func (this *NullCodeVisitor) VisitRepeatAsWhileEnd(val *Repeat)          {}
func (this *NullCodeVisitor) VisitRepeatAsDoWhileBegin(val *Repeat)      {}
func (this *NullCodeVisitor) VisitRepeatAsDoWhileEnd(val *Repeat)        {}
func (this *NullCodeVisitor) VisitFuncParamVarFirst(val *Var)            {}
func (this *NullCodeVisitor) VisitFuncParamVarNonFirstBegin()            {}
func (this *NullCodeVisitor) VisitFuncParamVarNonFirst(val *Var)         {}
func (this *NullCodeVisitor) VisitFuncParamVarNonFirstEnd()              {}
func (this *NullCodeVisitor) VisitFuncDeclare(val *Function)             {}
func (this *NullCodeVisitor) VisitFuncDefine(val *Function)              {}
func (this *NullCodeVisitor) VisitFuncNoReturn()                         {}
func (this *NullCodeVisitor) VisitFuncReturnFirst(val *Var)              {}
func (this *NullCodeVisitor) VisitFuncReturnNonFirst(val *Var)           {}
func (this *NullCodeVisitor) VisitMacroDefine(val *MacroDefine)          {}
