package Q118

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 定義一個 helper 函數計算樹的高度
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := height(node.Left)
		rightHeight := height(node.Right)
		// 如果任何子樹不平衡，或者左右子樹的高度差超過 1，則返回 -1
		if leftHeight == -1 || rightHeight == -1 || math.Abs(float64(leftHeight-rightHeight)) > 1 {
			return -1
		}
		return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
	}

	// 檢查整棵樹是否平衡
	return height(root) != -1
}
