package main

import "testing"

/**
两数平方和
Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5
*/
func TestJudgeSquareSum(t *testing.T) {
	if !judgeSquareSum(5) {
		t.Errorf("judgeSquareSum test failed. %d", 5)
	}
}
