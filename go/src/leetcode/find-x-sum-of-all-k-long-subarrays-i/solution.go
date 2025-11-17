// 3318. 计算子数组的 x-sum I
// https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-i/

package find_x_sum_of_all_k_long_subarrays_i

import "sort"

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	cnt := make(map[int]int)
	for i := 0; i < k; i++ {
		cnt[nums[i]]++
	}
	for i := 0; i < n-k+1; i++ {
		if i > 0 {
			cnt[nums[i-1]]--
			cnt[nums[i+k-1]]++
		}
		xx := make([][]int, 0, len(cnt))
		for kk, v := range cnt {
			xx = append(xx, []int{kk, v})
		}
		sort.Slice(xx, func(a, b int) bool {
			if xx[a][1] != xx[b][1] {
				return xx[a][1] > xx[b][1]
			}
			return xx[a][0] > xx[b][0]
		})
		for j := 0; j < x && j < len(xx); j++ {
			ans[i] += xx[j][0] * xx[j][1]
		}
	}
	return ans
}

func FindXSum(nums []int, k int, x int) []int {
	return findXSum(nums, k, x)
}
