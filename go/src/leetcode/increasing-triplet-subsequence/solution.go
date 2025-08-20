// 334. 递增的三元子序列
// https://leetcode.cn/problems/increasing-triplet-subsequence/

package increasing_triplet_subsequence

import "fmt"

func increasingTriplet(nums []int) bool {
	a, b, ap := nums[0], nums[0], nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		v := nums[i]
		if v > b {
			if b > a {
				fmt.Println(a, b, v)
				return true
			} else {
				b = v
			}
		} else if v < b {
			if v > a {
				b = v
			} else if v <= a {
				if a == b {
					a, b, ap = v, v, v
				} else if v < ap {
					ap = v
				} else if v > ap {
					a, b = ap, v
				}
			}
		}
	}
	fmt.Println(a, b, ap)
	return false
}

func IncreasingTriplet(nums []int) bool {
	return increasingTriplet(nums)
}
