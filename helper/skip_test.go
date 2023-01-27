package helper

import (
	"runtime"
	"testing"
)

// skip testing on func

func TestSkip(t *testing.T) {
	// fmt.Println("runtime.GOOS", runtime.GOOS)
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on Windows")
	}
}
