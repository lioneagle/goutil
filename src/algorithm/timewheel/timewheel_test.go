package timewheel

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/lioneagle/goutil/src/test"
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
	testdata := []struct {
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

	statWanted := &TimeWheelStat{
		Add:           1,
		AddOk:         1,
		InternalAdd:   1,
		InternalAddOk: 1,
	}

	tick := int64(10)
	tw := NewTimeWheel(3, []int{60, 60, 24}, tick, 0, 1000)

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			//t.Parallel()

			tw.RemoveAll()
			tw.stat.Clear()
			interval := v.sceond + v.minute*60 + v.hour*3600
			interval *= tick

			ret := tw.Add(interval, nil, nil)

			test.ASSERT_TRUE(t, ret >= 0, "")
			test.EXPECT_EQ(t, tw.allocator.Chunks[ret].data.wheel, v.wheel, "")
			test.EXPECT_EQ(t, tw.allocator.Chunks[ret].data.slot, v.slot, "")
			test.EXPECT_EQ(t, tw.stat, *statWanted, "")
		})
	}
}

func TestTimeWheelBinaryAddOk(t *testing.T) {
	testdata := []struct {
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

	statWanted := &TimeWheelStat{
		Add:           1,
		AddOk:         1,
		InternalAdd:   1,
		InternalAddOk: 1,
	}

	tick := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, tick, 0, 1000)

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			//t.Parallel()

			tw.RemoveAll()
			tw.stat.Clear()
			interval := v.sceond + v.minute*64 + v.hour*64*64
			interval *= tick

			ret := tw.Add(interval, nil, nil)

			test.ASSERT_TRUE(t, ret >= 0, "")
			test.EXPECT_EQ(t, tw.allocator.Chunks[ret].data.wheel, v.wheel, "")
			test.EXPECT_EQ(t, tw.allocator.Chunks[ret].data.slot, v.slot, "")
			test.EXPECT_EQ(t, tw.stat, *statWanted, "")
		})
	}
}

func TestTimeWheelAddNOk(t *testing.T) {
	tw := NewTimeWheel(3, []int{60, 60, 24}, 1, 0, 1000)

	ret := tw.Add(0, nil, nil)
	test.EXPECT_EQ(t, ret, int32(-2), "")

	statWanted1 := &TimeWheelStat{Add: 1, Expire: 1, ExpireBeforeAdd: 1}
	test.EXPECT_EQ(t, tw.stat, *statWanted1, "")

	tw.stat.Clear()

	ret = tw.Add(60*60*24, nil, nil)
	test.EXPECT_EQ(t, ret, int32(-1), "")

	statWanted2 := &TimeWheelStat{Add: 1}
	test.EXPECT_EQ(t, tw.stat, *statWanted2, "")

	tw.RemoveAll()
}

func TestTimeWheelBinaryAddNOk(t *testing.T) {
	delta := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, delta, 0, 1000)
	ret := tw.Add(0, nil, func(interface{}) {})

	test.EXPECT_EQ(t, ret, int32(-2), "")

	statWanted1 := &TimeWheelStat{
		Add:             1,
		Expire:          1,
		ExpireBeforeAdd: 1,
		Post:            1,
	}
	test.EXPECT_EQ(t, tw.stat, *statWanted1, "")

	tw.stat.Clear()

	ret = tw.Add((1<<16)*delta, nil, nil)
	test.EXPECT_EQ(t, ret, int32(-1), "")

	statWanted2 := &TimeWheelStat{Add: 1}
	test.EXPECT_EQ(t, tw.stat, *statWanted2, "")

	tw.RemoveAll()
}

func TestTimeWheelBinaryRemoveOk1(t *testing.T) {
	tick := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, tick, 0, 1000)

	tm := tw.Add(10*tick, nil, nil)
	test.EXPECT_TRUE(t, tm >= 0, "")

	ret := tw.Remove(tm)
	test.EXPECT_TRUE(t, ret, "")

	statWanted1 := &TimeWheelStat{
		Add:              1,
		AddOk:            1,
		InternalAdd:      1,
		InternalAddOk:    1,
		Remove:           1,
		RemoveOk:         1,
		InternalRemove:   1,
		InternalRemoveOk: 1,
	}
	test.EXPECT_EQ(t, tw.stat, *statWanted1, "")

	tw.RemoveAll()

}

func TestTimeWheelBinaryRemoveOk2(t *testing.T) {
	tick := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, tick, 0, 1000)

	tm1 := tw.Add(10*tick, nil, nil)
	test.EXPECT_TRUE(t, tm1 >= 0, "")

	tm2 := tw.Add(20*tick, nil, nil)
	test.EXPECT_TRUE(t, tm2 >= 0, "")

	ret := tw.Remove(tm1)
	test.EXPECT_TRUE(t, ret, "")

	statWanted1 := &TimeWheelStat{
		Add:              2,
		AddOk:            2,
		InternalAdd:      2,
		InternalAddOk:    2,
		Remove:           1,
		RemoveOk:         1,
		InternalRemove:   1,
		InternalRemoveOk: 1,
	}
	test.EXPECT_EQ(t, tw.stat, *statWanted1, "")

	tw.RemoveAll()
}

func TestTimeWheelBinaryRemoveOk3(t *testing.T) {
	tick := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, tick, 0, 1000)

	tm1 := tw.Add(10*tick, nil, nil)
	test.EXPECT_TRUE(t, tm1 >= 0, "")

	tm2 := tw.Add(10*tick, nil, nil)
	test.EXPECT_TRUE(t, tm2 >= 0, "")

	ret := tw.Remove(tm1)
	test.EXPECT_TRUE(t, ret, "")

	statWanted1 := &TimeWheelStat{
		Add:              2,
		AddOk:            2,
		InternalAdd:      2,
		InternalAddOk:    2,
		Remove:           1,
		RemoveOk:         1,
		InternalRemove:   1,
		InternalRemoveOk: 1,
	}
	test.EXPECT_EQ(t, tw.stat, *statWanted1, "")

	tw.RemoveAll()
}

func TestTimeWheelBinaryRemoveOk4(t *testing.T) {
	tick := int64(10)

	tw := NewTimeWheelBinaryBits(3, []int{6, 6, 4}, tick, 0, 1000)

	tm1 := tw.Add(10*tick, nil, nil)
	test.EXPECT_TRUE(t, tm1 >= 0, "")

	tm2 := tw.Add(10*tick, nil, nil)
	test.EXPECT_TRUE(t, tm2 >= 0, "")

	ret := tw.Remove(tm2)
	test.EXPECT_TRUE(t, ret, "")

	statWanted1 := &TimeWheelStat{
		Add:              2,
		AddOk:            2,
		InternalAdd:      2,
		InternalAddOk:    2,
		Remove:           1,
		RemoveOk:         1,
		InternalRemove:   1,
		InternalRemoveOk: 1,
	}
	test.EXPECT_EQ(t, tw.stat, *statWanted1, "")

	tw.RemoveAll()
}

func TestTimeWheelStep1(t *testing.T) {
	tick := int64(10)
	start := int64(234)

	tw := NewTimeWheel(3, []int{60, 60, 24}, tick, start, 1000)

	tm1 := tw.Add(1*tick, nil, nil)
	test.EXPECT_TRUE(t, tm1 >= 0, "")

	tm2 := tw.Add(1*tick, nil, nil)
	test.EXPECT_TRUE(t, tm2 >= 0, "")

	tw.Step(start + tick)

	statWanted1 := &TimeWheelStat{
		Add:              2,
		AddOk:            2,
		InternalAdd:      2,
		InternalAddOk:    2,
		InternalRemove:   2,
		InternalRemoveOk: 2,
		Expire:           2,
		Step:             1,
	}
	test.EXPECT_EQ(t, tw.stat, *statWanted1, "")

	tw.RemoveAll()

}

func TestTimeWheelStep2(t *testing.T) {
	tick := int64(10)

	tw := NewTimeWheel(3, []int{100, 10, 3}, tick, 0, 1000)

	tm1 := tw.Add(5*tick, nil, func(interface{}) {})
	test.EXPECT_TRUE(t, tm1 >= 0, "")

	tw.Step(3 * tick)

	tm2 := tw.Add(3*tick, nil, func(interface{}) {})
	test.EXPECT_TRUE(t, tm2 >= 0, "")

	tw.Step(4 * tick)

	tm3 := tw.Add(250*tick, nil, func(interface{}) {})
	test.EXPECT_TRUE(t, tm3 >= 0, "")

	tw.Step(100 * tick)

	tw.Step(254 * tick)

	//fmt.Println("tw =", tw)

	statWanted1 := &TimeWheelStat{
		Add:              3,
		AddOk:            3,
		InternalAdd:      4,
		InternalAddOk:    4,
		InternalRemove:   4,
		InternalRemoveOk: 4,
		Expire:           3,
		Post:             3,
		Step:             4,
		MoveWheels:       2,
		MoveSlot:         1,
	}
	test.EXPECT_EQ(t, tw.stat, *statWanted1, "")

	tw.RemoveAll()

	test.EXPECT_TRUE(t, tw.size >= 0, "")

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

func BenchmarkGoTimer1(b *testing.B) {
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
