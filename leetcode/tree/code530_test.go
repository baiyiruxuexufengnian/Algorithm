package tree

import (
	"fmt"
	"log"
	"math"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDiff := math.MaxInt64
	var pre *TreeNode
	var getMinimumDifferenceCore func(*TreeNode, *int)
	getMinimumDifferenceCore = func(cur *TreeNode, maxDiff *int) {
		if cur == nil {
			return
		}
		getMinimumDifferenceCore(cur.Left, maxDiff)
		if pre != nil && cur.Val - pre.Val < *maxDiff {
			fmt.Printf("max:%v\n", pre.Val)
			*maxDiff = cur.Val - pre.Val
		}
		pre = cur
		getMinimumDifferenceCore(cur.Right, maxDiff)
	}
	getMinimumDifferenceCore(root, &maxDiff)
	return maxDiff
}

type test struct {
	val int
}

func testFunc1(in *test) {
	log.Println("I am func1")
	testFunc2(in)
	in = &test{val: 90}
}

func testFunc2(in *test) {
	log.Println("I am func2")
}

func TestFunc(t *testing.T) {
	var in *test
	testFunc1(in)
	t.Log(in)
}

