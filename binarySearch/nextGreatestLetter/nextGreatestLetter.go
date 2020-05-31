package nextGreatestLetter

/**
744. 寻找比目标字母大的最小字母
给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。

在比较时，字母是依序循环出现的。举个例子：
如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'
*/

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	l, r := 0, n-1
	for l <= r {
		m := l + (r-l)/2
		if letters[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if l < n {
		return letters[l]
	}
	return letters[0]
}
