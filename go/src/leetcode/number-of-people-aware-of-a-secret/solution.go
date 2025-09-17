// 2327. 知道秘密的人数
// https://leetcode.cn/problems/number-of-people-aware-of-a-secret/

package number_of_people_aware_of_a_secret

func peopleAwareOfSecret(n int, delay int, forget int) int {
	mod := 1000000007

	dp0 := make([]int, forget)
	dp1 := make([]int, forget)
	dp0[0] = 1
	for i := 1; i < n; i++ {
		dp1[0] = 0
		for j := 1; j < forget; j++ {
			dp1[j] = dp0[j-1]
			if j >= delay {
				dp1[0] = (dp1[0] + dp1[j]) % mod
			}
		}
		dp0, dp1 = dp1, dp0
	}
	v := 0
	for i := 0; i < len(dp0); i++ {
		v = (v + dp0[i]) % mod
	}
	return v
}
func PeopleAwareOfSecret(n int, delay int, forget int) int {
	return peopleAwareOfSecret(n, delay, forget)
}
