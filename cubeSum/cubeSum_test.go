package cubeSum

import "testing"

func TestCubeSum(t *testing.T) {
	want := "No way"
	if got := CubeSum(); got != want {
		t.Errorf("CubeSum() = %q, want %q", got, want)
	}
}
