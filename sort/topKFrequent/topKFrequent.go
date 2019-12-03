package topKFrequent

import "sort"

/**
347. 前 K 个高频元素
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
*/

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	// 统计次数, key 为数字， value 为出现的次数
	rec := make(map[int]int, len(nums))

	for _, n := range nums {
		rec[n]++
	}
	// 对出现次数进行排序
	counts := make([]int, 0, len(rec))
	for _, c := range rec {
		counts = append(counts, c)
	}
	sort.Ints(counts)

	// 由于要找前 K 个高频元素，counts 又是从低频到高频的次数， 故 min 次以上才行
	min := counts[len(counts)-k]
	for n, c := range rec {
		if c >= min {
			res = append(res, n)
		}
	}

	return res
}
