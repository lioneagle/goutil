package chars

import (
	"testing"
)

func TestStringsEqual(t *testing.T) {
	funcName := "TestStringsEqual"

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
		ret := StringsEqual(v.lhs, v.rhs)
		if ret != v.ret {
			t.Errorf("%s[%d] failed, ret = %v, wanted = %v\n", funcName, i, ret, v.ret)
		}
	}
}

func TestContains(t *testing.T) {
	funcName := "TestContains"

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
		ret := Contains(v.src, v.filter)
		if ret != v.ret {
			t.Errorf("%s[%d] failed, ret = %v, wanted = %v\n", funcName, i, ret, v.ret)
		}
	}
}

func TestFilter(t *testing.T) {
	funcName := "TestFilter"

	testdata := []struct {
		src    []string
		filter []string
		ret    []string
	}{
		{[]string{"abc434", "asd123", "890"}, []string{"abc", "123", "df"}, []string{"abc434", "asd123"}},
	}

	for i, v := range testdata {
		ret := Filter(v.src, v.filter)
		if !StringsEqual(ret, v.ret) {
			t.Errorf("%s[%d] failed, ret = %v, wanted = %v\n", funcName, i, ret, v.ret)
		}
	}
}

func TestFilterReverse(t *testing.T) {
	funcName := "TestFilterReverse"

	testdata := []struct {
		src    []string
		filter []string
		ret    []string
	}{
		{[]string{"abc434", "asd123", "890"}, []string{"abc", "123", "df"}, []string{"890"}},
	}

	for i, v := range testdata {
		ret := FilterReverse(v.src, v.filter)
		if !StringsEqual(ret, v.ret) {
			t.Errorf("%s[%d] failed, ret = %v, wanted = %v\n", funcName, i, ret, v.ret)
		}
	}
}
