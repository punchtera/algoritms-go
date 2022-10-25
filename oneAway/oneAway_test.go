package oneAway

import (
	"fmt"
	"testing"
)

func TestIsStringOneAway(t *testing.T) {

	var tests = []struct {
		firstString, secondString string
		want                      bool
	}{
		{"palee", "paleb", true},
		{"ok", "oki", true},
		{"juan", "kuam", false},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%s,%s", tt.firstString, tt.secondString)
		t.Run(testname, func(t *testing.T) {
			result := isStringOneOperationAway(tt.firstString, tt.secondString)
			if result != tt.want {
				t.Errorf("IsPalindromePosible() = %t, want %t", result, true)
			}
		})
	}
}
