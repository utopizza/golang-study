package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	p, q := l1, l2
	increment := 0
	head := &ListNode{}
	r := head
	for p != nil || q != nil || increment != 0 {
		var pVal, qVal int
		if p != nil {
			pVal = p.Val
			p = p.Next
		}
		if q != nil {
			qVal = q.Val
			q = q.Next
		}
		sum := pVal + qVal + increment
		increment = sum / 10
		r.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		r = r.Next
	}
	return head.Next
}
