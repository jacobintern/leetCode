package Q124

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	count := 0
	traversalNode(root, &count)
	return count
}

func traversalNode(root *TreeNode, n *int) {
	if root == nil {
		return
	}
	if root != nil {
		*n++
		traversalNode(root.Left, n)
		traversalNode(root.Right, n)
	}
}
