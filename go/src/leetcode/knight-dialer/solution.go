// 935. 骑士拨号器
// https://leetcode.cn/problems/knight-dialer/

package knight_dialer

func knightDialer(n int) int {
	if n == 1 {
		return 10
	}
	const m = 1e9 + 7
	a := []int{1, 1, 1, 1, 1, 0, 1, 1, 1, 1}
	b := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for j := 1; j < n; j++ {
		b[0] = (a[4] + a[6]) % m
		b[1] = (a[6] + a[8]) % m
		b[2] = (a[7] + a[9]) % m
		b[3] = (a[4] + a[8]) % m
		b[4] = (a[0] + a[3] + a[9]) % m
		b[6] = (a[0] + a[1] + a[7]) % m
		b[7] = (a[2] + a[6]) % m
		b[8] = (a[1] + a[3]) % m
		b[9] = (a[2] + a[4]) % m
		a, b = b, a
	}
	r := 0
	for _, v := range a {
		r = (r + v) % m
	}
	return r
}

func KnightDialer(n int) int {
	return knightDialer(n)
}
