package backends

import (
	_ "io"
)

/* Config implement config interface for common, c, cpp, golang
 */
type Config struct {
	// common config
	ConfigCommonImpl

	// config for c and cpp
	CConfigImpl

	// config for golang
	GolangConfigImpl
}

func NewConfig() *Config {
	ret := &Config{}
	ret.Init()
	return ret
}

func (this *Config) Init() {
	this.ConfigCommonImpl.Init()
	this.CConfigImpl.Init()
	this.GolangConfigImpl.Init()
}
