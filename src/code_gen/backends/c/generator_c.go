package backends

import (
	//"fmt"
	"io"

	//"github.com/lioneagle/abnf/src/basic"
	//"github.com/lioneagle/goutil/src/chars"

	"github.com/lioneagle/goutil/src/code_gen/backends" // a
	//"github.com/lioneagle/goutil/src/code_gen/model"
)

type CGeneratorBase struct {
	backends.CLikeGeneratorBase
}

func NewCGeneratorBase(w io.Writer, config backends.CConfig) *CGeneratorBase {
	gen := &CGeneratorBase{}
	gen.CLikeGeneratorBase.Init(w, config)
	return gen
}
