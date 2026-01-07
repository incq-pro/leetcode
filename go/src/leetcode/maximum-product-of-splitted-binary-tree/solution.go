// 1339. 分裂二叉树的最大乘积
// https://leetcode.cn/problems/maximum-product-of-splitted-binary-tree

package maximum_product_of_splitted_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	afterOrder(root)
	total := int64(root.Val)
	m := int64(0)
	var preOrder func(root *TreeNode)
	preOrder = func(r *TreeNode) {
		if r.Left != nil {
			m = max(m, (total-int64(r.Left.Val))*int64(r.Left.Val))
			preOrder(r.Left)
		}
		if r.Right != nil {
			m = max(m, (total-int64(r.Right.Val))*int64(r.Right.Val))
			preOrder(r.Right)
		}
	}
	preOrder(root)
	return int(m % (1e9 + 7))
}

func afterOrder(root *TreeNode) {
	s := root.Val
	if root.Left != nil {
		afterOrder(root.Left)
		s += root.Left.Val
	}
	if root.Right != nil {
		afterOrder(root.Right)
		s += root.Right.Val
	}
	root.Val = s
}

func MaxProduct(root *TreeNode) int {
	return maxProduct(root)
}
