// 2438. 二的幂数组中查询范围内的乘积
// https://leetcode.cn/problems/range-product-queries-of-powers

package range_product_queries_of_powers

func productQueries(n int, queries [][]int) []int {
	var m int64 = 1000000007
	var powers []int
	for i := 0; i < 32; i++ {
		if n&(1<<i) != 0 {
			powers = append(powers, 1<<i)
		}
	}
	ret := make([]int, len(queries))
	for i, q := range queries {
		var v int64 = 1
		for j := q[0]; j <= q[1]; j++ {
			v = (v * int64(powers[j])) % m
		}
		ret[i] = int(v)
	}
	return ret
}

func ProductQueries(n int, queries [][]int) []int {
	return productQueries(n, queries)
}
