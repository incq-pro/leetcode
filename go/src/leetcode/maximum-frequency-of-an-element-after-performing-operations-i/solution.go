// 3346. 执行操作后元素的最高频率 I
// https://leetcode.cn/problems/maximum-frequency-of-an-element-after-performing-operations-i

package maximum_frequency_of_an_element_after_performing_operations_i

import (
	"slices"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	slices.Sort(nums)
	numCount := make(map[int]int, len(nums))
	for _, num := range nums {
		numCount[num] = numCount[num] + 1
	}

	leftBounder := func(value int) int {
		low, high := 0, len(nums)-1
		for low < high {
			mid := (low + high) / 2
			if nums[mid] < value {
				low = mid + 1
			} else {
				high = mid
			}
		}
		return low
	}

	rightBounder := func(value int) int {
		low, high := 0, len(nums)-1
		for low < high {
			mid := (low + high + 1) / 2
			if nums[mid] > value {
				high = mid - 1
			} else {
				low = mid
			}
		}
		return low
	}
	ans := 0
	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		l := leftBounder(i - k)
		r := rightBounder(i + k)
		tempAns := min(r-l+1, numOperations+numCount[i])
		ans = max(ans, tempAns)
	}
	return ans
}

func MaxFrequency(nums []int, k int, numOperations int) int {
	return maxFrequency(nums, k, numOperations)
}
