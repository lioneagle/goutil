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

	"github.com/pkg/errors"
)

func FileEqual(filename1, filename2 string) error {
	file1, err := ioutil.ReadFile(filename1)
	if err != nil {
		return err
	}

	file2, err := ioutil.ReadFile(filename2)
	if err != nil {
		return err
	}

	if len(file1) != len(file2) {
		return errors.Errorf("different len, len(%s) = %d, len(%s) = %d", filename1, len(file1), filename2, len(file2))
	}

	ret := bytes.Equal(file1, file2)
	if !ret {
		for i := 0; i < len(file1); i++ {
			if file1[i] != file2[i] {
				return errors.Errorf("first diffrent char is %d at position %d", file1[i], i)
			}
		}
	}

	return nil
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

func AppendFile(filename string, data []byte, perm os.FileMode) error {
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

func CopyFile(dstFileName, srcFileName string) (int64, error) {
	src, err := os.Open(srcFileName)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0x777)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
