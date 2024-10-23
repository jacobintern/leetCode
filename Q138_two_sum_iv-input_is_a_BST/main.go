package Q138

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return traverse(root, k, map[int]bool{})
}

func traverse(root *TreeNode, k int, m map[int]bool) bool {
	if root == nil {
		return false
	}

	if m[k-root.Val] {
		return true
	}

	m[root.Val] = true

	return traverse(root.Left, k, m) || traverse(root.Right, k, m)
}
