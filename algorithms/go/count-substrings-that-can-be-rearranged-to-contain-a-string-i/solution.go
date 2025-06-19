// 3298. 统计重新排列后包含另一个字符串的子字符串数目 II
// https://leetcode.cn/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-ii/

package count_substrings_that_can_be_rearranged_to_contain_a_string_i

func validSubstringCount(word1 string, word2 string) int64 {
	if len(word1) < len(word2) {
		return 0
	}
	word2map := make([]int, 26)
	for _, v := range word2 {
		word2map[key(byte(v))]++
	}
	word1map := make([]int, 26)
	var num int64
	start := 0
	left := 0
	right := 0
	n := len(word1)
	for ; right < n; right++ {
		v := word1[right]
		word1map[key(v)]++
		if isPrefix(word1map, word2map, v) {
			for ; left <= right; left++ {
				reduceKey := key(word1[left])
				if word1map[reduceKey] <= word2map[reduceKey] {
					break
				}
				word1map[reduceKey]--
			}
			num += int64((left - start + 1) * (n - right))
			word1map[key(word1[left])]--
			left++
			start = left
		}
	}
	return num
}

func key(x byte) int {
	return int(x - 'a')
}

func isPrefix(w1 []int, w2 []int, x byte) bool {
	k := key(x)
	if w1[k] < w2[k] {
		return false
	}
	for i := 0; i < len(w1); i++ {
		if w1[i] < w2[i] {
			return false
		}
	}
	return true
}

func ValidSubstringCount(word1 string, word2 string) int64 {
	return validSubstringCount(word1, word2)
}
