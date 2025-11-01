// 3289. 数字小镇中的捣蛋鬼
// https://leetcode.cn/problems/the-two-sneaky-numbers-of-digitville

package the_two_sneaky_numbers_of_digitville

func getSneakyNumbers(nums []int) []int {
	h := make([]int, len(nums))
	ans := make([]int, 0, 2)
	for _, num := range nums {
		if h[num] > 0 {
			ans = append(ans, num)
		} else {
			h[num] += 1
		}
	}
	return ans
}

func GetSneakyNumbers(nums []int) []int {
	return getSneakyNumbers(nums)
}
