package backends

import (
	//"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lioneagle/goutil/src/code_gen/backends"
	"github.com/lioneagle/goutil/src/code_gen/model"

	"github.com/lioneagle/goutil/src/file"
	"github.com/lioneagle/goutil/src/logger"
	"github.com/lioneagle/goutil/src/test"
)

func TestCGeneratorGenBlock(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames("../../test_data/", "test_standard", "test_output", "c_genertator_gen_block.c")

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	config := backends.NewCConfig()
	generator := NewCGeneratorBase(outputFile, config)

	block := model.NewBlock()
	sentence := model.NewSentence("test_gen_block();")
	block.AppendCode(sentence)

	block.Accept(generator)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorGenEnum(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames("../../test_data/", "test_standard", "test_output", "c_genertator_gen_enum.c")

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	config := backends.NewCConfig()
	generator := NewCGeneratorBase(outputFile, config)

	vars := []struct {
		name      string
		typeName  string
		initValue string
		comment   string
	}{
		{"name", "int", "100", "name of book"},
		{"value", "int", "100", "value of book"},
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

	configComment := config.ConfigCommonImpl.String()
	configComment += "\n" + config.CConfigImpl.String()

	generator.CLikeGeneratorBase.GenMultiLineComment(configComment)

	generator.PrintReturn(outputFile)

	constList.Accept(generator)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}
