package tree

import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	que := make([]*TreeNode, 0)
	ret := make([]int, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		maxVal := math.MinInt32
		for i := 0; i < size; i ++ {
			if que[i].Val > maxVal {
				maxVal = que[i].Val
			}
			if que[i].Left != nil {
				que = append(que, que[i].Left)
			}
			if que[i].Right != nil {
				que = append(que, que[i].Right)
			}
		}
		que = que[size:]
		ret = append(ret, maxVal)
	}
	return ret
}
