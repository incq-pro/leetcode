// 756. 金字塔转换矩阵
// https://leetcode.cn/problems/pyramid-transition-matrix

package pyramid_transition_matrix

import (
	"fmt"
	"time"
)

func pyramidTransition(bottom string, allowed []string) bool {
	next := make(map[string][]string, len(allowed))
	for _, al := range allowed {
		next[al[:2]] = append(next[al[:2]], al[2:])
	}
	var dfs func(string, string) bool
	dfs = func(b string, u string) bool {
		bl, ul := len(b), len(u)
		if bl == 1 {
			return true
		} else if ul == bl-1 {
			return dfs(u, "")
		}
		if n, ok := next[b[ul:ul+2]]; ok {
			for _, nn := range n {
				if dfs(b, u+nn) {
					return true
				}
			}
		}
		return false
	}
	return dfs(bottom, "")
}

func PyramidTransition(bottom string, allowed []string) bool {
	now := time.Now()
	defer func() {
		fmt.Println(time.Now().Unix() - now.Unix())
	}()
	return pyramidTransition(bottom, allowed)
}
