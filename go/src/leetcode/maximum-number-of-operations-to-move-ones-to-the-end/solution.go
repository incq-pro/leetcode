// 3228. 将 1 移动到末尾的最大操作次数
// https://leetcode.cn/problems/maximum-number-of-operations-to-move-ones-to-the-end/

package maximum_number_of_operations_to_move_ones_to_the_end

func maxOperations(s string) int {
	ones := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			for i+1 < len(s) && s[i+1] == '0' {
				i++
			}
			ans += ones
		} else {
			ones++
		}
	}
	return ans
}

func MaxOperations(s string) int {
	return maxOperations(s)
}
