package sqrt

/**
69. x 的平方根
实现 int sqrt(int x) 函数。
计算并返回 x 的平方根，其中 x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

由于 x 平方根的整数部分 ans 是满足 k^2 <= x 的最大 k 值，因此我们可以对 k 进行二分查找，从而得到答案。
二分查找的下界为 0，上界可以粗略地设定为 x。
在二分查找的每一步中，我们只需要比较中间元素 mid 的平方与 x 的大小关系，并通过比较的结果调整上下界的范围。
由于我们所有的运算都是整数运算，不会存在误差，因此在得到最终的答案 ans 后，
也就不需要再去尝试 ans+1 了。
*/

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	var l, r = 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
