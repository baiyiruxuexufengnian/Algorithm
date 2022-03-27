package dp

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 记忆化搜索
//  var mem = make(map[*TreeNode]int)
// func rob(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     if root.Left == nil && root.Right == nil {
//         return root.Val
//     }
//     if val, ok := mem[root]; ok {
//         return val
//     }
//     money := root.Val
//     if root.Left != nil {
//         money += rob(root.Left.Left) + rob(root.Left.Right)
//     }
//     if root.Right != nil {
//         money += rob(root.Right.Left) + rob(root.Right.Right)
//     }
//     money2 := rob(root.Left) + rob(root.Right)
//     if money > money2 {
//         mem[root] = money
//     } else {
//         mem[root] = money2
//     }
//     return mem[root]
// }

func robIII(root *TreeNode) int {
	result := robCore(root)
	return max(result[0], result[1])
}

// 动态规划解法
// 返回数组，index=0 代表不偷当前节点可以获利的最大值，index=1代表偷当前节点获利的最大值
func robCore(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := robCore(root.Left)
	right := robCore(root.Right)
	result := make([]int, 2)
	// 当前节点不偷 == > 回溯到该节点时，其左右孩子可以偷，但是要取其左右孩子作为根节点时，它们在自己可偷和不可以偷的情况下能获取的最大利润
	result[0] = max(left[0], left[1]) + max(right[0], right[1])
	// 当前节点要偷 == > 其左右孩子不可以偷，则只能取在左右孩子不能偷的情况下可以获利的最大值
	result[1] = root.Val + left[0] + right[0]
	return result
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
