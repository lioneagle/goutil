package backends

import (
	"fmt"
)

type CConfig interface {
	ConfigCommon

	BraceAtNextLine() bool
	SetBraceAtNextLine(val bool)
}

type CConfigImpl struct {
	braceAtNextLine bool
}

func (this *CConfigImpl) Init() {
	this.braceAtNextLine = true
}

func (this *CConfigImpl) String() string {
	return fmt.Sprintf("braceAtNextLine = %v", this.braceAtNextLine)
}

func (this *CConfigImpl) BraceAtNextLine() bool {
	return this.braceAtNextLine
}

func (this *CConfigImpl) SetBraceAtNextLine(val bool) {
	this.braceAtNextLine = val
}

func NewCConfig() *Config {
	ret := NewConfig()
	return ret
}
