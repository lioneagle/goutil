package mem

import (
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestArenaAllocatorStatPrint(t *testing.T) {
	stat := &ArenaAllocatorStat{}
	test.EXPECT_EQ(t, stat.String(), "", "")

	stat.allocNum = 1
	test.EXPECT_EQ(t, stat.String(), "alloc num: 1\n", "")

	stat.allocNum = 1
	stat.allocNumOk = 2
	stat.freeAllNum = 3
	stat.freePartNum = 4
	test.EXPECT_EQ(t, stat.AllocNum(), StatNumber(1), "")
	test.EXPECT_EQ(t, stat.AllocNumOk(), StatNumber(2), "")
	test.EXPECT_EQ(t, stat.FreeAllNum(), StatNumber(3), "")
	test.EXPECT_EQ(t, stat.FreePartNum(), StatNumber(4), "")
	test.EXPECT_EQ(t, stat.String(), "alloc num: 1\nalloc num ok: 2\nfree all num: 3\nfree part num: 4\n", "")

	stat.Init()
	test.EXPECT_EQ(t, stat.String(), "", "")
}
