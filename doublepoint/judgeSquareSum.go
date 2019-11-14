package main

import (
	"fmt"
	"math"
)

/**
两数平方和
Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5
*/
func judgeSquareSum(target int) bool {
	j := int(math.Sqrt(float64(target)))
	for i := 0; i <= j; {
		if i*i+j*j < target {
			i++
		} else if i*i+j*j > target {
			j--
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Printf("judgeSquareSum: %v", judgeSquareSum(8))
}
