package linux

import (
	"os"
	"testing"
)

func TestGetProcCGroup(t *testing.T) {
	if _, err := ReadProcCGroup(os.Getpid()); err == nil {
		t.Error("GetProcCGroup should have failed on Darwin")
	}
}
