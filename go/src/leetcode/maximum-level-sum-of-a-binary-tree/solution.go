// 1161. 最大层内元素和
// https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree

package maximum_level_sum_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 1
	}
	// 最大值的层级
	ml := 1
	// 最大值
	mv := root.Val
	l := 1
	b0 := []*TreeNode{root}
	var b1 []*TreeNode
	for len(b0) > 0 {
		s := 0
		for _, n := range b0 {
			s += n.Val
			if n.Left != nil {
				b1 = append(b1, n.Left)
			}
			if n.Right != nil {
				b1 = append(b1, n.Right)
			}
		}
		if s > mv {
			mv = s
			ml = l
		}
		b0, b1 = b1, b0[0:0]
		l++
	}
	return ml
}

func MaxLevelSum(root *TreeNode) int {
	return maxLevelSum(root)
}
