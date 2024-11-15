// 3162. 优质数对的总数 I
// https://leetcode.cn/problems/find-the-number-of-good-pairs-i/

package find_the_number_of_good_pairs_i

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	total := 0
	for _, y := range nums2 {
		yk := y * k
		for _, x := range nums1 {
			if x%yk == 0 {
				total++
			}
		}
	}
	return total
}

func NumberOfPairs(nums1 []int, nums2 []int, k int) int {
	return numberOfPairs(nums1, nums2, k)
}
