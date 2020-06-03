package firstBadVersion

/**
278. 第一个错误的版本
示例:
给定 n = 5，并且 version = 4 是第一个错误的版本。
调用 isBadVersion(3) -> false
调用 isBadVersion(5) -> true
调用 isBadVersion(4) -> true

所以，4 是第一个错误的版本。
*/

func firstBadVersion(n int) int {
	l, h := 0, n
	for l < h {
		m := l + (h-l)/2
		if ok := isBadVersion(m); ok {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}
