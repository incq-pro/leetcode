// 396. 旋转函数
// https://leetcode.cn/problems/rotate-function

package rotate_function

func maxRotateFunction(nums []int) int {
	n := len(nums)
	s := 0
	f0, f1 := 0, 0
	for i, num := range nums {
		s += num
		f0 += i * num
	}

	m := f0
	for i := 1; i < n; i++ {
		f1 = f0 + s - n*nums[n-i]
		if f1 > m {
			m = f1
		}
		f0 = f1
	}
	return m
}

func MaxRotateFunction(nums []int) int {
	return maxRotateFunction(nums)
}
