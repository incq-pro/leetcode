// 3494. 酿造药水需要的最少总时间
// https://leetcode.cn/problems/find-the-minimum-amount-of-time-to-brew-potions/

package find_the_minimum_amount_of_time_to_brew_potions

func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	m := len(mana)
	// 完成时间
	t0 := make([]int64, n)
	t1 := make([]int64, n)
	t0[0] = int64(skill[0] * mana[0])
	for i := 1; i < n; i++ {
		t0[i] = t0[i-1] + int64(skill[i]*mana[0])
	}
	for j := 1; j < m; j++ {
		t1[n-1] = t0[n-1]
		for i := n - 2; i >= 0; i-- {
			t1[i] = max(t0[i], t1[i+1]-int64(skill[i]*mana[j]))
		}
		t1[0] += int64(skill[0] * mana[j])
		for i := 1; i < n; i++ {
			t1[i] = t1[i-1] + int64(skill[i]*mana[j])
		}
		t1, t0 = t0, t1
	}
	return t0[n-1]
}

func MinTime(skill []int, mana []int) int64 {
	return minTime(skill, mana)
}
