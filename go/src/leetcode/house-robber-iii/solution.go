// 337. 打家劫舍 III
// https://leetcode.cn/problems/house-robber-iii/

package house_robber_iii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	v := dfs(root)
	return max(v[0], v[1])
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l, r := dfs(root.Left), dfs(root.Right)
	selectMax := root.Val + l[1] + r[1]
	notSelectMax := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selectMax, notSelectMax}
}

func Rob(root *TreeNode) int {
	return rob(root)
}
