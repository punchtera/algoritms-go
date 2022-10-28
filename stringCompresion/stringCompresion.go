package stringCompresion

import (
	"strconv"
)

func CompressString(stringOne string) string {
	numberOfTimes := 1
	result := ""
	stringOneLength := len(stringOne)
	for i := 0; i <= stringOneLength-1; i++ {
		currentChar := stringOne[i]
		var afterChar byte

		if i == stringOneLength-1 {
			afterChar = byte(0)
		} else {
			afterChar = stringOne[i+1]
		}

		if currentChar == afterChar {
			numberOfTimes = numberOfTimes + 1
		} else {
			result += string(currentChar) + strconv.Itoa(numberOfTimes)
			numberOfTimes = 1
		}

	}
	return result
}
