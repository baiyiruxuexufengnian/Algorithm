package tree

type Node1 struct {
    Val int
    Left *Node1
    Right *Node1
    Next *Node1
}

// 提交时记得将Node1换成Node
// 117题和这道题代码一模一样
func connect(root *Node1) *Node1 {
	if root == nil {
		return nil
	}
	que := make([]*Node1, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		for i := 0; i < size; i ++ {
			if i == size - 1 {
				que[i].Next = nil
			} else {
				que[i].Next = que[i + 1]
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
	return root
}
