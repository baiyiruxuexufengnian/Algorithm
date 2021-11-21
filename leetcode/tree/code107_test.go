package tree

func levelOrderBottom(root *TreeNode) [][]int {
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
	left := 0
	right := len(ret) - 1
	for left < right {
		tmp := ret[left]
		ret[left] = ret[right]
		ret[right] = tmp
		left ++
		right --
	}
	return ret
}
