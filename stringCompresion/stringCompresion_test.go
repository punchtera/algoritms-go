package stringCompresion

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	givenString := "taaacccccccoaa"
	expectedString := "t1a3c7o1a2"
	result := CompressString(givenString)

	if result != expectedString {
		t.Errorf("CompressString() = %s, want %s", result, expectedString)
	}
}
