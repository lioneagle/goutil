package algorithm

import (
	//"fmt"
	"reflect"
	"testing"
	"time"
)

type record struct {
	t1 time.Time
	t2 time.Time
}

func checkStat(t *testing.T, stat, wanted *TimeWheelStat, prefix string, index int) {
	lhs := reflect.ValueOf(*stat)
	rhs := reflect.ValueOf(*wanted)
	typeOf := lhs.Type()
	for i := 0; i < lhs.NumField(); i++ {
		f1 := lhs.Field(i)
		f2 := rhs.Field(i)
		if f1.Interface() != f2.Interface() {
			t.Errorf("%s[%d] failed: stat.%s = %v, wanted = %v\n", prefix, index, typeOf.Field(i).Name, f1.Interface(), f2.Interface())
		}
	}
}

func TestTimeWheelAddOk(t *testing.T) {
	wanted := []struct {
		sceond int64
		minute int64
		hour   int64
		wheel  int32
		slot   int32
	}{
		{1, 0, 0, 0, 1},
		{59, 0, 0, 0, 59},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 2, 1},
		{0, 10, 1, 2, 1},
		{32, 11, 1, 2, 1},
		{59, 59, 23, 2, 23},
	}

	statWanted := &TimeWheelStat{Add: 1, AddOk: 1, InternalAdd: 1, InternalAddOk: 1}
	prefix := "TestTimeWheelAddOk"

	tick := int64(10)

	tw := NewTimeWheel(3, []int{60, 60, 24}, tick, 0, 1000)

	for i, v := range wanted {
		tw.RemoveAll()
		tw.stat.Clear()
		interval := v.sceond + v.minute*60 + v.hour*3600
		interval *= tick

		ret := tw.Add(interval, nil, nil)

		if ret < 0 {
			t.Errorf("%s[%d] failed: ret = %d, wanted = 0\n", prefix, i, ret)
			continue
		}

		if tw.allocator.Chunks[ret].data.wheel != v.wheel {
			t.Errorf("%s[%d] failed: wheel = %d, wanted = %d\n", prefix, i, tw.allocator.Chunks[ret].data.wheel, v.wheel)
		}

		if tw.allocator.Chunks[ret].data.slot != v.slot {
			t.Errorf("%s[%d] failed: slot = %d, wanted = %d\n", prefix, i, tw.allocator.Chunks[ret].data.slot, v.slot)
		}

		checkStat(t, &tw.stat, statWanted, prefix, i)
	}

}

func TestTimeWheelBinaryAddOk(t *testing.T) {
	wanted := []struct {
		sceond int64
		minute int64
		hour   int64
		wheel  int32
		slot   int32
	}{
		{1, 0, 0, 0, 1},
		{63, 0, 0, 0, 63},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 2, 1},
		{0, 10, 1, 2, 1},
		{32, 11, 1, 2, 1},
		{63, 63, 15, 2, 15},
	}

	statWanted := &TimeWheelStat{Add: 1, AddOk: 1, InternalAdd: 1, InternalAddOk: 1}
	prefix := "TestTimeWheelBinaryAddOk"

	tick := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, tick, 0, 1000)

	for i, v := range wanted {
		tw.RemoveAll()
		tw.stat.Clear()
		interval := v.sceond + v.minute*64 + v.hour*64*64
		interval *= tick

		ret := tw.Add(interval, nil, nil)

		if ret < 0 {
			t.Errorf("%s[%d] failed: ret = %d, wanted = 0\n", prefix, i, ret)
			continue
		}

		if tw.allocator.Chunks[ret].data.wheel != v.wheel {
			t.Errorf("%s[%d] failed: wheel = %d, wanted = %d\n", prefix, i, tw.allocator.Chunks[ret].data.wheel, v.wheel)
		}

		if tw.allocator.Chunks[ret].data.slot != v.slot {
			t.Errorf("%s[%d] failed: slot = %d, wanted = %d\n", prefix, i, tw.allocator.Chunks[ret].data.slot, v.slot)
		}

		checkStat(t, &tw.stat, statWanted, prefix, i)
	}

}

func TestTimeWheelAddNOk(t *testing.T) {
	prefix := "TestTimeWheelBinaryAddNOk"
	tw := NewTimeWheel(3, []int{60, 60, 24}, 1, 0, 1000)

	ret := tw.Add(0, nil, nil)
	if ret != -2 {
		t.Errorf("%s failed: ret = %d, wanted = -2\n", prefix, ret)
	}

	statWanted1 := &TimeWheelStat{Add: 1, Expire: 1, ExpireBeforeAdd: 1}
	checkStat(t, &tw.stat, statWanted1, prefix, 0)

	tw.stat.Clear()

	ret = tw.Add(60*60*24, nil, nil)
	if ret != -1 {
		t.Errorf("%s failed: ret = %d, wanted = -1\n", prefix, ret)
	}

	statWanted2 := &TimeWheelStat{Add: 1}
	checkStat(t, &tw.stat, statWanted2, prefix, 1)

	tw.RemoveAll()
}

func TestTimeWheelBinaryAddNOk(t *testing.T) {
	prefix := "TestTimeWheelAddNOk"
	delta := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, delta, 0, 1000)
	ret := tw.Add(0, nil, func(interface{}) {})

	if ret != -2 {
		t.Errorf("%s failed: ret = %d, wanted = -2\n", prefix, ret)
	}

	statWanted1 := &TimeWheelStat{Add: 1, Expire: 1, ExpireBeforeAdd: 1, Post: 1}
	checkStat(t, &tw.stat, statWanted1, prefix, 0)

	tw.stat.Clear()

	ret = tw.Add((1<<16)*delta, nil, nil)
	if ret != -1 {
		t.Errorf("%s failed: ret = %d, wanted = -1\n", prefix, ret)
	}

	statWanted2 := &TimeWheelStat{Add: 1}
	checkStat(t, &tw.stat, statWanted2, prefix, 1)

	tw.RemoveAll()
}

func TestTimeWheelBinaryRemoveOk1(t *testing.T) {
	prefix := "TestTimeWheelBinaryRemoveOk"

	tick := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, tick, 0, 1000)

	tm := tw.Add(10*tick, nil, nil)
	if tm < 0 {
		t.Errorf("%s failed: tm = %d, wanted >=0\n", prefix, tm)
	}

	ret := tw.Remove(tm)
	if !ret {
		t.Errorf("%s failed: ret = false, wanted = true\n", prefix)
	}

	statWanted1 := &TimeWheelStat{Add: 1, AddOk: 1, InternalAdd: 1, InternalAddOk: 1, Remove: 1, RemoveOk: 1, InternalRemove: 1, InternalRemoveOk: 1}
	checkStat(t, &tw.stat, statWanted1, prefix, 0)

	tw.RemoveAll()

}

func TestTimeWheelBinaryRemoveOk2(t *testing.T) {
	prefix := "TestTimeWheelBinaryRemoveOk"

	tick := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, tick, 0, 1000)

	tm1 := tw.Add(10*tick, nil, nil)
	if tm1 < 0 {
		t.Errorf("%s failed: tm1 = %d, wanted >=0\n", prefix, tm1)
	}

	tm2 := tw.Add(20*tick, nil, nil)
	if tm2 < 0 {
		t.Errorf("%s failed: tm2 = %d, wanted >=0\n", prefix, tm2)
	}

	ret := tw.Remove(tm1)
	if !ret {
		t.Errorf("%s failed: ret = false, wanted = true\n", prefix)
	}

	statWanted1 := &TimeWheelStat{Add: 2, AddOk: 2, InternalAdd: 2, InternalAddOk: 2, Remove: 1, RemoveOk: 1, InternalRemove: 1, InternalRemoveOk: 1}
	checkStat(t, &tw.stat, statWanted1, prefix, 0)

	tw.RemoveAll()
}

func TestTimeWheelBinaryRemoveOk3(t *testing.T) {
	prefix := "TestTimeWheelBinaryRemoveOk"

	tick := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, tick, 0, 1000)

	tm1 := tw.Add(10*tick, nil, nil)
	if tm1 < 0 {
		t.Errorf("%s failed: tm1 = %d, wanted >=0\n", prefix, tm1)
	}

	tm2 := tw.Add(10*tick, nil, nil)
	if tm2 < 0 {
		t.Errorf("%s failed: tm2 = %d, wanted >=0\n", prefix, tm2)
	}

	ret := tw.Remove(tm1)
	if !ret {
		t.Errorf("%s failed: ret = false, wanted = true\n", prefix)
	}

	statWanted1 := &TimeWheelStat{Add: 2, AddOk: 2, InternalAdd: 2, InternalAddOk: 2, Remove: 1, RemoveOk: 1, InternalRemove: 1, InternalRemoveOk: 1}
	checkStat(t, &tw.stat, statWanted1, prefix, 0)

	tw.RemoveAll()
}

func TestTimeWheelBinaryRemoveOk4(t *testing.T) {
	prefix := "TestTimeWheelBinaryRemoveOk"

	tick := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, tick, 0, 1000)

	tm1 := tw.Add(10*tick, nil, nil)
	if tm1 < 0 {
		t.Errorf("%s failed: tm1 = %d, wanted >=0\n", prefix, tm1)
	}

	tm2 := tw.Add(10*tick, nil, nil)
	if tm2 < 0 {
		t.Errorf("%s failed: tm2 = %d, wanted >=0\n", prefix, tm2)
	}

	ret := tw.Remove(tm2)
	if !ret {
		t.Errorf("%s failed: ret = false, wanted = true\n", prefix)
	}

	statWanted1 := &TimeWheelStat{Add: 2, AddOk: 2, InternalAdd: 2, InternalAddOk: 2, Remove: 1, RemoveOk: 1, InternalRemove: 1, InternalRemoveOk: 1}
	checkStat(t, &tw.stat, statWanted1, prefix, 0)

	tw.RemoveAll()
}

func TestTimeWheelStep1(t *testing.T) {
	prefix := "TestTimeWheelStep1"

	tick := int64(10)
	start := int64(234)

	tw := NewTimeWheel(3, []int{60, 60, 24}, tick, start, 1000)

	tm1 := tw.Add(1*tick, nil, nil)
	if tm1 < 0 {
		t.Errorf("%s failed: tm = %d, wanted >=0\n", prefix, tm1)
	}

	tm2 := tw.Add(1*tick, nil, nil)
	if tm2 < 0 {
		t.Errorf("%s failed: tm = %d, wanted >=0\n", prefix, tm2)
	}

	tw.Step(start + tick)

	statWanted1 := &TimeWheelStat{Add: 2, AddOk: 2, InternalAdd: 2, InternalAddOk: 2, Expire: 2, Step: 1}
	checkStat(t, &tw.stat, statWanted1, prefix, 0)

	tw.RemoveAll()

}

func TestTimeWheelStep2(t *testing.T) {
	prefix := "TestTimeWheelStep1"

	tick := int64(10)

	tw := NewTimeWheel(3, []int{10, 10, 3}, tick, 0, 1000)

	tm1 := tw.Add(9*tick, nil, func(interface{}) {})
	if tm1 < 0 {
		t.Errorf("%s failed: tm = %d, wanted >=0\n", prefix, tm1)
	}

	tm2 := tw.Add(19*tick, nil, func(interface{}) {})
	if tm2 < 0 {
		t.Errorf("%s failed: tm = %d, wanted >=0\n", prefix, tm2)
	}

	tw.Step(19 * tick)

	statWanted1 := &TimeWheelStat{
		Add:              2,
		AddOk:            2,
		InternalAdd:      3,
		InternalAddOk:    3,
		InternalRemove:   1,
		InternalRemoveOk: 1,
		Expire:           2,
		Post:             2,
		Step:             1,
		MoveWheels:       1,
		MoveSlot:         1,
	}
	checkStat(t, &tw.stat, statWanted1, prefix, 0)

	tw.RemoveAll()

}

func TestGoTimer1(t *testing.T) {
	//fmt.Println("time.Now() =", time.Now())
	//fmt.Println("time.Second =", time.Second)
	ticker := time.NewTicker(1000000 * 1)

	times := make([]*record, 0)
	for i := 0; i < 10; i++ {
		t1 := <-ticker.C
		//fmt.Println(time.String())
		times = append(times, &record{t1, time.Now()})
	}
	ticker.Stop()

	/*for i, v := range times {
		fmt.Printf("[%d]: t1 = %s, t2 = %s\n", i, v.t1.String(), v.t2.String())
	}*/
}

func TestGoTimer2(t *testing.T) {
	/*start := time.Now()
	for i := 0; i < 100000000; i++ {
		time.Now()
	}
	end := time.Now()

	fmt.Printf("start: %s\n", start.String())
	fmt.Printf("end: %s\n", end.String())
	fmt.Printf("ns/op: %s\n", end.Sub(start).String())*/

}

func BenchmarkGoTimer(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ticker := time.NewTicker(100000000000)
		ticker.Stop()
	}
}

func BenchmarkGoTimer2(b *testing.B) {
	ticker := time.NewTicker(1000000 * 2)
	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-ticker.C
	}
	b.StopTimer()
	ticker.Stop()
}

func BenchmarkGoTimer3(b *testing.B) {
	ticker := time.NewTicker(1)
	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		<-ticker.C
	}
	b.StopTimer()
	ticker.Stop()
}

func BenchmarkGoTimer4(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		time.Now()
	}
	b.StopTimer()
}

func BenchmarkTimeWheelAddRemove1(b *testing.B) {
	b.StopTimer()
	tw := NewTimeWheel(5, []int{256, 64, 64, 64, 64}, 1, 0, 10000)
	//tw := NewTimeWheel(8, []int{64, 64, 64, 64, 64, 64, 64, 64}, 1, 10000)
	//tw := NewTimeWheel(4, []int{256, 256, 256, 256}, 1, 10000)
	//tw := NewTimeWheel(3, []int{1 << 11, 1 << 11, 1 << 10}, 1, 10000)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		e := tw.Add(20000000, 100, nil)
		//e := tw.Add(1, 100, nil)
		tw.Remove(e)
	}
}

func BenchmarkTimeWheelAddRemove2(b *testing.B) {
	b.StopTimer()
	tw := NewTimeWheel(5, []int{256, 64, 64, 64, 64}, 1, 0, 10000)
	//tw := NewTimeWheel(8, []int{64, 64, 64, 64, 64, 64, 64, 64}, 1, 10000)
	//tw := NewTimeWheel(4, []int{256, 256, 256, 256}, 1, 10000)
	//tw := NewTimeWheel(3, []int{1 << 11, 1 << 11, 1 << 10}, 1, 10000)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		e := tw.Add(1, 100, nil)
		tw.Remove(e)
	}
}

func BenchmarkTimeWheelBinaryAddRemove1(b *testing.B) {
	b.StopTimer()
	tw := NewTimeWheelBinaryBits(5, []int{8, 6, 6, 6, 6}, 1, 0, 10000)
	//tw := NewTimeWheel(8, []int{4, 4, 4, 4, 4, 4, 4, 4}, 1, 10000)
	//tw := NewTimeWheel(4, []int{8, 8, 8, 8}, 1, 10000)
	//tw := NewTimeWheel(3, []int{11, 11, 10}, 1, 10000)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		e := tw.Add(20000000, nil, nil)
		//e := tw.Add(1, 100, nil)
		tw.Remove(e)
	}
}

func BenchmarkTimeWheelBinaryAddRemove2(b *testing.B) {
	b.StopTimer()
	tw := NewTimeWheelBinaryBits(5, []int{8, 6, 6, 6, 6}, 1, 0, 10000)
	//tw := NewTimeWheel(8, []int{4, 4, 4, 4, 4, 4, 4, 4}, 1, 10000)
	//tw := NewTimeWheel(4, []int{8, 8, 8, 8}, 1, 10000)
	//tw := NewTimeWheel(3, []int{11, 11, 10}, 1, 10000)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		e := tw.Add(1, nil, nil)
		tw.Remove(e)
	}
}
