// 464. 我能赢吗
// https://leetcode.cn/problems/can-i-win/

package can_i_win

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	dp := make([]int8, 1<<maxChoosableInteger)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(int, int) int8
	dfs = func(usedNum, curTotal int) (res int8) {
		v := &dp[usedNum]
		if *v != -1 {
			return *v
		}
		defer func() { *v = res }()
		for i := 0; i < maxChoosableInteger; i++ {
			if ((usedNum>>i)&1) == 0 && (curTotal+i+1 >= desiredTotal || dfs(usedNum|(1<<i), curTotal+i+1) == 0) {
				return 1
			}
		}
		return
	}
	return dfs(0, 0) == 1
}

func CanIWin(maxChoosableInteger int, desiredTotal int) bool {
	return canIWin(maxChoosableInteger, desiredTotal)
}
