package list

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	res := fast
	var pre *ListNode

	for slow != nil && fast != nil {
		slowNext := fast.Next
		tmp := fast.Next
		fast.Next = slow
		slow.Next = tmp
		if pre != nil {
			pre.Next = fast
		}
		pre = slow

		slow = slowNext
		if slow != nil {
			fast = slow.Next
		}
	}
	return res
}
