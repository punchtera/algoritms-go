package oneAway

import "testing"

func TestIsStringOneAway(t *testing.T) {

	givenTestString1 := "pale"
	givenTestString2 := "pales"
	result := isStringOneOperationAway(givenTestString1, givenTestString2)

	if result != true {
		t.Errorf("IsPalindromePosible() = %t, want %t", result, true)
	}
}
