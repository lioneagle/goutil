package fswatcher

import (
	"fmt"

	"github.com/lioneagle/goutil/src/buffer"
)

// FsEvent represents a single file system notification.
type FsEvent struct {
	Path string // Relative path to the file or directory.
	Op   FsOp   // File operation that triggered the event.
}

func (this *FsEvent) String() string {
	return fmt.Sprintf("%q: %s", this.Path, this.Op.String())
}

type FsOp uint32

const (
	FS_CREATE FsOp = 1 << iota
	FS_REMOVE
	FS_WRITE
	FS_RENAME
	FS_CHMOD
)

func (this FsOp) IsFsCreate() bool {
	return this&FS_CREATE == FS_CREATE
}

func (this FsOp) IsFsRemove() bool {
	return this&FS_REMOVE == FS_REMOVE
}

func (this FsOp) IsFsWrite() bool {
	return this&FS_WRITE == FS_WRITE
}

func (this FsOp) IsFsRename() bool {
	return this&FS_RENAME == FS_RENAME
}

func (this FsOp) IsFsChmod() bool {
	return this&FS_CHMOD == FS_CHMOD
}

func (this FsOp) String() string {
	buf := buffer.NewByteBuffer(nil)

	if this.IsFsCreate() {
		buf.WriteString(" | FS_CREATE")
	}

	if this.IsFsRemove() {
		buf.WriteString(" | FS_REMOVE")
	}

	if this.IsFsWrite() {
		buf.WriteString(" | FS_WRITE")
	}

	if this.IsFsRename() {
		buf.WriteString(" | FS_RENAME")
	}

	if this.IsFsChmod() {
		buf.WriteString(" | FS_CHMOD")
	}

	if buf.Len() == 0 {
		return ""
	}

	return buf.String()[3:]
}
