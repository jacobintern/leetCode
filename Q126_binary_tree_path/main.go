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
	res = append(res, traverslNode(root.Left, path, res)...)
	res = append(res, traverslNode(root.Right, path, res)...)

	return res
}

func traverslNode(root *TreeNode, path string, res []string) []string {
	if root.Left == nil && root.Right == nil {
		res = append(res, path)
		return res
	}
	path = path + "->" + strconv.Itoa(root.Val)

	if root.Left != nil {
		return traverslNode(root.Left, path, res)
	}

	return traverslNode(root.Right, path, res)
}
