package draw

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestPath(t *testing.T) {
	rootPath := SystemInstace.GetRootPath()
	fmt.Println("rootPath =", rootPath)

	skinPath := SystemInstace.GetSkinPath()
	xmlPath := SystemInstace.GetXmlPath()

	test.EXPECT_EQ(t, skinPath, rootPath, "")

	fmt.Println("xmlPath =", xmlPath)
}
