package chars

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestStringsEqual(t *testing.T) {
	testdata := []struct {
		lhs []string
		rhs []string
		ret bool
	}{
		{[]string{"abc434", "asd123", "890"}, []string{"abc", "123", "df"}, false},
		{[]string{"abc434", "asd123", "890"}, []string{"abc434", "asd123", "890"}, true},

		{[]string{"abc434", "asd123", "890"}, []string{"abc434", "890"}, false},
		{[]string{"abc434", "asd123", "890"}, []string{}, false},
		{[]string{""}, []string{"abc434", "890"}, false},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			ret := StringsEqual(v.lhs, v.rhs)
			test.EXPECT_EQ(t, ret, v.ret, "")
		})
	}
}

func TestContains(t *testing.T) {
	testdata := []struct {
		src    string
		filter []string
		ret    bool
	}{
		{"abc1234", []string{"abc", "123", "df"}, true},
		{"abc1234", []string{"78", "ac", "df"}, false},

		{"abc1234", []string{""}, true},
		{"abc1234", []string{}, false},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			ret := Contains(v.src, v.filter)
			test.EXPECT_EQ(t, ret, v.ret, "")
		})
	}
}

func TestFilter(t *testing.T) {
	testdata := []struct {
		src    []string
		filter []string
		ret    []string
	}{
		{[]string{"abc434", "asd123", "890"}, []string{"abc", "123", "df"}, []string{"abc434", "asd123"}},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			ret := Filter(v.src, v.filter)
			test.EXPECT_TRUE(t, StringsEqual(ret, v.ret), "ret = %v, wanted = %v", ret, v.ret)
		})
	}
}

func TestFilterReverse(t *testing.T) {
	testdata := []struct {
		src    []string
		filter []string
		ret    []string
	}{
		{[]string{"abc434", "asd123", "890"}, []string{"abc", "123", "df"}, []string{"890"}},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			ret := FilterReverse(v.src, v.filter)
			test.EXPECT_TRUE(t, StringsEqual(ret, v.ret), "ret = %v, wanted = %v", ret, v.ret)
		})
	}
}

func TestPackSpace(t *testing.T) {
	testdata := []struct {
		src string
		ret string
	}{
		{"", ""},
		{" anb", "anb"},
		{" anb ", "anb"},
		{" an b ", "an b"},
		{" an  b ", "an b"},
		{" an \tb ", "an b"},
		{" an\t\tb ", "an b"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			ret := StringPackSpace(v.src)
			test.EXPECT_EQ(t, ret, v.ret, "")
		})
	}
}
