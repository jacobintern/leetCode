package Q191

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	minDiff := math.MaxInt32
	var prev *int // 用 *int 來記錄前一個節點的值
	inorder(root, &prev, &minDiff)
	return minDiff
}

func inorder(node *TreeNode, prev **int, minDiff *int) {
	if node == nil {
		return
	}

	inorder(node.Left, prev, minDiff)

	if *prev != nil {
		*minDiff = min(*minDiff, node.Val-**prev)
	}
	*prev = &node.Val // 更新 prev 為當前節點值

	inorder(node.Right, prev, minDiff)
}
