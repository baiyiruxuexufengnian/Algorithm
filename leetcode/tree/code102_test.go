package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		arr := make([]int, 0)
		for i := 0; i < size; i ++ {
			arr = append(arr, que[i].Val)
			if que[i].Left != nil {
				que = append(que, que[i].Left)
			}
			if que[i].Right != nil {
				que = append(que, que[i].Right)
			}
		}
		que = que[size:]
		ret = append(ret, arr)
	}
	return ret
}

//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//	que1 := make([]*TreeNode, 0)
//	que2 := make([]*TreeNode, 0)
//	ret := make([][]int, 0)
//	que1 = append(que1, root)
//	for len(que1) != 0 || len(que2) != 0 {
//		if len(que1) != 0 && len(que2) == 0 {
//			arr := make([]int, 0)
//			for len(que1) != 0 {
//				top := que1[0]
//				que1 = que1[1:]
//				arr = append(arr, top.Val)
//				if top.Left != nil {
//					que2 = append(que2, top.Left)
//				}
//				if top.Right != nil {
//					que2 = append(que2, top.Right)
//				}
//			}
//			ret = append(ret, arr)
//		}
//
//		if len(que1) == 0 && len(que2) != 0 {
//			arr := make([]int, 0)
//			for len(que2) != 0 {
//				top := que2[0]
//				que2 = que2[1:]
//				arr = append(arr, top.Val)
//				if top.Left != nil {
//					que1 = append(que1, top.Left)
//				}
//				if top.Right != nil {
//					que1 = append(que1, top.Right)
//				}
//			}
//			ret = append(ret, arr)
//		}
//	}
//	return ret
//}
