package singleelement

/**
540. 有序数组中的单一元素
给定一个只包含整数的有序数组，每个元素都会出现两次，唯有一个数只会出现一次，找出这个数。
示例 :
输入: [1,1,2,3,3,4,4,8,8]
输出: 2
输入: [3,3,7,7,10,11,11]
输出: 10
*/

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
