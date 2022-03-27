package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	// fast 速度是 slow的两倍
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	fast, slow.Next = slow.Next, nil
	left := sortList(head)
	right := sortList(fast)
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	virtualNode := &ListNode{}
	cur := virtualNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = cur.Next
		}
		cur = cur.Next
	}
	return virtualNode.Next
}
