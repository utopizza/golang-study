package leetcode

func swapPairs(head *ListNode) *ListNode {
	addedHead := &ListNode{
		Val:  0,
		Next: head,
	}
	h := addedHead
	for h != nil && h.Next != nil && h.Next.Next != nil {
		p, q := h.Next, h.Next.Next
		h.Next = q
		p.Next = q.Next
		q.Next = p
		h = p
	}
	return addedHead.Next
}
