package uniqueString

import "testing"

func TestUniqueString(t *testing.T) {
	givenString := "faake"
	want := true

	if result := UniqueString(givenString); result != want {
		t.Errorf("UniqueString() = %t, want %t", result, want)
	}
}
