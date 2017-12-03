package file

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func FileEqual(filename1, filename2 string) bool {
	file1, err := ioutil.ReadFile(filename1)
	if err != nil {
		fmt.Printf("ERROR: cannot open file %s\r\n", filename1)
		return false
	}

	file2, err := ioutil.ReadFile(filename2)
	if err != nil {
		fmt.Printf("ERROR: cannot open file %s\r\n", filename2)
		return false
	}

	return bytes.Equal(file1, file2)
}

func ReplaceFileSuffix(filename, newSuffix string) string {
	return fmt.Sprintf("%s.%s", RemoveFileSuffix(filename), newSuffix)
}

func RemoveFileSuffix(filename string) string {
	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(filename, ext)
}

func Ext(path string) string {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

func GetCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func RemoveExistFiles(filenames []string) error {
	for i := 0; i < len(filenames); i++ {
		err := RemoveExistFile(filenames[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveExistFile(filename string) error {
	ok, _ := PathOrFileIsExist(filename)
	if ok {
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func PathOrFileIsExist(pathOrFile string) (bool, error) {
	_, err := os.Stat(pathOrFile)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
