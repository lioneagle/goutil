package chars

import (
	"fmt"
	"io"
	"runtime"
)

type Return struct {
	returnString string
}

func (this *Return) Init() {
	if runtime.GOOS == "windows" {
		this.returnString = "\r\n"
	} else if runtime.GOOS == "windows" {
		this.returnString = "\r"
	} else {
		this.returnString = "\n"
	}
}

func (this *Return) SetReturnString(ret string) {
	this.returnString = ret
}

func PrintReturn(w io.Writer) {
	fmt.Fprint(w, g_return.returnString)
}
