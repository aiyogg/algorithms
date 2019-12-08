package frequencySort

import "sort"

/**
451. Sort Characters By Frequency
Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
*/

func frequencySort(s string) string {
	// r 记录各个字符出现的次数
	r := ['z' + 1]int{}
	for _, c := range s {
		r[c]++
	}

	ss := make([]string, 0, len(s))
	for i := range r {
		if r[i] == 0 {
			continue
		}
		ss = append(ss, makeString(byte(i), r[i]))
	}
	sort.Sort(segments(ss))

	res := ""
	for _, s := range ss {
		res += s
	}
	return res
}

// makeString 构造 n 个 b 组成的字符串
func makeString(b byte, n int) string {
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = b
	}
	return string(bytes)
}

// segments 实现了 sort.Interface 接口
type segments []string

func (s segments) Len() int {
	return len(s)
}

func (s segments) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	}
	return len(s[i]) > len(s[j])
}

func (s segments) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
