// 2552. 统计上升四元组
// https://leetcode.cn/problems/count-increasing-quadruplets

package count_increasing_quadruplets

func countQuadruplets(nums []int) int64 {
	n := len(nums)
	pre := make([]int, n+1)
	ans := int64(0)
	for j := 0; j < n; j++ {
		suf := 0
		for k := n - 1; k > j; k-- {
			if nums[k] < nums[j] {
				ans += int64(pre[nums[k]] * suf)
			} else {
				suf++
			}
		}
		for x := nums[j] + 1; x <= n; x++ {
			pre[x] = pre[x] + 1
		}
	}
	return ans
}

func CountQuadruplets(nums []int) int64 {
	return countQuadruplets(nums)
}
