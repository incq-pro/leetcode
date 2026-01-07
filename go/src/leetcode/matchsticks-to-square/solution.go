// 473. 火柴拼正方形
// https://leetcode.cn/problems/matchsticks-to-square

package matchsticks_to_square

import "fmt"

func makesquare(matchsticks []int) bool {
	total := 0
	for _, num := range matchsticks {
		total += num
	}
	if total%4 != 0 {
		return false
	}
	tLen := total / 4
	n := len(matchsticks)
	dp := make([]int, 1<<n)
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	for s := 1; s < len(dp); s++ {
		for j, num := range matchsticks {
			if (s & (1 << j)) == 0 {
				continue
			}
			s0 := s & ^(1 << j)
			if dp[s0] >= 0 && dp[s0]+num <= tLen {
				dp[s] = (dp[s0] + num) % tLen
			}
		}
	}
	return dp[len(dp)-1] == 0
}

func makesquareSlow(matchsticks []int) bool {
	total := 0
	for _, m := range matchsticks {
		total += m
	}
	if total%4 != 0 {
		return false
	}
	s := total / 4
	dp := make(map[string]bool)
	var dfs func([]int, int) bool
	dfs = func(sides []int, i int) bool {
		key := fmt.Sprintf("%d:%d:%d:%d:%d", sides[0], sides[1], sides[2], sides[3], i)
		if v, ok := dp[key]; ok {
			return v
		}
		if i == len(matchsticks) {
			for _, v := range sides {
				if v != s {
					dp[key] = false
					return false
				}
				dp[key] = true
				return true
			}
		}

		m := matchsticks[i]
		for j := 0; j < 4; j++ {
			if sides[j]+m <= s {
				sides[j] += m
				if dfs(sides, i+1) {
					dp[key] = true
					return true
				}
				sides[j] -= m
			}
		}
		dp[key] = false
		return false
	}
	return dfs(make([]int, 4), 0)
}

func Makesquare(matchsticks []int) bool {
	return makesquare(matchsticks)
}
