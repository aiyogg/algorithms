package sortColors

/**
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:不能使用代码库中的排序函数来解决这道题。

示例:
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
*/

func sortColors(nums []int) {
	// 三个指针分别对应 0 的右侧，当前元素，2 的左侧
	// for 循环中， nums[p0:curr] 中始终全是 1
	// 循环结束后，
	// nums[:p0] 中全是 0
	// nums[p2:] 中全是 2
	p0, curr, p2 := 0, 0, len(nums)-1
	for curr <= p2 {
		switch nums[curr] {
		// 若 nums[curr] = 0 ：交换第 curr个 和 第p0个 元素，并将指针都向右移。
		case 0:
			nums[p0], nums[curr] = nums[curr], nums[p0]
			p0++
			curr++
		// 若 nums[curr] = 1 ：将指针curr右移。
		case 1:
			curr++
		// 若 nums[curr] = 2 ：交换第 curr个和第 p2个元素，并将 p2指针左移 。
		case 2:
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
		}
	}
}
