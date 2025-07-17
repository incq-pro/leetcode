// 2. 两数相加
// https://leetcode.cn/problems/add-two-numbers

package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{0, nil}
	p := r
	c := 0
	for {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		v := v1 + v2 + c
		p.Val = v % 10
		c = v / 10
		if l1 == nil && l2 == nil {
			break
		}
		p.Next = &ListNode{0, nil}
		p = p.Next
	}
	if c > 0 {
		p.Next = &ListNode{c, nil}
	}
	return r
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}

func String(p *ListNode) string {
	return ""
}
