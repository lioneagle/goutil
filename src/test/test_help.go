package test

import (
	"path/filepath"
)

func GenTestFileNames(basicPath, standardPath, outputPath string, filename string) (StandardFileName, OutputFileName string) {
	standard_path := filepath.FromSlash(basicPath + standardPath)
	output_path := filepath.FromSlash(basicPath + outputPath)

	StandardFileName = filepath.FromSlash(standard_path + "/" + filename)
	OutputFileName = filepath.FromSlash(output_path + "/" + filename)

	return StandardFileName, OutputFileName
}
