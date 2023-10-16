package Q129

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	cur := strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		return cur
	}

	left := tree2str(root.Left)
	right := tree2str(root.Right)
	cur += "(" + left + ")"
	if right != "" {
		cur += "(" + right + ")"
	}

	return cur
}
