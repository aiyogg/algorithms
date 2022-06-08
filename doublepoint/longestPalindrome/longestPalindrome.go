/*
5. 最长回文子串 Medium
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

链接：https://leetcode.cn/problems/longest-palindromic-substring
*/

package longestpalindrome

// 1.暴力解法
// 穷举所有的子串，记录最长的子串和对应的起始位置
func longestPalindrome1(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	maxLen, start := 1, 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if validPalindromic(s, i, j) && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}

// 验证字符串是否为回文
func validPalindromic(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
