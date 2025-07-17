// 2239. 找到最接近 0 的数字
// https://leetcode.cn/problems/find-closest-number-to-zero/

package find_closest_number_to_zero

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func findClosestNumber(nums []int) int {
	v, d := nums[0], abs(nums[0])
	for _, x := range nums {
		d2 := abs(x)
		if d2 < d {
			v, d = x, d2
		} else if d2 == d {
			if v < x {
				v = x
			}
		}
	}
	return v
}

func FindClosestNumber(nums []int) int {
	return findClosestNumber(nums)
}
