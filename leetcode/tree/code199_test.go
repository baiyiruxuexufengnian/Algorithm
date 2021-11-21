package tree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		// 所有的节点都需遍历进队列，和层序遍历一样，只是加入到返回队列的时机变了
		// 在每一层上面，我们最后一个进入到节点的就是我们从右看到的节点
		for i := 0; i < size; i ++ {
			if i == size - 1 {
				ret = append(ret, que[i].Val)
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
	return ret
}
