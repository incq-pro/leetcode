// https://leetcode.cn/problems/partition-equal-subset-sum
// 416. 分割等和子集

package partition_equal_subset_sum

func canPartition(nums []int) bool {
	s := 0
	for _, num := range nums {
		s += num
	}
	if s%2 != 0 {
		return false
	}
	W := s / 2
	dp := make([]int, W+1)
	for _, num := range nums {
		for j := W; j > 0; j-- {
			if j-num >= 0 {
				dp[j] = max(dp[j], dp[j-num]+num)
			} else {
				break
			}
		}
	}
	return dp[W] == W
}

func CanPartition(nums []int) bool {
	return canPartition(nums)
}
