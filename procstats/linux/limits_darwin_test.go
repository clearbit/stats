package linux

import (
	"os"
	"testing"
)

func TestGetProcLimits(t *testing.T) {
	if _, err := ReadProcLimits(os.Getpid()); err == nil {
		t.Error("GetProcLimits should have failed on Darwin")
	}
}
