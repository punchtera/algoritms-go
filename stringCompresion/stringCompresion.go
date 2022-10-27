package stringCompresion

import (
	"fmt"
	"strconv"
)

func CompressString(stringOne string) string {
	numberOfTimes := 1
	result := ""
	stringOneLength := len(stringOne)
	for i := 0; i < stringOneLength-1; i++ {
		currentChar := stringOne[i]
		afterChar := stringOne[i+1]

		// var afterChar byte
		// if stringOneLength <= i {
		// 	afterChar = stringOne[i+1]
		// } else {
		// 	afterChar = byte(0)
		// }

		if currentChar == afterChar {
			numberOfTimes = numberOfTimes + 1
		} else {
			result += string(currentChar) + strconv.Itoa(numberOfTimes)
			numberOfTimes = 1
		}

	}
	fmt.Printf("this is the result :  %s \n", result)
	return " "
}
