package singleelement

func singleNonDuplicate(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h {
		m := l + (h-l)/2
		// 保证 l/h/m 都在偶数位，使得查找区间大小一直都是奇数
		if m%2 == 1 {
			m--
		}
		// 由于查找区间中元素个数为奇数，则当 nums[m] == nums[m+1] 时，查找区间转移到右侧
		if nums[m] == nums[m+1] {
			l = m + 2 // 跳过相等的一组
		} else {
			h = m
		}
	}
	return nums[l]
}
