package test

import (
	"os"
	"path/filepath"
	"testing"
)

type TestFile struct {
	Name string
	File *os.File
}

type TestFiles struct {
	Output   TestFile
	Standard TestFile
}

func GetTestFiles(t *testing.T, basicPath, filename string) (*TestFiles, error) {
	t.Helper()

	var err error

	ret := &TestFiles{}
	ret.Standard.Name, ret.Output.Name = GenTestFileNames(basicPath,
		"test_standard", "test_output", filename)

	ret.Output.File, err = os.Create(ret.Output.Name)

	return ret, err
}

func GenTestFileNames(basicPath, standardPath,
	outputPath string, filename string) (StandardFileName, OutputFileName string) {
	standard_path := filepath.FromSlash(basicPath + standardPath)
	output_path := filepath.FromSlash(basicPath + outputPath)

	StandardFileName = filepath.FromSlash(standard_path + "/" + filename)
	OutputFileName = filepath.FromSlash(output_path + "/" + filename)

	return StandardFileName, OutputFileName
}
