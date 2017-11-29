package algorithm

import (
	//"fmt"
	"testing"
	"time"
)

type record struct {
	t1 time.Time
	t2 time.Time
}

func TestGoTimer(t *testing.T) {
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

func BenchmarkTimeWheelAddRemove(b *testing.B) {
	b.StopTimer()
	tw := NewTimeWheel(5, []int{256, 64, 64, 64, 64}, 1, 10000)
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
