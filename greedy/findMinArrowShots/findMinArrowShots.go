package findMinArrowShots

import (
	"algorithms/common"
	"sort"
)

type Interval = common.Interval

func findMinArrowShots(points []Interval) int {
	if len(points) == 0 {
		return 0
	}

	sort.Sort(B(points))
	end := points[0].End
	res := 1

	for i := 1; i < len(points); i++ {
		if points[i].Start <= end {
			continue
		}
		res++
		end = points[i].End
	}
	return res
}

type B []Interval

func (b B) Len() int {
	return len(b)
}
func (b B) Less(i, j int) bool {
	return b[i].End < b[j].End
}
func (b B) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
