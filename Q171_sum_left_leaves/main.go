package Q171

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0

	if root.Left == nil && root.Right == nil {
		return sum
	}

	traverseNode(root, &sum)

	return sum
}

func traverseNode(root *TreeNode, sum *int) {
	if root.Left == nil && root.Right == nil {
		return
	}

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			*sum += root.Left.Val
		}
		traverseNode(root.Left, sum)
	}

	if root.Right != nil {
		traverseNode(root.Right, sum)
	}
}
