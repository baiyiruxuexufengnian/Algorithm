package tree

func minDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		depth ++
		for i := 0; i < size; i ++ {
			if que[i].Left == nil && que[i].Right == nil {
				return depth
			}
			if que[i].Left != nil {
				que = append(que, que[i].Left)
			}
			if que[i].Right != nil {
				que = append(que, que[i].Right)
			}
		}
		que = que[size:]
	}
	return depth
}
