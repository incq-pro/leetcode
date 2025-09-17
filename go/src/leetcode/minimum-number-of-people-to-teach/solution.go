// 1733. 需要教语言的最少人数
// https://leetcode.cn/problems/minimum-number-of-people-to-teach/

package minimum_number_of_people_to_teach

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	nocon := make(map[int]bool)
	for _, f := range friendships {
		mp := make(map[int]bool)
		for _, l := range languages[f[0]-1] {
			mp[l] = true
		}
		common := false
		for _, l := range languages[f[1]-1] {
			if mp[l] {
				common = true
				break
			}
		}
		if !common {
			nocon[f[0]] = true
			nocon[f[1]] = true
		}
	}

	maxcnt := 0
	cnt := make([]int, n+1)
	for p := range nocon {
		for _, l := range languages[p-1] {
			cnt[l]++
			maxcnt = max(maxcnt, cnt[l])
		}
	}
	return len(nocon) - maxcnt
}
func MinimumTeachings(n int, languages [][]int, friendships [][]int) int {
	return minimumTeachings(n, languages, friendships)
}
