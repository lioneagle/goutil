package profile

import (
	//"fmt"
	"testing"
	"time"

	"github.com/lioneagle/goutil/src/test"
)

func TestProfilerSortByName(t *testing.T) {
	g_profiler.entries["A"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(1)}
	g_profiler.entries["C"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(2)}
	g_profiler.entries["D"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(3)}
	g_profiler.entries["B"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(4)}

	results := g_profiler.SortByName(true)

	test.EXPECT_EQ(t, results[0].Name, "A", "")
	test.EXPECT_EQ(t, results[1].Name, "B", "")
	test.EXPECT_EQ(t, results[2].Name, "C", "")
	test.EXPECT_EQ(t, results[3].Name, "D", "")

	test.EXPECT_EQ(t, results[0].entry.TotalTime, time.Duration(1), "")
	test.EXPECT_EQ(t, results[1].entry.TotalTime, time.Duration(4), "")
	test.EXPECT_EQ(t, results[2].entry.TotalTime, time.Duration(2), "")
	test.EXPECT_EQ(t, results[3].entry.TotalTime, time.Duration(3), "")

	results = g_profiler.SortByName(false)

	test.EXPECT_EQ(t, results[0].Name, "D", "")
	test.EXPECT_EQ(t, results[1].Name, "C", "")
	test.EXPECT_EQ(t, results[2].Name, "B", "")
	test.EXPECT_EQ(t, results[3].Name, "A", "")

	test.EXPECT_EQ(t, results[0].entry.TotalTime, time.Duration(3), "")
	test.EXPECT_EQ(t, results[1].entry.TotalTime, time.Duration(2), "")
	test.EXPECT_EQ(t, results[2].entry.TotalTime, time.Duration(4), "")
	test.EXPECT_EQ(t, results[3].entry.TotalTime, time.Duration(1), "")
}

func TestProfilerSortByTotalTime(t *testing.T) {
	g_profiler.entries["A"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(1)}
	g_profiler.entries["C"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(2)}
	g_profiler.entries["D"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(3)}
	g_profiler.entries["B"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(4)}

	results := g_profiler.SortByTotalTime(true)

	test.EXPECT_EQ(t, results[0].Name, "A", "")
	test.EXPECT_EQ(t, results[1].Name, "C", "")
	test.EXPECT_EQ(t, results[2].Name, "D", "")
	test.EXPECT_EQ(t, results[3].Name, "B", "")

	test.EXPECT_EQ(t, results[0].entry.TotalTime, time.Duration(1), "")
	test.EXPECT_EQ(t, results[1].entry.TotalTime, time.Duration(2), "")
	test.EXPECT_EQ(t, results[2].entry.TotalTime, time.Duration(3), "")
	test.EXPECT_EQ(t, results[3].entry.TotalTime, time.Duration(4), "")

	results = g_profiler.SortByTotalTime(false)

	test.EXPECT_EQ(t, results[0].Name, "B", "")
	test.EXPECT_EQ(t, results[1].Name, "D", "")
	test.EXPECT_EQ(t, results[2].Name, "C", "")
	test.EXPECT_EQ(t, results[3].Name, "A", "")

	test.EXPECT_EQ(t, results[0].entry.TotalTime, time.Duration(4), "")
	test.EXPECT_EQ(t, results[1].entry.TotalTime, time.Duration(3), "")
	test.EXPECT_EQ(t, results[2].entry.TotalTime, time.Duration(2), "")
	test.EXPECT_EQ(t, results[3].entry.TotalTime, time.Duration(1), "")
}

func TestProfilerSortByAvgTime(t *testing.T) {
	g_profiler.entries["A"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(1)}
	g_profiler.entries["C"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(2)}
	g_profiler.entries["D"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(3)}
	g_profiler.entries["B"] = &ProfileEntry{Calls: 1, TotalTime: time.Duration(4)}

	results := g_profiler.SortByAvgTime(true)

	test.EXPECT_EQ(t, results[0].Name, "A", "")
	test.EXPECT_EQ(t, results[1].Name, "C", "")
	test.EXPECT_EQ(t, results[2].Name, "D", "")
	test.EXPECT_EQ(t, results[3].Name, "B", "")

	test.EXPECT_EQ(t, results[0].entry.TotalTime, time.Duration(1), "")
	test.EXPECT_EQ(t, results[1].entry.TotalTime, time.Duration(2), "")
	test.EXPECT_EQ(t, results[2].entry.TotalTime, time.Duration(3), "")
	test.EXPECT_EQ(t, results[3].entry.TotalTime, time.Duration(4), "")

	results = g_profiler.SortByAvgTime(false)

	test.EXPECT_EQ(t, results[0].Name, "B", "")
	test.EXPECT_EQ(t, results[1].Name, "D", "")
	test.EXPECT_EQ(t, results[2].Name, "C", "")
	test.EXPECT_EQ(t, results[3].Name, "A", "")

	test.EXPECT_EQ(t, results[0].entry.TotalTime, time.Duration(4), "")
	test.EXPECT_EQ(t, results[1].entry.TotalTime, time.Duration(3), "")
	test.EXPECT_EQ(t, results[2].entry.TotalTime, time.Duration(2), "")
	test.EXPECT_EQ(t, results[3].entry.TotalTime, time.Duration(1), "")
}

func TestProfilerEnter(t *testing.T) {
	token := Enter("test1")

	test.EXPECT_NE(t, token, nil, "")
	test.EXPECT_EQ(t, token.Name, "test1", "")
}
