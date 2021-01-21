package leetcode

// ListNode Defintion
type ListNode struct {
	Val  int
	Next *ListNode
}

func makeLinkedList(input []int) *ListNode {
	if input == nil {
		return nil
	}
	head := &ListNode{}
	p := head
	for _, i := range input {
		p.Next = &ListNode{
			Val:  i,
			Next: nil,
		}
		p = p.Next
	}
	return head.Next
}

func compareLinkedList(l1, l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val != l2.Val {
				return false
			}
			l1 = l1.Next
			l2 = l2.Next
		} else {
			return false
		}
	}
	return true
}

func compareIntegerSlice(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
