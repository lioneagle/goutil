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
