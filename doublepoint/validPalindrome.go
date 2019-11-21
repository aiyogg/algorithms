package main

/**
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
输入: "abca"
输出: True
解释: 你可以删除c字符。
*/

func validPalindrome(s string) bool {
	hasDeleted := false // 允许删除一个字符
	j := len(s) - 1
	for i := 0; i <= j; {
		if s[i] != s[j] {
			if hasDeleted {
				return false
			}
			i++
			if s[i] != s[j] {
				i--
				j--
				if s[i] != s[j] {
					return false
				}
				hasDeleted = true
			} else {
				hasDeleted = true
			}
		}
		i++
		j--
	}
	return true
}

// 方案二： 利用递归的实现
func validPalindromeRecursive(s string) bool {
	return helper([]byte(s), 0, len(s)-1, false)
}
func helper(bs []byte, lo, hi int, hasDelete bool) bool {
	for lo < hi {
		if bs[lo] != bs[hi] {
			if hasDelete {
				return false
			}
			return helper(bs, lo+1, hi, true) || // 删除 s[lo]
				helper(bs, lo, hi-1, true) // 删除 s[hi]
		}
		lo++
		hi--
	}
	return true
}

//func main() {
//	fmt.Printf("validPalindrome: %v \n", validPalindrome("abxcdcba"))
//	fmt.Printf("validPalindromeRecursive: %v", validPalindromeRecursive("abxcdcba"))
//}
