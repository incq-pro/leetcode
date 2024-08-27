package two_sum

func twoSum(nums []int, target int) []int {
	// m is nums value and index
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		left := target - num
		if preIndex, exits := m[left]; exits {
			return []int{preIndex, i}
		}
		m[num] = i
	}
	return nil
}

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}
