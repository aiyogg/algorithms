package main

import "fmt"

/*
有序数组的两数之和
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/

func twoSum(numbers []int, target int) []int {
	j := len(numbers) - 1
	for i := 0; i < j; {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			return []int{i, j}
		}
	}
	return []int{}
}

func main() {
	fmt.Printf("towSum: %v", twoSum([]int{2, 7, 11, 15}, 18))
}
