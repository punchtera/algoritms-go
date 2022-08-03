package cubeSum

import (
	"fmt"
	"math"
)

var fmtPrintf = fmt.Printf

func CubeSumO3() {
	limit := 50

	for indexA := 1; indexA <= limit; indexA++ {
		for indexB := 1; indexB <= limit; indexB++ {
			for indexC := 1; indexC <= limit; indexC++ {
				indexD := math.Round(math.Pow(math.Pow(float64(indexA), 3)+math.Pow(float64(indexB), 3)-math.Pow(float64(indexC), 3), 1/3))
				beforeEqual := math.Pow(float64(indexA), 3) + math.Pow(float64(indexB), 3)
				afterEqual := math.Pow(float64(indexC), 3) + math.Pow(float64(indexD), 3)

				if beforeEqual == afterEqual && indexD > 0 && indexD < float64(limit) {
					fmtPrintf("The results is a=%v, b=%v, c=%v, d=%v", indexA, indexB, indexC, indexD)
				}
			}
		}
	}
}
