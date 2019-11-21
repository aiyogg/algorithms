package main

/**
Merge Sorted Array
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	// 复制一份原始的 nums1
	copy(temp, nums1)

	j, k := 0, 0
	for i := 0; i < len(nums1); i++ {
		// nums2 遍历完了
		if k >= n {
			nums1[i] = temp[j]
			j++
			continue
		}
		// temp 遍历完了
		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}
		if temp[j] < nums2[k] {
			nums1[i] = temp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
}

//func main() {
//	nums1 := []int{1, 2, 3, 0, 0, 0}
//	nums2 := []int{2, 5, 6}
//	merge(nums1, 3, nums2, 3)
//	fmt.Printf("merge: %v", nums1)
//}
