// 2145. 统计隐藏数组数目
// https://leetcode.cn/problems/count-the-hidden-sequences

package count_the_hidden_sequences

func numberOfArrays(differences []int, lower int, upper int) int {
	l, u := lower, upper
	for _, d := range differences {
		l = max(lower, l-d)
		u = min(upper, u-d)
		if l > u {
			return 0
		}
	}
	return u - l + 1
}

func NumberOfArrays(differences []int, lower int, upper int) int {
	return numberOfArrays(differences, lower, upper)
}
