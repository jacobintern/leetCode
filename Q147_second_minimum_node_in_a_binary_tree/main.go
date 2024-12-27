package Q147

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var min, min2 int

func findSecondMinimumValue(root *TreeNode) int {
	min, min2 = root.Val, math.MaxInt
	traverse(root)

	if min2 == math.MaxInt {
		return -1
	}

	return min2
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Val < min {
		min = root.Val
	}
	if root.Val > min && root.Val < min2 {
		min2 = root.Val
	}

	traverse(root.Left)
	traverse(root.Right)
}
