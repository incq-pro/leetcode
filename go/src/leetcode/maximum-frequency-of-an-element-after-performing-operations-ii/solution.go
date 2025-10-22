// 3347. 执行操作后元素的最高频率 II
// https://leetcode.cn/problems/maximum-frequency-of-an-element-after-performing-operations-ii

package maximum_frequency_of_an_element_after_performing_operations_ii

import "slices"

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

	done := make(map[int]int, len(nums)*2)
	ansFunc := func(i int) int {
		if a, ok := done[i]; ok {
			return a
		}
		l := leftBounder(i - k)
		r := rightBounder(i + k)
		tempAns := min(r-l+1, numOperations+numCount[i])
		done[i] = tempAns
		return tempAns
	}

	ans := 0
	for _, num := range nums {
		ans = max(ans, ansFunc(num))
		if num-k >= nums[0] {
			ans = max(ans, ansFunc(num-k))
		}
		if num+k <= nums[len(nums)-1] {
			ans = max(ans, ansFunc(num+k))
		}
	}
	return ans
}

func MaxFrequency(nums []int, k int, numOperations int) int {
	return maxFrequency(nums, k, numOperations)
}
