package tree

import "math"

// 非递归版本
func maxDepthN1(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	que := make([]*Node, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		depth ++
		for i := 0; i < size; i ++ {
			for _, v := range que[i].Children {
				if v != nil {
					que = append(que, v)
				}
			}
		}
		que = que[size:]
	}
	return depth
}

func maxDepthN(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	for _, n := range root.Children {
		tmp := maxDepthN(n) + 1
		depth = int(math.Max(float64(depth), float64(tmp)))
	}
	return depth
}

