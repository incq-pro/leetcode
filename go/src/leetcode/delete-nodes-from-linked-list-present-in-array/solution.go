// 3217. 从链表中移除在数组中存在的节点
// https://leetcode.cn/problems/delete-nodes-from-linked-list-present-in-array

package delete_nodes_from_linked_list_present_in_array

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}
	var pp *ListNode = nil
	var ans *ListNode = nil
	var p = head
	for p != nil {
		if !m[p.Val] {
			if pp == nil {
				pp = p
				ans = p
			} else {
				pp.Next = p
				pp = p
			}
		}
		p = p.Next
	}
	if pp != nil {
		pp.Next = p
	}
	return ans
}

func modifiedListV2(nums []int, head *ListNode) *ListNode {
	exists := make(map[int]bool, len(nums))
	for _, num := range nums {
		exists[num] = true
	}
	s := &ListNode{Next: head}
	p := s
	for p.Next != nil {
		if exists[p.Next.Val] {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return s.Next
}

func ModifiedList(nums []int, head *ListNode) *ListNode {
	return nil
}
