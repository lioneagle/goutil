package model

type CodeVisitor interface {
	//VisitPacakge(val *Package)
	//VisitPacakgeList(val *PackageList)

	VisitComment(val *Comment)

	VisitCodesBegin(val *Codes)
	VisitCodesEnd(val *Codes)

	VisitBlockBegin(val *Block)
	VisitBlockEnd(val *Block)
	//VisitBlock(val *Block)

	VisitStructBegin(val *Struct)
	VisitStructRangePublicBegin(val *Struct)
	VisitStructRangePublicEnd(val *Struct)
	VisitStructRangeProtectedBegin(val *Struct)
	VisitStructRangeProtectedEnd(val *Struct)
	VisitStructRangePrivateBegin(val *Struct)
	VisitStructRangePrivateEnd(val *Struct)
	VisitStructEnd(val *Struct)

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

	VisitChoiceGroupBegin(val *ChoiceGroup)
	VisitChoiceGroupItemBegin(val *Choice)
	VisitChoiceGroupItemEnd(val *Choice)
	VisitChoiceGroupDefaultBegin(val Code)
	VisitChoiceGroupDefaultEnd(val Code)
	VisitChoiceGroupEnd(val *ChoiceGroup)

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
	VisitMacroUndefine(val *MacroUndefine)
	VisitMacroCode(val Code)

	VisitMacroMultiChoiceBegin(val *MultiChoice)
	VisitMacroChoiceFirstBegin(val *Choice)
	VisitMacroChoiceFirstEnd(val *Choice)
	VisitMacroChoiceNonFirstBegin(val *Choice)
	VisitMacroChoiceNonFirstEnd(val *Choice)
	VisitMacroMultiChoiceLastCode(val Code)
	VisitMacroMultiChoiceEnd(val *MultiChoice)

	VisitModuleImport(val *ModuleImport)
	VisitModuleImportListBegin(val *ModuleImportList)
	VisitModuleImportListEnd(val *ModuleImportList)

	VisitFileBegin(val *File)
	VisitFileEnd(val *File)
	//VisitFileList(val *FileList)
}

type NullCodeVisitor struct {
}

func (this *NullCodeVisitor) VisitComment(val *Comment) {}

//func (this *NullCodeVisitor) VisitPacakge(val *Package)         {}
//func (this *NullCodeVisitor) VisitPacakgeList(val *PackageList) {}
//func (this *NullCodeVisitor) VisitVar(val *Var)                 {}
//func (this *NullCodeVisitor) VisitConst(val *Const)             {}

func (this *NullCodeVisitor) VisitCodesBegin(val *Codes) {}
func (this *NullCodeVisitor) VisitCodesEnd(val *Codes)   {}

func (this *NullCodeVisitor) VisitBlockBegin(val *Block) {}

//func (this *NullCodeVisitor) VisitBlock(val *Block)             {}
func (this *NullCodeVisitor) VisitBlockEnd(val *Block) {}

//func (this *NullCodeVisitor) VisitFunction(val *Function)       {}
func (this *NullCodeVisitor) VisitSentence(val *Sentence) {}

func (this *NullCodeVisitor) VisitStructBegin(val *Struct)                {}
func (this *NullCodeVisitor) VisitStructRangePublicBegin(val *Struct)     {}
func (this *NullCodeVisitor) VisitStructRangePublicEnd(val *Struct)       {}
func (this *NullCodeVisitor) VisitStructRangeProtectedBegin(val *Struct)  {}
func (this *NullCodeVisitor) VisitStructRangeProtectedEnd(val *Struct)    {}
func (this *NullCodeVisitor) VisitStructRangePrivateBegin(val *Struct)    {}
func (this *NullCodeVisitor) VisitStructRangePrivateEnd(val *Struct)      {}
func (this *NullCodeVisitor) VisitStructEnd(val *Struct)                  {}
func (this *NullCodeVisitor) VisitStructFieldVarListBegin(val *VarList)   {}
func (this *NullCodeVisitor) VisitStructFieldVar(val *Var)                {}
func (this *NullCodeVisitor) VisitStructFieldVarListEnd(val *VarList)     {}
func (this *NullCodeVisitor) VisitMultiChoiceBegin(val *MultiChoice)      {}
func (this *NullCodeVisitor) VisitChoiceFirstBegin(val *Choice)           {}
func (this *NullCodeVisitor) VisitChoiceFirstEnd(val *Choice)             {}
func (this *NullCodeVisitor) VisitChoiceNonFirstBegin(val *Choice)        {}
func (this *NullCodeVisitor) VisitChoiceNonFirstEnd(val *Choice)          {}
func (this *NullCodeVisitor) VisitMultiChoiceLastCode(val Code)           {}
func (this *NullCodeVisitor) VisitMultiChoiceEnd(val *MultiChoice)        {}
func (this *NullCodeVisitor) VisitChoiceGroupBegin(val *ChoiceGroup)      {}
func (this *NullCodeVisitor) VisitChoiceGroupItemBegin(val *Choice)       {}
func (this *NullCodeVisitor) VisitChoiceGroupItemEnd(val *Choice)         {}
func (this *NullCodeVisitor) VisitChoiceGroupDefaultBegin(val Code)       {}
func (this *NullCodeVisitor) VisitChoiceGroupDefaultEnd(val Code)         {}
func (this *NullCodeVisitor) VisitChoiceGroupEnd(val *ChoiceGroup)        {}
func (this *NullCodeVisitor) VisitRepeatAsForBegin(val *Repeat)           {}
func (this *NullCodeVisitor) VisitRepeatAsForEnd(val *Repeat)             {}
func (this *NullCodeVisitor) VisitRepeatAsWhileBegin(val *Repeat)         {}
func (this *NullCodeVisitor) VisitRepeatAsWhileEnd(val *Repeat)           {}
func (this *NullCodeVisitor) VisitRepeatAsDoWhileBegin(val *Repeat)       {}
func (this *NullCodeVisitor) VisitRepeatAsDoWhileEnd(val *Repeat)         {}
func (this *NullCodeVisitor) VisitFuncParamVarFirst(val *Var)             {}
func (this *NullCodeVisitor) VisitFuncParamVarNonFirstBegin()             {}
func (this *NullCodeVisitor) VisitFuncParamVarNonFirst(val *Var)          {}
func (this *NullCodeVisitor) VisitFuncParamVarNonFirstEnd()               {}
func (this *NullCodeVisitor) VisitFuncDeclare(val *Function)              {}
func (this *NullCodeVisitor) VisitFuncDefine(val *Function)               {}
func (this *NullCodeVisitor) VisitFuncNoReturn()                          {}
func (this *NullCodeVisitor) VisitFuncReturnFirst(val *Var)               {}
func (this *NullCodeVisitor) VisitFuncReturnNonFirst(val *Var)            {}
func (this *NullCodeVisitor) VisitMacroDefine(val *MacroDefine)           {}
func (this *NullCodeVisitor) VisitMacroUndefine(val *MacroUndefine)       {}
func (this *NullCodeVisitor) VisitMacroCode(val Code)                     {}
func (this *NullCodeVisitor) VisitMacroMultiChoiceBegin(val *MultiChoice) {}
func (this *NullCodeVisitor) VisitMacroChoiceFirstBegin(val *Choice)      {}
func (this *NullCodeVisitor) VisitMacroChoiceFirstEnd(val *Choice)        {}
func (this *NullCodeVisitor) VisitMacroChoiceNonFirstBegin(val *Choice)   {}
func (this *NullCodeVisitor) VisitMacroChoiceNonFirstEnd(val *Choice)     {}
func (this *NullCodeVisitor) VisitMacroMultiChoiceLastCode(val Code)      {}
func (this *NullCodeVisitor) VisitMacroMultiChoiceEnd(val *MultiChoice)   {}
func (this *NullCodeVisitor) VisitModuleImport(val *ModuleImport)              {}
func (this *NullCodeVisitor) VisitModuleImportListBegin(val *ModuleImportList) {}
func (this *NullCodeVisitor) VisitModuleImportListEnd(val *ModuleImportList)   {}
func (this *NullCodeVisitor) VisitFileBegin(val *File)                         {}
func (this *NullCodeVisitor) VisitFileEnd(val *File)                           {}

//func (this *NullCodeVisitor) VisitFileList(val *FileList)                      {}
