package Q126

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	path := strconv.Itoa(root.Val)
	res := []string{}
	traverslNode(root.Left, path, res)
	traverslNode(root.Right, path, res)

	return res
}

func traverslNode(root *TreeNode, path string, res []string) {
	path = path + "->" + strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		res = append(res, path)
		return
	}

	if root.Left != nil {
		traverslNode(root.Left, path, res)
	}
	if root.Right != nil {
		traverslNode(root.Right, path, res)
	}
}
