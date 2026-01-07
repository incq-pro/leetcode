// 413. 等差数列划分
// https://leetcode.cn/problems/arithmetic-slices/

package arithmetic_slices

func numberOfArithmeticSlices(nums []int) int {
	l, r := 0, 1
	n := len(nums)
	ans := 0
	for r < n {
		delta := nums[r] - nums[l]
		for r = r + 1; r < n && nums[r]-nums[r-1] == delta; r++ {
		}
		length := r - l
		l = r - 1
		ans += (length - 1) * (length - 2) / 2
	}
	return ans
}

func NumberOfArithmeticSlices(nums []int) int {
	return numberOfArithmeticSlices(nums)
}
