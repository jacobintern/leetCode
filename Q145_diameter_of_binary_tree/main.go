package Q145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	traverse(root)
	return maxDiameter
}

func traverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := traverse(root.Left)
	right := traverse(root.Right)

	maxDiameter = max(maxDiameter, left+right)

	return max(left, right) + 1
}
