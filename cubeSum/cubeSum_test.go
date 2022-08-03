package cubeSum

import (
	"testing"
)

func TestCubeSumO3(t *testing.T) {
	// Save current function and restore at the end:
	old := fmtPrintf
	counter := 0
	defer func() { fmtPrintf = old }()

	fmtPrintf = func(format string, a ...any) (n int, err error) {
		counter = counter + 1
		return
	}

	// check for working implementation
	CubeSumO3()

	if counter != 101 {
		t.Errorf("CubeSumO3() if not working as expected %v", counter)
	}
}
