package linux

import (
	"os"
	"testing"
)

func TestGetProcStat(t *testing.T) {
	if _, err := ReadProcStat(os.Getpid()); err == nil {
		t.Error("GetProcStat should have failed on Darwin")
	}
}
