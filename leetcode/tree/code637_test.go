package tree

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	que := make([]*TreeNode, 0)
	ret := make([]float64, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		sum := 0
		for i := 0; i < size; i ++ {
			sum += que[i].Val
			if que[i].Left != nil {
				que = append(que, que[i].Left)
			}
			if que[i].Right != nil {
				que = append(que, que[i].Right)
			}
		}
		que = que[size:]
		ret = append(ret, float64(sum) / float64(size))
	}
	return ret
}

//func averageOfLevels(root *TreeNode) []float64 {
//	if root == nil {
//		return nil
//	}
//	que := make([]*TreeNode, 0)
//	ret := make([]float64, 0)
//	que = append(que, root)
//	for len(que) != 0 {
//		size := len(que)
//		sum := 0
//		for i := 0; i < size; i ++ {
//			top := que[0]
//			sum += top.Val
//			if top.Left != nil {
//				que = append(que, top.Left)
//			}
//			if top.Right != nil {
//				que = append(que, top.Right)
//			}
//			que = que[1:]
//		}
//		ret = append(ret, float64(sum) / float64(size))
//	}
//	return ret
//}
