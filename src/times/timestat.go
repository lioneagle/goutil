package times

import (
	"fmt"
	"io"
	"time"
)

type TimeStat struct {
	start time.Time
	end   time.Time
}

func NewTimeStat() *TimeStat {
	ret := &TimeStat{}
	ret.Start()
	ret.end = ret.start
	return ret
}

func (this *TimeStat) Start() {
	this.start = time.Now()
}

func (this *TimeStat) Stop() {
	this.end = time.Now()
}

func (this *TimeStat) Seconds() float64 {
	return this.end.Sub(this.start).Seconds()
}

func (this *TimeStat) Fprint(w io.Writer, prefix string) {
	fmt.Fprintf(w, "%s use time: %v", prefix, this.end.Sub(this.start))
	fmt.Fprintln(w)
}
