package tree

type Node struct {
	Val int
	Children []*Node
}

func levelOrderII(root *Node) [][]int {
	if root == nil {
		return nil
	}
	que := make([]*Node, 0)
	ret := make([][]int, 0)
	que = append(que, root)
	for len(que) != 0 {
		size := len(que)
		arr := make([]int, 0)
		for i := 0; i < size; i ++ {
			arr = append(arr, que[i].Val)
			for _, v := range que[i].Children {
				if v != nil {
					que = append(que, v)
				}
			}
		}
		que = que[size:]
		ret = append(ret, arr)
	}
	return ret
}
