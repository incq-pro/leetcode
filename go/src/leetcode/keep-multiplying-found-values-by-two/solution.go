// 2154. 将找到的值乘以 2
// https://leetcode.cn/problems/keep-multiplying-found-values-by-two

package keep_multiplying_found_values_by_two

func findFinalValue(nums []int, original int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}
	for {
		if m[original] {
			original *= 2
		} else {
			break
		}
	}
	return original
}

func FindFinalValue(nums []int, original int) int {
	return findFinalValue(nums, original)
}
