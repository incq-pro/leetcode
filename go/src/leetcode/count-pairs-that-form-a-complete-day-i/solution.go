// 3184. 构成整天的下标对数目 I
// https://leetcode.cn/problems/count-pairs-that-form-a-complete-day-i/

package count_pairs_that_form_a_complete_day_i

func countCompleteDayPairs(hours []int) int {
	n := len(hours)
	c := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (hours[i]+hours[j])%24 == 0 {
				c++
			}
		}
	}
	return c
}

func CountCompleteDayPairs(hours []int) int {
	return countCompleteDayPairs(hours)
}
