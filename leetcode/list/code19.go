package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n < 0 || head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	virtualHead := &ListNode{
		Val: -1,
		Next: head,
	}
	pre := virtualHead
	for fast != nil  && n > 0 {
		fast = fast.Next
		n --
	}

	for fast != nil {
		pre = slow
		fast = fast.Next
		slow = slow.Next
	}
	pre.Next = slow.Next
	return virtualHead.Next
}