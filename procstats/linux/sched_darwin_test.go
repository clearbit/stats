package linux

import (
	"os"
	"testing"
)

func TestGetProcSched(t *testing.T) {
	if _, err := ReadProcSched(os.Getpid()); err == nil {
		t.Error("GetProcSched should have failed on Darwin")
	}
}
