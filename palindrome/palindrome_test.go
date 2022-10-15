package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	givenString := "Tact coa"

	result := IsPalindromePosible(givenString)

	if result != true {
		t.Errorf("IsPalindromePosible() = %t, want %t", result, true)
	}
}
