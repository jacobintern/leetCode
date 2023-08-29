package Q108

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	empty := TreeNode{}
	if root == &empty {
		return 0
	}

	l := findTilt(root.Left)
	r := findTilt(root.Right)

	return int(math.Abs(float64(l - r)))
}
