package checkPossibility

/**
665. 非递减数列
给定一个长度为 n 的整数数组，你的任务是判断在最多改变 1 个元素的情况下，该数组能否变成一个非递减数列。
我们是这样定义一个非递减数列的： 对于数组中所有的 i (1 <= i < n)，满足 array[i] <= array[i + 1]。

示例 1:
输入: [4,2,3]
输出: True
解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
*/

func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 1; i < len(nums) && cnt < 2; i++ {
		if nums[i] >= nums[i-1] {
			continue
		}
		cnt++
		// 当 nums[i] < nums[i-2]时，
		// 修改 nums[i - 1] = nums[i] 不能使数组成为非递减数组，只能修改 nums[i] = nums[i - 1]
		if i-2 >= 0 && nums[i-2] > nums[i] {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[i]
		}
	}
	return cnt <= 1
}
