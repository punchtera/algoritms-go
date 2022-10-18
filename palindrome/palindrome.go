package palindrome

import (
	"strings"
)

func IsPalindromePosible(testString string) bool {
	testStringNoSpaces := strings.ToLower(strings.Replace(testString, " ", "", -1))
	testCharsMap := make(map[rune]int)

	for _, char := range testStringNoSpaces {
		if testCharsMap[char] == 0 {
			testCharsMap[char] = 1
		} else {
			testCharsMap[char] = testCharsMap[char] + 1
		}
	}

	numberOfOdds := 0
	for _, value := range testCharsMap {
		if value%2 != 0 {
			numberOfOdds += 1
		}
	}

	if numberOfOdds == 0 || numberOfOdds == 1 {
		return true
	}
	return false
}
