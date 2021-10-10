package list

// Definition for singly-linked list.

type ListNode struct {
	Val int
	Next *ListNode
}

//// 其实不算使用了虚拟头的特性
//func removeElements(head *ListNode, val int) *ListNode {
//	if head == nil {
//		return head
//	}
//	for head != nil && head.Val == val {
//		head = head.Next
//	}
//	virtualHead := &ListNode{Next: head,Val: -1}
//	for head != nil && head.Next != nil {
//		if head.Next.Val == val {
//			head.Next = head.Next.Next
//		} else {
//			head = head.Next
//		}
//	}
//	return virtualHead.Next
//}

// 使用虚拟头的特性
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	virtualHead := &ListNode{Next: head,Val: -1}
	cur := virtualHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return virtualHead.Next
}