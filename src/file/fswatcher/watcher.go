package fswatcher

import (
	_ "github.com/lioneagle/goutil/src/detailerr"
)

type watcher interface {
	Watch(path string, event FsOp)
}
