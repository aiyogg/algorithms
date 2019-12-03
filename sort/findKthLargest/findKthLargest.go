package sort

/**
215. 数组中的第K个最大元素
*/

import "container/heap"

func findKthLargest(nums []int, k int) int {
	temp := highHeap(nums)
	h := &temp
	heap.Init(h)

	if k == 1 {
		return (*h)[0]
	}
	for i := 1; i < k; i++ {
		heap.Remove(h, 0)
	}
	return (*h)[0]
}

// highHeap 实现了 heap 的接口
type highHeap []int

func (h highHeap) Len() int {
	return len(h)
}
func (h highHeap) Less(i, j int) bool {
	// h[i] > h[j] 使得 h[0] == max(h...)
	return h[i] > h[j]
}
func (h highHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *highHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *highHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}

/**
方法二： 快速选择
- 随机选择一个枢轴。
- 使用划分算法将枢轴放在数组中的合适位置 pos。将小于枢轴的元素移到左边，大于等于枢轴的元素移到右边。
- 比较 pos 和 N - k 以决定在哪边继续递归处理。
*/
func findKthLargest2(nums []int, k int) int {
	return kthLargest(nums, 0, len(nums)-1, k-1)
}

func kthLargest(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	pivot := partition(nums, left, right)
	if pivot == k {
		return nums[k]
	} else if pivot < k {
		return kthLargest(nums, pivot+1, right, k)
	}
	return kthLargest(nums, left, pivot-1, k)
}

/**
分区，将 slice 分为 分别是小于基准和大于等于基准的两部分
 */
func partition(nums []int, left, right int) int {
	pivot := left
	for i := left; i < right; i++ {
 		if nums[i] > nums[right] {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot++
		}
	}

	nums[pivot], nums[right] = nums[right], nums[pivot]
	return pivot
}
