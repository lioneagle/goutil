package profile

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type ProfileEntry struct {
	Calls     int
	TotalTime time.Duration
}

type ProfileToken struct {
	Name  string
	start time.Time
}

func (this *ProfileToken) Exit() {
	g_profiler.mutex.Lock()
	defer g_profiler.mutex.Unlock()

	var entry *ProfileEntry
	var ok bool

	entry, ok = g_profiler.entries[this.Name]
	if !ok {
		entry = &ProfileEntry{}
		g_profiler.entries[this.Name] = entry
	}
	entry.Calls++
	entry.TotalTime += time.Since(this.start)
}

type ProfileResult struct {
	Name  string
	entry *ProfileEntry
}

func (this *ProfileResult) Avg() time.Duration {
	return this.entry.TotalTime / time.Duration(this.entry.Calls)
}

func (this ProfileResult) String() string {
	return fmt.Sprintf("%64s %6d, %20s, %20s", this.Name, this.entry.Calls,
		this.entry.TotalTime, this.Avg())
}

type Profiler struct {
	mutex   sync.Mutex
	entries map[string]*ProfileEntry
}

var g_profiler = Profiler{entries: make(map[string]*ProfileEntry)}

func (this *Profiler) Enter(name string) *ProfileToken {
	return &ProfileToken{name, time.Now()}
}

func (this *Profiler) Results() (ret []*ProfileResult) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for key, entry := range this.entries {
		ret = append(ret, &ProfileResult{key, entry})
	}

	return ret
}

func (this *Profiler) SortByName(isAcsending bool) []*ProfileResult {
	ret := this.Results()
	if isAcsending {
		sort.Slice(ret, func(i, j int) bool { return ret[i].Name < ret[j].Name })
	} else {
		sort.Slice(ret, func(i, j int) bool { return ret[i].Name >= ret[j].Name })
	}
	return ret
}

func (this *Profiler) SortByTotalTime(isAcsending bool) []*ProfileResult {
	ret := this.Results()
	if isAcsending {
		sort.Slice(ret, func(i, j int) bool { return ret[i].entry.TotalTime < ret[j].entry.TotalTime })
	} else {
		sort.Slice(ret, func(i, j int) bool { return ret[i].entry.TotalTime >= ret[j].entry.TotalTime })
	}
	return ret
}

func (this *Profiler) SortByAvgTime(isAcsending bool) []*ProfileResult {
	ret := this.Results()
	if isAcsending {
		sort.Slice(ret, func(i, j int) bool { return ret[i].Avg() < ret[j].Avg() })
	} else {
		sort.Slice(ret, func(i, j int) bool { return ret[i].Avg() >= ret[j].Avg() })
	}
	return ret
}

func (this *Profiler) String() (ret string) {
	results := this.SortByAvgTime(false)
	ret = fmt.Sprintf("%64s %6s, %20s, %20s\n", "Name", "Calls", "Total Time", "Average")
	for _, v := range results {
		ret += fmt.Sprintf("%s\n", v)
	}
	return ret
}

func Enter(name string) *ProfileToken {
	return g_profiler.Enter(name)
}

func Results(name string) []*ProfileResult {
	return g_profiler.Results()
}

func SortByName(isAcsending bool) []*ProfileResult {
	return g_profiler.SortByName(isAcsending)
}

func SortByTotalTime(isAcsending bool) []*ProfileResult {
	return g_profiler.SortByTotalTime(isAcsending)
}

func SortByAvgTime(isAcsending bool) []*ProfileResult {
	return g_profiler.SortByAvgTime(isAcsending)
}
