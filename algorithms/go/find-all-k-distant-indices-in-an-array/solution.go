// 2200. 找出数组中的所有 K 近邻下标
// https://leetcode.cn/problems/find-all-k-distant-indices-in-an-array/

package find_all_k_distant_indices_in_an_array

func findKDistantIndices(nums []int, key int, k int) []int {
	res := make([]int, 0, len(nums))
	i := 0
	for j, x := range nums {
		if x == key {
			for i = max(i, j-k); i <= j+k; i++ {
				res = append(res, i)
			}
		}
	}
	return res
}

func FindKDistantIndices(nums []int, key int, k int) []int {
	return findKDistantIndices(nums, key, k)
}
