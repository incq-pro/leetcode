// 3171. 找到按位或最接近 K 的子数组
// https://leetcode.cn/problems/find-subarray-with-bitwise-or-closest-to-k/

package find_subarray_with_bitwise_or_closest_to_k

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	n := len(nums)
	bitsMaxPos := make([]int, 31)
	for i := range bitsMaxPos {
		bitsMaxPos[i] = -1
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		for j := 0; j <= 30; j++ {
			if nums[i]>>j&1 == 1 {
				bitsMaxPos[j] = i
			}
		}

		posToBit := make([][2]int, 0)
		for j := 0; j <= 30; j++ {
			if bitsMaxPos[j] != -1 {
				posToBit = append(posToBit, [2]int{bitsMaxPos[j], j})
			}
		}
		sort.Slice(posToBit, func(a, b int) bool {
			return posToBit[a][0] > posToBit[b][0]
		})

		val := 0
		for j, p := 0, 0; j < len(posToBit); p = j {
			for j < len(posToBit) && posToBit[j][0] == posToBit[p][0] {
				val |= 1 << posToBit[j][1]
				j++
			}
			res = min(res, int(math.Abs(float64(val-k))))
		}
	}
	return res
}

func MinimumDifference(nums []int, k int) int {
	return minimumDifference(nums, k)
}
