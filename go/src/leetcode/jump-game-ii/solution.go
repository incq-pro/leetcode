// 45. 跳跃游戏 II
// https://leetcode.cn/problems/jump-game-ii/

package jump_game_ii

import (
	"math"
)

func jump(nums []int) int {
	n := len(nums)
	steps := make([]int, n)
	for i := range steps {
		steps[i] = math.MaxInt32
	}
	steps[0] = 0
	for i := 0; i < n; i++ {
		s := steps[i] + 1
		longest := min(i+nums[i], n-1)
		for j := i + 1; j <= longest; j++ {
			steps[j] = min(s, steps[j])
		}
	}
	return steps[n-1]
}

func Jump(nums []int) int {
	return jump(nums)
}
