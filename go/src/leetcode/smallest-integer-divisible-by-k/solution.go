// 1015. 可被 K 整除的最小整数
// https://leetcode.cn/problems/smallest-integer-divisible-by-k

package smallest_integer_divisible_by_k

func smallestRepunitDivByK(k int) int {
	ans := 0
	a := 0
	seen := make(map[int]bool, k)
	for {
		a = (a*10 + 1) % k
		ans++
		if a == 0 {
			break
		} else if seen[a] {
			return -1
		}
		seen[a] = true
	}
	return ans
}

func SmallestRepunitDivByK(k int) int {
	return smallestRepunitDivByK(k)
}
