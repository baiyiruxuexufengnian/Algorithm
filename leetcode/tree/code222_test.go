package tree

// func countNodes(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     count := 0
//     que := make([]*TreeNode, 0)
//     que = append(que, root)
//     for len(que) != 0 {
//         size := len(que)
//         count += size
//         for i := 0; i < size; i ++ {
//             if que[i].Left != nil {
//                 que = append(que, que[i].Left)
//             }
//             if que[i].Right != nil {
//                 que = append(que, que[i].Right)
//             }
//         }
//         // pop
//         que = que[size:]
//     }
//     return count
// }

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := countNodes(root.Left)
	r := countNodes(root.Right)
	return l + r + 1
}
