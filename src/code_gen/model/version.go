package model

import (
	"fmt"
	//"io"
	"time"
)

const (
	VERSION              = "v0.1.0"
	TIMESTAMP_FMT string = "2006-01-02 15:04:05.000"
)

func GenerateVersion() string {
	return fmt.Sprintf("---------------- generated by abnf %s %s ----------------", VERSION, time.Now().Format(TIMESTAMP_FMT))
}