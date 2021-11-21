package tree

func maxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		for i := 0; i < size; i ++ {
			if que[i].Left != nil {
				que = append(que, que[i].Left)
			}
			if que[i].Right != nil {
				que = append(que, que[i].Right)
			}
		}
		depth ++
		que = que[size:]
	}
	return depth
}
