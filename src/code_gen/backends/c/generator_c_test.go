package backends

import (
	//"fmt"

	"testing"

	"github.com/lioneagle/goutil/src/code_gen/backends"
	"github.com/lioneagle/goutil/src/code_gen/model"

	"github.com/lioneagle/goutil/src/file"
	"github.com/lioneagle/goutil/src/test"
)

func TestCGeneratorGenBlock(t *testing.T) {
	testFiles, err := test.GetTestFiles(t, "../../test_data/c/", "c_genertator_gen_block.c")
	test.ASSERT_EQ(t, err, nil, "")
	defer testFiles.Output.File.Close()

	config := backends.NewCConfig()
	generator := NewCGeneratorBase(testFiles.Output.File, config)

	block := model.NewBlock()
	sentence := model.NewSentence("test_gen_block();")
	block.AppendCode(sentence)

	block.Accept(generator)

	err = file.FileEqual(testFiles.Output.Name, testFiles.Standard.Name)
	test.EXPECT_EQ(t, err, nil, "")
}

func TestCGeneratorGenEnum(t *testing.T) {
	testFiles, err := test.GetTestFiles(t, "../../test_data/c/", "c_genertator_gen_enum.c")
	test.ASSERT_EQ(t, err, nil, "")
	defer testFiles.Output.File.Close()

	config := backends.NewCConfig()
	generator := NewCGeneratorBase(testFiles.Output.File, config)

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

	generator.PrintReturn(testFiles.Output.File)

	constList.Accept(generator)

	err = file.FileEqual(testFiles.Output.Name, testFiles.Standard.Name)
	test.EXPECT_EQ(t, err, nil, "")
}
