package backends

type GolangConfig interface {
	ConfigCommon

	PackageName() string
	SetPackageName(val string)
}

type GolangConfigImpl struct {
	packageName string
}

func (this *GolangConfigImpl) Init() {
	this.packageName = "abnf"
}

func (this *GolangConfigImpl) PackageName() string {
	return this.packageName
}

func (this *GolangConfigImpl) SetPackageName(val string) {
	this.packageName = val
}

func NewGolangConfig() *Config {
	ret := NewConfig()
	ret.Indent().Switch = 0
	return ret
}
