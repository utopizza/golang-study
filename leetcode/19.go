package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 0 {
		return head
	}

	slow, fast := head, head
	for ; n > 0 && fast.Next != nil; n-- {
		fast = fast.Next
	}
	if n > 1 {
		return head
	}
	if n == 1 {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}

	return head
}
