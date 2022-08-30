package uniqueString

func UniqueString(testWord string) bool {

	m := make(map[string]int)

	for key, v := range testWord {
		m[string(v)] = key
	}

	return len(m) != len(testWord)
}
