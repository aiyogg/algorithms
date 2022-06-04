// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Given a string s, find the length of the longest substring without repeating characters.
// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	// 记录字符是否出现过
	m := map[byte]int{}
	n := len(s)
	j, res := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 删除已经出现过的字符
			delete(m, s[i-1])
		}
		for j+1 < n && m[s[j+1]] == 0 {
			// 字符未出现过，移动右指针，记录字符出现过
			m[s[j+1]]++
			j++
		}
		// 记录最大长度
		res = max(res, j-i+1)
	}

	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
