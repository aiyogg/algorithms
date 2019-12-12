package eraseOverlapIntervals

import (
	"algorithms/common"
	"sort"
)

/**
435. 无重叠区间
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
注意:
可以认为区间的终点总是大于它的起点。
区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

示例
输入: [ [1,2], [2,3], [3,4], [1,3] ]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。

思路
结尾越小，区间范围越小，故按结尾大小给各个区间排序，小区间在前
如果一个区间的开头比另一个区间的结尾小，则这两个区间有重叠，即需要一出区间数++
*/

type Interval = common.Interval

func eraseOverlapIntervals(intervals []Interval) int {
	end := -1 << 31
	res := 0
	if len(intervals) == 1 {
		return 0
	}

	sort.Sort(B(intervals))
	for _, interval := range intervals {
		if end <= interval.Start {
			end = interval.End
		} else {
			res++
		}
	}
	return res
}

// B 经过排序后， End 小的排在前面
type B []Interval

func (a B) Len() int           { return len(a) }
func (a B) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a B) Less(i, j int) bool { return a[i].End < a[j].End }
