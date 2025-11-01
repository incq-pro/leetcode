// 3354. 使数组元素等于零
// https://leetcode.cn/problems/make-array-elements-equal-to-zero/

package make_array_elements_equal_to_zero

func countValidSelections(nums []int) int {
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			left, right := preSum[i], preSum[n-1]-preSum[i]
			if left == right {
				cnt += 2
			} else if left == right+1 || left == right-1 {
				cnt += 1
			}
		}
	}
	return cnt
}

func CountValidSelections(nums []int) int {
	return countValidSelections(nums)
}
