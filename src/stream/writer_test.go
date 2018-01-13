package stream

import (
	"fmt"
	"testing"
)

func TestWriterWrite(t *testing.T) {
	prefix := "TestWriterWrite"

	w := &Writer{}
	w.Write("abc_%d", 101)
	wanted := "abc_101"

	if w.String() != wanted {
		t.Errorf("%s failed: w.String() = \"%s\", wanted = \"%s\"\n", prefix, w.String(), wanted)
	}
}

func TestWriterWriteln(t *testing.T) {
	prefix := "TestWriterWriteln"

	w := &Writer{}
	w.WriteByte(';')
	w.WriteString("cde")
	w.Writeln("abc_%d", 101)
	wanted := ";cdeabc_101" + fmt.Sprintln()

	if w.String() != wanted {
		t.Errorf("%s failed: w.String() = \"%s\", wanted = \"%s\"\n", prefix, w.String(), wanted)
	}
}
