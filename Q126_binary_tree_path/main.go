package Q126

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res = []string{}
	traverslNode(root, "")

	return res
}

func traverslNode(root *TreeNode, path string) string {
	if path == "" {
		path = strconv.Itoa(root.Val)
	} else {
		path = path + "->" + strconv.Itoa(root.Val)
	}

	if root.Left != nil {
		traverslNode(root.Left, path)
	}
	if root.Right != nil {
		traverslNode(root.Right, path)
	}

	if root.Left == nil && root.Right == nil {
		res = append(res, path)
	}

	return path
}
