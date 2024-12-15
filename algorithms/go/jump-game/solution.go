// 55. 跳跃游戏
// https://leetcode.cn/problems/jump-game/

package jump_game

func canJump(nums []int) bool {
	i := len(nums) - 1
	for i > 0 {
		found := false
		for j := i - 1; j >= 0; j-- {
			if j+nums[j] >= i {
				i = j
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return i == 0
}

func CanJump(nums []int) bool {
	return canJump(nums)
}
