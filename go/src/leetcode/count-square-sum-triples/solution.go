// 1925. 统计平方和三元组的数目
// https://leetcode.cn/problems/count-square-sum-triple

package count_square_sum_triples

func countTriples(n int) int {
	ans := 0
	for k := 1; k <= n; k++ {
		v := k * k
		for l, r := 0, k-1; l <= r; {
			m := l*l + r*r
			if m == v {
				if l != r {
					ans += 2
				} else {
					ans += 1
				}
				l++
				r--
			} else if m > v {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}

func CountTriples(n int) int {
	return countTriples(n)
}
