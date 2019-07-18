package backends

import (
	//"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lioneagle/goutil/src/code_gen/model"

	"github.com/lioneagle/goutil/src/file"
	"github.com/lioneagle/goutil/src/logger"
	"github.com/lioneagle/goutil/src/test"
)

func TestCLikeGeneratorGenBlock(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames("../test_data/", "test_standard", "test_output", "clike_genertator_gen_block.c")

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	config := NewCConfig()
	generator := NewCLikeGeneratorBase(outputFile, config)

	block := model.NewBlock()
	sentence := model.NewSentence("test_gen_block();")
	block.AppendCode(sentence)

	block.Accept(generator)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorGenEnum(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames("../test_data/", "test_standard", "test_output", "clike_genertator_gen_enum.c")

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	config := NewCConfig()
	generator := NewCLikeGeneratorBase(outputFile, config)

	vars := []struct {
		name      string
		typeName  string
		initValue string
		comment   string
	}{
		{"name", "int", "100", "name of book"},
		{"value", "int", "", "value of book"},
		{"note", "int", "10", "note of book"},
	}

	constList := model.NewConstList("ATTR_TYPE")

	for _, v := range vars {
		val := model.NewVar()
		val.SetName(v.name)
		val.SetTypeName(v.typeName)
		val.SetInitValue(v.initValue)
		val.SetComment(v.comment)
		constList.AppendConst(val)
	}

	constList.Accept(generator)

	config.Indent().Assign = 2
	config.Indent().Comment = 2
	config.SetVarUseSingleLineComment(false)
	config.SetBraceAtNextLine(false)

	generator.PrintReturn(outputFile)

	constList.Accept(generator)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorGenComment(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames("../test_data/", "test_standard", "test_output", "clike_genertator_gen_comment.c")

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	config := NewCConfig()
	generator := NewCLikeGeneratorBase(outputFile, config)

	comment := `
    NAME: parse
    PARAMS: 
    context -- context for parsing
    src     -- source to parse
    len     -- length of source
    RETURN:
    length parsed
    NOTE: create for test
    `

	generator.GenMultiLineComment(comment)
	generator.genSingleLineCommentWithoutIndent("test single line comment")

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorGenStruct(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames("../test_data/", "test_standard", "test_output", "clike_genertator_gen_struct.c")

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	config := NewCConfig()
	generator := NewCLikeGeneratorBase(outputFile, config)

	struct1 := model.NewStruct()
	struct1.SetName("ATTR_TYPE")

	vars := []struct {
		name     string
		typeName string
		comment  string
	}{
		{"name", "int", "name of book"},
		{"value", "int", "value of book"},
		{"note", "int", "note of book"},
	}

	for _, v := range vars {
		var1 := model.NewVar()
		var1.SetName(v.name)
		var1.SetTypeName(v.typeName)
		var1.SetComment(v.comment)
		struct1.AppendFieldPublic(var1)
	}

	struct1.Accept(generator)

	config.SetBraceAtNextLine(false)

	generator.PrintReturn(outputFile)
	struct1.Accept(generator)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorGenSingleChoice(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames("../test_data/", "test_standard", "test_output", "clike_genertator_gen_if.c")

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	config := NewCConfig()
	generator := NewCLikeGeneratorBase(outputFile, config)

	choice := model.NewSingleChoice()
	choice.SetCondition("i < 100")
	choice.SetComment("generate if, left brace at next line")
	choice.SetCodeTrue(model.NewSentence("x += 5;"))
	choice.Accept(generator)

	choice.SetCodeFalse(model.NewSentence("x -= 3;"))
	choice.Accept(generator)

	generator.PrintReturn(outputFile)

	config.SetBraceAtNextLine(false)

	choice.SetCodeTrue(model.NewSentence("y++;"))
	choice.SetComment("generate if, left brace at same line")
	choice.SetCodeFalse(nil)
	choice.Accept(generator)

	choice.SetCodeFalse(model.NewSentence("y -= 2;"))
	choice.Accept(generator)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorGenFor(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames("../test_data/", "test_standard", "test_output", "clike_genertator_gen_repeat_as_for.c")

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	config := NewCConfig()
	generator := NewCLikeGeneratorBase(outputFile, config)

	repeat := model.NewRepeat()
	repeat.SetComment("generate for, left brace at next line")
	repeat.SetCondition("i = 0; i < 100; i++")
	repeat.SetCode(model.NewSentence("x += 5;"))
	repeat.SetAcceptType(model.REPEAT_TYPE_FOR)
	repeat.Accept(generator)

	generator.PrintReturn(outputFile)

	config.SetBraceAtNextLine(false)

	repeat.SetComment("generate for, left brace at same line")
	repeat.SetCode(model.NewSentence("y++;"))
	repeat.Accept(generator)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorGenWhile(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames("../test_data/", "test_standard", "test_output", "clike_genertator_gen_repeat_as_while.c")

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	config := NewCConfig()
	generator := NewCLikeGeneratorBase(outputFile, config)

	repeat := model.NewRepeat()
	repeat.SetComment("generate while, left brace at next line")
	repeat.SetCondition("i < 100")
	repeat.SetCode(model.NewSentence("x += 5;"))
	repeat.SetAcceptType(model.REPEAT_TYPE_WHILE)
	repeat.Accept(generator)

	generator.PrintReturn(outputFile)

	config.SetBraceAtNextLine(false)

	repeat.SetComment("generate while, left brace at same line")
	repeat.SetCode(model.NewSentence("y++;"))
	repeat.Accept(generator)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorGenDoWhile(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames("../test_data/", "test_standard", "test_output", "clike_genertator_gen_repeat_as_do_while.c")

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	config := NewCConfig()
	generator := NewCLikeGeneratorBase(outputFile, config)

	repeat := model.NewRepeat()
	repeat.SetComment("generate do/while, left brace at next line")
	repeat.SetCondition("i < 200")
	repeat.SetCode(model.NewSentence("x += 5;"))
	repeat.SetAcceptType(model.REPEAT_TYPE_DO_WHILE)
	repeat.Accept(generator)

	generator.PrintReturn(outputFile)

	config.SetBraceAtNextLine(false)

	repeat.SetComment("generate do/while, left brace at same line")
	repeat.SetCode(model.NewSentence("y++;"))
	repeat.Accept(generator)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}
