package Q144

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var r int
var preNode *TreeNode

func getMinimumDifference(root *TreeNode) int {
	preNode = nil
	r = math.MaxInt32
	traverse(root)
	return r
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	if preNode != nil {
		r = int(min(r, abs(preNode.Val, root.Val)))
	}
	preNode = root
	traverse(root.Right)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
