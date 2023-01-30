package linux

import (
	"os"
	"testing"
)

func TestGetOpenFileCount(t *testing.T) {
	if _, err := ReadOpenFileCount(os.Getpid()); err == nil {
		t.Error("GetOpenFileCount should have failed on Darwin")
	}
}
