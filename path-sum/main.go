package Q25

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}
	ls := hasPathSum(root.Left, targetSum-root.Val)
	rs := hasPathSum(root.Right, targetSum-root.Val)
	return ls || rs
}
