package Q22

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkTree(root.Left, root.Right)
}

func checkTree(root1 *TreeNode, root2 *TreeNode) bool {
	// 一直比直到空
	if root1 == nil && root2 == nil {
		return true
	}
	if (root1 == nil && root2 != nil) ||
		(root1 != nil && root2 == nil) {
		return false
	} else if root1.Val != root2.Val {
		return false
	}
	return checkTree(root1.Left, root2.Right) && checkTree(root1.Right, root2.Left)
}
