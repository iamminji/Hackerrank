package easy

import (
	"fmt"
)

// Complete the easy function below.
func plusMinus(arr []int32) {

	pos := 0.0
	neg := 0.0
	zero := 0.0

	for _, n := range arr {
		if n == 0 {
			zero += 1
		} else if n > 0 {
			pos += 1
		} else {
			neg += 1
		}
	}

	size := float64(len(arr))

	fmt.Printf("%.6f\n", pos/size)
	fmt.Printf("%.6f\n", neg/size)
	fmt.Printf("%.6f\n", zero/size)

}
