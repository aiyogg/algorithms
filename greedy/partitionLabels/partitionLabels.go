package partitionLabels

/**
763. 划分字母区间

字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。返回一个表示每个字符串片段的长度的列表。

示例 1:
输入: S = "ababcbacadefegdehijhklij"
输出: [9,7,8]
解释:
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。

1. S的长度在[1, 500]之间。
2. S只包含小写字母'a'到'z'。

这题有点难。。。
 */

func partitionLabels(S string) []int {
	maxIndex := [26]int{}
	for i, b := range S {
		maxIndex[b-'a'] = i
	}

	begin := 0
	end := maxIndex[S[begin]-'a']
	res := make([]int, 0, len(S))

	for i, b := range S {
		if i < end {
			// 在 S[:i+1] 和 S[i:] 中存在相同的字母 S[end]
			// 所以此时不能分隔，仅更新 end
			end = max(end, maxIndex[b-'a'])
			continue
		}

		// 此时 S[begin:i+1] 中的所有字母都不会出现在其他片段中
		// 可以进行分隔
		res = append(res, i-begin+1)
		begin = i + 1 // 从 i+1 处作为新片段的起始点
		if begin < len(S) {
			// 及时更新 end
			end = maxIndex[S[begin]-'a']
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}