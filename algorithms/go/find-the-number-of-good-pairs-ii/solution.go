// 3164. 优质数对的总数 II
// https://leetcode.cn/problems/find-the-number-of-good-pairs-ii/

package find_the_number_of_good_pairs_ii

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	count := make(map[int]int)
	max1 := nums1[0]
	for _, x := range nums1 {
		count[x]++
		if x > max1 {
			max1 = x
		}
	}
	count2 := make(map[int]int)
	for _, x := range nums2 {
		count2[x]++
	}
	var total int64 = 0
	for a, cnt2 := range count2 {
		for b := a * k; b <= max1; b += a * k {
			if cnt1, ok := count[b]; ok {
				total += int64(cnt1 * cnt2)
			}
		}
	}
	return total
}

func NumberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	return numberOfPairs(nums1, nums2, k)
}
