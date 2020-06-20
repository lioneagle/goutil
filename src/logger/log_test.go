package logger

import (
	"testing"
)

type discardIo struct{}

func (this *discardIo) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func Benchmark10Fields(b *testing.B) {
	b.StopTimer()
	log := NewLogger(&discardIo{})
	log.SetLevel(INFO)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		log.Infof("Ten fields, passed at the log site. %d,%d,%d,%d,%d,%d,%d,%d,%d,%d",
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}

}
