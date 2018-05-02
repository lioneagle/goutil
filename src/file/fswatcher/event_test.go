package fswatcher

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestFsEventString(t *testing.T) {
	testdata := []struct {
		path string
		op   FsOp
		str  string
	}{
		{"test/file1", FS_CREATE | FS_WRITE, `"test/file1": FS_CREATE | FS_WRITE`},
		{"test/file2", FS_RENAME, `"test/file2": FS_RENAME`},
		{"test/file3", FS_REMOVE, `"test/file3": FS_REMOVE`},
		{"test/file4", FS_CHMOD | FS_WRITE, `"test/file4": FS_WRITE | FS_CHMOD`},
		{"test/file4", 0, `"test/file4": `},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			event := &FsEvent{Path: v.path, Op: v.op}
			test.EXPECT_EQ(t, event.String(), v.str, "")
		})
	}
}
