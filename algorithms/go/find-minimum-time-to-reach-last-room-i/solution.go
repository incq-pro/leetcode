// 3341. 到达最后一个房间的最少时间 I
// https://leetcode.cn/problems/find-minimum-time-to-reach-last-room-i/

package find_minimum_time_to_reach_last_room_i

func minTimeToReach(moveTime [][]int) int {
	n := len(moveTime)
	m := len(moveTime[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 1; i < m; i++ {
		dp[0][i] = max(dp[0][i-1], moveTime[0][i]) + 1
	}
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], moveTime[i][0]) + 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = max(moveTime[i][j], min(dp[i-1][j], dp[i][j-1])) + 1
		}
	}
	return dp[n-1][m-1]
}

func MinTimeToReach(moveTime [][]int) int {
	return minTimeToReach(moveTime)
}
