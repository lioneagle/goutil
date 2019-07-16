package backends

import (
	//"fmt"
	"io"

	//"github.com/lioneagle/abnf/src/basic"
	//"github.com/lioneagle/goutil/src/chars"

	"github.com/lioneagle/goutil/src/code_gen/backends"
	//"github.com/lioneagle/goutil/src/code_gen/model"
)

type CGeneratorH struct {
	CGeneratorBase
}

func NewCGeneratorH(w io.Writer, config backends.CConfig) *CGeneratorH {
	gen := &CGeneratorH{}
	gen.Init(w, config)
	return gen
}

func (this *CGeneratorH) Init(w io.Writer, config backends.CConfig) {
	this.CGeneratorBase.Init(w, config)
}
