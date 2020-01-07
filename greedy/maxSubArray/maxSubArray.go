package maxSubArray

/**
53. 最大子序和

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

// 贪心
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, maxSum := -1<<31, -1<<31
	for _, num := range nums {
		sum = max(sum+num, num)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 分治法
func maxSubArray2(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}
func helper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	p := (left + right) / 2
	leftSum := helper(nums, left, p)
	rightSum := helper(nums, p+1, right)
	crossSum := crossSum(nums, left, right, p)

	return max(max(leftSum, rightSum), crossSum)
}

func crossSum(nums []int, left, right, p int) int {
	if left == right {
		return nums[left]
	}
	leftSubsum := -1 << 31
	currSum := 0
	for i := p; i > left-1; i-- {
		currSum += nums[i]
		leftSubsum = max(leftSubsum, currSum)
	}

	rightSubsum := -1 << 31
	currSum = 0
	for i := p + 1; i < right; i++ {
		currSum += nums[i]
		rightSubsum = max(rightSubsum, currSum)
	}
	return leftSubsum + rightSubsum
}
