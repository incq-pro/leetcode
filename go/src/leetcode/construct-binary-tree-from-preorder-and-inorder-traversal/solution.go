// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(preorder []int, pl int, pr int, inorder []int, il int, ir int) *TreeNode {
	if pl > pr {
		return nil
	}
	p := &TreeNode{
		Val: preorder[pl],
	}
	i := indexOf(inorder, preorder[pl], il, ir)
	p.Left = build(preorder, pl+1, pl+(i-il), inorder, il, i-1)
	p.Right = build(preorder, pl+1+(i-il), pr, inorder, i+1, ir)
	return p
}

func indexOf(num []int, x int, l int, r int) int {
	for i := l; i <= r; i++ {
		if num[i] == x {
			return i
		}
	}
	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	p := build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
	return p
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	return buildTree(preorder, inorder)
}
