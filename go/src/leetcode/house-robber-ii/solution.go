// 213. 打家劫舍 II
// https://leetcode.cn/problems/house-robber-ii/

package house_robber_ii

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	n := len(nums)
	return max(_rob(nums[0:n-1]), _rob(nums[1:n]))
}

func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(first+v, second)
	}
	return second
}

func Rob(nums []int) int {
	return rob(nums)
}
