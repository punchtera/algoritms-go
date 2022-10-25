package oneAway

func isStringOneOperationAway(testString1 string, testString2 string) bool {
	charsMapDiff := make(map[rune]int)
	for _, char := range testString1 {
		if times, ok := charsMapDiff[char]; ok {
			charsMapDiff[char] = times + 1
		} else {
			charsMapDiff[char] = 1
		}
	}

	for _, char2 := range testString2 {
		if times, ok := charsMapDiff[char2]; ok && times == 1 {
			delete(charsMapDiff, char2)
		} else if ok && times > 1 {
			charsMapDiff[char2] = times - 1
		} else {
			charsMapDiff[char2] = 1
		}
	}

	if mapLength := len(charsMapDiff); mapLength <= 2 {
		// its a single operation or 2 different characters, a replce
		return true
	}

	return false
}
