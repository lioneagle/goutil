package backends

import (
	//"fmt"
	"io"

	//"github.com/lioneagle/abnf/src/basic"
	//"github.com/lioneagle/goutil/src/chars"

	"github.com/lioneagle/goutil/src/code_gen/backends"
	//"github.com/lioneagle/goutil/src/code_gen/model"
)

type CGeneratorC struct {
	CGeneratorBase
}

func NewCGeneratorC(w io.Writer, config backends.CConfig) *CGeneratorC {
	gen := &CGeneratorC{}
	gen.Init(w, config)
	return gen
}

func (this *CGeneratorC) Init(w io.Writer, config backends.CConfig) {
	this.CGeneratorBase.Init(w, config)
}
